package element

// https://developer.mozilla.org/fr/docs/Web/API/Element
import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/attr"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/domtokenlist"
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
		if elementinterface, err = baseobject.Get(js.Global(), "Element"); err != nil {
			elementinterface = js.Undefined()
		}
		baseobject.Register(elementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return elementinterface
}

type Element struct {
	node.Node
}

type ElementFrom interface {
	Element_() Element
}

func (e Element) Element_() Element {
	return e
}

func New() (Element, error) {
	var err error
	var e Element
	if ei := GetInterface(); !ei.IsUndefined() {
		e.BaseObject = e.SetObject(ei.New())

	} else {
		err = ErrNotImplemented
	}

	return e, err
}

func NewFromJSObject(obj js.Value) (Element, error) {
	var e Element
	var err error
	if ei := GetInterface(); !ei.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ei) {
				e.BaseObject = e.SetObject(obj)

			} else {
				err = ErrNotAnElement
			}
		}

	} else {
		err = ErrNotImplemented
	}

	return e, err
}

func ItemFromHTMLCollection(collection htmlcollection.HtmlCollection, index int) (Element, error) {
	var elem Element
	var err error
	var item interface{}
	if item, err = collection.Item(index); err == nil {
		elem, err = NewFromJSObject(item.(baseobject.ObjectFrom).JSObject())
	}
	return elem, err

}

func (e Element) getAttributeElement(attribute string) (Element, error) {
	var nodeObject js.Value
	var newElement Element
	var err error

	if nodeObject, err = e.Get(attribute); err == nil {

		if nodeObject.IsUndefined() {
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

	if obj, err = e.Get("attributes"); err == nil {
		namednmap, err = namednodemap.NewFromJSObject(obj)
	}
	return namednmap, err
}

func (e Element) ChildElementCount() (int, error) {
	return e.GetAttributeInt("childElementCount")
}

func (e Element) Children() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = e.Get("children"); err == nil {

		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (e Element) ClassList() (domtokenlist.DOMTokenList, error) {
	var err error
	var obj js.Value
	var dlist domtokenlist.DOMTokenList

	if obj, err = e.Get("classList"); err == nil {

		dlist, err = domtokenlist.NewFromJSObject(obj)
	}

	return dlist, err
}

func (e Element) ClassName() (string, error) {

	return e.GetAttributeString("className")
}

func (e Element) SetClassName(value string) error {

	return e.SetAttributeString("className", value)
}

func (e Element) ClientHeight() (int, error) {

	return e.GetAttributeInt("clientHeight")
}

func (e Element) ClientLeft() (int, error) {

	return e.GetAttributeInt("clientLeft")
}

func (e Element) ClientTop() (int, error) {

	return e.GetAttributeInt("clienTop")
}

func (e Element) ClientWidth() (int, error) {

	return e.GetAttributeInt("clienWidth")
}

func (e Element) ComputedName() (string, error) {

	return e.GetAttributeString("computedName")
}

func (e Element) ComputedRole() (string, error) {

	return e.GetAttributeString("computedRole")
}

func (e Element) ID() (string, error) {

	return e.GetAttributeString("id")
}

func (e Element) SetID(value string) error {

	return e.SetAttributeString("id", value)
}

func (e Element) InnerHTML() (string, error) {

	return e.GetAttributeString("innerHTML")
}

func (e Element) SetInnerHTML(value string) error {

	return e.SetAttributeString("innerHTML", value)
}

func (e Element) LocalName() (string, error) {

	return e.GetAttributeString("localname")
}

func (e Element) NamespaceURI() (string, error) {

	return e.GetAttributeString("namespaceURI")
}

func (e Element) NextElementSibling() (Element, error) {
	return e.getAttributeElement("nextElementSibling")
}

func (e Element) OuterHTML() (string, error) {

	return e.GetAttributeString("outerHTML")
}

func (e Element) SetOuterHTML(value string) error {

	return e.SetAttributeString("outerHTML", value)
}

func (e Element) Prefix() (string, error) {

	return e.GetAttributeString("prefix")
}

func (e Element) PreviousElementSibling() (Element, error) {
	return e.getAttributeElement("previousElementSibling")
}

func (e Element) ScrollHeight() (int, error) {

	return e.GetAttributeInt("scrollHeight")
}

func (e Element) ScrollLeft() (int, error) {

	return e.GetAttributeInt("scrollLeft")
}

func (e Element) ScrollTop() (int, error) {

	return e.GetAttributeInt("scrollTop")
}

func (e Element) ScrollWidth() (int, error) {

	return e.GetAttributeInt("scrollWidth")
}

func (e Element) TagName() (string, error) {

	return e.GetAttributeString("tagName")
}

func OwnerElementForAttr(a attr.Attr) (Element, error) {
	var elemObject js.Value
	var newElement Element
	var err error

	if elemObject, err = a.Get("ownerElement"); err == nil {

		if elemObject.IsUndefined() {
			err = attr.ErrNoOwnerElement
		} else {

			newElement, err = NewFromJSObject(elemObject)

		}

	}

	return newElement, err
}
