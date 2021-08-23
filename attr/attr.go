package attr

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var attrinterface js.Value

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if attrinterface, err = js.Global().GetWithErr("Attr"); err != nil {
			attrinterface = js.Null()
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
	Attr() Attr
}

func (a Attr) Attr() Attr {
	return a
}

func New() (Attr, error) {

	var a Attr
	var err error
	if ai := GetInterface(); !ai.IsNull() {
		a.BaseObject = a.SetObject(ai.New())

	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func NewFromJSObject(obj js.Value) (Attr, error) {
	var a Attr
	var err error
	if ai := GetInterface(); !ai.IsNull() {
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

	return a.GetAttributeString("localName")
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
