package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/hogosurudebug"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/response"
)

func main() {
	hogosuru.Init()
	hogosurudebug.EnableDebug()

	f, p, _ := hogosuru.LoadWasm("console.wasm")

	f.Then(func(r response.Response) *promise.Promise {

		println("wasm download success")
		return nil
	}, func(e error) {
		hogosuru.AssertErr(e)
	})

	p.Then(func(o interface{}) *promise.Promise {

		println("test wasm load successfull")
		return nil
	}, func(e error) {
		hogosuru.AssertErr(e)
	})

	ch := make(chan struct{})
	<-ch
}
