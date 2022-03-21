package worker

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/eventtarget"
	"github.com/realPy/hogosuru/initinterface"
	"github.com/realPy/hogosuru/messageevent"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var workerinterface js.Value

//GetInterface get the JS interface of serviceworkerregistration
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if workerinterface, err = baseobject.Get(js.Global(), "Worker"); err != nil {
			workerinterface = js.Undefined()
		}
		baseobject.Register(workerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		messageevent.GetInterface()

	})

	return workerinterface
}

type Worker struct {
	eventtarget.EventTarget
}

type WorkerFrom interface {
	Worker_() Worker
}

func (w Worker) Worker_() Worker {
	return w
}

func NewFromJSObject(obj js.Value) (Worker, error) {
	var w Worker

	if wi := GetInterface(); !wi.IsUndefined() {
		if obj.InstanceOf(wi) {
			w.BaseObject = w.SetObject(obj)
			return w, nil

		}
	}

	return w, ErrNotImplemented
}

func New(url string, opts ...map[string]interface{}) (Worker, error) {

	var arrayJS []interface{}
	var w Worker
	var err error
	var obj js.Value

	arrayJS = append(arrayJS, js.ValueOf(url))

	if len(opts) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(opts[0]))
	}

	if workeri := GetInterface(); !workeri.IsUndefined() {

		if obj, err = baseobject.New(workeri, arrayJS...); err == nil {
			w.BaseObject = w.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return w, err
}

func (w Worker) PostMessage(message string, transfer ...array.Array) error {

	var arrayJS []interface{}

	var err error

	arrayJS = append(arrayJS, js.ValueOf(message))

	if len(transfer) > 0 {
		arrayJS = append(arrayJS, transfer[0].JSObject())
	}

	_, err = w.Call("postMessage", arrayJS...)

	return err

}

func (w Worker) Terminate() error {

	var err error
	_, err = w.Call("terminate")

	return err

}
