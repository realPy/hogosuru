package components

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmlprogresselement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type Loader struct {
	node     node.Node
	progress htmlprogresselement.HtmlProgressElement
}

func (l *Loader) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	if style, err := d.CreateElement("style"); hogosuru.AssertErr(err) {

		if head, err := d.Head(); hogosuru.AssertErr(err) {

			style.SetInnerHTML(loadercss)
			head.AppendChild(style.Node)

		}

	}

	if loader, err := htmldivelement.New(d); hogosuru.AssertErr(err) {
		loader.SetID("loader-container")
		loader.SetClassName("loader")

		loader.SetInnerHTML(loaderhtml)
		l.node = loader.Node
	}
	return nil, nil
}

func (l *Loader) Node() node.Node {

	return l.node
}

func (l *Loader) OnUnload() {

}
