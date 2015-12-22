package apis

type route struct {
	context *C
	method  string
	url     string
	actions []action
	bCtrl   []action
	aCtrl   []action
}

func (r *route) Set(actions ...action) *route {
	_actions := []action{}
	if len(r.bCtrl) > 0 {
		_actions = r.bCtrl

	}

	if len(_actions) == 0 {
		_actions = actions
	} else {
		for i := range actions {
			_actions = append(_actions, actions[i])
		}
	}

	if len(r.aCtrl) > 0 {
		for i := range r.aCtrl {
			_actions = append(_actions, r.aCtrl[i])
		}
	}

	r.actions = _actions

	r.context.addRoute(r)

	return r
}
