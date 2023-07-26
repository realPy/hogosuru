package permissionstatus

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

var permissionstatusinterface js.Value

// GetJSInterface get the JS interface of PermissionStatus
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if permissionstatusinterface, err = baseobject.Get(js.Global(), "PermissionStatus"); err != nil {
			permissionstatusinterface = js.Undefined()
		}

		baseobject.Register(permissionstatusinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return permissionstatusinterface
}

type PermissionStatus struct {
	eventtarget.EventTarget
}

type PermissionStatusFrom interface {
	PermissionStatus_() PermissionStatus
}

func (p PermissionStatus) PermissionStatus_() PermissionStatus {
	return p
}

func NewFromJSObject(obj js.Value) (PermissionStatus, error) {
	var p PermissionStatus

	if psi := GetInterface(); !psi.IsUndefined() {
		if obj.InstanceOf(psi) {
			p.BaseObject = p.SetObject(obj)
			return p, nil

		}
	}

	return p, ErrNotImplemented
}

func (p PermissionStatus) Name() (string, error) {

	return p.GetAttributeString("name")
}

func (p PermissionStatus) State() (string, error) {

	return p.GetAttributeString("state")
}

// deprecated
func (p PermissionStatus) Status() (string, error) {

	return p.GetAttributeString("status")
}
