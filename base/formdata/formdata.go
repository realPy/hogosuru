package formdata

// https://developer.mozilla.org/fr/docs/Web/API/FormData

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/htmlformelement"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/iterator"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var formadatainterface js.Value

// FormData struct
type FormData struct {
	baseobject.BaseObject
}

type FormDataFrom interface {
	FormData_() FormData
}

func (f FormData) FormData_() FormData {
	return f
}

// GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if formadatainterface, err = baseobject.Get(js.Global(), "FormData"); err != nil {
			formadatainterface = js.Undefined()
		}

	})

	return formadatainterface
}

func NewFromJSObject(obj js.Value) (FormData, error) {
	var f FormData
	var err error
	if fi := GetInterface(); !fi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(fi) {
				f.BaseObject = f.SetObject(obj)

			} else {
				err = ErrNotAFormData
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return f, err
}

func New(f ...htmlformelement.HtmlFormElement) (FormData, error) {

	var formdata FormData
	var obj js.Value
	var err error
	var opt []interface{}

	if fci := GetInterface(); !fci.IsUndefined() {
		if len(f) > 0 {
			opt = append(opt, f[0].JSObject())
		}
		if obj, err = baseobject.New(fci, opt...); err == nil {
			formdata.BaseObject = formdata.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return formdata, err
}

func (f FormData) Append(key string, value interface{}) error {
	var err error
	_, err = f.Call("append", js.ValueOf(key), baseobject.GetJsValueOf(value))
	return err
}

func (f FormData) Delete(key string) error {
	var err error

	_, err = f.Call("delete", js.ValueOf(key))
	return err
}

func (f FormData) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = f.Call("entries"); err == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (f FormData) Get(key string) (interface{}, error) {

	var err error
	var obj js.Value
	var result interface{}

	if obj, err = f.Call("get", js.ValueOf(key)); err == nil {
		if obj.IsNull() {
			err = ErrNotAFormValueNotFound
		} else {
			result, err = baseobject.GoValue(obj)
		}

	}
	return result, err
}

func (f FormData) Has(key string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = f.Call("has", js.ValueOf(key)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (f FormData) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = f.Call("keys"); err == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (f FormData) Set(key string, value interface{}) error {
	var err error
	_, err = f.Call("set", js.ValueOf(key), baseobject.GetJsValueOf(value))
	return err
}

func (f FormData) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = f.Call("values"); err == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}
