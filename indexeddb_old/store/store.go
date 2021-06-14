package store

import (
	"fmt"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/object"
)

type Store struct {
	baseobject.BaseObject
}

func OnSuccessFailure(awaitable js.Value) chan SuccessFailure {
	ch := make(chan SuccessFailure)
	cbok := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- SuccessFailure{Success: true, Payload: args}
		return nil
	})
	cberror := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- SuccessFailure{Success: false, Payload: args}
		return nil
	})
	awaitable.Set("onsuccess", cbok)
	awaitable.Set("onerror", cberror)
	return ch
}

type SuccessFailure struct {
	Success bool
	Payload []js.Value
}

func getEventTargetResult(ev js.Value) (js.Value, error) {
	if target, err := ev.GetWithErr("target"); err == nil {
		if result, err := target.GetWithErr("result"); err == nil {
			return result, nil
		} else {
			return js.Value{}, fmt.Errorf("result not found")
		}
	} else {
		return js.Value{}, fmt.Errorf("target not found")
	}
}

func NewFromJSObject(obj js.Value) (Store, error) {

	var s Store
	s.BaseObject = s.SetObject(obj)
	return s, nil
}

func (s Store) callWaitableMethod(method string, args ...interface{}) (js.Value, error) {
	var waitable, obj js.Value
	var err error

	if waitable, err = s.JSObject().CallWithErr(method, args...); err == nil {
		ch := OnSuccessFailure(waitable)
		results := <-ch
		if results.Success {
			ev := results.Payload[0]
			if obj, err = getEventTargetResult(ev); err == nil {
				return obj, err
			}
		} else {
			err = fmt.Errorf("error store:%s", method)
			// recuperer error https://developer.mozilla.org/fr/docs/Web/API/IDBRequest/error
		}
	}

	return obj, err
}

//Add a value in store and return the index  of the new inserted element
func (s Store) Add(value map[string]interface{}) (int, error) {
	var err error
	var obj js.Value
	var indexret int

	if obj, err = s.callWaitableMethod("add", js.ValueOf(value)); err == nil {
		if obj.Type() == js.TypeNumber {
			indexret = obj.Int()
		} else {
			err = baseobject.ErrObjectNotNumber
		}
	}
	return indexret, err
}

func (s Store) Clear() error {
	_, err := s.callWaitableMethod("clear")
	return err

}

func (s Store) Count() (int, error) {
	var err error
	var obj js.Value
	var count int
	if obj, err = s.callWaitableMethod("count"); err == nil {

		if obj.Type() == js.TypeNumber {
			count = obj.Int()
		} else {
			err = baseobject.ErrObjectNotNumber
		}
	}
	return count, err

}

func (s Store) Delete(index int) error {
	var err error

	if _, err = s.callWaitableMethod("delete", js.ValueOf(index)); err == nil {
		return nil
	}
	return err

}

func (s Store) DeleteIndex(nameIndex string) error {
	var err error

	if _, err = s.callWaitableMethod("deleteIndex", js.ValueOf(nameIndex)); err == nil {
		return nil
	}
	return err

}

func (s Store) CreateIndex(nameIndex string, nameKey string, option map[string]interface{}) error {
	_, err := s.callWaitableMethod("createIndex", js.ValueOf(nameIndex), js.ValueOf(nameKey), js.ValueOf(option))
	return err
}

func (s Store) Get(key int) (object.Object, error) {
	var jsobj js.Value
	var err error
	var obj object.Object

	if jsobj, err = s.callWaitableMethod("get", js.ValueOf(key)); err == nil {
		obj, err = object.NewFromJSObject(jsobj)
	}

	return obj, err
}

func (s Store) GetAllKeys() (baseobject.GOArray, error) {

	var arraykeys baseobject.GOArray
	var err error
	var arrayObject js.Value

	if arrayObject, err = s.callWaitableMethod("getAllKeys"); err == nil {
		arraykeys, err = baseobject.Array(arrayObject), nil
	}
	return arraykeys, err
}

func (s Store) GetAll() ([]baseobject.GOMap, error) {

	var err error
	var arraysObject js.Value
	var arrayGoMap []baseobject.GOMap

	if arraysObject, err = s.callWaitableMethod("getAll"); err == nil {

		baseobject.ParseArray(arraysObject, func(v js.Value) {
			arrayGoMap = append(arrayGoMap, baseobject.Map(v))
		})
	}
	return arrayGoMap, err
}

//Put a value in store and return the index  of the new inserted element
func (s Store) Put(value map[string]interface{}) (int, error) {
	var err error
	var obj js.Value
	if obj, err = s.callWaitableMethod("put", js.ValueOf(value)); err == nil {
		if value := baseobject.NewGOValue(obj); value.IsInt() {
			return value.Int(), nil
		} else {
			return 0, nil
		}

	}
	return 0, err
}
