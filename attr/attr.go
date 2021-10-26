package attr

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var attrinterface js.Value

//GetInterface get the JS interface Attr
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if attrinterface, err = baseobject.Get(js.Global(), "Attr"); err != nil {
			attrinterface = js.Undefined()
		}
		baseobject.Register(attrinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return attrinterface
}

type Attr struct {
	node.Node
}

type AttrFrom interface {
	Attr_() Attr
}

func (a Attr) Attr_() Attr {
	return a
}

func NewFromJSObject(obj js.Value) (Attr, error) {
	var a Attr
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.InstanceOf(ai) {
			a.BaseObject = a.SetObject(obj)

		} else {
			err = ErrNotAnAttr
		}

	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func (a Attr) Name() (string, error) {

	return a.GetAttributeString("name")
}

func (a Attr) NamespaceURI() (string, error) {

	return a.GetAttributeString("namespaceURI")
}

func (a Attr) LocalName() (string, error) {

	return a.GetAttributeString("localName")
}

func (a Attr) Prefix() (string, error) {

	return a.GetAttributeString("prefix")
}

func (a Attr) Value() (string, error) {

	return a.GetAttributeString("value")
}

//use element.OwnerElementForAttr
//func (a Attr) OwnerElementObjet()
