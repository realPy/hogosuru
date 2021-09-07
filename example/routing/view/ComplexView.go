package view

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmlspanelement"
	"github.com/realPy/hogosuru/node"
)

type ComplexComponents struct {
	divcomplex *htmlspanelement.HtmlSpanElement
}

func (w *ComplexComponents) Node() node.Node {

	return w.divcomplex.Node
}

func (w *ComplexComponents) OnLoad(d document.Document, n node.Node, route string) []hogosuru.Rendering {

	if divcomplex, err := htmlspanelement.New(d); err == nil {
		w.divcomplex = &divcomplex
		divcomplex.SetTextContent("Complex")
		divcomplex.SetDataset("toto", "value")

	}
	return nil
}
func (w *ComplexComponents) OnUnload() {

}
