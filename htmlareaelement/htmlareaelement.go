package htmlareaelement

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

var htmlareaelementinterface js.Value

//HtmlAreaElement struct
type HtmlAreaElement struct {
	htmlelement.HtmlElement
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlareaelementinterface, err = js.Global().GetWithErr("HTMLAreaElement"); err != nil {
			htmlareaelementinterface = js.Null()
		}

	})

	baseobject.Register(htmlareaelementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmlareaelementinterface
}

func New(d document.Document) (HtmlAreaElement, error) {
	var err error

	var h HtmlAreaElement
	var e element.Element

	if e, err = d.CreateElement("area"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlAreaElement, error) {
	var h HtmlAreaElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlAreaElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlAreaElement, error) {
	var h HtmlAreaElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlAreaElement
}

func (h HtmlAreaElement) Alt() (string, error) {
	return h.GetAttributeString("alt")
}

func (h HtmlAreaElement) SetAlt(value string) error {
	return h.SetAttributeString("alt", value)
}

func (h HtmlAreaElement) Coords() (string, error) {
	return h.GetAttributeString("coords")
}

func (h HtmlAreaElement) SetCoords(value string) error {
	return h.SetAttributeString("coords", value)
}

func (h HtmlAreaElement) Download() (string, error) {
	return h.GetAttributeString("download")
}

func (h HtmlAreaElement) SetDownload(value string) error {
	return h.SetAttributeString("download", value)
}

func (h HtmlAreaElement) Hash() (string, error) {
	return h.GetAttributeString("hash")
}

func (h HtmlAreaElement) SetHash(value string) error {
	return h.SetAttributeString("hash", value)
}

func (h HtmlAreaElement) Host() (string, error) {
	return h.GetAttributeString("host")
}

func (h HtmlAreaElement) SetHost(value string) error {
	return h.SetAttributeString("host", value)
}

func (h HtmlAreaElement) Hostname() (string, error) {
	return h.GetAttributeString("hostname")
}

func (h HtmlAreaElement) SetHostname(value string) error {
	return h.SetAttributeString("hostname", value)
}

func (h HtmlAreaElement) Href() (string, error) {
	return h.GetAttributeString("href")
}

func (h HtmlAreaElement) SetHref(value string) error {
	return h.SetAttributeString("href", value)
}

func (h HtmlAreaElement) Origin() (string, error) {
	return h.GetAttributeString("origin")
}

func (h HtmlAreaElement) Password() (string, error) {
	return h.GetAttributeString("password")
}

func (h HtmlAreaElement) SetPassword(value string) error {
	return h.SetAttributeString("password", value)
}

func (h HtmlAreaElement) Pathname() (string, error) {
	return h.GetAttributeString("pathname")
}

func (h HtmlAreaElement) SetPathname(value string) error {
	return h.SetAttributeString("pathname", value)
}

func (h HtmlAreaElement) Port() (string, error) {
	return h.GetAttributeString("port")
}

func (h HtmlAreaElement) SetPort(value string) error {
	return h.SetAttributeString("port", value)
}

func (h HtmlAreaElement) Protocol() (string, error) {
	return h.GetAttributeString("protocol")
}

func (h HtmlAreaElement) SetProtocol(value string) error {
	return h.SetAttributeString("protocol", value)
}

func (h HtmlAreaElement) ReferrerPolicy() (string, error) {
	return h.GetAttributeString("referrerPolicy")
}

func (h HtmlAreaElement) SetReferrerPolicy(value string) error {
	return h.SetAttributeString("referrerPolicy", value)
}

func (h HtmlAreaElement) Rel() (string, error) {
	return h.GetAttributeString("rel")
}

func (h HtmlAreaElement) SetRel(value string) error {
	return h.SetAttributeString("rel", value)
}

func (h HtmlAreaElement) RelList() (domtokenlist.DOMTokenList, error) {
	var err error
	var obj js.Value
	var dlist domtokenlist.DOMTokenList

	if obj, err = h.JSObject().GetWithErr("relList"); err == nil {

		dlist, err = domtokenlist.NewFromJSObject(obj)
	}

	return dlist, err
}

func (h HtmlAreaElement) Search() (string, error) {
	return h.GetAttributeString("search")
}

func (h HtmlAreaElement) SetSearch(value string) error {
	return h.SetAttributeString("search", value)
}

func (h HtmlAreaElement) TabIndex() (int, error) {
	return h.GetAttributeInt("tabIndex")
}

func (h HtmlAreaElement) SetIndex(value int) error {
	return h.SetAttributeInt("tabIndex", value)
}

func (h HtmlAreaElement) Target() (string, error) {
	return h.GetAttributeString("target")
}

func (h HtmlAreaElement) SetTarget(value string) error {
	return h.SetAttributeString("target", value)
}

func (h HtmlAreaElement) Shape() (string, error) {
	return h.GetAttributeString("shape")
}

func (h HtmlAreaElement) SetShape(value string) error {
	return h.SetAttributeString("shape", value)
}

func (h HtmlAreaElement) Username() (string, error) {
	return h.GetAttributeString("username")
}

func (h HtmlAreaElement) SetUsername(value string) error {
	return h.SetAttributeString("username", value)
}
