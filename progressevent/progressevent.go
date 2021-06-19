package progressevent

// https://developer.mozilla.org/en-US/docs/Web/API/ProgressEvent

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var progresseeventinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if progresseeventinterface, err = js.Global().GetWithErr("ProgressEvent"); err != nil {
			progresseeventinterface = js.Null()
		}

	})

	baseobject.Register(progresseeventinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return progresseeventinterface
}

type ProgressEvent struct {
	event.Event
}

func New() (ProgressEvent, error) {

	var p ProgressEvent

	if pei := GetInterface(); !pei.IsNull() {
		p.BaseObject = p.SetObject(pei.New())

		return p, nil
	}
	return p, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (ProgressEvent, error) {
	var p ProgressEvent

	if pei := GetInterface(); !pei.IsNull() {
		if obj.InstanceOf(pei) {
			p.BaseObject = p.SetObject(obj)

			return p, nil
		}
	}

	return p, ErrNotAnProgressEvent
}

func (p ProgressEvent) LengthComputable() (bool, error) {
	var err error
	var result bool
	var obj js.Value

	if obj, err = p.JSObject().GetWithErr("lengthComputable"); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}
	return result, err
}

func (p ProgressEvent) Loaded() (int, error) {
	var loadedObject js.Value
	var err error
	if loadedObject, err = p.JSObject().GetWithErr("loaded"); err == nil {
		if loadedObject.Type() == js.TypeNumber {
			return loadedObject.Int(), nil
		} else {
			return 0, baseobject.ErrObjectNotNumber
		}

	}
	return 0, err

}

func (p ProgressEvent) Total() (int, error) {
	var loadedObject js.Value
	var err error
	if loadedObject, err = p.JSObject().GetWithErr("total"); err == nil {
		if loadedObject.Type() == js.TypeNumber {
			return loadedObject.Int(), nil
		} else {
			return 0, baseobject.ErrObjectNotNumber
		}

	}
	return 0, err

}
