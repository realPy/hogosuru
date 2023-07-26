package serviceworkercontainer

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/eventtarget"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/serviceworker"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var serviceworkercontainerinterface js.Value

// GetInterface get the JS interface of serviceworkercontainer
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if serviceworkercontainerinterface, err = baseobject.Get(js.Global(), "ServiceWorkerContainer"); err != nil {
			serviceworkercontainerinterface = js.Undefined()
		}
		baseobject.Register(serviceworkercontainerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		promise.GetInterface()
		serviceworker.GetInterface()

	})

	return serviceworkercontainerinterface
}

type ServiceWorkerContainer struct {
	eventtarget.EventTarget
}

type ServiceWorkerContainerFrom interface {
	ServiceWorkerContainer_() ServiceWorkerContainer
}

func (s ServiceWorkerContainer) ServiceWorkerContainer_() ServiceWorkerContainer {
	return s
}

func NewFromJSObject(obj js.Value) (ServiceWorkerContainer, error) {
	var s ServiceWorkerContainer

	if si := GetInterface(); !si.IsUndefined() {
		if obj.InstanceOf(si) {
			s.BaseObject = s.SetObject(obj)
			return s, nil

		}
	}

	return s, ErrNotImplemented
}

func (s ServiceWorkerContainer) Controller() (serviceworker.ServiceWorker, error) {

	var err error
	var obj interface{}
	var sw serviceworker.ServiceWorker
	var ok bool

	if obj, err = s.GetAttributeGlobal("controller"); err == nil {

		if obj == nil {
			return sw, ErrControllerNotDefined
		} else {
			if sw, ok = obj.(serviceworker.ServiceWorker); !ok {

				err = serviceworker.ErrNotAServiceWorker

			}
		}

	}

	return sw, err

}

func (s ServiceWorkerContainer) Ready() (promise.Promise, error) {
	var err error
	var obj interface{}
	var p promise.Promise
	var ok bool

	if obj, err = s.GetAttributeGlobal("ready"); err == nil {

		if p, ok = obj.(promise.Promise); !ok {
			err = promise.ErrNotAPromise
		}

	}

	return p, err

}

func (s ServiceWorkerContainer) GetRegistration(clientURL string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = s.Call("getRegistration", js.ValueOf(clientURL)); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (s ServiceWorkerContainer) Register(url string, options ...map[string]interface{}) (promise.Promise, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var p promise.Promise

	arrayJS = append(arrayJS, js.ValueOf(url))

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj, err = s.Call("register", arrayJS...); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}

	return p, err
}
