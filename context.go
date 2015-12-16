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

	routes *routes
	router *httprouter.Router
}

type routes []*route

func (c *context) setMgo(addr string) error {
	mgoSession, err := mgo.Dial(addr)

	c.mgo = mgoSession

	return err
}

// https://www.reddit.com/r/golang/comments/283vpk/help_with_slices_and_passbyreference/
// The reference to the underlying array is a value, and in Go, everything is passed by value (even pointers/references).
// So your function that receives a slice is receiving a copy of the slice header.
// The slice header contains important information like starting address and size.
// When you append to the slice in your other function, the copy of the slice header gets modified, but the original calling function doesn't see that copy, it still has its own. That's why functions like append return the new value, which is the modified slice header.
func (c *context) addRoute(r *route) *context {
	_rs := *c.routes
	_rs = append(_rs, r)

	c.routes = &_rs
	return c
}

func (c *context) attachRoutes() {
	c.router = httprouter.New()

	_routes := *c.routes
	for i := range _routes {
		d := &D{c: c, actions: _routes[i].actions}
		switch _routes[i].method {
		case "GET":
			c.router.GET(_routes[i].url, d.Call)
			break
		case "POST":
			c.router.POST(_routes[i].url, d.Call)
			break
		case "PUT":
			c.router.PUT(_routes[i].url, d.Call)
			break
		case "DELETE":
			c.router.DELETE(_routes[i].url, d.Call)
			break
		default:
			break
		}
	}
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
