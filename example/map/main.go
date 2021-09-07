package main

import (
	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/objectmap"
)

//eval code

func main() {
	a := array.New_(array.New_("a", "b"), array.New_("c"))
	a.Export("po")
	m := objectmap.New_(a)
	m.Export("test")

	ch := make(chan struct{})
	<-ch

}
