package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBIndex

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonIDBIndex sync.Once

var idbindexinterface js.Value

//GetIDBIndexInterface get teh JS interface of broadcast channel
func GetIDBIndexInterface() js.Value {

	singletonIDBIndex.Do(func() {
		var err error
		if idbindexinterface, err = js.Global().GetWithErr("IDBIndex"); err != nil {
			idbindexinterface = js.Null()
		}
	})
	return idbindexinterface
}

//IDBIndex struct
type IDBIndex struct {
	baseobject.BaseObject
}

func IDBDIndexNewFromJSObject(obj js.Value) (IDBIndex, error) {
	var i IDBIndex
	var err error
	if ai := GetIDBIndexInterface(); !ai.IsNull() {
		if obj.InstanceOf(ai) {
			i.BaseObject = i.SetObject(obj)
		} else {
			err = ErrNotAnIDBIndex
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}
