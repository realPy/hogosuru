package document

import (
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/nodelist"
)

func New_() Document {
	d, _ := New()
	return d
}

func (d Document) Body_() htmlelement.HtmlElement {
	body, _ := d.Body()
	return body
}

func (d Document) QuerySelectorAll_(selector string) nodelist.NodeList {

	n, _ := d.QuerySelectorAll(selector)
	return n
}
