package document

import (
	"syscall/js"

	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/node"
)

func (d Document) QuerySelector(selector string) (js.Value, error) {

	return d.JSObject().CallWithErr("querySelector", js.ValueOf(selector))
}

func (d Document) CreateElement(tagname string) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.JSObject().CallWithErr("createElement", js.ValueOf(tagname)); err == nil {

		elem = element.NewFromJSObject(obj)
	}

	return elem, err
}

func (d Document) CreateTextNode(text string) (node.Node, error) {

	var err error
	var obj js.Value
	var nod node.Node

	if obj, err = d.JSObject().CallWithErr("createTextNode", js.ValueOf(text)); err == nil {

		nod = node.NewFromJSObject(obj)
	}

	return nod, err
}
