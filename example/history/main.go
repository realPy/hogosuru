package main

import (
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/history"
)

func main() {
	d := document.New_()

	//hist, _ := history.New()
	historyObj, _ := history.GetHistory()
	//println("Back to previous page")

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

	// if clickbutton, err := d.GetElementById("forward"); err == nil {

	// 	clickbutton.OnClick(func(e event.Event) {
	// 		println("Back to previous page")
	// 		if err := historyObj.Forward(); err != nil {
	// 			println("Error %s" + err.Error())
	// 		}
	// 	})
	// } else {
	// 	println("error")
	// }

	// if clickbutton, err := d.GetElementById("go"); err == nil {

	// 	clickbutton.OnClick(func(e event.Event) {
	// 		println("Back to previous page")
	// 		if err := historyObj.Go(-3); err != nil {
	// 			println("Error %s" + err.Error())
	// 		}
	// 	})
	// } else {
	// 	println("error")
	// }

	// if obj := history.HistoryBlob.State(); err != nil {
	// 	println("Error", err)
	// } else {
	// 	println(obj.String())
	// }
	//println(history.State())

	ch := make(chan struct{})
	<-ch
}
