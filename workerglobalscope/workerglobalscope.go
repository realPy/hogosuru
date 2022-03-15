package workerglobalscope

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/eventtarget"
	"github.com/realPy/hogosuru/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var workerglobalscopeinterface js.Value

//GetInterface get the JS interface of serviceworkerregistration
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if workerglobalscopeinterface, err = baseobject.Get(js.Global(), "WorkerGlobalScope"); err != nil {
			workerglobalscopeinterface = js.Undefined()
		}
		baseobject.Register(workerglobalscopeinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return workerglobalscopeinterface
}

type WorkerGlobalScope struct {
	eventtarget.EventTarget
}

type WorkerGlobalScopeFrom interface {
	WorkerGlobalScope_() WorkerGlobalScope
}

func (w WorkerGlobalScope) WorkerGlobalScope_() WorkerGlobalScope {
	return w
}

func NewFromJSObject(obj js.Value) (WorkerGlobalScope, error) {
	var w WorkerGlobalScope

	if wi := GetInterface(); !wi.IsUndefined() {
		if obj.InstanceOf(wi) {
			w.BaseObject = w.SetObject(obj)
			return w, nil

		}
	}

	return w, ErrNotImplemented
}

func (w WorkerGlobalScope) ImportsScripts(values ...string) error {

	var err error
	var arrayJS []interface{}

	for _, value := range values {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}
	_, err = w.Call("importScripts", arrayJS...)

	return err
}
