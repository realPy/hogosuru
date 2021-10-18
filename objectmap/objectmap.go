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
		if mapinterface, err = js.Global().GetWithErr("Map"); err != nil {
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
		if obj.InstanceOf(ai) {
			o.BaseObject = o.SetObject(obj)

		} else {
			err = ErrNotAMap
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

	var arrayJS []interface{}

	for _, value := range values {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

	}

	if omi := GetInterface(); !omi.IsUndefined() {

		o.BaseObject = o.SetObject(omi.New(arrayJS...))

	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func (o ObjectMap) Clear() error {
	var err error
	_, err = o.JSObject().CallWithErr("clear")
	return err
}

func (o ObjectMap) Delete(key interface{}) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	var globalKeyObj interface{}

	if objGo, ok := key.(baseobject.ObjectFrom); ok {
		globalKeyObj = objGo.JSObject()
	} else {
		globalKeyObj = js.ValueOf(key)
	}

	if obj, err = o.JSObject().CallWithErr("delete", globalKeyObj); err == nil {
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

	if obj, err = o.JSObject().CallWithErr("entries"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) ForEach(f func(value, index interface{})) error {
	var err error

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(baseobject.GoValue(args[0]), baseobject.GoValue(args[1]))
		return nil
	})

	_, err = o.JSObject().CallWithErr("forEach", jsfunc)
	jsfunc.Release()
	return err
}

func (o ObjectMap) Get(key interface{}) (interface{}, error) {

	var err error
	var obj js.Value
	var result interface{}

	var globalKeyObj interface{}

	if objGo, ok := key.(baseobject.ObjectFrom); ok {
		globalKeyObj = objGo.JSObject()
	} else {
		globalKeyObj = js.ValueOf(key)
	}

	if obj, err = o.JSObject().CallWithErr("get", globalKeyObj); err == nil {
		result = baseobject.GoValue(obj)
	}
	return result, err
}

func (o ObjectMap) Has(key interface{}) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	var globalKeyObj interface{}

	if objGo, ok := key.(baseobject.ObjectFrom); ok {
		globalKeyObj = objGo.JSObject()
	} else {
		globalKeyObj = js.ValueOf(key)
	}

	if obj, err = o.JSObject().CallWithErr("has", globalKeyObj); err == nil {
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

	if obj, err = o.JSObject().CallWithErr("keys"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) Set(key interface{}, value interface{}) error {
	var err error
	var globalKeyObj interface{}
	var globalValueObj interface{}

	if objGo, ok := key.(baseobject.ObjectFrom); ok {
		globalKeyObj = objGo.JSObject()
	} else {
		globalKeyObj = js.ValueOf(key)
	}

	if objGo, ok := value.(baseobject.ObjectFrom); ok {
		globalValueObj = objGo.JSObject()
	} else {
		globalValueObj = js.ValueOf(value)
	}

	_, err = o.JSObject().CallWithErr("set", globalKeyObj, globalValueObj)
	return err
}

func (o ObjectMap) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = o.JSObject().CallWithErr("values"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) Size() (int, error) {
	return o.GetAttributeInt("size")
}
