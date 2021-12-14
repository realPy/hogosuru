package location

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var locationinterface js.Value

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if locationinterface, err = baseobject.Get(js.Global(), "Location"); err != nil {
			locationinterface = js.Undefined()
		}
		baseobject.Register(locationinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return locationinterface
}

type Location struct {
	baseobject.BaseObject
}

type LocationFrom interface {
	Location_() Location
}

func (l Location) Location_() Location {
	return l
}

func NewFromJSObject(obj js.Value) (Location, error) {
	var l Location

	if li := GetInterface(); !li.IsUndefined() {
		if obj.InstanceOf(li) {
			l.BaseObject = l.SetObject(obj)
			return l, nil

		}
	}

	return l, ErrNotImplemented
}

func (l Location) Hash() (string, error) {

	return l.GetAttributeString("hash")
}

func (l Location) Host() (string, error) {

	return l.GetAttributeString("host")
}

func (l Location) Hostname() (string, error) {

	return l.GetAttributeString("hostname")
}

func (l Location) Href() (string, error) {

	return l.GetAttributeString("href")
}

func (l Location) Origin() (string, error) {

	return l.GetAttributeString("origin")
}

func (l Location) Pathname() (string, error) {

	return l.GetAttributeString("pathname")
}

func (l Location) Port() (string, error) {

	return l.GetAttributeString("port")
}

func (l Location) Protocol() (string, error) {

	return l.GetAttributeString("protocol")
}

func (l Location) Search() (string, error) {

	return l.GetAttributeString("search")
}

func (l Location) Username() (string, error) {

	return l.GetAttributeString("username")
}

func (l Location) Password() (string, error) {

	return l.GetAttributeString("password")
}

func (l Location) Assign(url string) error {
	var err error
	_, err = l.Call("assign", js.ValueOf(url))
	return err
}

func (l Location) Reload(value bool) error {
	var err error
	_, err = l.Call("reload", js.ValueOf(value))
	return err
}

func (l Location) Replace(url string) error {
	var err error
	_, err = l.Call("replace", js.ValueOf(url))
	return err
}
