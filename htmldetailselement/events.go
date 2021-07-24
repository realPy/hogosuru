package htmldetailselement

import "github.com/realPy/hogosuru/event"

func (h HtmlDetailsElement) OnToggle(handler func(e event.Event)) error {

	return h.AddEventListener("toggle", handler)
}
