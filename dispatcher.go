package apis

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"encoding/gob"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// D is getter for context, C is setter for context
type D struct {
	c *context

	actions []action

	req       http.Request
	body      []byte
	urlParams httprouter.Params
}

type DB *mgo.Session

func (d *D) DbQuery(dbQuery func(DB, error) error) error {
	if d.c.mgo == nil {
		d.LogInfo("Mongo DB not initiated")
		return dbQuery(nil, errors.New("DB not initiated"))
	}

	mgoClone := d.c.mgo.Clone()
	defer mgoClone.Close()

	return dbQuery(mgoClone, nil)
}

func (d D) call(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	req.Close = true

	d.urlParams = params

	var resp interface{}
	var err error
	var status int

	for _, action := range d.actions {
		resp, err, status = action.Call(&d)
		if err != nil {
			d.Write(rw, resp, err, status)
			return
		}
		err = d.SetBody(resp)
		if err != nil {
			d.Write(rw, nil, err, 400)
			return
		}
	}

	d.Write(rw, resp, err, status)
	return
}

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

func (d *D) URLParam(key string) string {
	return d.urlParams.ByName(key)
}

func (d *D) QueryParam(key string) string {
	return d.req.URL.Query().Get(key)
}

func (d *D) Request() http.Request {
	return d.req
}

func (d *D) Body(i interface{}) error {
	body, err := d.reqToByteArray()
	if err != nil {
		d.LogInfo("Cannot convert Request body to byte array", err.Error())
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		d.LogInfo("data couldn't be marshalled", err.Error())
	}

	return err
}
func (d *D) SetBody(i interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(i)
	if err != nil {
		return err
	}

	d.body = buf.Bytes()
	return nil
}

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

func (d *D) LogTrace(trace ...interface{}) {
	d.c.traceLog.Println(trace)
}

func (d *D) LogInfo(info ...interface{}) {
	d.c.infoLog.Println(info)
}

func (d *D) LogWarning(warning ...interface{}) {
	d.c.warningLog.Println(warning)
}

func (d *D) LogError(err ...interface{}) {
	d.c.errorLog.Println(err)
}
