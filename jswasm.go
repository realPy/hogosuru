package jswasm

import "github.com/realPy/jswasm/js"

func promise() (p js.Value, set func(js.Value)) {

	ch := make(chan js.Value)
	resolver := make(chan js.Value, 1)
	go func() {
		result := <-ch
		resolve := <-resolver
		resolve.Invoke(result)
	}()
	promise := js.Global().Get("Promise")
	p = promise.New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolver <- args[0]
		return nil
	}))
	set = func(v js.Value) {
		ch <- v
	}
	return
}

func await(awaitable js.Value) chan []js.Value {
	ch := make(chan []js.Value)
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- args
		return nil
	})
	awaitable.Call("then", cb)
	return ch
}
