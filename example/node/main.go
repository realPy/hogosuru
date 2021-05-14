package main

import "github.com/realPy/hogosuru/document"

func main() {

	d, _ := document.New()

	d1 := d.FirstChild()

	println("--->" + d1.Object.String())

	ch := make(chan struct{})
	<-ch

}
