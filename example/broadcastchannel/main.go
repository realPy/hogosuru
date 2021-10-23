package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/broadcastchannel"
	"github.com/realPy/hogosuru/messageevent"
)

func main() {
	hogosuru.Init()
	if c, err := broadcastchannel.New("channel1"); hogosuru.AssertErr(err) {

		c.OnMessage(func(m messageevent.MessageEvent) {

			if b, err := m.Data(); hogosuru.AssertErr(err) {
				println("receive message " + b.(baseobject.ObjectFrom).BaseObject_().String())

			}

		})

		c.PostMessage("New tab opened")
	}

	ch := make(chan struct{})
	<-ch

}
