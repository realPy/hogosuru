package main

import (
	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/messageevent"
	"github.com/realPy/hogosuru/websocket"
)

func main() {

	//connect on echo websocket
	if ws, err := websocket.New("wss://echo.websocket.org"); err == nil {

		/*

			func(w websocket.WebSocket, message interface{}) {

				if a, ok := message.(arraybuffer.ArrayBuffer); ok {
					println("Ws receive arraybuffer:" + a.String())
				}
				if s, ok := message.(string); ok {
					println("Ws receivea string:" + s)
				}

			}*/

		ws.OnMessage(func(m messageevent.MessageEvent) {

			if o, err := m.Data(); err == nil {
				if a, ok := o.(arraybuffer.ArrayBuffer); ok {

					println("Ws receive arraybuffer:" + a.String())
				} else {

					if b, ok := o.(baseobject.BaseObject); ok {

						println("Ws receive data :" + b.String())
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
