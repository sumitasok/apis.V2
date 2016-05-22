package apis

import (
	"fmt"
)

func ExampleInit() {
	c := Init()

	c.Get("/urls").Set(DummyController{})

	c.Listen(7000)
	// Output:
	// INFO: 2016/05/22 18:07:38 Application Started at 7000
}
