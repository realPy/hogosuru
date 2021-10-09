package htmlselectelement

import (
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

func (h HtmlSelectElement) OnInput(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("input", handler)
}
