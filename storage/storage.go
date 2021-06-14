package storage

// https://developer.mozilla.org/fr/docs/Mozilla/Add-ons/WebExtensions/API/storage

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var storageinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if storageinterface, err = js.Global().GetWithErr("Storage"); err != nil {
			storageinterface = js.Null()
		}

	})

	return storageinterface
}

type Storage struct {
	baseobject.BaseObject
}

func NewFromJSObject(obj js.Value) (Storage, error) {
	var s Storage

	if si := GetInterface(); !si.IsNull() {
		if obj.InstanceOf(si) {
			s.BaseObject = s.SetObject(obj)
			return s, nil
		}
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
		return baseobject.ToStringWithErr(itemObject)
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
		return baseobject.ToStringWithErr(itemObject)
	}
	return "", err
}
