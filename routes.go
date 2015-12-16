package apis

type route struct {
	context *C
	method  string
	url     string
	actions []action
}

func (r *route) Set(actions ...action) {
	r.actions = actions

	r.context.addRoute(r)

	return
}
