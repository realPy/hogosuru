package arraybuffer

//partial implemented (herited from function)
// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var arraybufferinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if arraybufferinterface, err = baseobject.Get(js.Global(), "ArrayBuffer"); err != nil {
			arraybufferinterface = js.Undefined()
		}
		baseobject.Register(arraybufferinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return arraybufferinterface
}

//ArrayBuffer struct
type ArrayBuffer struct {
	baseobject.BaseObject
}

type ArrayBufferFrom interface {
	ArrayBuffer_() ArrayBuffer
}

func (a ArrayBuffer) ArrayBuffer_() ArrayBuffer {
	return a
}

func New(size int) (ArrayBuffer, error) {

	var a ArrayBuffer

	if ai := GetInterface(); !ai.IsUndefined() {

		a.BaseObject = a.SetObject(ai.New(js.ValueOf(size)))
		return a, nil
	}

	return a, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (ArrayBuffer, error) {
	var a ArrayBuffer

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.InstanceOf(ai) {
			a.BaseObject = a.SetObject(obj)
			return a, nil
		}
	}

	return a, ErrNotAnArrayBuffer
}

func (a ArrayBuffer) ByteLength() (int64, error) {

	return a.GetAttributeInt64("byteLength")
}

func (a ArrayBuffer) Slice(begin int, end ...int) (ArrayBuffer, error) {

	var optjs []interface{}
	var err error
	var obj js.Value
	var ret ArrayBuffer

	optjs = append(optjs, js.ValueOf(begin))
	if len(end) > 0 {
		optjs = append(optjs, js.ValueOf(end[0]))
	}

	if obj, err = a.JSObject().CallWithErr("slice", optjs...); err == nil {

		ret, err = NewFromJSObject(obj)
	}
	return ret, err
}

func IsView(i interface{}) (bool, error) {

	var objjs interface{}
	var ret bool
	var err error
	var obj js.Value

	if objGo, ok := i.(baseobject.ObjectFrom); ok {
		objjs = objGo.JSObject()
	} else {
		objjs = js.ValueOf(i)
	}

	if ai := GetInterface(); !ai.IsUndefined() {

		if obj, err = ai.CallWithErr("isView", objjs); err == nil {

			if obj.Type() == js.TypeBoolean {
				ret = obj.Bool()
			} else {
				err = baseobject.ErrObjectNotBool
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return ret, err
}
