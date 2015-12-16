package apis

import (
	"fmt"
	assert "github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestContext(t *testing.T) {
	assert := assert.New(t)

	assert.True(true)
}

func BenchmarkContext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Init()
	}
}

func BenchmarkContextNewDispatcher(b *testing.B) {
	c := Init()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.NewDispatcher()
	}
}

type DummyAction struct{}

func TestRoute(t *testing.T) {
	assert := assert.New(t)

	c := Init()

	c.Get("/url").Set(DummyAction{})

	assert.True(true)
}

func BenchmarkGetRoute(b *testing.B) {
	c := Init()
	urls := []string{}
	for i := 0; i < b.N; i++ {
		urls = append(urls, fmt.Sprintf("/url%s", strconv.Itoa(i)))
	}

	b.ResetTimer()

	for _, url := range urls {
		c.Get(url).Set(DummyAction{})
	}
}
