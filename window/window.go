package window

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/eventtarget"
	"github.com/realPy/hogosuru/history"
	"github.com/realPy/hogosuru/indexeddb"
	"github.com/realPy/hogosuru/location"
	"github.com/realPy/hogosuru/storage"
)

var singleton sync.Once

var windowinterface js.Value

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if windowinterface, err = js.Global().GetWithErr("Window"); err != nil {
			windowinterface = js.Undefined()
		}
		baseobject.Register(windowinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return windowinterface
}

type Window struct {
	eventtarget.EventTarget
}

type WindowFrom interface {
	Window_() Window
}

func (w Window) Window_() Window {
	return w
}

func NewFromJSObject(obj js.Value) (Window, error) {
	var w Window

	if wi := GetInterface(); !wi.IsUndefined() {
		if obj.InstanceOf(wi) {
			w.BaseObject = w.SetObject(obj)
			return w, nil

		}
	}

	return w, ErrNotImplemented
}

func New() (Window, error) {

	var err error
	var w Window
	var windowObj js.Value
	if windowObj, err = js.Global().GetWithErr("window"); err == nil {

		w, err = NewFromJSObject(windowObj)

	}
	return w, err
}

func (w Window) Document() (document.Document, error) {
	var err error
	var obj js.Value
	var h document.Document

	if obj, err = w.JSObject().GetWithErr("document"); err == nil {
		h, err = document.NewFromJSObject(obj)
	}

	return h, err
}

func (w Window) History() (history.History, error) {
	var err error
	var obj js.Value
	var h history.History

	if obj, err = w.JSObject().GetWithErr("history"); err == nil {
		h, err = history.NewFromJSObject(obj)
	}

	return h, err
}

func (w Window) Location() (location.Location, error) {
	var err error
	var obj js.Value
	var l location.Location

	if obj, err = w.JSObject().GetWithErr("location"); err == nil {
		l, err = location.NewFromJSObject(obj)
	}

	return l, err
}

func (w Window) LocalStorage() (storage.Storage, error) {
	var err error
	var obj js.Value
	var s storage.Storage

	if obj, err = w.JSObject().GetWithErr("localStorage"); err == nil {
		s, err = storage.NewFromJSObject(obj)
	}

	return s, err
}

func (w Window) SessionStorage() (storage.Storage, error) {
	var err error
	var obj js.Value
	var s storage.Storage

	if obj, err = w.JSObject().GetWithErr("sessionStorage"); err == nil {
		s, err = storage.NewFromJSObject(obj)
	}

	return s, err
}

func (w Window) IndexdedDB() (indexeddb.IDBFactory, error) {
	var err error
	var obj js.Value
	var i indexeddb.IDBFactory

	if obj, err = w.JSObject().GetWithErr("indexedDB"); err == nil {
		i, err = indexeddb.IDBFactoryNewFromJSObject(obj)
	}

	return i, err
}
