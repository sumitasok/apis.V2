package apis

import (
	"fmt"
	"os"
)

func Init() *C {
	c := &C{&context{}}
	c.SetLogger()

	c.routes = &routes{}

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

type action interface{}

func (c *C) Get(url string) *route {
	return &route{context: c, method: "GET", url: url}
}

func (c *C) Post(url string) *route {
	return &route{context: c, method: "POST", url: url}
}

func (c *C) Put(url string) *route {
	return &route{context: c, method: "PUT", url: url}
}

func (c *C) Delete(url string) *route {
	return &route{context: c, method: "DELETE", url: url}
}
