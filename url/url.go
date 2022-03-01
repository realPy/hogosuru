package url

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
	"github.com/realPy/hogosuru/urlsearchparams"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var urlinterface js.Value

//GetInterface get the JS interface of URL
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if urlinterface, err = baseobject.Get(js.Global(), "URL"); err != nil {
			urlinterface = js.Undefined()
		}
		baseobject.Register(urlinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return urlinterface
}

type URL struct {
	baseobject.BaseObject
}

type URLFrom interface {
	URL_() URL
}

func (u URL) Location_() URL {
	return u
}

func New(value string) (URL, error) {
	var u URL
	var err error
	var obj js.Value

	if ui := GetInterface(); !ui.IsUndefined() {

		if obj, err = baseobject.New(ui, js.ValueOf(value)); err == nil {
			u.BaseObject = u.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}

	return u, err
}

func NewFromJSObject(obj js.Value) (URL, error) {
	var u URL

	if ui := GetInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil

		}
	}

	return u, ErrNotImplemented
}

func (u URL) Hash() (string, error) {

	return u.GetAttributeString("hash")
}

func (u URL) SetHash(hash string) error {

	return u.SetAttributeString("hash", hash)
}

func (u URL) Host() (string, error) {

	return u.GetAttributeString("host")
}

func (u URL) SetHost(host string) error {

	return u.SetAttributeString("host", host)
}

func (u URL) Hostname() (string, error) {

	return u.GetAttributeString("hostname")
}

func (u URL) SetHostname(hostname string) error {

	return u.SetAttributeString("hostname", hostname)
}

func (u URL) Href() (string, error) {

	return u.GetAttributeString("href")
}

func (u URL) SetHref(href string) error {

	return u.SetAttributeString("href", href)
}

func (u URL) Origin() (string, error) {

	return u.GetAttributeString("origin")
}

func (u URL) Pathname() (string, error) {

	return u.GetAttributeString("pathname")
}

func (u URL) SetPathname(pathname string) error {

	return u.SetAttributeString("pathname", pathname)
}

func (u URL) Port() (string, error) {

	return u.GetAttributeString("port")
}

func (u URL) SetPort(port string) error {

	return u.SetAttributeString("port", port)
}

func (u URL) Protocol() (string, error) {

	return u.GetAttributeString("protocol")
}

func (u URL) SetProtocol(protocol string) error {

	return u.SetAttributeString("protocol", protocol)
}

func (u URL) Username() (string, error) {

	return u.GetAttributeString("username")
}

func (u URL) SetUsername(username string) error {

	return u.SetAttributeString("username", username)
}

func (u URL) Password() (string, error) {

	return u.GetAttributeString("password")
}

func (u URL) SetPassword(password string) error {

	return u.SetAttributeString("password", password)
}

func (u URL) Search() (string, error) {

	return u.GetAttributeString("search")
}

func (u URL) SetSearch(search string) error {

	return u.SetAttributeString("search", search)
}

func (u URL) SearchParams() (urlsearchparams.URLSearchParams, error) {
	var err error
	var obj js.Value
	var params urlsearchparams.URLSearchParams

	if obj, err = u.Get("searchParams"); err == nil {

		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrNotAnObject

		} else {

			params, err = urlsearchparams.NewFromJSObject(obj)
		}
	}

	return params, err
}

func CreateObjectURL(object interface{}) (string, error) {
	var err error
	var obj js.Value
	var ret string
	if ui := GetInterface(); !ui.IsUndefined() {

		if objGo, ok := object.(baseobject.ObjectFrom); ok {

			if obj, err = baseobject.Call(ui, "createObjectURL", objGo.JSObject()); err == nil {
				if obj.Type() == js.TypeString {
					ret = obj.String()
				} else {
					err = baseobject.ErrObjectNotString
				}

			}

		}
	} else {
		err = ErrNotImplemented
	}

	return ret, err
}

func RevokeObjectURL(objecturl string) error {
	var err error

	if ui := GetInterface(); !ui.IsUndefined() {
		_, err = baseobject.Call(ui, "revokeObjectURL", objecturl)
	} else {
		err = ErrNotImplemented
	}

	return err
}

func (u URL) ToJSON() (string, error) {

	var err error
	var obj js.Value
	var ret string

	if obj, err = u.Call("toJSON"); err == nil {
		if obj.Type() == js.TypeString {
			ret = obj.String()
		} else {
			err = baseobject.ErrObjectNotString
		}

	}
	return ret, err
}
