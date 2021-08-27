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
	parentNode node.Node
	node       node.Node
	loader     components.Loader
	long       components.Long
	long2      components.Long
}

func (w *LoadingGlobalContainer) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	w.parentNode = n
	w.long.WaitingTime = 500
	w.long2.WaitingTime = 1000

	if global, err := htmldivelement.New(d); err == nil {

		global.SetID("global-container")
		w.node = global.Node

	}

	//if no promise return we dont wait all childs to append

	return nil, []hogosuru.Rendering{&w.loader, &w.long, &w.long2}
}

func (w *LoadingGlobalContainer) Node() node.Node {

	return w.node
}

func (w *LoadingGlobalContainer) OnEndChildRendering(r hogosuru.Rendering) {
	if r == &w.loader {

	}

	if r == &w.long2 {

		w.loader.Hide(true)
	}
	if r == &w.long {
		w.loader.SetProgressValue(50)
	}
}

func (w *LoadingGlobalContainer) OnEndChildsRendering(tree node.Node) {

	w.parentNode.AppendChild(tree)
}

func (w *LoadingGlobalContainer) OnUnload() {

}
