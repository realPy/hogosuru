package permissions

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

var permissionsinterface js.Value

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if permissionsinterface, err = baseobject.Get(js.Global(), "Permissions"); err != nil {
			permissionsinterface = js.Undefined()
		}
		baseobject.Register(permissionsinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return permissionsinterface
}

type Permissions struct {
	baseobject.BaseObject
}

type PermissionsFrom interface {
	Permissions_() Permissions
}

func (p Permissions) Permissions_() Permissions {
	return p
}

func NewFromJSObject(obj js.Value) (Permissions, error) {
	var p Permissions

	if pi := GetInterface(); !pi.IsUndefined() {
		if obj.InstanceOf(pi) {
			p.BaseObject = p.SetObject(obj)
			return p, nil

		}
	}

	return p, ErrNotImplemented
}

func (p Permissions) Query(permissiondescriptor map[string]interface{}) (promise.Promise, error) {

	var err error
	var obj js.Value
	var pro promise.Promise

	if obj, err = p.Call("query", js.ValueOf(permissiondescriptor)); err == nil {

		pro, err = promise.NewFromJSObject(obj)
	}

	return pro, err

}
