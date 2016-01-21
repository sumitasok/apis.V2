package apis

import (
	"fmt"
	"os"

	"net/http"
	"strconv"
)

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

func (c *C) LogRequest(status bool) {
	c.context.logRequest = status
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

func (c *C) setLogger() *C {
	c.setLog()

	return c
}

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

func (c *C) Get(url string) *route {
	return &route{context: c, method: "GET", url: url}
}

func (c *C) Post(url string) *route {
	return &route{context: c, method: "POST", url: url}
}

func (c *C) Put(url string) *route {
	return &route{context: c, method: "PUT", url: url}
}

func (c *C) Del(url string) *route {
	return &route{context: c, method: "DELETE", url: url}
}
