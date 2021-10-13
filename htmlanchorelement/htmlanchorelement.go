package htmlanchorelement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/domtokenlist"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmlanchorlementinterface js.Value

//HtmlAnchorElement struct
type HtmlAnchorElement struct {
	htmlelement.HtmlElement
}

type HtmlAnchorElementFrom interface {
	HtmlAnchorElement_() HtmlAnchorElement
}

func (h HtmlAnchorElement) HtmlAnchorElement_() HtmlAnchorElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlanchorlementinterface, err = js.Global().GetWithErr("HTMLAnchorElement"); err != nil {
			htmlanchorlementinterface = js.Undefined()
		}
		baseobject.Register(htmlanchorlementinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return htmlanchorlementinterface
}

func New(d document.Document) (HtmlAnchorElement, error) {
	var err error

	var h HtmlAnchorElement
	var e element.Element

	if e, err = d.CreateElement("a"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlAnchorElement, error) {
	var h HtmlAnchorElement
	var err error

	if hci := GetInterface(); !hci.IsUndefined() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlAnchorElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlAnchorElement, error) {
	var h HtmlAnchorElement

	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlAnchorElement
}

func (h HtmlAnchorElement) Download() (string, error) {
	return h.GetAttributeString("download")
}

func (h HtmlAnchorElement) SetDownload(value string) error {
	return h.SetAttributeString("download", value)
}

func (h HtmlAnchorElement) Hash() (string, error) {
	return h.GetAttributeString("hash")
}

func (h HtmlAnchorElement) SetHash(value string) error {
	return h.SetAttributeString("hash", value)
}

func (h HtmlAnchorElement) Host() (string, error) {
	return h.GetAttributeString("host")
}

func (h HtmlAnchorElement) SetHost(value string) error {
	return h.SetAttributeString("host", value)
}

func (h HtmlAnchorElement) Hostname() (string, error) {
	return h.GetAttributeString("hostname")
}

func (h HtmlAnchorElement) SetHostname(value string) error {
	return h.SetAttributeString("hostname", value)
}

func (h HtmlAnchorElement) Href() (string, error) {
	return h.GetAttributeString("href")
}

func (h HtmlAnchorElement) SetHref(value string) error {
	return h.SetAttributeString("href", value)
}

func (h HtmlAnchorElement) Hreflang() (string, error) {
	return h.GetAttributeString("hreflang")
}

func (h HtmlAnchorElement) SetHreflang(value string) error {
	return h.SetAttributeString("hreflang", value)
}

func (h HtmlAnchorElement) Origin() (string, error) {
	return h.GetAttributeString("origin")
}

func (h HtmlAnchorElement) Password() (string, error) {
	return h.GetAttributeString("password")
}

func (h HtmlAnchorElement) SetPassword(value string) error {
	return h.SetAttributeString("password", value)
}

func (h HtmlAnchorElement) Pathname() (string, error) {
	return h.GetAttributeString("pathname")
}

func (h HtmlAnchorElement) SetPathname(value string) error {
	return h.SetAttributeString("pathname", value)
}

func (h HtmlAnchorElement) Port() (string, error) {
	return h.GetAttributeString("port")
}

func (h HtmlAnchorElement) SetPort(value string) error {
	return h.SetAttributeString("port", value)
}

func (h HtmlAnchorElement) Protocol() (string, error) {
	return h.GetAttributeString("protocol")
}

func (h HtmlAnchorElement) SetProtocol(value string) error {
	return h.SetAttributeString("protocol", value)
}

func (h HtmlAnchorElement) ReferrerPolicy() (string, error) {
	return h.GetAttributeString("referrerPolicy")
}

func (h HtmlAnchorElement) SetReferrerPolicy(value string) error {
	return h.SetAttributeString("referrerPolicy", value)
}

func (h HtmlAnchorElement) Rel() (string, error) {
	return h.GetAttributeString("rel")
}

func (h HtmlAnchorElement) SetRel(value string) error {
	return h.SetAttributeString("rel", value)
}

func (h HtmlAnchorElement) RelList() (domtokenlist.DOMTokenList, error) {
	var err error
	var obj js.Value
	var dlist domtokenlist.DOMTokenList

	if obj, err = h.JSObject().GetWithErr("relList"); err == nil {

		dlist, err = domtokenlist.NewFromJSObject(obj)
	}

	return dlist, err
}

func (h HtmlAnchorElement) Search() (string, error) {
	return h.GetAttributeString("search")
}

func (h HtmlAnchorElement) SetSearch(value string) error {
	return h.SetAttributeString("search", value)
}

func (h HtmlAnchorElement) TabIndex() (int, error) {
	return h.GetAttributeInt("tabIndex")
}

func (h HtmlAnchorElement) SetIndex(value int) error {
	return h.SetAttributeInt("tabIndex", value)
}

func (h HtmlAnchorElement) Target() (string, error) {
	return h.GetAttributeString("target")
}

func (h HtmlAnchorElement) SetTarget(value string) error {
	return h.SetAttributeString("target", value)
}

func (h HtmlAnchorElement) Text() (string, error) {
	return h.GetAttributeString("text")
}

func (h HtmlAnchorElement) SetText(value string) error {
	return h.SetAttributeString("text", value)
}

func (h HtmlAnchorElement) Type() (string, error) {
	return h.GetAttributeString("type")
}

func (h HtmlAnchorElement) SetType(value string) error {
	return h.SetAttributeString("type", value)
}

func (h HtmlAnchorElement) Username() (string, error) {
	return h.GetAttributeString("username")
}

func (h HtmlAnchorElement) SetUsername(value string) error {
	return h.SetAttributeString("username", value)
}
