package view

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/node"
)

type HelloView struct {
	view *htmldivelement.HtmlDivElement
}

func (w *HelloView) Node() node.Node {

	return w.view.Node
}

func (w *HelloView) OnLoad(d document.Document, n node.Node, route string) []hogosuru.Rendering {
	var err error
	var elem htmldivelement.HtmlDivElement
	if elem, err = htmldivelement.New(d); err == nil {
		w.view = &elem
		w.view.SetTextContent("Hello")

	}

	return nil
}

func (w *HelloView) OnUnload() {
	if w.view != nil {
		p, _ := w.view.ParentNode()
		p.RemoveChild(w.view.Node)
	}
}
