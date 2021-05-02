package progressevent

import (
	"errors"

	"github.com/realPy/jswasm/js"

	"github.com/realPy/jswasm/object"
)

var (
	//ErrNotAnMEv ErrNotAnMEv error
	ErrNotAnPEv = errors.New("The given value must be an Progress Event")
)

func NewProgressEvent(obj js.Value) (object.GOMap, error) {
	var m map[string]object.GOValue = make(map[string]object.GOValue)

	if object.String(obj) == "[object ProgressEvent]" {
		if value, err := obj.GetWithErr("loaded"); err == nil {
			m["loaded"] = object.NewGOValue(value)
		}
		if value, err := obj.GetWithErr("total"); err == nil {
			m["total"] = object.NewGOValue(value)
		}
		if value, err := obj.GetWithErr("lengthComputable"); err == nil {
			m["lengthComputable"] = object.NewGOValue(value)
		}
		return object.NewGoMap(m), nil
	}
	return object.NewGoMap(m), ErrNotAnPEv
}
