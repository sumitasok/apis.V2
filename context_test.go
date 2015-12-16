package apis

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
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

func BenchmarkSetRoute(b *testing.B) {
	c := Init()
	url := "/url"

	r := c.Get(url)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Set(DummyAction{})
	}
}

func BenchmarkHttpRouter(b *testing.B) {
	router := httprouter.New()
	dispatcher := Init().NewDispatcher()
	urls := []string{}
	for i := 0; i < b.N; i++ {
		urls = append(urls, fmt.Sprintf("/url%s", strconv.Itoa(i)))
	}

	b.ResetTimer()

	for _, url := range urls {
		router.GET(url, dispatcher.Call)
	}
}

func BenchmarkAddRoute(b *testing.B) {
	c := Init()
	r := route{context: c, method: "GET", url: "/url", actions: []action{DummyAction{}}}

	for i := 0; i < b.N; i++ {
		c.addRoute(&r)
	}
}

func TestGetSetAddRoute(t *testing.T) {
	assert := assert.New(t)

	c := Init()

	c.Get("/url").Set(DummyAction{})

	assert.Len(*c.routes, 1)

	c.Get("/static").Set(DummyAction{})

	assert.Len(*c.routes, 2)

	assert.True(true)
}

func BenchmarkAttachRouter(b *testing.B) {
	c := Init()
	for i := 0; i < b.N; i++ {
		r := &route{context: c, method: "GET", url: fmt.Sprintf("/url%s", strconv.Itoa(i)), actions: []action{DummyAction{}}}
		c.addRoute(r)
	}

	for i := 0; i < 1; i++ {
		c.attachRoutes()
	}
}
