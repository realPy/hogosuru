package animationevent

//https://developer.mozilla.org/fr/docs/Web/API/AnimationEvent

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var animationeventinterface js.Value

// AnimationEvent AnimationEvent struct
type AnimationEvent struct {
	event.Event
}

type AnimationEventFrom interface {
	AnimationEvent() AnimationEvent
}

func (a AnimationEvent) AnimationEvent_() AnimationEvent {
	return a
}

// GetInterface get the JS interface animationEvent
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if animationeventinterface, err = baseobject.Get(js.Global(), "AnimationEvent"); err != nil {
			animationeventinterface = js.Undefined()

		}

		baseobject.Register(animationeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return animationeventinterface
}

func NewFromJSObject(obj js.Value) (AnimationEvent, error) {
	var a AnimationEvent
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			if obj.InstanceOf(di) {
				a.BaseObject = a.SetObject(obj)

			} else {
				err = ErrNotAnAnimationEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return a, err
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
