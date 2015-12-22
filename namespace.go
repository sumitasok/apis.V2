package apis

type NameSpace struct {
	c      *C
	prefix string
	aCtrl  []action
	bCtrl  []action
}

func (c *C) NameSpace(prefix string, beforeControllers ...action) *NameSpace {
	ns := &NameSpace{
		c:      c,
		prefix: prefix,
		bCtrl:  beforeControllers,
	}

	return ns
}

func (n *NameSpace) Serve(dispatchers func(*NameSpace), afterControllers ...action) *NameSpace {
	if len(n.aCtrl) > 0 { // inherited from previous namespace
		for i := range n.aCtrl {
			afterControllers = append(afterControllers, n.aCtrl[i])
		}
	}
	n.aCtrl = afterControllers

	dispatchers(n)

	return n
}

func (n *NameSpace) Get(url string) *route {
	return &route{context: n.c, method: "GET", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}

func (n *NameSpace) Post(url string) *route {
	return &route{context: n.c, method: "POST", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}
func (n *NameSpace) Put(url string) *route {
	return &route{context: n.c, method: "PUT", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}
func (n *NameSpace) Delete(url string) *route {
	return &route{context: n.c, method: "DELETE", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}

func (n NameSpace) NameSpace(prefix string, beforeControllers ...action) *NameSpace {
	if len(n.bCtrl) > 0 {
		for i := range beforeControllers {
			n.bCtrl = append(n.bCtrl, beforeControllers[i])
		}
	}

	ns := &NameSpace{
		c:      n.c,
		prefix: n.prefix + prefix,
		bCtrl:  n.bCtrl,
		aCtrl:  n.aCtrl,
	}

	return ns

}

type NS struct {
	*NameSpace
}
