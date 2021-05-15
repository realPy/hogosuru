package main

import "github.com/realPy/hogosuru/document"

func main() {

	d, _ := document.New()

	d1 := d.FirstChild().FirstChild()

	d1.Export("d1")
	d1.NextSibling().Export("d2")
	ch := make(chan struct{})
	<-ch

}
