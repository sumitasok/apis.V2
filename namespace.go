package apis

type NS struct {
	c      *C
	prefix string
	aCtrl  []action
	bCtrl  []action
}

func (c *C) NameSpace(prefix string, beforeControllers ...action) *NS {
	ns := &NS{
		c:      c,
		prefix: prefix,
		bCtrl:  beforeControllers,
	}

	return ns
}

func (c *C) NS(prefix string, beforeControllers ...action) *NS {
	ns := &NS{
		c:      c,
		prefix: prefix,
		bCtrl:  beforeControllers,
	}

	return ns
}

func (n *NS) Serve(dispatchers func(*NS), afterControllers ...action) *NS {
	if len(n.aCtrl) > 0 { // inherited from previous namespace
		for i := range n.aCtrl {
			afterControllers = append(afterControllers, n.aCtrl[i])
		}
	}
	n.aCtrl = afterControllers

	dispatchers(n)

	return n
}

func (n *NS) Get(url string) *route {
	return &route{context: n.c, method: "GET", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}

func (n *NS) Post(url string) *route {
	return &route{context: n.c, method: "POST", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}
func (n *NS) Put(url string) *route {
	return &route{context: n.c, method: "PUT", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}
func (n *NS) Delete(url string) *route {
	return &route{context: n.c, method: "DELETE", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}

func (n NS) NameSpace(prefix string, beforeControllers ...action) *NS {
	if len(n.bCtrl) > 0 {
		for i := range beforeControllers {
			n.bCtrl = append(n.bCtrl, beforeControllers[i])
		}
	}

	ns := &NS{
		c:      n.c,
		prefix: n.prefix + prefix,
		bCtrl:  n.bCtrl,
		aCtrl:  n.aCtrl,
	}

	return ns

}
