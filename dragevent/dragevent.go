package dragevent

// https://developer.mozilla.org/en-US/docs/Web/API/DragEvent

import (
	"sync"

	"syscall/js"

	datatransfert "github.com/realPy/hogosuru/datatransfer"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var drageventinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//DragEvent DragEvent struct
type DragEvent struct {
	event.Event
}

//GetJSInterface get teh JS interface of event
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var drageventinstance JSInterface
		var err error
		if drageventinstance.objectInterface, err = js.Global().GetWithErr("DragEvent"); err == nil {
			drageventinterface = &drageventinstance
		}
	})

	return drageventinterface
}

func NewFromJSObject(obj js.Value) (DragEvent, error) {
	var e DragEvent

	if object.String(obj) == "[object DragEvent]" {
		e.Object = e.SetObject(obj)
		return e, nil
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
