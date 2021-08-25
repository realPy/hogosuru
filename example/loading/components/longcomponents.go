package components

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type Long struct {
	node node.Node
}

func (l *Long) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	var p promise.Promise
	l.node = n
	p, _ = promise.SetTimeout(func() (interface{}, error) {
		println("ended")
		return nil, nil
	}, 3000)

	return &p, []hogosuru.Rendering{&ButtonD{}}
}

func (l *Long) Node() node.Node {

	return l.node
}

func (l *Long) OnUnload() {

}
