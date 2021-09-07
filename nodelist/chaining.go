package nodelist

import "github.com/realPy/hogosuru/node"

func (n NodeList) Item_(index int) node.Node {
	node, _ := n.Item(index)
	return node
}
