package apis

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestDispatcherLoggerShouldBeInitiated(t *testing.T) {
	assert := assert.New(t)

	assert.NotPanics(func() { Init().NewDispatcher().LogInfo("Info") })

	assert.True(true)
}
