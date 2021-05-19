package attr

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var elementinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var elementinstance JSInterface
		var err error
		if elementinstance.objectInterface, err = js.Global().GetWithErr("Element"); err == nil {
			elementinterface = &elementinstance
		}
	})

	return elementinterface
}

type Attr struct {
	node.Node
}

func New() Attr {

	var a Attr
	if ai := GetJSInterface(); ai != nil {
		a.Object = a.SetObject(ai.objectInterface.New())
		return a
	}

	a.Error = &ErrNotImplemented
	return a
}

func NewFromJSObject(obj js.Value) Attr {
	var a Attr

	if ai := GetJSInterface(); ai != nil {
		if obj.InstanceOf(ai.objectInterface) {
			a.Object = a.SetObject(obj)
			return a
		}

	}

	a.Error = &ErrNotAnAttr
	return a
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

func (a Attr) OwnerElement() element.Element {
	var elemObject js.Value
	var newElement element.Element
	var err error

	newElement.Error = a.Error
	if a.Error == nil {
		if elemObject, err = a.JSObject().GetWithErr("ownerElement"); err == nil {

			if elemObject.IsNull() {
				err = ErrNoOwnerElement

			} else {

				newElement = element.NewFromJSObject(elemObject)

			}

		} else {
			newElement.Error = &err
		}

	}

	return newElement
}
