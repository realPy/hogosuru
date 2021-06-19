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
