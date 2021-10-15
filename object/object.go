package object

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/objectmap"
)

var singleton sync.Once

var objectinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if objectinterface, err = js.Global().GetWithErr("Object"); err != nil {
			objectinterface = js.Undefined()
		}
		baseobject.Register(objectinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return objectinterface
}

//Object struct
type Object struct {
	baseobject.BaseObject
}

type ObjectFrom interface {
	Object_() Object
}

func (o Object) Object_() Object {
	return o
}

func NewFromJSObject(obj js.Value) (Object, error) {
	var o Object
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.InstanceOf(ai) {
			o.BaseObject = o.SetObject(obj)

		} else {
			err = ErrNotAnObject
		}
	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func (o Object) Keys() (array.Array, error) {

	var err error
	var obj js.Value
	var newArr array.Array

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj, err = ai.CallWithErr("keys", o.JSObject()); err == nil {
			newArr, err = array.NewFromJSObject(obj)

		}

	}

	return newArr, err
}

func (o Object) Values() (array.Array, error) {

	var err error
	var obj js.Value
	var newArr array.Array

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj, err = ai.CallWithErr("values", o.JSObject()); err == nil {
			newArr, err = array.NewFromJSObject(obj)

		}

	}

	return newArr, err
}

func (o Object) Map() (objectmap.ObjectMap, error) {
	var err error
	var obj js.Value
	var newMap objectmap.ObjectMap

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj, err = ai.CallWithErr("entries", o.JSObject()); err == nil {
			newMap, err = objectmap.New(obj)

		}

	}
	return newMap, err
}
