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

//GetInterface get the JS interface ArrayBuffer
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
	var ai, obj js.Value
	var err error
	if ai = GetInterface(); ai.IsUndefined() {
		return a, ErrNotImplemented
	}
	if obj, err = baseobject.New(ai, js.ValueOf(size)); err != nil {
		return a, err
	}
	a.BaseObject = a.SetObject(obj)
	return a, nil
}

func NewFromJSObject(obj js.Value) (ArrayBuffer, error) {
	var a ArrayBuffer
	var ai js.Value
	if ai = GetInterface(); ai.IsUndefined() {
		return a, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return a, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(ai) {
		return a, ErrNotAnArrayBuffer
	}
	a.BaseObject = a.SetObject(obj)
	return a, nil
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
	if obj, err = a.Call("slice", optjs...); err == nil {
		return NewFromJSObject(obj)
	}
	return ret, err
}

func IsView(i interface{}) (bool, error) {
	var err error
	var ai, obj js.Value
	if ai = GetInterface(); ai.IsUndefined() {
		return false, ErrNotImplemented
	}
	if obj, err = baseobject.Call(ai, "isView", baseobject.GetJsValueOf(i)); err != nil {
		return false, ErrNotImplemented
	}
	if obj.Type() != js.TypeBoolean {
		return false, baseobject.ErrObjectNotBool
	}
	return obj.Bool(), nil
}
