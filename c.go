package apis

import (
	"fmt"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Init() *C {
	c := &C{&context{}}
	c.SetLogger()

	c.router = httprouter.New()

	return c
}

// C is setter for context, D is geter for context
type C struct {
	*context
}

func (c *C) NewDispatcher() *D {
	return &D{c: c.context}
}

func (c *C) SetMgo(addr string) *C {
	err := c.setMgo(addr)

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(0)
	}

	return c
}

func (c *C) SetLogger() *C {
	c.setLog()

	return c
}

type route struct {
	context *C
	method  string
	url     string
	actions []action
}

func (r *route) Set(actions ...action) *C {
	// _routes := append(*r.context.routes, r)
	// r.context.routes = &_routes
	r.actions = actions

	d := r.context.NewDispatcher()

	r.context.router.GET(r.url, d.Call)

	return r.context
}

type action interface{}

func (c *C) Get(url string) *route {
	return &route{context: c, method: "GET", url: url}
}
