package event

import "github.com/realPy/jswasm/js"

type JSEvent struct {
	Name string
}

type JSCustomEvent struct {
	Name string
}

func (j JSEvent) DispatchRootEvent() {
	var event = js.Global().Get("Event")
	ev := event.New(j.Name)
	doc := js.Global().Get("document")
	doc.Call("dispatchEvent", ev)
}

func (j JSCustomEvent) DispatchRootEvent(detail string) {

	var event = js.Global().Get("CustomEvent")
	ev := event.New(j.Name, js.ValueOf(map[string]interface{}{"detail": detail}))
	doc := js.Global().Get("document")
	doc.Call("dispatchEvent", ev)
}
