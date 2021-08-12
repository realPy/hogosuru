package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
)

func main() {
	p, _ := hogosuru.LoadWasm("console.wasm")

	p.Async(func(bo baseobject.BaseObject) *promise.Promise {

		println("test wasm load successfull")
		return nil
	}, func(e error) {
		hogosuru.AssertErr(e)
	})

	ch := make(chan struct{})
	<-ch
}
