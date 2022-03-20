package serviceworkercontainer

import (
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

func (s ServiceWorkerContainer) OnControllerChange(handler func(e event.Event)) (js.Func, error) {

	return s.AddEventListener("controllerchange", handler)
}

func (s ServiceWorkerContainer) OnMessage(handler func(e event.Event)) (js.Func, error) {

	return s.AddEventListener("message", handler)
}

func (s ServiceWorkerContainer) OnError(handler func(e event.Event)) (js.Func, error) {

	return s.AddEventListener("error", handler)
}
