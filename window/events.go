package window

import (
	"github.com/realPy/hogosuru/event"
)

func (w Window) OnHashChange(handler func(e event.Event)) error {

	return w.AddEventListener("hashchange", handler)
}

func (w Window) OnPopState(handler func(e event.Event)) error {

	return w.AddEventListener("popstate", handler)
}