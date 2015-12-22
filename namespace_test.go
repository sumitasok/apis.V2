package apis

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestNamespace(t *testing.T) {
	assert := assert.New(t)

	assert.True(true)
}

func BenchmarkNamespace(b *testing.B) {
	c := Init()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.NameSpace("/sales", DummyAction{}).Serve(func(ns *NS) {
			ns.Get("").Set(DummyAction{})
		}, DummyAction{})
	}
}

func TestNameSpaceRoute(t *testing.T) {
	assert := assert.New(t)

	c := Init()
	r := &route{}
	c.NameSpace("/sales", DummyAction{}).Serve(func(ns *NS) {
		r = ns.Get("").Set(DummyAction{})
	}, DummyAction{})

	assert.Len(r.actions, 3)

	assert.True(true)
}

func TestNestedNameSpaceRoute(t *testing.T) {
	assert := assert.New(t)

	c := Init()
	r := &route{}
	c.NameSpace("/sales", DummyAction{Index: 1}).Serve(func(ns *NS) {
		ns.Get("").Set(DummyAction{Index: 2})
		ns.NameSpace("/reports", DummyAction{Index: 2}).Serve(func(ns1 *NS) {
			ns1.Put("").Set(DummyAction{Index: 3})
			ns1.NameSpace("/reports", DummyAction{Index: 3}).Serve(func(ns2 *NS) {
				r = ns2.Post("").Set(DummyAction{Index: 4})
				ns2.Delete("").Set(DummyAction{Index: 4})
			}, DummyAction{Index: 5})
		}, DummyAction{Index: 6})
	}, DummyAction{Index: 7})

	assert.Len(r.actions, 7)
	assert.Equal(r.actions, []action{
		DummyAction{Index: 1}, DummyAction{Index: 2}, DummyAction{Index: 3}, DummyAction{Index: 4}, DummyAction{Index: 5}, DummyAction{Index: 6}, DummyAction{Index: 7},
	})

	assert.Len(*c.routes, 4)

	assert.True(true)
}

func BenchmarkNestedNameSpace(b *testing.B) {

	c := Init()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.NameSpace("/sales", DummyAction{Index: 1}).Serve(func(ns *NS) {
			ns.NameSpace("/reports", DummyAction{Index: 2}).Serve(func(ns1 *NS) {
				ns1.NameSpace("/events", DummyAction{Index: 3}).Serve(func(ns2 *NS) {
					ns2.Get("").Set(DummyAction{Index: 4})
				}, DummyAction{Index: 5})
			}, DummyAction{Index: 6})
		}, DummyAction{Index: 7})
	}
}
