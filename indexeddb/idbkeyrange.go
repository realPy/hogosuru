package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBIndex

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonIDBKeyRange sync.Once

var idbkeyrangeinterface js.Value

//GetIDBIndexInterface get teh JS interface of broadcast channel
func GetIDBKeyRangeInterface() js.Value {

	singletonIDBIndex.Do(func() {
		var err error
		if idbkeyrangeinterface, err = js.Global().GetWithErr("IDBKeyRange"); err != nil {
			idbkeyrangeinterface = js.Null()
		}
		baseobject.Register(idbkeyrangeinterface, func(v js.Value) (interface{}, error) {
			return IDBDKeyRangeNewFromJSObject(v)
		})
	})
	return idbkeyrangeinterface
}

//IDBKeyRange struct
type IDBKeyRange struct {
	baseobject.BaseObject
}

func IDBDKeyRangeNewFromJSObject(obj js.Value) (IDBKeyRange, error) {
	var i IDBKeyRange
	var err error
	if ai := GetIDBKeyRangeInterface(); !ai.IsNull() {
		if obj.InstanceOf(ai) {
			i.BaseObject = i.SetObject(obj)
		} else {
			err = ErrNotAnIDBKeyRange
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func newKeyRange(method string, values ...interface{}) (IDBKeyRange, error) {
	var i IDBKeyRange
	var err error

	if ii := GetIDBKeyRangeInterface(); !ii.IsNull() {
		i.BaseObject = i.SetObject(ii.New(values...))
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBKeyRange) getAttributeBool(attribute string) (bool, error) {

	var err error
	var obj js.Value
	var ret bool

	if obj, err = i.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return ret, err
}

func Bound(values ...interface{}) (IDBKeyRange, error) {
	return newKeyRange("bound", values...)
}

func LowerBound(values ...interface{}) (IDBKeyRange, error) {
	return newKeyRange("lowerBound", values...)
}

func UpperBound(values ...interface{}) (IDBKeyRange, error) {
	return newKeyRange("upperBound", values...)
}

func Only(value interface{}) (IDBKeyRange, error) {
	return newKeyRange("only", value)
}

func (i IDBKeyRange) Includes(value interface{}) (bool, error) {
	var obj js.Value
	var err error
	var ret bool

	if obj, err = i.JSObject().CallWithErr("includes", value); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return ret, err
}

func (i IDBKeyRange) lowerOpen() (bool, error) {
	return i.getAttributeBool("lowerOpen")
}

func (i IDBKeyRange) upperOpen() (bool, error) {
	return i.getAttributeBool("upperOpen")
}

func (i IDBKeyRange) Lower() (interface{}, error) {
	return i.JSObject().GetWithErr("lower")
}

func (i IDBKeyRange) Upper() (interface{}, error) {
	return i.JSObject().GetWithErr("upper")
}
