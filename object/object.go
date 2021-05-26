package object

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var objectinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var objectinstance JSInterface
		var err error
		if objectinstance.objectInterface, err = js.Global().GetWithErr("Object"); err == nil {
			objectinterface = &objectinstance
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
	if ai := GetJSInterface(); ai != nil {
		if obj.InstanceOf(ai.objectInterface) {
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

	if ai := GetJSInterface(); ai != nil {
		if obj, err = ai.objectInterface.CallWithErr("keys", o.JSObject()); err == nil {
			newArr, err = array.NewFromJSObject(obj)

		}

	}

	return newArr, err
}

func Values(o Object) (array.Array, error) {

	var err error
	var obj js.Value
	var newArr array.Array

	if ai := GetJSInterface(); ai != nil {
		if obj, err = ai.objectInterface.CallWithErr("values", o.JSObject()); err == nil {
			newArr, err = array.NewFromJSObject(obj)

		}

	}

	return newArr, err
}
