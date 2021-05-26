package attr

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var attrinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var attrinstance JSInterface
		var err error
		if attrinstance.objectInterface, err = js.Global().GetWithErr("Attr"); err == nil {
			attrinterface = &attrinstance
		}
	})

	return attrinterface
}

type Attr struct {
	node.Node
}

func New() (Attr, error) {

	var a Attr
	var err error
	if ai := GetJSInterface(); ai != nil {
		a.BaseObject = a.SetObject(ai.objectInterface.New())

	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func NewFromJSObject(obj js.Value) (Attr, error) {
	var a Attr
	var err error
	if ai := GetJSInterface(); ai != nil {
		if obj.InstanceOf(ai.objectInterface) {
			a.BaseObject = a.SetObject(obj)

		} else {
			err = ErrNotAnAttr
		}

	} else {
		err = ErrNotImplemented
	}

	return a, err
}

func (a Attr) getStringAttribute(attribute string) (string, error) {

	var err error
	var obj js.Value
	var valueStr = ""

	if a.Error == nil {
		if obj, err = a.JSObject().GetWithErr(attribute); err == nil {
			if obj.IsNull() {
				valueStr = ""

			} else {

				valueStr = obj.String()
			}
		}
	} else {
		err = *a.Error
	}
	return valueStr, err

}

func (a Attr) Name() (string, error) {

	return a.getStringAttribute("name")
}

func (a Attr) NamespaceURI() (string, error) {

	return a.getStringAttribute("localName")
}

func (a Attr) LocalName() (string, error) {

	return a.getStringAttribute("localName")
}

func (a Attr) Prefix() (string, error) {

	return a.getStringAttribute("prefix")
}

func (a Attr) Value() (string, error) {

	return a.getStringAttribute("value")
}

//use element.OwnerElementForAttr
//func (a Attr) OwnerElementObjet()
