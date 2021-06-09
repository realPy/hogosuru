package dragevent

// https://developer.mozilla.org/en-US/docs/Web/API/DragEvent

import (
	"sync"

	"syscall/js"

	datatransfert "github.com/realPy/hogosuru/datatransfer"
	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var drageventinterface js.Value

//DragEvent DragEvent struct
type DragEvent struct {
	event.Event
}

//GetInterface get teh JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if drageventinterface, err = js.Global().GetWithErr("DragEvent"); err != nil {
			drageventinterface = js.Null()
		}

	})

	return drageventinterface
}

func NewFromJSObject(obj js.Value) (DragEvent, error) {
	var e DragEvent

	if di := GetInterface(); !di.IsNull() {
		if obj.InstanceOf(di) {
			e.BaseObject = e.SetObject(obj)
			return e, nil
		}
	}
	return e, ErrNotAnDragEvent
}

func (d DragEvent) DataTransfer() (datatransfert.DataTransfer, error) {

	var err error
	var obj js.Value

	if obj, err = d.JSObject().GetWithErr("dataTransfer"); err == nil {

		return datatransfert.NewFromJSObject(obj)
	}
	return datatransfert.DataTransfer{}, err

}
