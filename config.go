package apis

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	// ConfigPathSeperator '.' if config path is 'data.db.addr' is dot seperated
	ConfigPathSeperator = "."
)

// Config returns the config of the application
func (c *C) Config() *Config {
	return c.context.appConfig
}

// ConfigHttpResp gets the config from an http server which returns a json formated config data.
func (c *C) ConfigHttpResp(resp *http.Response, err error) *C {
	if err != nil {
		c.infoLog.Fatal(err)
		return c
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(ExitConfigError)
	}

	c.ConfigBytes(data)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	return c
}

// ConfigFile reads the config from a json file
func (c *C) ConfigFile(filepath string) *C {

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		os.Exit(ExitConfigError)
	}

	c.ConfigBytes(data)

	return c
}

// ConfigBytes returns the config from Int data
func (c *C) ConfigBytes(data []byte) *C {
	j, err := simplejson.NewJson(data)
	if err != nil {
		os.Exit(ExitConfigError)
	}

	c.appConfig = &Config{data: j}

	return c
}

// Root brings the node and sets the root node to the level specified by parameters eg: ("data", "db")
func (c *C) Root(nodes ...string) *C {
	for i := range nodes {
		if j, ok := c.appConfig.data.CheckGet(nodes[i]); ok {
			c.appConfig.data = j
		} else {
			os.Exit(ExitConfigError)
		}
	}

	return c
}

// Config holds the details about the config
type Config struct {
	allowFallback bool
	data          *simplejson.Json
}

// AllowFallback is set to true takes the fallback value passed as second paramter if the input config doesnot have the value
// if this is set, you might forget adding config to your config file or server.
func (c *Config) AllowFallback(status bool) {
	c.allowFallback = status
}

// Array returns the value if it is an array
func (c *Config) Array(path string, fallback []interface{}) []interface{} {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.Array()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}

// Bool returns the value if it is a Bool
func (c *Config) Bool(path string, fallback bool) bool {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.Bool()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}

// Bytes returns the value if it is a Bytes
func (c *Config) Bytes(path string, fallback []byte) []byte {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.Bytes()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}

// Float returns the value if it is a Float
func (c *Config) Float(path string, fallback float64) float64 {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.Float64()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}

// String returns the value if it is a String
func (c *Config) String(path string, fallback string) string {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.String()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}

func lastConfigNode(path string, node simplejson.Json) (*simplejson.Json, bool) {
	keys := strings.Split(path, ConfigPathSeperator)

	_node, ok := nextConfigNode(keys, &node)

	return _node, ok
}

func nextConfigNode(keys []string, node *simplejson.Json) (*simplejson.Json, bool) {
	if len(keys) == 0 {
		return nil, false
	}

	if len(keys) == 1 {
		return node.CheckGet(keys[0])
	}

	if node, ok := node.CheckGet(keys[0]); ok {
		return nextConfigNode(keys[1:], node)
	}

	return nil, false
}

// Int returns the value if it is a Int
func (c *Config) Int(path string, fallback int) int {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.Int()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}

// Int64 returns the value if it is a Int64
func (c *Config) Int64(path string, fallback int64) int64 {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.Int64()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}

// Interface returns the value if it is a Interface
func (c *Config) Interface(path string, fallback interface{}) interface{} {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val := node.Interface()
	return val
}

// Map returns the value if it is a Map
func (c *Config) Map(path string, fallback map[string]interface{}) map[string]interface{} {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.Map()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}

// StringArray returns the value if it is a StringArray
func (c *Config) StringArray(path string, fallback []string) []string {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.StringArray()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}

// Uint64 returns the value if it is a Uint64
func (c *Config) Uint64(path string, fallback uint64) uint64 {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	val, err := node.Uint64()
	if err != nil {
		if c.allowFallback {
			return fallback
		}
		panic("config " + path + " not found")
	}

	return val
}
