package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/example/routing/view"
)

func main() {

	hogosuru.Router().DefaultRendering(&view.GlobalContainer{})
	hogosuru.Router().Add("", &view.WebMain{})
	hogosuru.Router().Add("hello", &view.HelloView{})
	hogosuru.Router().Start()
	ch := make(chan struct{})
	<-ch

}
