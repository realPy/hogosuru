package serviceworker

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/eventtarget"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var serviceworkerinterface js.Value

// GetInterface get the JS interface of serviceworker
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if serviceworkerinterface, err = baseobject.Get(js.Global(), "ServiceWorker"); err != nil {
			serviceworkerinterface = js.Undefined()
		}
		baseobject.Register(serviceworkerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		//promise.GetInterface()

	})

	return serviceworkerinterface
}

type ServiceWorker struct {
	eventtarget.EventTarget
}

type ServiceWorkerFrom interface {
	ServiceWorker_() ServiceWorker
}

func (s ServiceWorker) ServiceWorker_() ServiceWorker {
	return s
}

func NewFromJSObject(obj js.Value) (ServiceWorker, error) {
	var s ServiceWorker

	if si := GetInterface(); !si.IsUndefined() {
		if obj.InstanceOf(si) {
			s.BaseObject = s.SetObject(obj)
			return s, nil

		}
	}

	return s, ErrNotImplemented
}

func (s ServiceWorker) ScriptURL() (string, error) {

	return s.GetAttributeString("scriptURL")
}

func (s ServiceWorker) State() (string, error) {

	return s.GetAttributeString("state")
}
