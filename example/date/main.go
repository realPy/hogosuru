package main

import (
	"syscall/js"
	"time"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/date"
	"github.com/realPy/hogosuru/promise"
)

func main() {
	hogosuru.Init()
	if d, err := date.New(); hogosuru.AssertErr(err) {
		value, _ := d.GetMilliseconds()
		println("-->", value)

	}

	value, _ := date.UTC(2012, 11, 20, 3, 0, 0)

	if d2, err := date.New(value); hogosuru.AssertErr(err) {
		d2.Export("oto")
		if t, _ := d2.ValueOf(); t == value {
			println("Ok")
		}
	}

	println("------>", value)
	d1, _ := date.New()
	ret, _ := d1.ToLocaleString("en-GB", map[string]interface{}{"timeZone": "UTC"})
	println(ret)

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

	p3, _ := promise.Any(p1, p2)

	data, _ := p3.Await()

	if obj, ok := data.(baseobject.ObjectFrom); ok {
		println("First elem response", obj.BaseObject_().String())
	}

	ch := make(chan struct{})
	<-ch

}
