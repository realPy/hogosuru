package namednodemap

// https://developer.mozilla.org/fr/docs/Web/API/NamedNodeMap

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/attr"
	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var namednodemapinterface js.Value

//NamedNodeMap struct
type NamedNodeMap struct {
	baseobject.BaseObject
}

type NamedNodeMapFrom interface {
	NamedNodeMap_() NamedNodeMap
}

func (n NamedNodeMap) NamedNodeMap_() NamedNodeMap {
	return n
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if namednodemapinterface, err = baseobject.Get(js.Global(), "NamedNodeMap"); err != nil {
			namednodemapinterface = js.Undefined()
		}
		baseobject.Register(namednodemapinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return namednodemapinterface
}

func NewFromJSObject(obj js.Value) (NamedNodeMap, error) {
	var n NamedNodeMap

	if nli := GetInterface(); !nli.IsUndefined() {
		if obj.InstanceOf(nli) {
			n.BaseObject = n.SetObject(obj)
			return n, nil
		}
	}
	return n, ErrNotANamedNodeMap
}

func (n NamedNodeMap) Item(index int) (attr.Attr, error) {

	return attr.NewFromJSObject(n.JSObject().Index(index))
}

func (n NamedNodeMap) GetNamedItem(name string) (attr.Attr, error) {
	var elemObject js.Value
	var newAttr attr.Attr
	var err error

	if elemObject, err = n.JSObject().CallWithErr("getNamedItem", js.ValueOf(name)); err == nil {

		if elemObject.IsUndefined() {
			err = ErrNotNameAttr

		} else {

			newAttr, err = attr.NewFromJSObject(elemObject)

		}
	}

	return newAttr, err
}

func (n NamedNodeMap) SetNamedItem(a attr.Attr) error {
	var err error
	_, err = n.JSObject().CallWithErr("setNamedItem", a.JSObject())
	return err
}

func (n NamedNodeMap) RemoveNamedItem(name string) error {
	var err error
	_, err = n.JSObject().CallWithErr("removeNamedItem", js.ValueOf(name))
	return err
}

func (n NamedNodeMap) GetNamedItemNS(namespace string, name string) (attr.Attr, error) {
	var err error
	var elemObject js.Value
	var newAttr attr.Attr

	if elemObject, err = n.JSObject().CallWithErr("getNamedItemNS", js.ValueOf(namespace), js.ValueOf(name)); err == nil {

		if elemObject.IsUndefined() {
			err = ErrNotNameAttr

		} else {

			newAttr, err = attr.NewFromJSObject(elemObject)

		}
	}

	return newAttr, err

}

func (n NamedNodeMap) SetNamedItemNS(namespace string, a attr.Attr) error {
	var err error
	_, err = n.JSObject().CallWithErr("setNamedItemNS", js.ValueOf(namespace), a.JSObject())
	return err
}

func (n NamedNodeMap) RemoveNamedItemNS(namespace string, name string) error {
	var err error
	_, err = n.JSObject().CallWithErr("removeNamedItemNS", js.ValueOf(namespace), js.ValueOf(name))
	return err
}
