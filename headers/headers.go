package headers

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/iterator"
)

// https://developer.mozilla.org/en-US/docs/Web/API/Headers

var singleton sync.Once

var headersinterface js.Value

//History struct
type Headers struct {
	baseobject.BaseObject
}

type HeadersFrom interface {
	Headers() Headers
}

func (h Headers) Headers() Headers {
	return h
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if headersinterface, err = js.Global().GetWithErr("Headers"); err != nil {
			headersinterface = js.Null()
		}

	})

	baseobject.Register(headersinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return headersinterface
}

func NewFromJSObject(obj js.Value) (Headers, error) {
	var h Headers

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHeaders
}

func (h Headers) Append(name, value string) error {
	var err error
	_, err = h.JSObject().CallWithErr("append", js.ValueOf(name), js.ValueOf(value))
	return err
}

func (h Headers) Delete(name string) error {
	var err error
	_, err = h.JSObject().CallWithErr("delete", js.ValueOf(name))
	return err
}

func (h Headers) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = h.JSObject().CallWithErr("entries"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (h Headers) Get(name string) (string, error) {

	var err error
	var obj js.Value
	var result string

	if obj, err = h.JSObject().CallWithErr("get", js.ValueOf(name)); err == nil {

		if obj.Type() == js.TypeString {
			result = obj.String()
		} else {
			err = baseobject.ErrObjectNotBool
		}

	}

	return result, err
}

func (h Headers) Has(name string) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj, err = h.JSObject().CallWithErr("has", js.ValueOf(name)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (h Headers) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = h.JSObject().CallWithErr("keys"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (h Headers) Set(name, value string) error {
	var err error
	_, err = h.JSObject().CallWithErr("set", js.ValueOf(name), js.ValueOf(value))
	return err
}

func (h Headers) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = h.JSObject().CallWithErr("values"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}
