package object

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var objectinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if objectinterface, err = js.Global().GetWithErr("Object"); err != nil {
			objectinterface = js.Null()
		}

	})

	return objectinterface
}

//Object struct
type Object struct {
	baseobject.BaseObject
}

func NewFromJSObject(obj js.Value) (Object, error) {
	var o Object
	var err error
	if ai := GetInterface(); !ai.IsNull() {
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

func Keys(o Object) (array.Array, error) {

	var err error
	var obj js.Value
	var newArr array.Array

	if ai := GetInterface(); !ai.IsNull() {
		if obj, err = ai.CallWithErr("keys", o.JSObject()); err == nil {
			newArr, err = array.NewFromJSObject(obj)

		}

	}

	return newArr, err
}

func Values(o Object) (array.Array, error) {

	var err error
	var obj js.Value
	var newArr array.Array

	if ai := GetInterface(); !ai.IsNull() {
		if obj, err = ai.CallWithErr("values", o.JSObject()); err == nil {
			newArr, err = array.NewFromJSObject(obj)

		}

	}

	return newArr, err
}
