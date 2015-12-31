package apis

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	ConfigPathSeperator = "."
)

func (c *C) ConfigHttpResp(resp *http.Response, err error) *C {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(ExitConfigError)
	}

	c.ConfigBytes(data)

	return c
}

func (c *C) ConfigFile(filepath string) *C {

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		os.Exit(ExitConfigError)
	}

	c.ConfigBytes(data)

	return c
}

func (c *C) ConfigBytes(data []byte) *C {
	j, err := simplejson.NewJson(data)
	if err != nil {
		os.Exit(ExitConfigError)
	}

	c.appConfig = &Config{data: j}

	return c
}

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

type Config struct {
	data *simplejson.Json
}

func (c *Config) Array(path string, fallback []interface{}) []interface{} {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.Array()
	if err != nil {
		return fallback
	}

	return val
}

func (c *Config) Bool(path string, fallback bool) bool {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.Bool()
	if err != nil {
		return fallback
	}

	return val
}

func (c *Config) Bytes(path string, fallback []byte) []byte {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.Bytes()
	if err != nil {
		return fallback
	}

	return val
}

func (c *Config) Float(path string, fallback float64) float64 {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.Float64()
	if err != nil {
		return fallback
	}

	return val
}

func (c *Config) String(path string, fallback string) string {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.String()
	if err != nil {
		return fallback
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
	} else {
		return nil, false
	}
}

func (c *Config) Int(path string, fallback int) int {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.Int()
	if err != nil {
		return fallback
	}

	return val
}

func (c *Config) Int64(path string, fallback int64) int64 {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.Int64()
	if err != nil {
		return fallback
	}

	return val
}

func (c *Config) Interface(path string, fallback interface{}) interface{} {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val := node.Interface()
	return val
}

func (c *Config) Map(path string, fallback map[string]interface{}) map[string]interface{} {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.Map()
	if err != nil {
		return fallback
	}

	return val
}

func (c *Config) StringArray(path string, fallback []string) []string {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.StringArray()
	if err != nil {
		return fallback
	}

	return val
}

func (c *Config) Uint64(path string, fallback uint64) uint64 {
	node, ok := lastConfigNode(path, *c.data)
	if !ok {
		return fallback
	}

	val, err := node.Uint64()
	if err != nil {
		return fallback
	}

	return val
}
