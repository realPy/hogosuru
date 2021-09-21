package animationevent

//https://developer.mozilla.org/fr/docs/Web/API/AnimationEvent

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var animationeventinterface js.Value

//AnimationEvent AnimationEvent struct
type AnimationEvent struct {
	event.Event
}

type AnimationEventFrom interface {
	AnimationEvent() AnimationEvent
}

func (a AnimationEvent) AnimationEvent() AnimationEvent {
	return a
}

//GetInterface get the JS interface of animationEvent
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if animationeventinterface, err = js.Global().GetWithErr("AnimationEvent"); err != nil {
			animationeventinterface = js.Null()
		}
		baseobject.Register(animationeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return animationeventinterface
}

func NewFromJSObject(obj js.Value) (AnimationEvent, error) {
	var e AnimationEvent

	if di := GetInterface(); !di.IsNull() {
		if obj.InstanceOf(di) {
			e.BaseObject = e.SetObject(obj)
			return e, nil
		}
	}
	return e, ErrNotAnAnimationEvent
}

func (a AnimationEvent) AnimationName() (string, error) {
	return a.GetAttributeString("animationName")
}

func (a AnimationEvent) ElapsedTime() (float64, error) {
	return a.GetAttributeDouble("elapsedTime")
}

func (a AnimationEvent) PseudoElement() (string, error) {
	return a.GetAttributeString("pseudoElement")
}
