package main

import (
	"github.com/realPy/hogosuru/document"
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
	ch := make(chan struct{})
	<-ch

}
