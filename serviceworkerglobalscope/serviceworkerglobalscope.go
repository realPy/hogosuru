package serviceworkerglobalscope

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/serviceworkerregistration"
	"github.com/realPy/hogosuru/workerglobalscope"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var serviceworkerglobalscopeinterface js.Value

//GetInterface get the JS interface of serviceworkerregistration
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if serviceworkerglobalscopeinterface, err = baseobject.Get(js.Global(), "ServiceWorkerGlobalScope"); err != nil {
			serviceworkerglobalscopeinterface = js.Undefined()
		}
		baseobject.Register(serviceworkerglobalscopeinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		serviceworkerregistration.GetInterface()
		promise.GetInterface()
	})

	return serviceworkerglobalscopeinterface
}

type ServiceWorkerGlobalScope struct {
	workerglobalscope.WorkerGlobalScope
}

type WorkerGlobalScopeFrom interface {
	ServiceWorkerGlobalScope_() ServiceWorkerGlobalScope
}

func (w ServiceWorkerGlobalScope) ServiceWorkerGlobalScope_() ServiceWorkerGlobalScope {
	return w
}

func NewFromJSObject(obj js.Value) (ServiceWorkerGlobalScope, error) {
	var w ServiceWorkerGlobalScope

	if wi := GetInterface(); !wi.IsUndefined() {
		if obj.InstanceOf(wi) {
			w.BaseObject = w.SetObject(obj)
			return w, nil

		}
	}

	return w, ErrNotImplemented
}

func (w ServiceWorkerGlobalScope) Registration() (serviceworkerregistration.ServiceWorkerRegistration, error) {

	var err error
	var obj interface{}
	var s serviceworkerregistration.ServiceWorkerRegistration
	var ok bool

	if obj, err = w.GetAttributeGlobal("registration"); err == nil {
		if s, ok = obj.(serviceworkerregistration.ServiceWorkerRegistration); !ok {
			err = serviceworkerregistration.ErrNotAServiceWorkerRegistration
		}
	}

	return s, err
}

func (w ServiceWorkerGlobalScope) SkipWaiting() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = w.Call("skipWaiting"); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}
