package main

import (
	"fmt"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/dedicatedworkerglobalscope"
	"github.com/realPy/hogosuru/base/messageevent"
)

func main() {
	hogosuru.Init()
	fmt.Printf("Get self\n")
	if instance, err := baseobject.Self(); hogosuru.AssertErr(err) {

		if d, ok := instance.(dedicatedworkerglobalscope.DedicatedWorkerGlobalScope); ok {
			fmt.Printf("Install handler\n")
			d.PostMessage("installok")
			d.OnMessage(func(m messageevent.MessageEvent) {
				d.PostMessage("testok")
			})
		}

	}

	ch := make(chan struct{})
	<-ch

}
