package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/example/loading/loadingview"
)

func main() {
	hogosuru.Init()
	hogosuru.Router().DefaultRendering(&loadingview.LoadingGlobalContainer{})
	hogosuru.Router().Add("/app/", &loadingview.RedSquare{})
	hogosuru.Router().Start(hogosuru.STDROUTE)
	ch := make(chan struct{})
	<-ch

}
