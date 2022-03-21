package dedicatedworkerglobalscope

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
	"github.com/realPy/hogosuru/messageevent"
	"github.com/realPy/hogosuru/workerglobalscope"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var dedicatedworkerglobalscopeinterface js.Value

//GetInterface get the JS interface of serviceworkerregistration
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if dedicatedworkerglobalscopeinterface, err = baseobject.Get(js.Global(), "DedicatedWorkerGlobalScope"); err != nil {
			dedicatedworkerglobalscopeinterface = js.Undefined()
		}
		baseobject.Register(dedicatedworkerglobalscopeinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		messageevent.GetInterface()

	})

	return dedicatedworkerglobalscopeinterface
}

type DedicatedWorkerGlobalScope struct {
	workerglobalscope.WorkerGlobalScope
}

type DedicatedWorkerGlobalScopeFrom interface {
	DedicatedWorkerGlobalScope_() DedicatedWorkerGlobalScope
}

func (d DedicatedWorkerGlobalScope) DedicatedWorkerGlobalScope_() DedicatedWorkerGlobalScope {
	return d
}

func NewFromJSObject(obj js.Value) (DedicatedWorkerGlobalScope, error) {
	var d DedicatedWorkerGlobalScope

	if di := GetInterface(); !di.IsUndefined() {
		if obj.InstanceOf(di) {
			d.BaseObject = d.SetObject(obj)
			return d, nil

		}
	}

	return d, ErrNotImplemented
}

func (d DedicatedWorkerGlobalScope) PostMessage(message string, transfer ...array.Array) error {

	var arrayJS []interface{}

	var err error

	arrayJS = append(arrayJS, js.ValueOf(message))

	if len(transfer) > 0 {
		arrayJS = append(arrayJS, transfer[0].JSObject())
	}

	_, err = d.Call("postMessage", arrayJS...)

	return err

}

func (d DedicatedWorkerGlobalScope) Name() (string, error) {

	return d.GetAttributeString("name")
}

func (d DedicatedWorkerGlobalScope) Close() error {

	var err error
	_, err = d.Call("close")

	return err

}
