package main

import (
	apis "github.com/sumitasok/apis.V2"
)

type DummyController struct {
}

func (d DummyController) Config() *apis.Config {
	return &apis.Config{}
}

func (d DummyController) Call(dispatcher *apis.D) (interface{}, error, int) {
	return nil, nil, 200
}
