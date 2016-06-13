package apis

import (
	"fmt"
	assert "github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func ExampleConfig() {
	c := Init()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:7007/config", nil)

	c.ConfigHttpResp(client.Do(req)).Root("data", "sales", "dev")

	fmt.Println(c.Config().String("payment_url", ""))
	fmt.Println(c.Config().String("db_tables.order_db", ""))
	// Output
}

func TestConfigInit(t *testing.T) {
	assert := assert.New(t)

	c := Init()
	c.ConfigBytes([]byte(`{"apis": {"dev": {"name": "Sumit", "dob": {"month": "nov", "date": 26}}}}`))
	c.Config().AllowFallback(true)

	_, ok := c.context.appConfig.data.CheckGet("apis")

	assert.True(ok)

	_, ok = c.context.appConfig.data.CheckGet("name")

	assert.False(ok)

	c.Root("apis", "dev")

	_, ok = c.context.appConfig.data.CheckGet("name")

	assert.True(ok)

	val := c.context.appConfig.String("name", "apis")

	assert.Equal("Sumit", val)
	assert.NotEqual("apis", val)

	val = c.context.appConfig.String("dob.month", "apis")

	assert.Equal("nov", val)
	assert.NotEqual("apis", val)

	val = c.context.appConfig.String("dob.year", "1987")

	assert.Equal("1987", val)

	valDate := c.context.appConfig.Int64("dob.date", 25)

	assert.EqualValues(26, valDate)

	assert.True(true)
}

func TestConfigFile(t *testing.T) {
	assert := assert.New(t)

	c := Init()

	c.ConfigFile("./samples/config.json")
	c.Config().AllowFallback(true)

	_, ok := c.context.appConfig.data.CheckGet("apis")

	assert.True(ok)

	_, ok = c.context.appConfig.data.CheckGet("name")

	assert.False(ok)

	c.Root("apis", "dev")

	_, ok = c.context.appConfig.data.CheckGet("name")

	assert.True(ok)

	val := c.context.appConfig.String("name", "apis")

	assert.Equal("Sumit", val)
	assert.NotEqual("apis", val)

	val = c.context.appConfig.String("dob.month", "apis")

	assert.Equal("nov", val)
	assert.NotEqual("apis", val)

	val = c.context.appConfig.String("dob.year", "1987")

	assert.Equal("1987", val)

	valDate := c.context.appConfig.Int64("dob.date", 25)

	assert.EqualValues(26, valDate)

	assert.True(true)
}
