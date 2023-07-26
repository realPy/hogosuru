package worker

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/messageevent"
)

func (w Worker) OnMessage(handler func(m messageevent.MessageEvent)) (js.Func, error) {

	return w.AddEventListener("message", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

func (w Worker) OnMessageError(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("onmessageerror", handler)
}
