package indexeddb

import (
	"github.com/realPy/jswasm/js"
)

type Store struct {
	storeObject js.Value
}

func CreateStore(dbObject js.Value, name string, schema map[string]interface{}) (Store, error) {
	var store Store
	if storeObject, err := dbObject.CallWithErr("createObjectStore", js.ValueOf(name), schema); err == nil {
		store.storeObject = storeObject
		return store, nil
	} else {
		return Store{}, err
	}
}

func (s Store) CreateIndex(nameIndex string, nameKey string, option map[string]interface{}) error {
	_, err := s.storeObject.CallWithErr("createIndex", js.ValueOf(nameIndex), js.ValueOf(nameKey), js.ValueOf(option))
	return err
}
