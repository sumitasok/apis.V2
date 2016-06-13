package apis

import (
	"fmt"
	"os"

	"net/http"
	"strconv"
)

// Init initiates a context for the web application
// Init() -> returns the context on which the whole app is going to be built on
func Init() *C {
	c := &C{&context{}}
	c.setLogger()
	c.context.logRequest = false

	c.routes = &routes{}

	return c
}

// C is setter for context, D is geter for context
type C struct {
	*context
}

// LogRequest set true to enable logging, if in production you want to disable logging for performance, use this methodd to set it false
func (c *C) LogRequest(status bool) {
	c.context.logRequest = status
}

// NewDispatcher returns a new dispatcher with context set
func (c *C) NewDispatcher() *D {
	return &D{c: c.context}
}

// SetMgo sets the mongodb when address is passed. Tries to connect and exits if no connection was found.
func (c *C) SetMgo(addr string) *C {
	err := c.setMgo(addr)

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(0)
	}

	return c
}

func (c *C) setLogger() *C {
	c.setLog()

	return c
}

// Listen call listen with port number to start your server at this port number.
func (c *C) Listen(port int) {
	// Routes
	c.attachRoutes()

	// Mongo DB
	if c.mgo != nil {
		defer c.mgo.Close()
	}

	c.infoLog.Println("Application Started at", strconv.Itoa(port))

	http.ListenAndServe(":"+strconv.Itoa(port), c.router)
}

// Get basic Get route
func (c *C) Get(url string) *route {
	return &route{context: c, method: "GET", url: url}
}

// Post basic post route
func (c *C) Post(url string) *route {
	return &route{context: c, method: "POST", url: url}
}

// Put basic Put route
func (c *C) Put(url string) *route {
	return &route{context: c, method: "PUT", url: url}
}

// Del basic del route
func (c *C) Del(url string) *route {
	return &route{context: c, method: "DELETE", url: url}
}
