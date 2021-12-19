package objectmap

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/iterator"
)

var singleton sync.Once

var mapinterface js.Value

//GetInterface get the JS interface of object channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if mapinterface, err = baseobject.Get(js.Global(), "Map"); err != nil {
			mapinterface = js.Undefined()
		}
		baseobject.Register(mapinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return mapinterface
}

//ObjectMap
type ObjectMap struct {
	baseobject.BaseObject
}

type ObjectMapFrom interface {
	ObjectMap_() ObjectMap
}

func (o ObjectMap) ObjectMap_() ObjectMap {
	return o
}

func NewFromJSObject(obj js.Value) (ObjectMap, error) {
	var o ObjectMap
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				o.BaseObject = o.SetObject(obj)

			} else {
				err = ErrNotAMap
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func NewFromBaseObject(b baseobject.BaseObject) (ObjectMap, error) {

	return New(b)
}

func New(values ...interface{}) (ObjectMap, error) {

	var o ObjectMap
	var err error
	var obj js.Value
	var arrayJS []interface{}

	for _, value := range values {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}

	if omi := GetInterface(); !omi.IsUndefined() {

		if obj, err = baseobject.New(omi, arrayJS...); err == nil {
			o.BaseObject = o.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func (o ObjectMap) Clear() error {
	var err error
	_, err = o.Call("clear")
	return err
}

func (o ObjectMap) Delete(key interface{}) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = o.Call("delete", baseobject.GetJsValueOf(key)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (o ObjectMap) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = o.Call("entries"); err == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) ForEach(f func(value, index interface{})) error {
	var err error

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		f(baseobject.GoValue_(args[0]), baseobject.GoValue_(args[1]))
		return nil
	})

	_, err = o.Call("forEach", jsfunc)
	jsfunc.Release()
	return err
}

func (o ObjectMap) Get(key interface{}) (interface{}, error) {

	var err error
	var obj js.Value
	var result interface{}

	if obj, err = o.Call("get", baseobject.GetJsValueOf(key)); err == nil {
		result, err = baseobject.GoValue(obj)
	}
	return result, err
}

func (o ObjectMap) Has(key interface{}) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = o.Call("has", baseobject.GetJsValueOf(key)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (o ObjectMap) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = o.Call("keys"); err == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) Set(key interface{}, value interface{}) error {
	var err error
	_, err = o.Call("set", baseobject.GetJsValueOf(key), baseobject.GetJsValueOf(value))
	return err
}

func (o ObjectMap) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = o.Call("values"); err == nil {
		iter, err = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) Size() (int, error) {
	return o.GetAttributeInt("size")
}
