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

type ProgressEventFrom interface {
	ProgressEvent() ProgressEvent
}

func (p ProgressEvent) ProgressEvent() ProgressEvent {
	return p
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
	return p.CallBool("lengthComputable")
}

func (p ProgressEvent) Loaded() (int, error) {
	return p.GetAttributeInt("loaded")
}

func (p ProgressEvent) Total() (int, error) {
	return p.GetAttributeInt("total")
}
