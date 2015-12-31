package apis

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigInit(t *testing.T) {
	assert := assert.New(t)

	c := Init()
	c.ConfigBytes([]byte(`{"apis": {"dev": {"name": "Sumit", "dob": {"month": "nov"}}}}`))

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

	assert.True(true)
}
