package element

import (
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/domrect"
	"github.com/realPy/hogosuru/object"
)

func (e Element) attachShadow() {
	//TODO IMPLEMENT
}

func (e Element) animate() {
	//TODO IMPLEMENT
}

func (e Element) Closest() (Element, error) {
	var err error
	var obj js.Value
	var elem Element

	if obj, err = e.JSObject().CallWithErr("closest"); err == nil {

		elem, err = NewFromJSObject(obj)
	}

	return elem, err
}

func (e Element) computedStyleMap() {
	//TODO IMPLEMENT
}

func (e Element) getAnimations() {
	//TODO IMPLEMENT
}

func (e Element) GetAttribute(attributename string) (object.Object, error) {

	var err error
	var obj js.Value
	var newobj object.Object

	if obj, err = e.JSObject().CallWithErr("getAttribute", js.ValueOf(attributename)); err == nil {
		if obj.IsNull() {
			err = ErrAttributeEmpty
		} else {
			newobj, err = object.NewFromJSObject(obj)
		}

	}
	return newobj, err
}

func (e Element) GetAttributeNames() (array.Array, error) {

	var err error
	var obj js.Value
	var arr array.Array

	if obj, err = e.JSObject().CallWithErr("getAttributeNames"); err == nil {
		if obj.IsNull() {
			err = ErrAttributeEmpty
		} else {
			arr, err = array.NewFromJSObject(obj)
		}

	}
	return arr, err
}
func (e Element) GetAttributeNS(namespace, name string) (object.Object, error) {
	var err error
	var obj js.Value
	var newobj object.Object

	if obj, err = e.JSObject().CallWithErr("getAttributeNS", js.ValueOf(namespace), js.ValueOf(name)); err == nil {
		if obj.IsNull() {
			err = ErrAttributeEmpty
		} else {
			newobj, err = object.NewFromJSObject(obj)
		}

	}
	return newobj, err
}

func (e Element) GetBoundingClientRect() (domrect.DOMRect, error) {
	var err error
	var obj js.Value
	var newdomrect domrect.DOMRect

	if obj, err = e.JSObject().CallWithErr("getBoundingClientRect"); err == nil {

		newdomrect, err = domrect.NewFromJSObject(obj)

	}
	return newdomrect, err
}

//retourne un DOMRectList
func (e Element) GetClientRects() (array.Array, error) {
	var err error
	var obj js.Value
	var arr array.Array

	if obj, err = e.JSObject().CallWithErr("getClientRects"); err == nil {

		arr, err = array.NewFromJSObject(obj)

	}
	return arr, err
}
