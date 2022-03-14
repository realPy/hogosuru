package navigationpreloadmanager

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
	"github.com/realPy/hogosuru/promise"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var navigationpreloadmanagerinterface js.Value

//GetInterface get the JS interface navigationptpreloadmanager
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if navigationpreloadmanagerinterface, err = baseobject.Get(js.Global(), "NavigationPreloadManager"); err != nil {
			navigationpreloadmanagerinterface = js.Undefined()
		}
		baseobject.Register(navigationpreloadmanagerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		promise.GetInterface()

	})

	return navigationpreloadmanagerinterface
}

type NavigationPreloadManager struct {
	baseobject.BaseObject
}

type NavigationPreloadManagerFrom interface {
	NavigationPreloadManager_() NavigationPreloadManager
}

func (n NavigationPreloadManager) NavigationPreloadManager_() NavigationPreloadManager {
	return n
}

func NewFromJSObject(obj js.Value) (NavigationPreloadManager, error) {
	var n NavigationPreloadManager

	if ni := GetInterface(); !ni.IsUndefined() {
		if obj.InstanceOf(ni) {
			n.BaseObject = n.SetObject(obj)
			return n, nil

		}
	}

	return n, ErrNotImplemented
}

func (n NavigationPreloadManager) Enable() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = n.Call("enable"); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (n NavigationPreloadManager) Disable() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = n.Call("disable"); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (n NavigationPreloadManager) SetHeaderValue(value string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = n.Call("setHeaderValue", js.ValueOf(value)); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (n NavigationPreloadManager) GetState() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = n.Call("getState"); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}
