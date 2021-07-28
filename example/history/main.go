package main

import (
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/history"
)

func main() {
	d := document.New_()

	historyObj, _ := history.GetHistory()

	if clickbutton, err := d.GetElementById("back"); err == nil {

		clickbutton.OnClick(func(e event.Event) {
			println("Back to previous page")
			if err := historyObj.Back(); err != nil {
				println("Error %s" + err.Error())
			}
		})
	} else {
		println("error")
	}

	if clickbutton2, err := d.GetElementById("forward"); err == nil {

		clickbutton2.OnClick(func(e event.Event) {
			println("Forward to next page")
			if err := historyObj.Forward(); err != nil {
				println("Error %s" + err.Error())
			}
		})
	} else {
		println("error")
	}

	if clickbutton, err := d.GetElementById("go"); err == nil {

		clickbutton.OnClick(func(e event.Event) {
			println("Back to previous page")
			if err := historyObj.Go(-3); err != nil {
				println("Error %s" + err.Error())
			}
		})
	} else {
		println("error")
	}

	obj, _ := historyObj.State()

	if clickbutton, err := d.GetElementById("replace"); err == nil {

		clickbutton.OnClick(func(e event.Event) {
			println("Replace")
			if err := historyObj.ReplaceState(obj, "Test", "/node.html"); err != nil {
				println("Error %s" + err.Error())
			}
			println(obj)
		})
	} else {
		println("error")
	}

	if clickbutton, err := d.GetElementById("push"); err == nil {

		clickbutton.OnClick(func(e event.Event) {
			println("Push")
			if err := historyObj.PushState(obj, "Test", "/node.html"); err != nil {
				println("Error %s" + err.Error())
			}
			println(obj)
		})
	} else {
		println("error")
	}

	obj, err := historyObj.Length()
	println(obj, err)

	obj2, err := historyObj.State()
	println(obj2, err)

	ch := make(chan struct{})
	<-ch
}
