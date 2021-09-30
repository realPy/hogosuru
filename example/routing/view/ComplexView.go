package view

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmlspanelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type ComplexComponents struct {
	divcomplex *htmlspanelement.HtmlSpanElement
}

func (w *ComplexComponents) Node(r hogosuru.Rendering) node.Node {

	return w.divcomplex.Node
}

func (w *ComplexComponents) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	if divcomplex, err := htmlspanelement.New(d); err == nil {
		w.divcomplex = &divcomplex
		divcomplex.SetTextContent("Complex")
		divcomplex.SetDataset("toto", "value")

	}
	return nil, nil
}

func (w *ComplexComponents) OnEndChildsRendering() {

}
func (w *ComplexComponents) OnEndChildRendering(r hogosuru.Rendering) {

}

func (w *ComplexComponents) OnUnload() {

}
