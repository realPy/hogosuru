package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/example/routing/view"
)

func main() {

	hogosuru.Router().DefaultRendering(&view.GlobalContainer{})
	hogosuru.Router().Add("/app/", &view.WebMain{})
	hogosuru.Router().Add("/app/hello", &view.HelloView{})
	hogosuru.Router().Start(hogosuru.STDROUTE)
	ch := make(chan struct{})
	<-ch

}
