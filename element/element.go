package element

// https://developer.mozilla.org/fr/docs/Web/API/Element
import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/object"
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

type Element struct {
	node.Node
}

func New() Element {

	var e Element
	if ei := GetJSInterface(); ei != nil {
		e.Object = e.SetObject(ei.objectInterface.New())
		return e
	}

	e.Error = &ErrNotImplemented
	return e
}

func NewFromJSObject(obj js.Value) Element {
	var e Element

	if ei := GetJSInterface(); ei != nil {
		if obj.InstanceOf(ei.objectInterface) {
			e.Object = e.SetObject(obj)
			return e
		}

	}

	e.Error = &ErrNotAnElement
	return e
}

func (e Element) getStringAttribute(attribute string) (string, error) {

	var err error
	var obj js.Value
	var valueStr = ""

	if e.Error == nil {
		if obj, err = e.JSObject().GetWithErr(attribute); err == nil {

			valueStr = obj.String()
		}
	} else {
		err = *e.Error
	}
	return valueStr, err

}

func (e Element) SetStringAttribute(attribute string, value string) error {
	var err error
	if e.Error == nil {
		if err = e.JSObject().SetWithErr(attribute, js.ValueOf(value)); err != nil {

			err = *e.Error
		}
	}
	return err
}

func (e Element) getAttributeNumber(attribute string) (float64, error) {

	var err error
	var obj js.Value
	var ret float64

	if obj, err = e.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeNumber {
			ret = obj.Float()
		} else {
			err = object.ErrObjectNotNumber
		}
	}
	return ret, err
}

func (e Element) getAttributeInt(attribute string) (int, error) {

	var err error
	var obj js.Value
	var ret int

	if obj, err = e.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeNumber {
			ret = obj.Int()
		} else {
			err = object.ErrObjectNotNumber
		}
	}
	return ret, err
}

func (e Element) getAttributeElement(attribute string) Element {
	var nodeObject js.Value
	var newElement Element
	var err error

	if e.Error != nil {
		return e
	}

	newElement.Error = e.Error
	if e.Error == nil {
		if nodeObject, err = e.JSObject().GetWithErr(attribute); err == nil {

			if nodeObject.IsNull() {
				err = ErrElementNoChilds

			} else {

				newElement = NewFromJSObject(nodeObject)

			}

		} else {
			newElement.Error = &err
		}

	}

	return newElement
}

func (e Element) ClassName() (string, error) {

	return e.getStringAttribute("className")
}

func (e Element) SetClassName(value string) error {

	return e.SetStringAttribute("className", value)
}

func (e Element) ClientHeight() (int, error) {

	return e.getAttributeInt("clientHeight")
}

func (e Element) ClientLeft() (int, error) {

	return e.getAttributeInt("clientLeft")
}

func (e Element) ClientTop() (int, error) {

	return e.getAttributeInt("clienTop")
}

func (e Element) ClientWidth() (int, error) {

	return e.getAttributeInt("clienWidth")
}

func (e Element) ComputedName() (string, error) {

	return e.getStringAttribute("computedName")
}

func (e Element) ComputedRole() (string, error) {

	return e.getStringAttribute("computedRole")
}

func (e Element) ID() (string, error) {

	return e.getStringAttribute("id")
}

func (e Element) SetID(value string) error {

	return e.SetStringAttribute("id", value)
}

func (e Element) InnerHTML() (string, error) {

	return e.getStringAttribute("innerHTML")
}

func (e Element) SetInnerHTML(value string) error {

	return e.SetStringAttribute("innerHTML", value)
}

func (e Element) LocalName() (string, error) {

	return e.getStringAttribute("localname")
}

func (e Element) NamespaceURI() (string, error) {

	return e.getStringAttribute("namespaceURI")
}

func (e Element) NextElementSibling() Element {
	return e.getAttributeElement("nextElementSibling")
}

func (e Element) OuterHTML() (string, error) {

	return e.getStringAttribute("outerHTML")
}

func (e Element) SetOuterHTML(value string) error {

	return e.SetStringAttribute("outerHTML", value)
}

func (e Element) Prefix() (string, error) {

	return e.getStringAttribute("prefix")
}

func (e Element) PreviousElementSibling() Element {
	return e.getAttributeElement("previousElementSibling")
}
