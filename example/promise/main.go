package main

import (
	"syscall/js"
	"time"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
)

func main() {
	hogosuru.Init()
	p1, _ := promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		println("Waiting p1")
		time.Sleep(8 * time.Second)
		println("End p1")
		return js.ValueOf("p1"), nil
	})

	p2, _ := promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		println("Waiting p2")
		time.Sleep(3 * time.Second)
		println("End p2")
		return js.ValueOf("p2"), nil
	})
	p2.Export("test")
	p3, _ := promise.Any(p1, p2)

	//data, _ := p3.Await()

	p3.Then(func(i interface{}) *promise.Promise {

		if b, ok := i.(baseobject.ObjectFrom); ok {
			println("First elem response", b.JSObject().String())
		}

		return nil
	}, func(e error) {

	})

	p3.Finally(func() {
		println("finished")
	})

	ch := make(chan struct{})
	<-ch

}
