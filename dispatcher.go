package apis

import (
	"errors"
	"io/ioutil"
	"net/http"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// D is getter for context, C is setter for context
type D struct {
	c *context

	actions []action

	req http.Request
}

type DB *mgo.Session

func (d *D) DbQuery(dbQuery func(DB, error)) {
	if d.c.mgo == nil {
		d.LogInfo("Mongo DB not initiated")
		dbQuery(nil, errors.New("DB not initiated"))
		return
	}

	mgoClone := d.c.mgo.Clone()
	defer mgoClone.Close()

	dbQuery(mgoClone, nil)
	return
}

func (d D) call(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

func (d *D) reqToByteArray() ([]byte, error) {
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
