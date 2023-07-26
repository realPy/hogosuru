package element

import (
	"errors"
	"syscall/js"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/domrect"
	"github.com/realPy/hogosuru/base/domrectlist"
	"github.com/realPy/hogosuru/base/htmlcollection"
	"github.com/realPy/hogosuru/base/node"
	"github.com/realPy/hogosuru/base/nodelist"
	"github.com/realPy/hogosuru/base/object"
)

func (e Element) attachShadow() {
	//TODO IMPLEMENT
}

func (e Element) Animate(keyframes, options interface{}) error {
	var argCall []interface{}

	var err error
	if keyframesObject, ok := keyframes.(array.ArrayFrom); ok {
		argCall = append(argCall, keyframesObject.Array_().JSObject())

	}

	if keyframesObject, ok := keyframes.(object.ObjectFrom); ok {
		argCall = append(argCall, keyframesObject.Object_().JSObject())
	}

	if optionsObject, ok := keyframes.(object.ObjectFrom); ok {
		argCall = append(argCall, optionsObject.Object_().JSObject())
	} else {
		argCall = append(argCall, js.ValueOf(options))
	}
	_, err = e.Call("animate")
	return err
}

func (e Element) After(elements ...Element) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elements {
		arrayJS = append(arrayJS, elem.JSObject())
	}

	_, err = e.Call("after", arrayJS...)

	return err
}

func (e Element) Append(elements ...Element) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elements {
		arrayJS = append(arrayJS, elem.JSObject())
	}

	_, err = e.Call("append", arrayJS...)

	return err
}

func (e Element) Before(elements ...Element) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elements {
		arrayJS = append(arrayJS, elem.JSObject())
	}

	_, err = e.Call("before", arrayJS...)

	return err
}

func (e Element) Closest(query string) (Element, error) {
	var err error
	var obj js.Value
	var elem Element

	if obj, err = e.Call("closest", js.ValueOf(query)); err == nil {

		elem, err = NewFromJSObject(obj)
	}

	return elem, err
}

func (e Element) computedStyleMap() {
	//TODO IMPLEMENT
}

func (e Element) getAnimations() {
	//TODO IMPLEMENT
}

func (e Element) GetAttribute(attributename string) (string, error) {

	var err error
	var obj js.Value
	var newstr string

	if obj, err = e.Call("getAttribute", js.ValueOf(attributename)); err == nil {
		if obj.IsNull() {
			err = ErrAttributeEmpty
		} else {
			newstr = obj.String()
		}

	}
	return newstr, err
}

func (e Element) GetAttributeNames() (array.Array, error) {

	var err error
	var obj js.Value
	var arr array.Array

	if obj, err = e.Call("getAttributeNames"); err == nil {
		if obj.IsNull() {
			err = ErrAttributeEmpty
		} else {
			arr, err = array.NewFromJSObject(obj)
		}

	}
	return arr, err
}
func (e Element) GetAttributeNS(namespace, name string) (object.Object, error) {
	var err error
	var obj js.Value
	var newobj object.Object

	if obj, err = e.Call("getAttributeNS", js.ValueOf(namespace), js.ValueOf(name)); err == nil {
		if obj.IsNull() {
			err = ErrAttributeEmpty
		} else {
			newobj, err = object.NewFromJSObject(obj)
		}

	}
	return newobj, err
}

func (e Element) GetBoundingClientRect() (domrect.DOMRect, error) {
	var err error
	var obj js.Value
	var newdomrect domrect.DOMRect

	if obj, err = e.Call("getBoundingClientRect"); err == nil {

		newdomrect, err = domrect.NewFromJSObject(obj)

	}
	return newdomrect, err
}

// retourne un DOMRectList
func (e Element) GetClientRects() (domrectlist.DOMRectList, error) {
	var err error
	var obj js.Value
	var arr domrectlist.DOMRectList

	if obj, err = e.Call("getClientRects"); err == nil {

		arr, err = domrectlist.NewFromJSObject(obj)

	}
	return arr, err
}

func (e Element) GetElementsByClassName(classname string) (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = e.Call("getElementsByClassName", js.ValueOf(classname)); err == nil {

		if !obj.IsNull() {
			collection, err = htmlcollection.NewFromJSObject(obj)
		} else {
			err = ErrElementsNotFound
		}

	}

	return collection, err
}

func (e Element) GetElementsByTagName(tagname string) (htmlcollection.HtmlCollection, error) {

	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = e.Call("getElementsByTagName", js.ValueOf(tagname)); err == nil {
		if obj.IsNull() || obj.IsUndefined() {
			err = ErrElementsNotFound

		} else {
			collection, err = htmlcollection.NewFromJSObject(obj)
		}
	}

	return collection, err
}

func (e Element) GetElementsByTagNameNS(namespace, tagname string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = e.Call("getElementsByTagNameNS", js.ValueOf(namespace), js.ValueOf(tagname)); err == nil {
		if obj.IsNull() || obj.IsUndefined() {
			err = ErrElementsNotFound
		} else {
			collection, err = htmlcollection.NewFromJSObject(obj)

		}
	}

	return collection, err
}

func (e Element) HasAttribute(attributename string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = e.Call("hasChildNodes", js.ValueOf(attributename)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err

}

func (e Element) HasPointerCapture(pointerid int) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = e.Call("hasPointerCapture", js.ValueOf(pointerid)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}
	return result, err
}

func (e Element) InsertAdjacentElement(position string, elem Element) (Element, error) {
	var elemObject js.Value
	var newelem Element
	var err error

	if elemObject, err = e.Call("insertAdjacentElement", js.ValueOf(position), elem.JSObject()); err == nil {

		if elemObject.IsNull() {
			err = ErrInsertAdjacent

		} else {
			newelem = elem
		}

	}
	return newelem, err
}

func (e Element) InsertAdjacentHTML(position string, textHTML string) error {

	var err error

	_, err = e.Call("insertAdjacentHTML", js.ValueOf(position), js.ValueOf(textHTML))
	return err
}

func (e Element) InsertAdjacentText(position string, text string) error {

	var err error

	_, err = e.Call("insertAdjacentText", js.ValueOf(position), js.ValueOf(text))
	return err
}

func (e Element) Matches(selector string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = e.Call("matches", js.ValueOf(selector)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}
	return result, err
}
func (e Element) pseudo() {
	//TODO IMPLEMENT
}

func (e Element) Prepend(elements ...Element) error {
	var err error
	var arrayJS []interface{}

	for _, elem := range elements {
		arrayJS = append(arrayJS, elem.JSObject())
	}

	_, err = e.Call("prepend", arrayJS...)

	return err
}

func (e Element) QuerySelector(selector string) (Element, error) {

	var err error
	var obj js.Value
	var nod Element

	if obj, err = e.Call("querySelector", js.ValueOf(selector)); err == nil {
		if !obj.IsNull() {
			nod, err = NewFromJSObject(obj)
		} else {
			err = errors.New(ErrElementNotFound.Error() + " " + selector)
		}
	}
	return nod, err
}

func (e Element) QuerySelectorAll(selector string) (nodelist.NodeList, error) {

	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj, err = e.Call("querySelectorAll", js.ValueOf(selector)); err == nil {
		if !obj.IsNull() {
			nlist, err = nodelist.NewFromJSObject(obj)
		} else {
			err = errors.New(ErrElementsNotFound.Error() + " " + selector)
		}
	}
	return nlist, err
}

func (e Element) ReleasePointerCapture(pointerid int) error {
	var err error
	_, err = e.Call("releasePointerCapture", js.ValueOf(pointerid))
	return err
}
func (e Element) Remove() error {
	var err error
	_, err = e.Call("remove")
	return err
}

func (e Element) RemoveAttribute(attrname string) error {
	var err error
	_, err = e.Call("removeAttribute", js.ValueOf(attrname))
	return err
}

func (e Element) RemoveAttributeNS(namespace, attrname string) error {
	var err error
	_, err = e.Call("removeAttributeNS", js.ValueOf(namespace), js.ValueOf(attrname))
	return err
}

func (e Element) ReplaceChildren(params ...interface{}) error {
	var err error
	var arrayJS []interface{}
	for _, param := range params {
		switch p := param.(type) {
		case node.Node:
			arrayJS = append(arrayJS, p.JSObject())
		case string:
			arrayJS = append(arrayJS, js.ValueOf(p))
		default:
			return ErrSendUnknownType
		}
	}

	_, err = e.Call("replaceChildren", arrayJS...)

	return err
}

func (e Element) RequestFullscreen() error {
	var err error
	_, err = e.Call("requestFullscreen")
	return err
}

func (e Element) RequestPointerLock() error {
	var err error
	_, err = e.Call("requestPointerLock")
	return err
}

func (e Element) Scroll(x, y int, opts ...map[string]interface{}) error {
	var err error
	var optJSValue []interface{}

	optJSValue = append(optJSValue, js.ValueOf(x))
	optJSValue = append(optJSValue, js.ValueOf(y))
	if opts != nil && len(opts) == 1 {
		optJSValue = append(optJSValue, js.ValueOf(opts[0]))
	}
	_, err = e.Call("scroll", optJSValue...)
	return err
}
func (e Element) ScrollTo(x, y int, opts ...map[string]interface{}) error {
	var err error
	var optJSValue []interface{}

	optJSValue = append(optJSValue, js.ValueOf(x))
	optJSValue = append(optJSValue, js.ValueOf(y))
	if opts != nil && len(opts) == 1 {
		optJSValue = append(optJSValue, js.ValueOf(opts[0]))
	}
	_, err = e.Call("scrollTo", optJSValue...)
	return err
}

func (e Element) SetAttribute(name, value string) error {
	var err error
	_, err = e.Call("setAttribute", js.ValueOf(name), js.ValueOf(value))
	return err
}
func (e Element) SetAttributeNS(namespace, name, value string) error {
	var err error
	_, err = e.Call("setAttributeNS", js.ValueOf(namespace), js.ValueOf(name), js.ValueOf(value))
	return err
}

func (e Element) SetPointerCapture(pointerid int) error {
	var err error
	_, err = e.Call("setPointerCapture", js.ValueOf(pointerid))
	return err
}

func (e Element) ToggleAttribute(name string, opts ...interface{}) (bool, error) {
	var err error
	var optJSValue []interface{}
	var obj js.Value
	var result bool

	optJSValue = append(optJSValue, js.ValueOf(name))
	if opts != nil && len(opts) == 1 {
		optJSValue = append(optJSValue, js.ValueOf(opts[0]))
	}

	if obj, err = e.Call("toggleAttribute", optJSValue...); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}
	return result, err
}
