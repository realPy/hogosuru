package messageevent

import (
	"errors"

	"github.com/realPy/hogosuru/js"
	"github.com/realPy/hogosuru/object"
)

var (

	//ErrNotAnMEv ErrNotAnMEv error
	ErrNotAnMEv = errors.New("The given value must be an Message Event")
)

type MessageEvent struct {
	object.Object
}

func NewFromJSObject(obj js.Value) (MessageEvent, error) {
	var m MessageEvent

	if object.String(obj) == "[object MessageEvent]" {
		m.Object = m.SetObject(obj)
		return m, nil
	}

	return m, ErrNotAMessageEvent
}

func (m MessageEvent) Data() (js.Value, error) {
	return m.JSObject().GetWithErr("data")
}

func (m MessageEvent) Source() (js.Value, error) {
	return m.JSObject().GetWithErr("source")
}
func (m MessageEvent) Origin() (string, error) {
	var err error
	var originObject js.Value

	if originObject, err = m.JSObject().GetWithErr("origin"); err == nil {
		return originObject.String(), nil
	}
	return "", err
}

func (m MessageEvent) LastEventId() (string, error) {
	var err error
	var originObject js.Value

	if originObject, err = m.JSObject().GetWithErr("lastEventId"); err == nil {
		return originObject.String(), nil
	}
	return "", err
}
