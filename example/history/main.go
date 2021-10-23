package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/window"
)

func main() {

	hogosuru.Init()

	d := document.New_()

	if w, err := window.New(); hogosuru.AssertErr(err) {

		if historyObj, err := w.History(); err == nil {
			if clickbutton, err := d.GetElementById("back"); hogosuru.AssertErr(err) {

				clickbutton.OnClick(func(e event.Event) {
					println("Back to previous page")
					if err := historyObj.Back(); err != nil {
						println("Error %s" + err.Error())
					}
				})
			}

			if clickbutton2, err := d.GetElementById("forward"); hogosuru.AssertErr(err) {

				clickbutton2.OnClick(func(e event.Event) {
					println("Forward to next page")
					if err := historyObj.Forward(); err != nil {
						println("Error %s" + err.Error())
					}
				})
			}

			if clickbutton, err := d.GetElementById("go"); hogosuru.AssertErr(err) {

				clickbutton.OnClick(func(e event.Event) {
					println("Back to previous page")
					if err := historyObj.Go(-3); err != nil {
						println("Error %s" + err.Error())
					}
				})
			}

			obj, _ := historyObj.State()

			if clickbutton, err := d.GetElementById("replace"); hogosuru.AssertErr(err) {

				clickbutton.OnClick(func(e event.Event) {
					println("Replace")
					if err := historyObj.ReplaceState(obj, "Test", "/node.html"); err != nil {
						println("Error %s" + err.Error())
					}
					println(obj)
				})
			}

			if clickbutton, err := d.GetElementById("push"); hogosuru.AssertErr(err) {

				clickbutton.OnClick(func(e event.Event) {
					println("Push")
					if err := historyObj.PushState(obj, "Test", "/node.html"); err != nil {
						println("Error %s" + err.Error())
					}
					println(obj)
				})
			}

			obj, err := historyObj.Length()
			println(obj, err)

			obj2, err := historyObj.State()
			println(obj2, err)

		}
	}

	ch := make(chan struct{})
	<-ch
}
