package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/messageevent"
	"github.com/realPy/hogosuru/websocket"
)

func main() {
	hogosuru.Init()
	//connect on echo websocket
	if ws, err := websocket.New("ws://localhost:9090/echo"); hogosuru.AssertErr(err) {

		ws.OnMessage(func(m messageevent.MessageEvent) {

			if o, err := m.Data(); hogosuru.AssertErr(err) {
				if a, ok := o.(arraybuffer.ArrayBuffer); ok {

					println("Ws receive arraybuffer:" + a.String())
				} else {

					if b, ok := o.(baseobject.ObjectFrom); ok {

						println("Ws receive data :" + b.BaseObject_().String())
					}
				}
			}

		})

		ws.OnOpen(func(e event.Event) {
			ws.Send("hello")
		})

	}

	ch := make(chan struct{})
	<-ch

}
