package namednodemap

// https://developer.mozilla.org/fr/docs/Web/API/NamedNodeMap

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/attr"
	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var namednodemapinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//NamedNodeMap struct
type NamedNodeMap struct {
	baseobject.BaseObject
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var nodelistinstance JSInterface
		var err error
		if nodelistinstance.objectInterface, err = js.Global().GetWithErr("NamedNodeMap"); err == nil {
			namednodemapinterface = &nodelistinstance
		}
	})

	return namednodemapinterface
}

func NewFromJSObject(obj js.Value) (NamedNodeMap, error) {
	var n NamedNodeMap

	if nli := GetJSInterface(); nli != nil {
		if obj.InstanceOf(nli.objectInterface) {
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

		if elemObject.IsNull() {
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

		if elemObject.IsNull() {
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