package serviceworker

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/event"
)

func (s ServiceWorker) OnStateChange(handler func(e event.Event)) (js.Func, error) {

	return s.AddEventListener("statechange", handler)
}

func (s ServiceWorker) OnError(handler func(e event.Event)) (js.Func, error) {

	return s.AddEventListener("error", handler)
}
