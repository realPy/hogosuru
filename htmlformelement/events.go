package htmlformelement

import (
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

func (h HtmlFormElement) OnFormData(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("formdata", handler)
}

func (h HtmlFormElement) OnReset(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("reset", handler)
}

func (h HtmlFormElement) OnSubmit(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("submit", handler)
}
