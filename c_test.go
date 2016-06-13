package apis

import ()

type DummyController struct {
}

func (d DummyController) Config() *Config {
	return &Config{}
}

func (d DummyController) Call(dispatcher *D) (interface{}, error, int) {
	return nil, nil, 200
}

func PendingExampleInit() {
	c := Init()

	c.Get("/urls").Set(DummyController{})

	c.Listen(7000)
	// Output:
	// INFO: 2016/05/22 18:07:38 Application Started at 7000
}
