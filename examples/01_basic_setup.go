package main

import (
	apis "github.com/sumitasok/apis.V2"
)

type DummyController struct {
}

func main() {
	c := apis.Init()

	c.Get("/urls").Set(DummyController{})

	c.Listen(7000)
}
