package main

import (
	"github.com/realPy/hogosuru/document"
)

func main() {

	d := document.New()

	d1 := d.FirstChild().FirstChild()

	d1.Export("d1")
	d1.NextSibling().Export("d2")

	nod := d.Body()

	if text := nod.TextContent(); nod.Error == nil {
		println("<--" + text + "-->")
	}

	ch := make(chan struct{})
	<-ch

}
