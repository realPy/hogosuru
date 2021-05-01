package object

import (
	"github.com/realPy/jswasm/js"
)

type MessageEvent struct {
	Data        js.Value
	Source      js.Value
	Origin      js.Value
	LastEventId js.Value
}

func NewMessageEvent(obj js.Value) (MessageEvent, error) {
	var m MessageEvent
	if String(obj) == "[object MessageEvent]" {
		m.Data, _ = obj.GetWithErr("data")
		m.Source, _ = obj.GetWithErr("source")
		m.Origin, _ = obj.GetWithErr("origin")
		m.LastEventId, _ = obj.GetWithErr("lastEventId")

	}
	return m, ErrNotAnMEv
}
func DataFromMessageEvent(obj js.Value) (js.Value, error) {

	if String(obj) == "[object MessageEvent]" {
		return obj.GetWithErr("data")
	}

	return js.Value{}, ErrNotAnMEv
}
