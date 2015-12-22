package apis

type NameSpace struct {
	c                 *C
	prefix            string
	AfterControllers  []action
	BeforeControllers []action
}

func (c *C) NameSpace(prefix string, beforeControllers ...action) *NameSpace {
	ns := &NameSpace{
		c:                 c,
		prefix:            prefix,
		BeforeControllers: beforeControllers,
	}

	return ns
}

func (n *NameSpace) Serve(dispatchers func(*NameSpace), afterControllers ...action) *NameSpace {
	if len(n.AfterControllers) > 0 { // inherited from previous namespace
		for i := range n.AfterControllers {
			afterControllers = append(afterControllers, n.AfterControllers[i])
		}
	}
	n.AfterControllers = afterControllers

	dispatchers(n)

	return n
}

func (n *NameSpace) Get(url string) *route {
	return &route{context: n.c, method: "GET", url: n.prefix + url, bCtrl: n.BeforeControllers, aCtrl: n.AfterControllers}
}

func (n NameSpace) NameSpace(prefix string, beforeControllers ...action) *NameSpace {
	if len(n.BeforeControllers) > 0 {
		for i := range beforeControllers {
			n.BeforeControllers = append(n.BeforeControllers, beforeControllers[i])
		}
	}

	ns := &NameSpace{
		c:                 n.c,
		prefix:            n.prefix + prefix,
		BeforeControllers: n.BeforeControllers,
		AfterControllers:  n.AfterControllers,
	}

	return ns

}
