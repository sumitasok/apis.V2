package apis

// NS NameSpace which holds common part of urls for avoiding redundancy
// as well as allows chaining methods which will act before actual controllers
// a CORS controller can be set on an NS
type NS struct {
	c      *C
	prefix string
	aCtrl  []action
	bCtrl  []action
}

// NameSpace which holds common part of urls for avoiding redundancy
// as well as allows chaining methods which will act before actual controllers
// a CORS controller can be set on an NS
func (c *C) NameSpace(prefix string, beforeControllers ...action) *NS {
	ns := &NS{
		c:      c,
		prefix: prefix,
		bCtrl:  beforeControllers,
	}

	return ns
}

// NS Namespace which holds common part of urls for avoiding redundancy
// as well as allows chaining methods which will act before actual controllers
// before controllers, are controllers set in NS which will act on the route before controllers in Serve
// a CORS controller can be set on an NS
func (c *C) NS(prefix string, beforeControllers ...action) *NS {
	ns := &NS{
		c:      c,
		prefix: prefix,
		bCtrl:  beforeControllers,
	}

	return ns
}

// Serve holds the main controllers chain, this can have after controllers
// after controllers are controllers which are set in Serve methods which will act after the main controllers are run
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

// Get basic Get route
func (n *NS) Get(url string) *route {
	return &route{context: n.c, method: "GET", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}

// Post basic Post route
func (n *NS) Post(url string) *route {
	return &route{context: n.c, method: "POST", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}

// Put basic Put route
func (n *NS) Put(url string) *route {
	return &route{context: n.c, method: "PUT", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}

// Del basic Del route
func (n *NS) Del(url string) *route {
	return &route{context: n.c, method: "DELETE", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}

// Options basic Options route
func (n *NS) Options(url string) *route {
	return &route{context: n.c, method: "OPTIONS", url: n.prefix + url, bCtrl: n.bCtrl, aCtrl: n.aCtrl}
}

// NS NameSpace called on another NameSpace
func (n NS) NS(prefix string, beforeControllers ...action) *NS {
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
