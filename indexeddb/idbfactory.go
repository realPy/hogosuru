package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBFactory

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonIDBFactory sync.Once

var idbfactoryinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetIDBFactoryInterface() js.Value {

	singletonIDBFactory.Do(func() {
		var err error
		if idbfactoryinterface, err = js.Global().GetWithErr("IDBFactory"); err != nil {
			idbfactoryinterface = js.Null()
		}
	})
	return idbfactoryinterface
}

//IDBFactory struct
type IDBFactory struct {
	baseobject.BaseObject
}

type IDBFactoryFrom interface {
	IDBFactory_() IDBFactory
}

func (i IDBFactory) IDBFactory_() IDBFactory {
	return i
}

func IDBFactoryNewFromJSObject(obj js.Value) (IDBFactory, error) {
	var i IDBFactory
	var err error
	if ai := GetIDBFactoryInterface(); !ai.IsNull() {
		if obj.InstanceOf(ai) {
			i.BaseObject = i.SetObject(obj)
		} else {
			err = ErrNotAnIDBFactory
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (f IDBFactory) genericRequest(method string, dbname string, option ...string) (IDBOpenDBRequest, error) {
	var err error
	var i IDBOpenDBRequest
	var idbobj js.Value

	var arrayJS []interface{}
	arrayJS = append(arrayJS, js.ValueOf(dbname))

	if len(option) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(option[0]))
	}

	if idbobj, err = f.JSObject().CallWithErr(method, arrayJS...); err == nil {
		i, err = IDBOpenDBRequestNewFromJSObject(idbobj)

	}

	return i, err

}

func (f IDBFactory) Open(dbname string, option ...string) (IDBOpenDBRequest, error) {

	return f.genericRequest("open", dbname, option...)
}

func (f IDBFactory) DeleteDatabase(dbname string, option ...string) (IDBOpenDBRequest, error) {
	return f.genericRequest("deleteDatabase", dbname, option...)

}

func (f IDBFactory) Databases() {
	//not support in firefox

}
