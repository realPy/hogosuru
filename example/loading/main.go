package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/example/loading/loadingview"
)

func main() {

	hogosuru.Router().DefaultRendering(&loadingview.LoadingGlobalContainer{})
	hogosuru.Router().Start(hogosuru.STDROUTE)
	ch := make(chan struct{})
	<-ch

}
