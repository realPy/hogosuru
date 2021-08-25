package loadingview

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/example/loading/components"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type LoadingGlobalContainer struct {
	node   node.Node
	loader components.Loader
	long   components.Long
}

func (w *LoadingGlobalContainer) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	if global, err := htmldivelement.New(d); err == nil {

		global.SetID("global-container")
		w.node = global.Node

	}

	return nil, []hogosuru.Rendering{&w.loader, &w.long}
}

func (w *LoadingGlobalContainer) Node() node.Node {

	return w.node
}

func (w *LoadingGlobalContainer) OnUnload() {

}
