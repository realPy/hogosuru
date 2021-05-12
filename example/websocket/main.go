package main

import (
	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/websocket"
)

func main() {

	//connect on echo websocket
	if ws, err := websocket.New("wss://echo.websocket.org"); err == nil {
		ws.SetOnMessage(func(w websocket.WebSocket, message interface{}) {

			if a, ok := message.(arraybuffer.ArrayBuffer); ok {
				println("Ws receive arraybuffer:" + a.String())
			}
			if s, ok := message.(string); ok {
				println("Ws receivea string:" + s)
			}

		})

		ws.SetOnOpen(func(ws websocket.WebSocket) {

			ws.Send("hello")
		})

	}

	ch := make(chan struct{})
	<-ch

}
