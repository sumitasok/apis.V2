package middlewares

import (
	apis "github.com/sumitasok/apis.V2"
)

type CORS struct {
}

func (c CORS) Config() *apis.Config {
	return &apis.Config{}
}

func (c CORS) Call(d *apis.D) (interface{}, error, int) {
	req := d.Request()
	if origin := req.Header.Get("Origin"); origin != "" {
		d.ResponseWriter().Header().Set("Access-Control-Allow-Origin", origin)
		d.ResponseWriter().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		d.ResponseWriter().Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if req.Method == "OPTIONS" {
		return nil, nil, 200
	}

	return nil, nil, 200
}
