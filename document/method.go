package document

import (
	"syscall/js"

	"github.com/realPy/hogosuru/attr"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/htmlelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/nodelist"
)

func (d Document) AdoptNode(externalNode node.Node) (interface{}, error) {

	var err error
	var obj js.Value
	var r interface{}

	if obj, err = d.Call("adoptNode", externalNode.JSObject()); err == nil {
		r, err = baseobject.Discover(obj)
	}
	return r, err

}

func (d Document) Append(i interface{}) error {

	var err error
	var obji interface{}

	if objb, ok := i.(baseobject.ObjectFrom); ok {

		obji = objb.JSObject()

	} else {
		obji = js.ValueOf(i)

	}

	_, err = d.Call("append", obji)

	return err
}

func (d Document) CreateAttribute(name string) (attr.Attr, error) {

	var err error
	var obj js.Value
	var attribute attr.Attr

	if obj, err = d.Call("createAttribute", js.ValueOf(name)); err == nil {

		attribute, err = attr.NewFromJSObject(obj)
	}

	return attribute, err
}

func (d Document) CreateComment(comment string) (node.Node, error) {
	var err error
	var obj js.Value
	var nod node.Node

	if obj, err = d.Call("createComment", js.ValueOf(comment)); err == nil {

		nod, err = node.NewFromJSObject(obj)
	}

	return nod, err
}

func (d Document) CreateDocumentFragment() (node.Node, error) {
	var err error
	var obj js.Value
	var nod node.Node

	if obj, err = d.Call("createDocumentFragment"); err == nil {

		nod, err = node.NewFromJSObject(obj)
	}

	return nod, err
}

func (d Document) CreateHTMLElement(tagname string) (htmlelement.HtmlElement, error) {
	var err error
	var htmlelm htmlelement.HtmlElement
	var elem element.Element

	if elem, err = d.CreateElement(tagname); err == nil {
		htmlelm, err = htmlelement.NewFromElement(elem)

	}
	return htmlelm, err
}

func (d Document) CreateElement(tagname string) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.Call("createElement", js.ValueOf(tagname)); err == nil {

		elem, err = element.NewFromJSObject(obj)
	}

	return elem, err
}

func (d Document) CreateElementNS(namespaceURI string, qualifiedName string) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.Call("createElementNS", js.ValueOf(namespaceURI), js.ValueOf(qualifiedName)); err == nil {

		elem, err = element.NewFromJSObject(obj)
	}

	return elem, err
}

func (d Document) CreateEvent(eventtype string) (event.Event, error) {

	var err error
	var obj js.Value
	var ev event.Event

	if obj, err = d.Call("createEvent", js.ValueOf(eventtype)); err == nil {

		ev, err = event.NewFromJSObject(obj)
	}

	return ev, err
}

func (d Document) createNodeIterator() {
	//TO IMPLEMENT
}

func (d Document) createProcessingInstruction() {
	//TO IMPLEMENT
}

func (d Document) createRange() {
	//TO IMPLEMENT
}

func (d Document) createTreeWalker() {
	//TO IMPLEMENT
}

func (d Document) CreateTextNode(text string) (node.Node, error) {

	var err error
	var obj js.Value
	var nod node.Node

	if obj, err = d.Call("createTextNode", js.ValueOf(text)); err == nil {

		nod, err = node.NewFromJSObject(obj)
	}

	return nod, err
}

func (d Document) ElementFromPoint(x, y int) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.Call("elementFromPoint", js.ValueOf(x), js.ValueOf(y)); err == nil {

		elem, err = element.NewFromJSObject(obj)
	}

	return elem, err
}

func (d Document) ElementsFromPoint(x, y int) ([]element.Element, error) {

	var err error
	var obj js.Value
	var elems []element.Element

	if obj, err = d.Call("elementsFromPoint", js.ValueOf(x), js.ValueOf(y)); err == nil {

		for i := 0; i < obj.Length(); {
			if el, err := element.NewFromJSObject(obj.Index(i)); err == nil {
				elems = append(elems, el)
			}

		}

	}

	return elems, err
}

func (d Document) exitPictureInPicture() {
	//TO IMPLEMENT
}

func (d Document) ExitPointerLock() error {
	_, err := d.Call("exitPointerLock")
	return err
}

func (d Document) getAnimations() {
	//TO IMPLEMENT
}

func (d Document) GetElementsByClassName(classname string) (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = d.Call("getElementsByClassName", js.ValueOf(classname)); err == nil {

		if !obj.IsUndefined() && !obj.IsNull() {
			collection, err = htmlcollection.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}

	}

	return collection, err
}

func (d Document) GetElementsByTagName(tagname string) (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = d.Call("getElementsByTagName", js.ValueOf(tagname)); err == nil {

		if !obj.IsUndefined() && !obj.IsNull() {
			collection, err = htmlcollection.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}

	}

	return collection, err
}

func (d Document) GetElementsByTagNameNS(namespace, tagname string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = d.Call("getElementsByTagNameNS", js.ValueOf(namespace), js.ValueOf(tagname)); err == nil {

		if !obj.IsUndefined() && !obj.IsNull() {
			collection, err = htmlcollection.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}

	}

	return collection, err
}

func (d Document) ImportNode(externalNode node.Node, deep bool) (interface{}, error) {
	var err error
	var obj js.Value
	var r interface{}

	if obj, err = d.Call("importNode", externalNode.JSObject(), js.ValueOf(deep)); err == nil {
		r, err = baseobject.Discover(obj)
	}
	return r, err
}

func (d Document) ReleaseCapture() error {
	_, err := d.Call("releaseCapture")
	return err
}

func (d Document) GetElementById(id string) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.Call("getElementById", js.ValueOf(id)); err == nil {
		if !obj.IsUndefined() && !obj.IsNull() {
			elem, err = element.NewFromJSObject(obj)
		} else {
			err = ErrElementNotFound
		}

	}

	return elem, err
}

func (d Document) QuerySelector(selector string) (element.Element, error) {

	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.Call("querySelector", js.ValueOf(selector)); err == nil {
		if !obj.IsUndefined() && !obj.IsNull() {
			elem, err = element.NewFromJSObject(obj)
		} else {
			err = ErrElementNotFound
		}
	}
	return elem, err
}

func (d Document) QuerySelectorAll(selector string) (nodelist.NodeList, error) {

	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj, err = d.Call("querySelectorAll", js.ValueOf(selector)); err == nil {
		if !obj.IsUndefined() && !obj.IsNull() {
			nlist, err = nodelist.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}
	}
	return nlist, err
}
