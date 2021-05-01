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

func (s Store) CreateIndex(nameIndex string, nameKey string, option map[string]interface{}) error {
	_, err := s.objstore.CallWithErr("createIndex", js.ValueOf(nameIndex), js.ValueOf(nameKey), js.ValueOf(option))
	return err
}

func (s Store) Add(value map[string]interface{}) error {
	_, err := s.objstore.CallWithErr("add", js.ValueOf(value))
	return err
}

func (s Store) Get(key int) (object.GOMap, error) {
	var waitable, arrayObject js.Value
	var err error
	var mapobj object.GOMap

	if waitable, err = s.objstore.CallWithErr("get", js.ValueOf(key)); err == nil {
		ch := OnSuccessFailure(waitable)
		results := <-ch
		if results.Success {
			ev := results.Payload[0]
			if arrayObject, err = getEventTargetResult(ev); err == nil {
				if entries, err := s.objinterface.Entries(arrayObject); err == nil {
					mapobj = object.Map(entries)
				}
			}
		} else {
			err = fmt.Errorf("erreur store:get")
			// recuperer error https://developer.mozilla.org/fr/docs/Web/API/IDBRequest/error
		}
	}
	return mapobj, err
}

func (s Store) GetAllKeys() (object.GOArray, error) {

	var arraykeys object.GOArray
	var err error
	var waitable, arrayObject js.Value

	if waitable, err = s.objstore.CallWithErr("getAllKeys"); err == nil {
		ch := OnSuccessFailure(waitable)
		results := <-ch
		if results.Success {
			ev := results.Payload[0]
			if arrayObject, err = getEventTargetResult(ev); err == nil {
				arraykeys, err = object.Array(arrayObject), nil
				//arraykeys, err = array.GoArrayInt(arrayObject), nil
			}
		} else {
			err = fmt.Errorf("erreur getAllkeys")
			// recuperer error https://developer.mozilla.org/fr/docs/Web/API/IDBRequest/error
		}

	}

	return arraykeys, err
}
