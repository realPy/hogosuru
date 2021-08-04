package view

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/node"
)

type WebMain struct {
	divmain *htmldivelement.HtmlDivElement
}

func (w *WebMain) Node() node.Node {

	return w.divmain.Node
}

func (w *WebMain) OnLoad(d document.Document, n node.Node, route string) []hogosuru.Rendering {

	if divmain, err := htmldivelement.New(d); err == nil {
		divmain.SetID("MainView")

		w.divmain = &divmain
		if b, err := htmlbuttonelement.New(d); err == nil {
			b.SetTextContent("Go Hello")
			w.divmain.AppendChild(b.Node)
			w.divmain.InsertAdjacentText("beforeend", "sssssss")
			b.OnClick(func(e event.Event) {

				p, _ := w.divmain.ParentNode()
				p.RemoveChild(w.divmain.Node)

				w.divmain = nil

			})

		}

	}
	//components will be added to Node()
	return []hogosuru.Rendering{&ComplexComponents{}}
}

func (w *WebMain) OnUnload() {
	if w.divmain != nil {
		p, _ := w.divmain.ParentNode()
		p.RemoveChild(w.divmain.Node)
	}
}
