package view

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type GlobalContainer struct {
	node node.Node
}

func (w *GlobalContainer) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {
	if global, err := htmldivelement.New(d); err == nil {

		global.SetID("main-container")
		w.node = global.Node

	}
	return nil, nil
}

func (w *GlobalContainer) Node(r hogosuru.Rendering) node.Node {

	return w.node
}

func (w *GlobalContainer) OnEndChildsRendering() {

}
func (w *GlobalContainer) OnEndChildRendering(r hogosuru.Rendering) {

}

func (w *GlobalContainer) OnUnload() {

}

type WebMain struct {
	divmain *htmldivelement.HtmlDivElement
}

func (w *WebMain) Node(r hogosuru.Rendering) node.Node {

	return w.divmain.Node
}

func (w *WebMain) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	if divmain, err := htmldivelement.New(d); err == nil {
		divmain.SetID("MainView")

		w.divmain = &divmain

		if b, err := htmlbuttonelement.New(d); err == nil {
			b.SetTextContent("Go Hello")
			w.divmain.AppendChild(b.Node)
			b.OnClick(func(e event.Event) {
				hogosuru.Router().Go("/app/hello")
			})

		}

	}
	//components will be added to Node()
	return nil, []hogosuru.Rendering{&ComplexComponents{}}
}
func (w *WebMain) OnEndChildsRendering() {

}
func (w *WebMain) OnEndChildRendering(r hogosuru.Rendering) {

}

func (w *WebMain) OnUnload() {
	if w.divmain != nil {
		p, _ := w.divmain.ParentNode()
		p.RemoveChild(w.divmain.Node)
	}
}
