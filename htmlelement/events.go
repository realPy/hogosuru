package htmlelement

import (
	"github.com/realPy/hogosuru/animationevent"
	"github.com/realPy/hogosuru/event"
)

func (h HtmlElement) addAnimationEventListener(name string, handler func(a animationevent.AnimationEvent)) error {

	var err, err2 error

	err = h.AddEventListener(name, func(e event.Event) {
		var a animationevent.AnimationEvent
		if a, err2 = animationevent.NewFromJSObject(e.JSObject()); err2 == nil {
			handler(a)
		}
	})

	if err2 != nil {
		return err2
	}

	return err
}

func (h HtmlElement) OnAnimationCancel(handler func(a animationevent.AnimationEvent)) error {

	return h.addAnimationEventListener("animationcancel", handler)
}

func (h HtmlElement) OnAnimationEnd(handler func(a animationevent.AnimationEvent)) error {

	return h.addAnimationEventListener("animationend", handler)
}

func (h HtmlElement) OnAnimationStart(handler func(a animationevent.AnimationEvent)) error {

	return h.addAnimationEventListener("animationstart", handler)
}

func (h HtmlElement) OnAnimationIteration(handler func(a animationevent.AnimationEvent)) error {

	return h.addAnimationEventListener("animationiteration", handler)
}

func (h HtmlElement) OnBeforeInput(handler func(e event.Event)) error {

	return h.AddEventListener("beforeinput", handler)
}

func (h HtmlElement) OnChange(handler func(e event.Event)) error {

	return h.AddEventListener("change", handler)
}

func (h HtmlElement) OnGotPointerCapture(handler func(e event.Event)) error {

	return h.AddEventListener("gotpointercapture", handler)
}

func (h HtmlElement) OnInput(handler func(e event.Event)) error {

	return h.AddEventListener("input", handler)
}

func (h HtmlElement) OnLostPointerCapture(handler func(e event.Event)) error {

	return h.AddEventListener("lostpointercapture", handler)
}

func (h HtmlElement) OnPointerCancel(handler func(e event.Event)) error {

	return h.AddEventListener("pointercancel", handler)
}

func (h HtmlElement) OnPointerDown(handler func(e event.Event)) error {

	return h.AddEventListener("pointerdown", handler)
}

func (h HtmlElement) OnPointerEnter(handler func(e event.Event)) error {

	return h.AddEventListener("pointerenter", handler)
}

func (h HtmlElement) OnPointerLeave(handler func(e event.Event)) error {

	return h.AddEventListener("pointerleave", handler)
}

func (h HtmlElement) OnPointerMove(handler func(e event.Event)) error {

	return h.AddEventListener("pointermove", handler)
}

func (h HtmlElement) OnPointerOut(handler func(e event.Event)) error {

	return h.AddEventListener("pointerout", handler)
}

func (h HtmlElement) OnPointerOver(handler func(e event.Event)) error {

	return h.AddEventListener("pointerover", handler)
}

func (h HtmlElement) OnPointerUp(handler func(e event.Event)) error {

	return h.AddEventListener("pointerup", handler)
}

func (h HtmlElement) OnTransitionCancel(handler func(e event.Event)) error {

	return h.AddEventListener("transitioncancel", handler)
}

func (h HtmlElement) OnTransitionEnd(handler func(e event.Event)) error {

	return h.AddEventListener("transitionend", handler)
}

func (h HtmlElement) OnTransitionRun(handler func(e event.Event)) error {

	return h.AddEventListener("transitionrun", handler)
}
func (h HtmlElement) OnTransitionStart(handler func(e event.Event)) error {

	return h.AddEventListener("transitionstart", handler)
}
