package apis

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// D is getter for context, C is setter for context
type D struct {
	c *context

	mgoClone *mgo.Session
	rw       *http.ResponseWriter
	req      *http.Request
}

func (d *D) GetMgoSession() *mgo.Session {
	if d.mgoClone != nil {
		return d.mgoClone
	}

	d.mgoClone = d.c.mgo.Clone()
	return d.mgoClone
}

func (d *D) Call(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {

	if d.mgoClone != nil {
		d.mgoClone.Close()
	}
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
