package element

import (
	"syscall/js"

	"github.com/realPy/hogosuru/object"
)

func (e Element) attachShadow() {
	//TODO IMPLEMENT
}

func (e Element) animate() {
	//TODO IMPLEMENT
}

func (e Element) Closest() Element {
	var err error
	var obj js.Value
	var elem Element

	if obj, err = e.JSObject().CallWithErr("closest"); err == nil {

		elem = NewFromJSObject(obj)
	}
	elem.Error = &err
	return elem
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
