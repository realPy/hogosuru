package htmlformelement

import "github.com/realPy/hogosuru/event"

func (h HtmlFormElement) OnFormData(handler func(e event.Event)) error {

	return h.AddEventListener("formdata", handler)
}

func (h HtmlFormElement) OnReset(handler func(e event.Event)) error {

	return h.AddEventListener("reset", handler)
}

func (h HtmlFormElement) OnSubmit(handler func(e event.Event)) error {

	return h.AddEventListener("submit", handler)
}
