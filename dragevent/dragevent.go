package dragevent

// https://developer.mozilla.org/en-US/docs/Web/API/DragEvent

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/datatransfer"
	"github.com/realPy/hogosuru/mouseevent"
)

var singleton sync.Once

var drageventinterface js.Value

//DragEvent DragEvent struct
type DragEvent struct {
	mouseevent.MouseEvent
}

type DragEventFrom interface {
	DragEvent_() DragEvent
}

func (d DragEvent) DragEvent_() DragEvent {
	return d
}

//GetInterface get teh JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if drageventinterface, err = baseobject.Get(js.Global(), "DragEvent"); err != nil {
			drageventinterface = js.Undefined()
		}
		baseobject.Register(drageventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return drageventinterface
}

func New(typeevent string, init ...map[string]interface{}) (DragEvent, error) {

	var d DragEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if pei := GetInterface(); !pei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(init) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(init[0]))
		}
		if obj, err = baseobject.New(pei, arrayJS...); err == nil {
			d.BaseObject = d.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func NewFromJSObject(obj js.Value) (DragEvent, error) {
	var e DragEvent
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				e.BaseObject = e.SetObject(obj)

			} else {
				err = ErrNotAnDragEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func (d DragEvent) DataTransfer() (datatransfer.DataTransfer, error) {

	var err error
	var obj js.Value

	if obj, err = d.Get("dataTransfer"); err == nil {

		return datatransfer.NewFromJSObject(obj)
	}
	return datatransfer.DataTransfer{}, err

}
