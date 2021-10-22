package indexeddb

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/eventtarget"
)

//IDBCursor struct
type IDBCursor struct {
	eventtarget.EventTarget
}

type IDBCursorFrom interface {
	IDBCursor_() IDBCursor
}

func (i IDBCursor) IDBCursor_() IDBCursor {
	return i
}

var singletonIDBCursor sync.Once

var idbcursorinterface js.Value

func IDBCursorGetInterface() js.Value {

	singletonIDBDatabase.Do(func() {

		var err error
		if idbcursorinterface, err = baseobject.Get(js.Global(), "IDBCursor"); err != nil {
			idbcursorinterface = js.Undefined()
		}
		baseobject.Register(idbcursorinterface, func(v js.Value) (interface{}, error) {
			return IDBDatabaseNewFromJSObject(v)
		})
	})

	return idbcursorinterface
}

func IDBCursorNewFromJSObject(obj js.Value) (IDBCursor, error) {
	var i IDBCursor
	var err error
	if ai := IDBDatabaseGetInterface(); !ai.IsUndefined() {
		if obj.InstanceOf(ai) {
			i.BaseObject = i.SetObject(obj)
		} else {
			err = ErrNotAnIDBCursor
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBCursor) Direction() (string, error) {
	return i.GetAttributeString("direction")
}

func (i IDBCursor) Key() (interface{}, error) {
	return i.GetAttributeGlobal("key")
}

func (i IDBCursor) PrimaryKey() (interface{}, error) {
	return i.GetAttributeGlobal("primaryKey")
}

func (i IDBCursor) Advance(count int) error {
	var err error
	_, err = i.JSObject().CallWithErr("advance", js.ValueOf(count))
	return err
}

func (i IDBCursor) Request() (IDBRequest, error) {
	var err error
	var obj js.Value
	var request IDBRequest

	if obj, err = i.JSObject().CallWithErr("request"); err == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}
	return request, err
}

func (i IDBCursor) Source() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	if obj, err = i.JSObject().CallWithErr("source"); err == nil {

		bobj, err = baseobject.Discover(obj)
	}
	return bobj, err
}

func (i IDBCursor) Continue(option ...interface{}) error {
	var err error
	var arrayJS []interface{}

	if len(option) > 0 {
		if objGo, ok := option[0].(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(option[0]))
		}
	}

	_, err = i.JSObject().CallWithErr("continue", arrayJS...)
	return err
}

func (i IDBCursor) Delete(count int) (IDBRequest, error) {
	var err error
	var obj js.Value
	var request IDBRequest

	if obj, err = i.JSObject().CallWithErr("delete"); err == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}
	return request, err
}

func (i IDBCursor) Update(value interface{}) (IDBRequest, error) {
	var err error
	var obj js.Value
	var request IDBRequest
	var arrayJS []interface{}

	if objGo, ok := value.(baseobject.ObjectFrom); ok {
		arrayJS = append(arrayJS, objGo.JSObject())
	} else {
		arrayJS = append(arrayJS, js.ValueOf(value))
	}

	if obj, err = i.JSObject().CallWithErr("update", arrayJS...); err == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}
	return request, err
}
