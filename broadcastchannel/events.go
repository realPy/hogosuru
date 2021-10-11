package broadcastchannel

import (
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/messageevent"
)

func (c BroadcastChannel) OnMessage(handler func(m messageevent.MessageEvent)) (js.Func, error) {

	return c.AddEventListener("message", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

func (c BroadcastChannel) OnMessageError(handler func(e event.Event)) (js.Func, error) {

	return c.AddEventListener("messageerror", handler)
}
