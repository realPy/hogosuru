package htmltextareaelement

import "github.com/realPy/hogosuru/event"

func (h HtmlTextAreaElement) OnInput(handler func(e event.Event)) error {

	return h.AddEventListener("input", handler)
}
