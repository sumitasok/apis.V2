package apis

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// D holds the complete details about an http-Request as well as the response application is making
// Also has a reference to the application context
// dispatcher is the collections of http.Request and chained controllers
type D struct {
	c *context

	actions []action

	req       *http.Request
	rw        http.ResponseWriter
	body      []byte
	urlParams httprouter.Params
	// data shared between chained controllers in a single dispatcher
	data map[string]interface{}
}

// SetData sets the data into the dispatcher scope
func (d *D) SetData(key string, val interface{}) {
	d.data[key] = val
}

// GetData gets the data from the dispatcher scope
func (d *D) GetData(key string) (interface{}, bool) {
	val, ok := d.data[key]
	return val, ok
}

// BodyByte every dispatcher has an http request, and a POST of PUT method will have data in the request.Body
// Body bytes returns the values in byte format
func (d *D) BodyByte() []byte {
	return d.body
}

// DB stands for mgo Sessions
// Will be soon replaced with a universal DB interface which will enable changing the DB but writing new implementations
// and without impacting the main code.
type DB *mgo.Session

// DbQuery takes an anonymous function and gives itself a copy of the db session
// and takes care of closing the same. Thus encapsulating db connections.
func (d *D) DbQuery(dbQuery func(DB, error) error) error {
	if d.c.mgo == nil {
		d.LogInfo("Mongo DB not initiated")
		return dbQuery(nil, errors.New("DB not initiated"))
	}

	mgoClone := d.c.mgo.Clone()
	defer mgoClone.Close()

	err := dbQuery(mgoClone, nil)

	return err
}

// call is used internally to recieve the http.Request, loop it through each chained controllers
// and return the data as response.
// in each chaining, if the controller returns an error, it returns the error as response
// If you want to exit and respond at a particular condition in controller, return error
func (d D) call(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var receivedAt, respondedAt time.Time
	var actualBody []byte

	if d.c.logRequest {
		receivedAt = time.Now()
	}

	req.Close = true

	d.urlParams = params
	d.req = req
	d.rw = rw

	if req.Method == "POST" || req.Method == "PUT" {
		body, _ := d.reqToByteArray()
		d.body = body
		actualBody = body
	}

	var resp interface{}
	var err error
	var status int

	for _, action := range d.actions {
		resp, err, status = action.Call(&d)
		if err != nil {
			d.Write(rw, resp, err, status)
			return
		}

		err = d.SetBody(resp) // if response is not nil, set it to resquest Body.
		if err != nil {
			d.Write(rw, nil, err, status)
			return
		}
	}

	d.Write(rw, resp, err, status)

	if d.c.logRequest {
		respondedAt = time.Now()
		timeTaken := respondedAt.Sub(receivedAt)
		d.LogInfo("Req: ", timeTaken.Seconds(), "sec", req.Method, req.Host, req.URL, string(actualBody), req.Referer(), "Resp", string(d.body))
	}

	return
}

// Write is used to write the response as a http.Response
// by marshalling it as json
// adding the error string
// adding the status code.
func (d *D) Write(rw http.ResponseWriter, resp interface{}, err error, status int) {
	response := map[string]interface{}{
		"data": resp,
	}
	response["status"] = status

	if err != nil {
		response["error"] = err.Error()
	}

	jData, jsonErr := json.Marshal(response)
	if jsonErr != nil {
		d.LogError("Json Marshalling failed", response)
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	rw.Write(jData)
}

// URLParam returns the parameters in URL /blog/:id -> key is "id"
func (d *D) URLParam(key string) string {
	return d.urlParams.ByName(key)
}

// QueryParam returns the parameters in Query /blog/:id?comment=1234 => key id "comment"
func (d *D) QueryParam(key string) string {
	return d.req.URL.Query().Get(key)
}

// Request returns the actual request, if program wants to use the actual request
// for any customisation/feature that the framework doesn't give
func (d *D) Request() *http.Request {
	return d.req
}

// ResponseWriter returns the reqponse writer
// application can use this to customise the response
func (d *D) ResponseWriter() http.ResponseWriter {
	return d.rw
}

// Body is used to unmarshal a json data comming as bytes in response.Body into a struct whole address is passed
func (d *D) Body(i interface{}) error {
	body := d.body
	if body == nil {
		d.LogInfo("body is empty")
	}

	err := json.Unmarshal(body, &i)
	if err != nil {
		d.LogInfo("data couldn't be marshalled", err.Error())
	}

	return err
}

// SetBody helps to replace the value that is coming in http.Request with the data being passed.
func (d *D) SetBody(i interface{}) error {
	// if response is nil, keeep the previous response intact.
	if i == nil {
		return nil
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err := enc.Encode(i)
	if err != nil {
		return err
	}

	d.body = buf.Bytes()
	return nil
}

// reqToByteArray converts the response.Body to byte data
func (d *D) reqToByteArray() ([]byte, error) {
	if len(d.body) > 0 {
		return d.body, nil
	}

	body, err := ioutil.ReadAll(d.req.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// LogTrace logs the trace
func (d *D) LogTrace(trace ...interface{}) {
	d.c.traceLog.Println(trace)
}

// LogInfo logs the info
func (d *D) LogInfo(info ...interface{}) {
	d.c.infoLog.Println(info)
}

// LogWarning logs warning
func (d *D) LogWarning(warning ...interface{}) {
	d.c.warningLog.Println(warning)
}

// LogError logs the Error
func (d *D) LogError(err ...interface{}) {
	d.c.errorLog.Println(err)
}
