package pushmanager

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/promise"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var pushmanagerinterface js.Value

// GetInterface get the JS interface pushmanager
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if pushmanagerinterface, err = baseobject.Get(js.Global(), "PushManager"); err != nil {
			pushmanagerinterface = js.Undefined()
		}
		baseobject.Register(pushmanagerinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		array.GetInterface()
		promise.GetInterface()
	})

	return pushmanagerinterface
}

type PushManager struct {
	baseobject.BaseObject
}

type PushManagerFrom interface {
	PushManagerManager_() PushManager
}

func (p PushManager) PushManager_() PushManager {
	return p
}

func NewFromJSObject(obj js.Value) (PushManager, error) {
	var p PushManager

	if pi := GetInterface(); !pi.IsUndefined() {
		if obj.InstanceOf(pi) {
			p.BaseObject = p.SetObject(obj)
			return p, nil

		}
	}

	return p, ErrNotImplemented
}

func (p PushManager) SupportedContentEncodings() (array.Array, error) {

	var err error
	var obj interface{}
	var a array.Array
	var ok bool

	if obj, err = p.GetAttributeGlobal("supportedContentEncodings"); err == nil {
		if a, ok = obj.(array.Array); !ok {
			err = array.ErrNotAnArray
		}
	}

	return a, err
}

func (p PushManager) GetSubscription() (promise.Promise, error) {
	var err error
	var obj js.Value
	var pr promise.Promise

	if obj, err = p.Call("getSubscription"); err == nil {

		pr, err = promise.NewFromJSObject(obj)

	}
	return pr, err

}

func (p PushManager) PermissionState() (promise.Promise, error) {
	var err error
	var obj js.Value
	var pr promise.Promise

	if obj, err = p.Call("permissionState"); err == nil {

		pr, err = promise.NewFromJSObject(obj)

	}
	return pr, err

}

func (p PushManager) Subscribe() (promise.Promise, error) {
	var err error
	var obj js.Value
	var pr promise.Promise

	if obj, err = p.Call("subscribe"); err == nil {

		pr, err = promise.NewFromJSObject(obj)

	}
	return pr, err

}
