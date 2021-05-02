package indexeddb

import (
	"fmt"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

type Store struct {
	objinterface object.ObjectInterface
	objstore     js.Value
}

func CreateStore(dbObject js.Value, name string, schema map[string]interface{}) (Store, error) {
	var store Store
	if storeObject, err := dbObject.CallWithErr("createObjectStore", js.ValueOf(name), schema); err == nil {
		store.objstore = storeObject
		return store, nil
	} else {
		return Store{}, err
	}
}

func (s Store) callWaitableMethod(method string, args ...interface{}) (js.Value, error) {
	var waitable, obj js.Value
	var err error
	if waitable, err = s.objstore.CallWithErr(method, args...); err == nil {
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
	if obj, err = s.callWaitableMethod("add", js.ValueOf(value)); err == nil {
		if value := object.NewGOValue(obj); value.IsInt() {
			return value.Int(), nil
		} else {
			return 0, nil
		}

	}
	return 0, err
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
		if value := object.NewGOValue(obj); value.IsInt() {
			return value.Int(), nil
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

func (s Store) Get(key int) (object.GOMap, error) {
	var arrayObject js.Value
	var err error
	var mapobj object.GOMap

	if arrayObject, err = s.callWaitableMethod("get", js.ValueOf(key)); err == nil {
		if entries, err := s.objinterface.Entries(arrayObject); err == nil {
			mapobj = object.Map(entries)
		}
	}

	return mapobj, err
}

func (s Store) GetAllKeys() (object.GOArray, error) {

	var arraykeys object.GOArray
	var err error
	var arrayObject js.Value

	if arrayObject, err = s.callWaitableMethod("getAllKeys"); err == nil {
		arraykeys, err = object.Array(arrayObject), nil
	}
	return arraykeys, err
}

func (s Store) GetAll() ([]object.GOMap, error) {

	var err error
	var arraysObject js.Value
	var arrayGoMap []object.GOMap

	if arraysObject, err = s.callWaitableMethod("getAll"); err == nil {

		object.ParseArray(arraysObject, func(v js.Value) {
			arrayGoMap = append(arrayGoMap, object.Map(v))
		})
	}
	return arrayGoMap, err
}

//Put a value in store and return the index  of the new inserted element
func (s Store) Put(value map[string]interface{}) (int, error) {
	var err error
	var obj js.Value
	if obj, err = s.callWaitableMethod("put", js.ValueOf(value)); err == nil {
		if value := object.NewGOValue(obj); value.IsInt() {
			return value.Int(), nil
		} else {
			return 0, nil
		}

	}
	return 0, err
}
