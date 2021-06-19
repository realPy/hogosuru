package element

// https://developer.mozilla.org/fr/docs/Web/API/Element
import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/attr"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/htmlcollection"

	"github.com/realPy/hogosuru/namednodemap"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var elementinterface js.Value

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if elementinterface, err = js.Global().GetWithErr("Element"); err != nil {
			elementinterface = js.Null()
		}

	})

	baseobject.Register(elementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return elementinterface
}

type Element struct {
	node.Node
}

func New() (Element, error) {
	var err error
	var e Element
	if ei := GetInterface(); !ei.IsNull() {
		e.BaseObject = e.SetObject(ei.New())

	} else {
		err = ErrNotImplemented
	}

	return e, err
}

func NewFromJSObject(obj js.Value) (Element, error) {
	var e Element
	var err error
	if ei := GetInterface(); !ei.IsNull() {
		if obj.InstanceOf(ei) {
			e.BaseObject = e.SetObject(obj)

		} else {
			err = ErrNotAnElement
		}

	} else {
		err = ErrNotImplemented
	}

	return e, err
}

func ItemFromHTMLCollection(collection htmlcollection.HTMLCollection, index int) (Element, error) {

	return NewFromJSObject(collection.Item(index))

}

func (e Element) getAttributeString(attribute string) (string, error) {

	var err error
	var obj js.Value
	var valueStr = ""

	if obj, err = e.JSObject().GetWithErr(attribute); err == nil {

		valueStr = obj.String()
	}

	return valueStr, err

}

func (e Element) setAttributeString(attribute string, value string) error {

	return e.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (e Element) getAttributeNumber(attribute string) (float64, error) {

	var err error
	var obj js.Value
	var ret float64

	if obj, err = e.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeNumber {
			ret = obj.Float()
		} else {
			err = baseobject.ErrObjectNotNumber
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
			err = baseobject.ErrObjectNotNumber
		}
	}
	return ret, err
}

func (e Element) getAttributeElement(attribute string) (Element, error) {
	var nodeObject js.Value
	var newElement Element
	var err error

	if nodeObject, err = e.JSObject().GetWithErr(attribute); err == nil {

		if nodeObject.IsNull() {
			err = ErrElementNoChilds

		} else {

			newElement, err = NewFromJSObject(nodeObject)

		}

	}

	return newElement, err
}

func (e Element) Attributes() (namednodemap.NamedNodeMap, error) {

	var err error
	var obj js.Value
	var namednmap namednodemap.NamedNodeMap

	if obj, err = e.JSObject().GetWithErr("attributes"); err == nil {
		namednmap, err = namednodemap.NewFromJSObject(obj)
	}
	return namednmap, err
}

func (e Element) ClassName() (string, error) {

	return e.getAttributeString("className")
}

func (e Element) SetClassName(value string) error {

	return e.setAttributeString("className", value)
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

	return e.getAttributeString("computedName")
}

func (e Element) ComputedRole() (string, error) {

	return e.getAttributeString("computedRole")
}

func (e Element) ID() (string, error) {

	return e.getAttributeString("id")
}

func (e Element) SetID(value string) error {

	return e.setAttributeString("id", value)
}

func (e Element) InnerHTML() (string, error) {

	return e.getAttributeString("innerHTML")
}

func (e Element) SetInnerHTML(value string) error {

	return e.setAttributeString("innerHTML", value)
}

func (e Element) LocalName() (string, error) {

	return e.getAttributeString("localname")
}

func (e Element) NamespaceURI() (string, error) {

	return e.getAttributeString("namespaceURI")
}

func (e Element) NextElementSibling() (Element, error) {
	return e.getAttributeElement("nextElementSibling")
}

func (e Element) OuterHTML() (string, error) {

	return e.getAttributeString("outerHTML")
}

func (e Element) SetOuterHTML(value string) error {

	return e.setAttributeString("outerHTML", value)
}

func (e Element) Prefix() (string, error) {

	return e.getAttributeString("prefix")
}

func (e Element) PreviousElementSibling() (Element, error) {
	return e.getAttributeElement("previousElementSibling")
}

func (e Element) ScrollHeight() (int, error) {

	return e.getAttributeInt("scrollHeight")
}

func (e Element) ScrollLeft() (int, error) {

	return e.getAttributeInt("scrollLeft")
}

func (e Element) ScrollTop() (int, error) {

	return e.getAttributeInt("scrollTop")
}

func (e Element) ScrollWidth() (int, error) {

	return e.getAttributeInt("scrollWidth")
}

func (e Element) TagName() (string, error) {

	return e.getAttributeString("tagName")
}

func OwnerElementForAttr(a attr.Attr) (Element, error) {
	var elemObject js.Value
	var newElement Element
	var err error

	if elemObject, err = a.JSObject().GetWithErr("ownerElement"); err == nil {

		if elemObject.IsNull() {
			err = attr.ErrNoOwnerElement
		} else {

			newElement, err = NewFromJSObject(elemObject)

		}

	}

	return newElement, err
}
