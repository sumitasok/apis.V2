package main

import (
	"encoding/gob"
	"github.com/sumitasok/apis.V2"
)

type Sales struct {
	PaymentURL string `json:"payment_url"`

	DBTables struct {
		OrderDB string `json:"order_db"`
	} `json:"db_tables"`
}

func main() {
	c := apis.Init()

	c.Get("/config").Set(Config{})

	c.Listen(7007)
}

type Config struct{}

func (c Config) Config() *apis.Config {
	return &apis.Config{}
}

func (c Config) Call(d *apis.D) (interface{}, error, int) {
	s := Sales{
		PaymentURL: "http://localhost:7003/status/update",
	}

	gob.Register(Sales{})

	appConfig := map[string]map[string]interface{}{
		"sales": map[string]interface{}{
			"dev": s,
		},
	}

	appConfig["sales"]["prod"] = s

	return appConfig, nil, 200
}
