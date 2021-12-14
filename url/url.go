package url

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

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

func (u URL) SetOrigin(origin string) error {

	return u.SetAttributeString("origin", origin)
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

/*
todo
searchParams
*/

func (u URL) CreateObjectURL(object interface{}) (string, error) {
	var err error
	var obj js.Value
	var ret string
	if objGo, ok := object.(baseobject.ObjectFrom); ok {

		if obj, err = u.Call("createObjectURL", objGo.JSObject()); err == nil {
			if obj.Type() == js.TypeString {
				ret = obj.String()
			} else {
				err = baseobject.ErrObjectNotString
			}

		}

	}

	return ret, err
}

func (u URL) RevokeObjectURL(objecturl string) error {
	var err error
	_, err = u.Call("revokeObjectURL", objecturl)

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
