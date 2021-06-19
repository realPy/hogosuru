package main

import (
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlinputelement"
)

func main() {

	d := document.New_()

	nod := d.Body_()

	if text, err := nod.TextContent(); err == nil {
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

	nodelist := d.QuerySelectorAll_(".pictureContainer")
	println("Found", nodelist.Length(), "elements")
	nodelist.Item_(0).Export("node1")
	/*
		d.AddEventListener("mousemove", func(e event.Event) {
			println("mouse move", e.JSObject().Get("clientX").String(), e.JSObject().Get("clientY").String())
		})
	*/

	if clickbutton, err := d.GetElementById("clickme"); err == nil {

		clickbutton.OnClick(func(e event.Event) {

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

				if str, err := testinput.GetAttribute("type"); err == nil {
					println("Second method type->" + str)
				} else {
					println("erreur" + err.Error())
				}

			}

		})
	}

	p, _ := d.CreateElement("input")
	p.SetAttribute("type", "checkbox")
	//	h, _ := htmlelement.NewFromElement(p)
	h, _ := htmlinputelement.NewFromElement(p)
	h.SetChecked(true)
	nod.AppendChild(h.Node)
	h.Focus()

	//h.SetHidden(true)
	h.Export("mat")

	h.SetDataset("toto", "value")

	v, _ := h.Dataset("toto")
	println(v.(string))
	ch := make(chan struct{})
	<-ch

}
