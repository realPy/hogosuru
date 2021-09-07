package components

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type Long struct {
	parentNode  node.Node
	node        node.Node
	WaitingTime int
}

func (l *Long) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	l.parentNode = n
	var p promise.Promise
	l.node, _ = d.CreateDocumentFragment()

	p, _ = promise.SetTimeout(func() (interface{}, error) {
		return nil, nil
	}, l.WaitingTime)

	return &p, []hogosuru.Rendering{&ButtonD{}}
}

func (w *Long) OnEndChildRendering(r hogosuru.Rendering) {

}

func (w *Long) OnEndChildsRendering(tree node.Node) {
	w.parentNode.AppendChild(tree)
}

func (l *Long) Node() node.Node {

	return l.node
}

func (l *Long) OnUnload() {

}
