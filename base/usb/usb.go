package usb

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/eventtarget"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/promise"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var usbinterface js.Value

// USB struct
type USB struct {
	eventtarget.EventTarget
}

type USBFrom interface {
	USB_() USB
}

func (u USB) Clipboard_() USB {
	return u
}

// GetJSInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if usbinterface, err = baseobject.Get(js.Global(), "USB"); err != nil {
			usbinterface = js.Undefined()
		}

		baseobject.Register(usbinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)

		})
		promise.GetInterface()

	})

	return usbinterface
}

func NewFromJSObject(obj js.Value) (USB, error) {
	var u USB

	if ci := GetInterface(); !ci.IsUndefined() {
		if obj.InstanceOf(ci) {
			u.BaseObject = u.SetObject(obj)
			return u, nil

		}
	}

	return u, ErrNotImplemented
}

func (u USB) GetDevices() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = u.Call("getDevices"); err == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}

func (u USB) RequestDevices(filter map[string]interface{}) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = u.Call("requestDevice", js.ValueOf(filter)); err == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}
