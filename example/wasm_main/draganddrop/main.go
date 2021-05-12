package main

import (
	"github.com/realPy/hogosuru/dragevent"
	"github.com/realPy/hogosuru/js"
)

func dropHandler() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("drop here")
		if e, err := dragevent.NewFromJSObject(args[0]); err == nil {
			e.PreventDefault()
		}

		return nil
	})
}

func dragOverHandler() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if e, err := dragevent.NewFromJSObject(args[0]); err == nil {
			e.PreventDefault()
		}
		return nil
	})
}

func main() {

	js.Global().Set("dropHandler", dropHandler())

	js.Global().Set("dragOverHandler", dragOverHandler())

	println("loaded")
	ch := make(chan struct{})
	<-ch

}
