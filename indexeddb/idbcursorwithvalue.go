package indexeddb

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

//IDBCursorWithValue struct
type IDBCursorWithValue struct {
	IDBCursor
}

type IDBCursorWithValueFrom interface {
	IDBCursorWithValue_() IDBCursorWithValue
}

func (i IDBCursorWithValue) IDBCursorWithValue_() IDBCursorWithValue {
	return i
}

var singletonIDBCursorWithValue sync.Once

var idbcursorinterfacewithvalue js.Value

func IDBCursorWithValueGetInterface() js.Value {

	singletonIDBCursorWithValue.Do(func() {

		var err error
		if idbcursorinterfacewithvalue, err = baseobject.Get(js.Global(), "IDBCursorWithValue"); err != nil {
			idbcursorinterfacewithvalue = js.Undefined()
		}

		baseobject.Register(idbcursorinterfacewithvalue, func(v js.Value) (interface{}, error) {
			return IDBCursorWithValueNewFromJSObject(v)
		})
	})

	return idbcursorinterfacewithvalue
}

func IDBCursorWithValueNewFromJSObject(obj js.Value) (IDBCursorWithValue, error) {
	var i IDBCursorWithValue
	var err error
	if ai := IDBCursorWithValueGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			if obj.InstanceOf(ai) {
				i.BaseObject = i.SetObject(obj)
			} else {
				err = ErrNotAnIDBCursor
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}
