package apis

type action interface {
	Call(*D) (interface{}, error, int) // json data, error, http status code
	Config() *Config
}
