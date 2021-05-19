package document

import (
	"syscall/js"

	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/object"
)

func (d Document) getAttributeElement(attribute string) element.Element {
	var elem element.Element
	var elemObject js.Value
	var err error

	elem.Error = d.Error
	if d.Error == nil {
		if elemObject, err = d.JSObject().GetWithErr(attribute); err == nil {

			elem = element.NewFromJSObject(elemObject)

		} else {
			elem.Error = &err
		}

	}

	return elem
}

func (d Document) getAttributeString(attribute string) (string, error) {

	var err error
	var obj js.Value
	var cs string = ""
	if obj, err = d.JSObject().GetWithErr(attribute); err == nil {

		cs = obj.String()
	}
	return cs, err
}

func (d Document) setAttributeString(attribute string, value string) error {

	return d.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (d Document) getAttributeHTMLCollection(attribute string) (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = d.JSObject().GetWithErr(attribute); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) getAttributeBool(attribute string) (bool, error) {

	var err error
	var obj js.Value
	var ret bool

	if obj, err = d.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return ret, err
}

func (d Document) ActiveElement() element.Element {

	return d.getAttributeElement("activeElement")

}

func (d Document) Body() node.Node {
	var body node.Node
	var bodyObject js.Value
	var err error

	body.Error = d.Error
	if d.Error == nil {
		if bodyObject, err = d.JSObject().GetWithErr("body"); err == nil {

			body = node.NewFromJSObject(bodyObject)

		} else {
			body.Error = &err
		}

	}

	return body
}

func (d Document) CharacterSet() (string, error) {
	return d.getAttributeString("characterSet")
}

func (d Document) ChildElementCount() int {
	var err error
	var obj js.Value

	if obj, err = d.JSObject().GetWithErr("childElementCount"); err == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		}
	}

	return 0
}

func (d Document) Children() (htmlcollection.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("children")
}

func (d Document) CompatMode() (string, error) {
	return d.getAttributeString("compatMode")
}

func (d Document) ContentType() (string, error) {
	return d.getAttributeString("contentType")
}

func (d *Document) Doctype() {
	//TO IMPLEMENT
}

func (d Document) DocumentElement() element.Element {
	return d.getAttributeElement("documentElement")
}

func (d *Document) DocumentURI() (string, error) {
	return d.getAttributeString("documentURI")
}

func (d Document) Embeds() (htmlcollection.HTMLCollection, error) {

	return d.getAttributeHTMLCollection("embeds")
}

func (d Document) FirstElementChild() element.Element {
	return d.getAttributeElement("firstElementChild")
}

func (d Document) Fonts() {
	//TO IMPLEMENT
}

func (d Document) Forms() (htmlcollection.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("forms")
}

func (d Document) FullscreenElement() element.Element {
	return d.getAttributeElement("fullscreenElement")
}

func (d Document) Head() (htmlcollection.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("head")
}

func (d Document) Hidden() (bool, error) {

	return d.getAttributeBool("hidden")
}

func (d Document) Images() (htmlcollection.HTMLCollection, error) {

	return d.getAttributeHTMLCollection("images")
}
func (d Document) Implementation() {
	//TO IMPLEMENT
}

func (d Document) LastElementChild() element.Element {
	return d.getAttributeElement("lastElementChild")
}

func (d Document) Links() (htmlcollection.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("links")
}

func (d Document) PictureInPictureElement() element.Element {
	return d.getAttributeElement("pictureInPictureElement")
}

func (d Document) PictureInPictureEnabled() (bool, error) {
	return d.getAttributeBool("pictureInPictureEnabled")
}

func (d Document) Plugins() (htmlcollection.HTMLCollection, error) {
	return d.getAttributeHTMLCollection("plugins")
}

func (d Document) PointerLockElement() element.Element {
	return d.getAttributeElement("pointerLockElement")
}

func (d Document) Scripts() (htmlcollection.HTMLCollection, error) {

	return d.getAttributeHTMLCollection("scripts")
}

func (d Document) ScrollingElement() element.Element {
	return d.getAttributeElement("scrollingElement")
}

func (d Document) VisibilityState() (string, error) {

	return d.getAttributeString("visibilityState")
}

func (d Document) Domain() (string, error) {

	return d.getAttributeString("domain")
}

func (d Document) LastModified() (string, error) {

	return d.getAttributeString("lastModified")
}

func (d Document) SetDomain(domain string) error {

	return d.setAttributeString("domain", domain)
}

func (d Document) ReadyState() (string, error) {

	return d.getAttributeString("readyState")

}

func (d Document) Referrer() (string, error) {

	return d.getAttributeString("referrer")
}

func (d Document) Title() (string, error) {

	return d.getAttributeString("title")
}

func (d Document) URL() (string, error) {

	return d.getAttributeString("URL")

}

func (d Document) Cookie() (string, error) {
	return d.getAttributeString("cookie")
}

func (d Document) SetCookie(cookie string) error {
	return d.setAttributeString("cookie", cookie)
}
