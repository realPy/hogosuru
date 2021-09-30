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

func (b *ButtonD) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	b.parentNode = n
	if button, err := htmlbuttonelement.New(d); hogosuru.AssertErr(err) {
		button.SetID("pouet")
		button.SetTextContent("Loading Ok :)")
		b.button = button
	}

	return nil, nil
}

func (b *ButtonD) OnEndChildRendering(r hogosuru.Rendering) {

}

func (b *ButtonD) OnEndChildsRendering() {
	b.parentNode.AppendChild(b.button.Node)
}

func (b *ButtonD) Node(r hogosuru.Rendering) node.Node {

	return b.button.Node
}

func (l *ButtonD) OnUnload() {

}
