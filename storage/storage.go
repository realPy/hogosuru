package storage

import (
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

type Storage struct {
	object js.Value
}

func GetLocalStorage(typeStorage string) (Storage, error) {
	var err error
	var localstorage Storage
	var localstorageobject, window js.Value
	if window, err = js.Global().GetWithErr("window"); err == nil {
		var strType string = "undefined"
		switch typeStorage {
		case "local":
			strType = "localStorage"
		case "session":
			strType = "sessionStorage"
		}
		if localstorageobject, err = window.GetWithErr(strType); err == nil {
			localstorage.object = localstorageobject
		}

	}
	return localstorage, err
}

func (l Storage) SetItem(key, value string) error {
	var err error
	_, err = l.object.CallWithErr("setItem", js.ValueOf(key), js.ValueOf(value))
	return err
}

func (l Storage) GetItem(key string) (string, error) {
	var err error
	var itemObject js.Value
	if itemObject, err = l.object.CallWithErr("getItem", js.ValueOf(key)); err == nil {
		return object.StringWithErr(itemObject)
	}
	return "", err
}

func (l Storage) RemoveItem(key string) error {
	var err error
	_, err = l.object.CallWithErr("removeItem", js.ValueOf(key))
	return err
}

func (l Storage) Clear() error {
	var err error
	_, err = l.object.CallWithErr("clear")
	return err
}

func (l Storage) Length() int {
	return l.object.Length()
}

func (l Storage) Key(index int) (string, error) {
	var err error
	var itemObject js.Value
	if itemObject, err = l.object.CallWithErr("key", js.ValueOf(index)); err == nil {
		return object.StringWithErr(itemObject)
	}
	return "", err
}
