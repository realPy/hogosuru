package clipboardevent

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/datatransfer"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var clipboardeventinterface js.Value

//ClipboardEvent ClipboardEvent struct
type ClipboardEvent struct {
	event.Event
}

type ClipboardEventFrom interface {
	ClipboardEvent_() ClipboardEvent
}

func (c ClipboardEvent) ClipboardEvent_() ClipboardEvent {
	return c
}

//GetInterface get the JS interface of ClipboardEvent
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if clipboardeventinterface, err = baseobject.Get(js.Global(), "ClipboardEvent"); err != nil {
			clipboardeventinterface = js.Undefined()
		}

		baseobject.Register(clipboardeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)

		})
		datatransfer.GetInterface()
	})

	return clipboardeventinterface
}

func NewFromJSObject(obj js.Value) (ClipboardEvent, error) {
	var c ClipboardEvent
	var err error

	if bi := GetInterface(); !bi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(bi) {
				c.BaseObject = c.SetObject(obj)

			} else {
				err = ErrNotAnCustomEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return c, err
}

//New Create a ClipboardEvent
func New(data datatransfer.DataTransfer) (ClipboardEvent, error) {
	var event ClipboardEvent
	var obj js.Value
	var err error
	if eventi := GetInterface(); !eventi.IsUndefined() {
		if obj, err = baseobject.New(eventi, data.JSObject()); err == nil {
			event.BaseObject = event.SetObject(obj)
		}
	} else {
		err = ErrNotImplemented
	}
	return event, err
}

func (c ClipboardEvent) ClipboardData() (datatransfer.DataTransfer, error) {
	var obj interface{}
	var err error
	var d datatransfer.DataTransfer
	var ok bool
	if obj, err = c.GetAttributeGlobal("clipboardData"); err == nil {
		if d, ok = obj.(datatransfer.DataTransfer); !ok {

			err = datatransfer.ErrNotADataTransfer

		}

	}
	return d, err
}
