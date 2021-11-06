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
		if progresseeventinterface, err = baseobject.Get(js.Global(), "ProgressEvent"); err != nil {
			progresseeventinterface = js.Undefined()
		}
		baseobject.Register(progresseeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return progresseeventinterface
}

type ProgressEvent struct {
	event.Event
}

type ProgressEventFrom interface {
	ProgressEvent_() ProgressEvent
}

func (p ProgressEvent) ProgressEvent_() ProgressEvent {
	return p
}

func New(typeevent string, opts ...map[string]interface{}) (ProgressEvent, error) {

	var p ProgressEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if pei := GetInterface(); !pei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(opts) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(opts[0]))
		}
		if obj, err = baseobject.New(pei, arrayJS...); err == nil {
			p.BaseObject = p.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return p, err
}

func NewFromJSObject(obj js.Value) (ProgressEvent, error) {
	var p ProgressEvent
	var err error
	if pei := GetInterface(); !pei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(pei) {
				p.BaseObject = p.SetObject(obj)

			} else {
				err = ErrNotAnProgressEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return p, err
}

func (p ProgressEvent) LengthComputable() (bool, error) {
	return p.GetAttributeBool("lengthComputable")
}

func (p ProgressEvent) Loaded() (int, error) {
	return p.GetAttributeInt("loaded")
}

func (p ProgressEvent) Total() (int, error) {
	return p.GetAttributeInt("total")
}
