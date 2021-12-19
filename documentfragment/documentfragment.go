package documentfragment

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/nodelist"
)

var singleton sync.Once

var documentfragementinterface js.Value

type DocumentFragment struct {
	node.Node
}

type DocumentFragmentFrom interface {
	DocumentFragment_() DocumentFragment
}

func (d DocumentFragment) DocumentFragment_() DocumentFragment {
	return d
}

func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if documentfragementinterface, err = baseobject.Get(js.Global(), "DocumentFragment"); err != nil {
			documentfragementinterface = js.Undefined()
		}
		baseobject.Register(documentfragementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return documentfragementinterface
}

func New() (DocumentFragment, error) {
	var d DocumentFragment
	var err error
	var di, obj js.Value
	if di = GetInterface(); di.IsUndefined() {
		return d, nil
	}
	if obj, err = baseobject.New(di); err != nil {
		return d, ErrNotImplemented
	}
	d.BaseObject = d.SetObject(obj)
	return d, nil
}

func NewFromJSObject(obj js.Value) (DocumentFragment, error) {
	var d DocumentFragment
	var dci js.Value
	if dci = GetInterface(); dci.IsUndefined() {
		return d, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return d, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(dci) {
		return d, ErrNotADocumentFragment
	}
	d.BaseObject = d.SetObject(obj)
	return d, nil
}

func (d DocumentFragment) ChildElementCount() (int, error) {
	return d.GetAttributeInt("childElementCount")
}

func (e DocumentFragment) Children() (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection
	if obj, err = e.Get("children"); err == nil {
		return htmlcollection.NewFromJSObject(obj)
	}
	return collection, err
}

func (e DocumentFragment) getAttributeElement(attribute string) (element.Element, error) {
	var nodeObject js.Value
	var newElement element.Element
	var err error
	if nodeObject, err = e.Get(attribute); err != nil {
		return newElement, err
	}
	if nodeObject.IsUndefined() {
		return newElement, element.ErrElementNoChilds
	}
	return element.NewFromJSObject(nodeObject)
}

func (d DocumentFragment) FirstElementChild() (element.Element, error) {
	return d.getAttributeElement("firstElementChild")
}

func (d DocumentFragment) LastElementChild() (element.Element, error) {
	return d.getAttributeElement("lastElementChild")
}

func (d DocumentFragment) nodesMethod(method string, elems ...interface{}) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elems {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(elem))
	}
	_, err = d.Call(method, arrayJS...)
	return err

}

func (d DocumentFragment) Prepend(elems ...interface{}) error {
	return d.nodesMethod("prepend", elems...)
}

func (d DocumentFragment) Append(elems ...interface{}) error {
	return d.nodesMethod("append", elems...)
}

func (d DocumentFragment) QuerySelector(selector string) (node.Node, error) {
	var err error
	var obj js.Value
	var nod node.Node
	if obj, err = d.Call("querySelector", js.ValueOf(selector)); err == nil {
		return node.NewFromJSObject(obj)
	}
	return nod, err
}

func (d DocumentFragment) QuerySelectorAll(selector string) (nodelist.NodeList, error) {
	var err error
	var obj js.Value
	var nlist nodelist.NodeList
	if obj, err = d.Call("querySelectorAll", js.ValueOf(selector)); err == nil {
		return nodelist.NewFromJSObject(obj)
	}
	return nlist, err
}

func (d DocumentFragment) ReplaceChild(new, old node.Node) (node.Node, error) {
	var err error
	_, err = d.Call("replaceChild", new.JSObject(), old.JSObject())
	return old, err

}

func (d DocumentFragment) GetElementById(id string) (element.Element, error) {
	var err error
	var obj js.Value
	var elem element.Element
	if obj, err = d.Call("getElementById", js.ValueOf(id)); err == nil {
		return element.NewFromJSObject(obj)
	}
	return elem, err
}
