package main

import (
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
)

func main() {

	d := document.New()

	nod := d.Body()

	if text := nod.TextContent(); nod.Error == nil {
		println("<--" + text + "-->")
	}

	if elem, err := d.CreateElement("b"); err == nil {

		if t, err := d.CreateTextNode("Hello"); err == nil {

			elem.AppendChild(t)
			elem.Export("manu")
		} else {
			println(err.Error())
		}

		nod.AppendChild(elem.Node)
	} else {
		println(err.Error())
	}

	if elem, err := d.CreateElement("p"); err == nil {

		elem.SetInnerHTML("<b>World</b>")
		nod.AppendChild(elem.Node)
	} else {
		println(err.Error())
	}

	nodelist, _ := d.QuerySelectorAll(".pictureContainer")
	println("Found", nodelist.Length(), "elements")
	nodelist.Item(0).Export("node1")
	/*
		d.AddEventListener("mousemove", func(e event.Event) {
			println("mouse move", e.JSObject().Get("clientX").String(), e.JSObject().Get("clientY").String())
		})
	*/

	if clickbutton, err := d.GetElementById("clickme"); err == nil {

		clickbutton.AddEventListener("click", func(e event.Event) {

			if testinput, err := d.GetElementById("test"); err == nil {
				attributes, _ := testinput.Attributes()

				if attr, err := attributes.GetNamedItem("type"); err == nil {
					if str, err := attr.Value(); err == nil {
						println("type->" + str)
					}

				} else {
					println("erreur" + err.Error())
				}

				//easy method

				if obj, err := testinput.GetAttribute("type"); err == nil {
					println("Second method type->" + obj.Value())
				} else {
					println("erreur" + err.Error())
				}

			}

		})
	}

	ch := make(chan struct{})
	<-ch

}
