package nodelist

import "github.com/realPy/hogosuru/base/node"

func (n NodeList) Item_(index int) node.Node {
	node, _ := n.Item(index)
	return node
}
