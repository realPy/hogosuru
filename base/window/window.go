package window

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/eventtarget"
	"github.com/realPy/hogosuru/base/history"
	"github.com/realPy/hogosuru/base/indexeddb"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/location"
	"github.com/realPy/hogosuru/base/navigator"
	"github.com/realPy/hogosuru/base/storage"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var windowinterface js.Value

// GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if windowinterface, err = baseobject.Get(js.Global(), "Window"); err != nil {
			windowinterface = js.Undefined()
		}
		baseobject.Register(windowinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
		navigator.GetInterface()
		history.GetInterface()
		location.GetInterface()
		storage.GetInterface()

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
	if windowObj, err = baseobject.Get(js.Global(), "window"); err == nil {

		w, err = NewFromJSObject(windowObj)

	}
	return w, err
}

func (w Window) Document() (document.Document, error) {
	var err error
	var obj js.Value
	var d document.Document

	if obj, err = w.Get("document"); err == nil {
		d, err = document.NewFromJSObject(obj)
	}

	return d, err
}

func (w Window) History() (history.History, error) {
	var err error
	var obj js.Value
	var h history.History

	if obj, err = w.Get("history"); err == nil {
		h, err = history.NewFromJSObject(obj)
	}

	return h, err
}

func (w Window) Location() (location.Location, error) {
	var err error
	var obj js.Value
	var l location.Location

	if obj, err = w.Get("location"); err == nil {
		l, err = location.NewFromJSObject(obj)
	}

	return l, err
}

func (w Window) LocalStorage() (storage.Storage, error) {
	var err error
	var obj js.Value
	var s storage.Storage

	if obj, err = w.Get("localStorage"); err == nil {
		s, err = storage.NewFromJSObject(obj)
	}

	return s, err
}

func (w Window) SessionStorage() (storage.Storage, error) {
	var err error
	var obj js.Value
	var s storage.Storage

	if obj, err = w.Get("sessionStorage"); err == nil {
		s, err = storage.NewFromJSObject(obj)
	}

	return s, err
}

func (w Window) IndexdedDB() (indexeddb.IDBFactory, error) {
	var err error
	var obj js.Value
	var i indexeddb.IDBFactory

	if obj, err = w.Get("indexedDB"); err == nil {
		i, err = indexeddb.IDBFactoryNewFromJSObject(obj)
	}

	return i, err
}

func (w Window) Navigator() (navigator.Navigator, error) {
	var err error
	var obj js.Value
	var n navigator.Navigator

	if obj, err = w.Get("navigator"); err == nil {
		n, err = navigator.NewFromJSObject(obj)
	}

	return n, err
}
