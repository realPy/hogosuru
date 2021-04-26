package localstorage

import (
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

type LocalStorage struct {
	object js.Value
}

func GetLocalStorage() (LocalStorage, error) {
	var err error
	var localstorage LocalStorage
	var localstorageobject, window js.Value
	if window, err = js.Global().GetWithErr("window"); err == nil {

		if localstorageobject, err = window.GetWithErr("localStorage"); err == nil {
			localstorage.object = localstorageobject
		}

	}
	return localstorage, err
}

func (l LocalStorage) SetItem(key, value string) error {
	var err error
	_, err = l.object.CallWithErr("setItem", js.ValueOf(key), js.ValueOf(value))
	return err
}

func (l LocalStorage) GetItem(key string) (string, error) {
	var err error
	var itemObject js.Value
	if itemObject, err = l.object.CallWithErr("getItem", js.ValueOf(key)); err == nil {
		return object.String(itemObject)
	}
	return "", err
}

func (l LocalStorage) RemoveItem(key string) error {
	var err error
	_, err = l.object.CallWithErr("removeItem", js.ValueOf(key))
	return err
}

func (l LocalStorage) Clear() error {
	var err error
	_, err = l.object.CallWithErr("clear")
	return err
}

func (l LocalStorage) Length() int {
	return l.object.Length()
}

func (l LocalStorage) Key(index int) (string, error) {
	var err error
	var itemObject js.Value
	if itemObject, err = l.object.CallWithErr("key", js.ValueOf(index)); err == nil {
		return object.String(itemObject)
	}
	return "", err
}
