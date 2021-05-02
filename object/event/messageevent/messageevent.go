package messageevent

import (
	"errors"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

var (

	//ErrNotAnMEv ErrNotAnMEv error
	ErrNotAnMEv = errors.New("The given value must be an Message Event")
)

type MessageEvent struct {
	Data        js.Value
	Source      js.Value
	Origin      js.Value
	LastEventID js.Value
}

/*
func NewMessageEvent(obj js.Value) (MessageEvent, error) {
	var m MessageEvent
	if object.String(obj) == "[object MessageEvent]" {
		m.Data, _ = obj.GetWithErr("data")
		m.Source, _ = obj.GetWithErr("source")
		m.Origin, _ = obj.GetWithErr("origin")
		m.LastEventID, _ = obj.GetWithErr("lastEventId")

	}
	return m, ErrNotAnMEv
}*/

func NewMessageEvent(obj js.Value) (object.GOMap, error) {
	var m map[string]object.GOValue = make(map[string]object.GOValue)

	if object.String(obj) == "[object MessageEvent]" {
		if value, err := obj.GetWithErr("data"); err == nil {
			m["data"] = object.NewGOValue(value)
		}
		if value, err := obj.GetWithErr("source"); err == nil {
			m["source"] = object.NewGOValue(value)
		}
		if value, err := obj.GetWithErr("origin"); err == nil {
			m["origin"] = object.NewGOValue(value)
		}

		if value, err := obj.GetWithErr("lastEventId"); err == nil {
			m["lastEventId"] = object.NewGOValue(value)
		}
		return object.NewGoMap(m), nil
	}
	return object.NewGoMap(m), ErrNotAnMEv
}

func DataFromMessageEvent(obj js.Value) (js.Value, error) {

	if object.String(obj) == "[object MessageEvent]" {
		return obj.GetWithErr("data")
	}

	return js.Value{}, ErrNotAnMEv
}
