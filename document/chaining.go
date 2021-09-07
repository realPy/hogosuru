package document

import (
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/nodelist"
)

func New_() Document {
	d, _ := New()
	return d
}

func (d Document) Body_() node.Node {
	body, _ := d.Body()
	return body
}

func (d Document) QuerySelectorAll_(selector string) nodelist.NodeList {

	n, _ := d.QuerySelectorAll(selector)
	return n
}
