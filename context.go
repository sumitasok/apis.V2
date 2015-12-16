package apis

import (
	"log"
	"os"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

type context struct {
	mgo *mgo.Session

	traceLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger

	router *httprouter.Router
}

func (c *context) setMgo(addr string) error {
	mgoSession, err := mgo.Dial(addr)

	c.mgo = mgoSession

	return err
}

func (c *context) setLog() {
	c.traceLog = log.New(os.Stdout,
		"TRACE: ",
		log.Ldate|log.Ltime) //|log.Lshortfile

	c.infoLog = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime) //|log.Lshortfile

	c.warningLog = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime) //|log.Lshortfile

	c.errorLog = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime) //|log.Lshortfile
}
