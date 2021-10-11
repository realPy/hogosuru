package htmldetailselement

import (
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

func (h HtmlDetailsElement) OnToggle(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("toggle", handler)
}
