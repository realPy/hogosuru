package document

import (
	"syscall/js"

	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/node"
)

func (d Document) getAttributeElement(attribute string) (element.Element, error) {
	var elem element.Element
	var elemObject js.Value
	var err error

	if elemObject, err = d.JSObject().GetWithErr(attribute); err == nil {

		elem, err = element.NewFromJSObject(elemObject)

	}

	return elem, err
}

func (d Document) getAttributeHTMLCollection(attribute string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = d.JSObject().GetWithErr(attribute); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) ActiveElement() (element.Element, error) {

	return d.getAttributeElement("activeElement")

}

func (d Document) Body() (node.Node, error) {
	var body node.Node
	var bodyObject js.Value
	var err error

	if bodyObject, err = d.JSObject().GetWithErr("body"); err == nil {

		body, err = node.NewFromJSObject(bodyObject)

	}

	return body, err
}

func (d Document) CharacterSet() (string, error) {
	return d.GetAttributeString("characterSet")
}

func (d Document) ChildElementCount() (int, error) {
	return d.GetAttributeInt("childElementCount")
}

func (d Document) Children() (htmlcollection.HtmlCollection, error) {
	return d.getAttributeHTMLCollection("children")
}

func (d Document) CompatMode() (string, error) {
	return d.GetAttributeString("compatMode")
}

func (d Document) ContentType() (string, error) {
	return d.GetAttributeString("contentType")
}

func (d *Document) Doctype() {
	//TO IMPLEMENT
}

func (d Document) DocumentElement() (element.Element, error) {
	return d.getAttributeElement("documentElement")
}

func (d *Document) DocumentURI() (string, error) {
	return d.GetAttributeString("documentURI")
}

func (d Document) Embeds() (htmlcollection.HtmlCollection, error) {

	return d.getAttributeHTMLCollection("embeds")
}

func (d Document) FirstElementChild() (element.Element, error) {
	return d.getAttributeElement("firstElementChild")
}

func (d Document) Fonts() {
	//TO IMPLEMENT
}

func (d Document) Forms() (htmlcollection.HtmlCollection, error) {
	return d.getAttributeHTMLCollection("forms")
}

func (d Document) FullscreenElement() (element.Element, error) {
	return d.getAttributeElement("fullscreenElement")
}

func (d Document) Head() (element.Element, error) {
	return d.getAttributeElement("head")
}

func (d Document) Hidden() (bool, error) {

	return d.GetAttributeBool("hidden")
}

func (d Document) Images() (htmlcollection.HtmlCollection, error) {

	return d.getAttributeHTMLCollection("images")
}
func (d Document) Implementation() {
	//TO IMPLEMENT
}

func (d Document) LastElementChild() (element.Element, error) {
	return d.getAttributeElement("lastElementChild")
}

func (d Document) Links() (htmlcollection.HtmlCollection, error) {
	return d.getAttributeHTMLCollection("links")
}

func (d Document) PictureInPictureElement() (element.Element, error) {
	return d.getAttributeElement("pictureInPictureElement")
}

func (d Document) PictureInPictureEnabled() (bool, error) {
	return d.GetAttributeBool("pictureInPictureEnabled")
}

func (d Document) Plugins() (htmlcollection.HtmlCollection, error) {
	return d.getAttributeHTMLCollection("plugins")
}

func (d Document) PointerLockElement() (element.Element, error) {
	return d.getAttributeElement("pointerLockElement")
}

func (d Document) Scripts() (htmlcollection.HtmlCollection, error) {

	return d.getAttributeHTMLCollection("scripts")
}

func (d Document) ScrollingElement() (element.Element, error) {
	return d.getAttributeElement("scrollingElement")
}

func (d Document) VisibilityState() (string, error) {

	return d.GetAttributeString("visibilityState")
}

func (d Document) Domain() (string, error) {

	return d.GetAttributeString("domain")
}

func (d Document) LastModified() (string, error) {

	return d.GetAttributeString("lastModified")
}

func (d Document) SetDomain(domain string) error {

	return d.SetAttributeString("domain", domain)
}

func (d Document) ReadyState() (string, error) {

	return d.GetAttributeString("readyState")

}

func (d Document) Referrer() (string, error) {

	return d.GetAttributeString("referrer")
}

func (d Document) Title() (string, error) {

	return d.GetAttributeString("title")
}

func (d Document) SetTitle(title string) error {

	return d.SetAttributeString("title", title)
}

func (d Document) URL() (string, error) {

	return d.GetAttributeString("URL")

}

func (d Document) Cookie() (string, error) {
	return d.GetAttributeString("cookie")
}

func (d Document) SetCookie(cookie string) error {
	return d.SetAttributeString("cookie", cookie)
}
