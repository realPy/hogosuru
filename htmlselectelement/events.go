package htmlselectelement

import "github.com/realPy/hogosuru/event"

func (h HtmlSelectElement) OnInput(handler func(e event.Event)) error {

	return h.AddEventListener("input", handler)
}
