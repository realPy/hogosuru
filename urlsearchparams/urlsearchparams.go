package urlsearchparams

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
	"github.com/realPy/hogosuru/iterator"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var urlsearchparamsinterface js.Value

//URLSearchParams struct
type URLSearchParams struct {
	baseobject.BaseObject
}

type URLSearchParamsFrom interface {
	URLSearchParams_() URLSearchParams
}

func (u URLSearchParams) URLSearchParams_() URLSearchParams {
	return u
}

//GetInterface get the JS interface URLSearchParams
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if urlsearchparamsinterface, err = baseobject.Get(js.Global(), "URLSearchParams"); err != nil {
			urlsearchparamsinterface = js.Undefined()
		}
		baseobject.Register(urlsearchparamsinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return urlsearchparamsinterface
}

func New(s ...string) (URLSearchParams, error) {

	var u URLSearchParams
	var err error
	var obj js.Value
	var arrayJS []interface{}

	if len(s) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(s[0]))
	}
	if hci := GetInterface(); !hci.IsUndefined() {

		if obj, err = baseobject.New(hci, arrayJS...); err == nil {
			u.BaseObject = u.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return u, err
}

func NewFromJSObject(obj js.Value) (URLSearchParams, error) {
	var u URLSearchParams
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				u.BaseObject = u.SetObject(obj)

			} else {
				err = ErrNotAnURLSearchParams
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return u, err
}

func (u URLSearchParams) Append(name, value string) error {
	var err error
	_, err = u.Call("append", js.ValueOf(name), js.ValueOf(value))
	return err
}

func (u URLSearchParams) Delete(name string) error {
	var err error
	_, err = u.Call("delete", js.ValueOf(name))
	return err
}

func (u URLSearchParams) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = u.Call("entries"); err == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (u URLSearchParams) Get(name string) (string, error) {

	var err error
	var obj js.Value
	var result string

	if obj, err = u.Call("get", js.ValueOf(name)); err == nil {

		if obj.Type() == js.TypeString {
			result = obj.String()
		} else {
			err = baseobject.ErrUndefinedValue
		}

	}

	return result, err
}

func (u URLSearchParams) Has(name string) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj, err = u.Call("has", js.ValueOf(name)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (u URLSearchParams) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = u.Call("keys"); err == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (u URLSearchParams) Set(name, value string) error {
	var err error
	_, err = u.Call("set", js.ValueOf(name), js.ValueOf(value))
	return err
}

func (u URLSearchParams) Sort() error {
	var err error
	_, err = u.Call("sort")
	return err
}

func (u URLSearchParams) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = u.Call("values"); err == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}
