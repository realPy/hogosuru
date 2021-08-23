package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
)

func main() {
	f, p, _ := hogosuru.LoadWasm("console.wasm")

	f.Async(func(bo baseobject.BaseObject) *promise.Promise {
		println("wasm download success")
		return nil
	}, func(e error) {
		hogosuru.AssertErr(e)
	})

	p.Async(func(bo baseobject.BaseObject) *promise.Promise {

		println("test wasm load successfull")
		return nil
	}, func(e error) {
		hogosuru.AssertErr(e)
	})

	ch := make(chan struct{})
	<-ch
}
