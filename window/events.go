package window

import "github.com/realPy/hogosuru/event"

func (w Window) OnHashChange(handler func(e event.Event)) error {

	return w.AddEventListener("hashchange", handler)
}
