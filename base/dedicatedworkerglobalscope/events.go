package dedicatedworkerglobalscope

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/messageevent"
)

func (d DedicatedWorkerGlobalScope) OnMessage(handler func(m messageevent.MessageEvent)) (js.Func, error) {

	return d.AddEventListener("message", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

func (d DedicatedWorkerGlobalScope) OnMessageError(handler func(e event.Event)) (js.Func, error) {

	return d.AddEventListener("onmessageerror", handler)
}
