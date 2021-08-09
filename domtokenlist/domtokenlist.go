package domtokenlist

//

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/iterator"
)

var singleton sync.Once

var domtokenlistinterface js.Value

//DOMRectLists struct
type DOMTokenList struct {
	baseobject.BaseObject
}

type DOMTokenListFrom interface {
	DOMTokenList() DOMTokenList
}

func (d DOMTokenList) DOMTokenList() DOMTokenList {
	return d
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if domtokenlistinterface, err = js.Global().GetWithErr("DOMTokenList"); err != nil {
			domtokenlistinterface = js.Null()
		}
	})
	baseobject.Register(domtokenlistinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return domtokenlistinterface
}

func NewFromJSObject(obj js.Value) (DOMTokenList, error) {
	var d DOMTokenList
	var err error
	if dli := GetInterface(); !dli.IsNull() {
		if obj.InstanceOf(dli) {
			d.BaseObject = d.SetObject(obj)

		} else {
			err = ErrNotAnDOMTokenList
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DOMTokenList) Item(index int) js.Value {
	var obj js.Value
	obj = d.JSObject().Index(index)
	return obj
}

func (d DOMTokenList) methodGetValue(method string, value string) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj, err = d.JSObject().CallWithErr(method, js.ValueOf(value)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (d DOMTokenList) Contains(search string) (bool, error) {

	return d.methodGetValue("contains", search)
}

func (d DOMTokenList) method(method string, tokens ...string) error {
	var err error
	var arrayJS []interface{}

	for _, token := range tokens {
		arrayJS = append(arrayJS, js.ValueOf(token))
	}

	_, err = d.JSObject().CallWithErr(method, arrayJS...)

	return err

}

func (d DOMTokenList) Add(tokens ...string) error {
	return d.method("add", tokens...)
}

func (d DOMTokenList) Remove(tokens ...string) error {
	return d.method("add", tokens...)
}

func (d DOMTokenList) Replace(oldtoken, newtoken string) error {
	return d.method("replace", oldtoken, newtoken)
}

func (d DOMTokenList) Toggle(token string, force ...bool) (bool, error) {
	var err error
	var arrayJS []interface{}
	var result bool
	var obj js.Value

	arrayJS = append(arrayJS, js.ValueOf(token))
	if len(force) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(force[0]))
	}

	if obj, err = d.JSObject().CallWithErr("toggle", arrayJS...); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err

}

func (d DOMTokenList) Supports(token string) (bool, error) {
	return d.methodGetValue("supports", token)
}

func (d DOMTokenList) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = d.JSObject().CallWithErr("entries"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (d DOMTokenList) ForEach(f func(string, string)) error {
	var err error

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0].String(), args[0].String())
		return nil
	})

	_, err = d.JSObject().CallWithErr("forEach", jsfunc)
	jsfunc.Release()
	return err
}

func (d DOMTokenList) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = d.JSObject().CallWithErr("keys"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (d DOMTokenList) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = d.JSObject().CallWithErr("values"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}
