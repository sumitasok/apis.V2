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

type data struct {
	value int
}
type pointer struct {
	data *data
	nonP int
}

func (p pointer) call() int {
	return p.data.value
}

func (p pointer) alt() {
	p.nonP = 2
	p.data.value = 3
}

func TestPointerData(t *testing.T) {
	assert := assert.New(t)

	p := pointer{data: &data{value: 1}}
	q := &p
	r := p

	assert.Equal(1, p.call())
	assert.Equal(0, p.nonP)

	p.data.value = 2

	assert.Equal(2, p.call())
	assert.Equal(0, p.nonP)

	assert.Equal(2, q.call())

	p.alt()
	assert.Equal(3, p.call())
	assert.Equal(0, p.nonP)
	assert.Equal(3, q.call())
	assert.Equal(3, r.call())

	assert.True(true)
}
