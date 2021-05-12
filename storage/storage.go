package storage

// https://developer.mozilla.org/fr/docs/Mozilla/Add-ons/WebExtensions/API/storage

import (
	"syscall/js"

	"github.com/realPy/hogosuru/object"
)

type Storage struct {
	object.Object
}

func NewFromJSObject(obj js.Value) (Storage, error) {
	var s Storage

	if object.String(obj) == "[object Storage]" {
		s.Object = s.SetObject(obj)
		return s, nil
	}

	return s, ErrNotAnLocalStorage
}

func New(typeStorage string) (Storage, error) {
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

			localstorage, err = NewFromJSObject(localstorageobject)
		}

	}
	return localstorage, err
}

func (l Storage) SetItem(key, value string) error {
	var err error
	_, err = l.JSObject().CallWithErr("setItem", js.ValueOf(key), js.ValueOf(value))
	return err
}

func (l Storage) GetItem(key string) (string, error) {
	var err error
	var itemObject js.Value
	if itemObject, err = l.JSObject().CallWithErr("getItem", js.ValueOf(key)); err == nil {
		return object.StringWithErr(itemObject)
	}
	return "", err
}

func (l Storage) RemoveItem(key string) error {
	var err error
	_, err = l.JSObject().CallWithErr("removeItem", js.ValueOf(key))
	return err
}

func (l Storage) Clear() error {
	var err error
	_, err = l.JSObject().CallWithErr("clear")
	return err
}
func (l Storage) Key(index int) (string, error) {
	var err error
	var itemObject js.Value
	if itemObject, err = l.JSObject().CallWithErr("key", js.ValueOf(index)); err == nil {
		return object.StringWithErr(itemObject)
	}
	return "", err
}
