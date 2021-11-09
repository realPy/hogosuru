package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBFactory

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
)

var singletonIDBFactory sync.Once

var idbfactoryinterface js.Value

//GetInterface get the JS interface
func GetIDBFactoryInterface() js.Value {

	singletonIDBFactory.Do(func() {
		var err error
		if idbfactoryinterface, err = baseobject.Get(js.Global(), "IDBFactory"); err != nil {
			idbfactoryinterface = js.Undefined()
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
	if ai := GetIDBFactoryInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.BaseObject = i.SetObject(obj)
			} else {
				err = ErrNotAnIDBFactory
			}
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

	if idbobj, err = f.Call(method, arrayJS...); err == nil {
		i, err = IDBOpenDBRequestNewFromJSObject(idbobj)

	}

	return i, err

}

func (f IDBFactory) Cmp(a, b interface{}) (int, error) {

	var arrayJS []interface{}
	var err error
	var obj js.Value
	var result int

	if objGo, ok := a.(baseobject.ObjectFrom); ok {
		arrayJS = append(arrayJS, objGo.JSObject())
	} else {
		arrayJS = append(arrayJS, js.ValueOf(a))
	}

	if objGo, ok := b.(baseobject.ObjectFrom); ok {
		arrayJS = append(arrayJS, objGo.JSObject())
	} else {
		arrayJS = append(arrayJS, js.ValueOf(b))
	}

	if obj, err = f.Call("cmp", arrayJS...); err == nil {
		if obj.Type() == js.TypeNumber {
			result = obj.Int()
		} else {
			err = baseobject.ErrObjectNotNumber
		}
	}
	return result, err

}

func (f IDBFactory) Open(dbname string, option ...string) (IDBOpenDBRequest, error) {

	return f.genericRequest("open", dbname, option...)
}

func (f IDBFactory) DeleteDatabase(dbname string, option ...string) (IDBOpenDBRequest, error) {
	return f.genericRequest("deleteDatabase", dbname, option...)

}

func (f IDBFactory) Databases() (promise.Promise, error) {
	//not support in firefox
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = f.Call("databases"); err == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}
