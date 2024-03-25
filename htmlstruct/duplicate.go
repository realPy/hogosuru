package htmlstruct

import (
	"errors"

	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/node"
)

func ClonableStruct(doc document.Document, root node.Node, i interface{}) (element.Element, error) {
	clone, err := doc.ImportNode(root.Node_(), true)
	if err != nil {
		return element.Element{}, err
	}
	el, ok := clone.(element.ElementFrom)
	if !ok {
		return element.Element{}, errors.New("can't clone struct: not an element")
	}

	clonelement := el.Element_()
	Unmarshal(clonelement, i)

	return clonelement, nil
}
