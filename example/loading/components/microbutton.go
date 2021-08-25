package components

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type ButtonD struct {
	node node.Node
}

func (l *ButtonD) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	if button, err := htmlbuttonelement.New(d); hogosuru.AssertErr(err) {
		button.SetID("pouet")
		button.SetTextContent("test")
		l.node = button.Node
	}

	return nil, nil
}

func (l *ButtonD) Node() node.Node {

	return l.node
}

func (l *ButtonD) OnUnload() {

}
