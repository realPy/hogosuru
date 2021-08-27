package components

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type ButtonD struct {
	parentNode node.Node
	button     htmlbuttonelement.HtmlButtonElement
}

func (l *ButtonD) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	l.parentNode = n
	if button, err := htmlbuttonelement.New(d); hogosuru.AssertErr(err) {
		button.SetID("pouet")
		button.SetTextContent("Loading Ok :)")
		l.button = button
	}

	return nil, nil
}

func (w *ButtonD) OnEndChildRendering(r hogosuru.Rendering) {

}

func (w *ButtonD) OnEndChildsRendering(tree node.Node) {
	w.parentNode.AppendChild(tree)
}

func (l *ButtonD) Node() node.Node {

	return l.button.Node
}

func (l *ButtonD) OnUnload() {

}
