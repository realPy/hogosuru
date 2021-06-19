package element

import "github.com/realPy/hogosuru/event"

func (e Element) OnAuxClick(handler func(e event.Event)) error {

	return e.AddEventListener("auxclick", handler)
}

func (e Element) OnBlur(handler func(e event.Event)) error {

	return e.AddEventListener("blur", handler)
}

func (e Element) OnClick(handler func(e event.Event)) error {

	return e.AddEventListener("click", handler)
}

func (e Element) OnCompositionEnd(handler func(e event.Event)) error {

	return e.AddEventListener("compositionend", handler)
}

func (e Element) OnCompositionStart(handler func(e event.Event)) error {

	return e.AddEventListener("compositionstart", handler)
}

func (e Element) OnCompositionUpdate(handler func(e event.Event)) error {

	return e.AddEventListener("compositionupdate", handler)
}

func (e Element) OnContextMenu(handler func(e event.Event)) error {

	return e.AddEventListener("contextmenu", handler)
}

func (e Element) OnCopy(handler func(e event.Event)) error {

	return e.AddEventListener("copy", handler)
}

func (e Element) OnCut(handler func(e event.Event)) error {

	return e.AddEventListener("cut", handler)
}

func (e Element) OnDblClick(handler func(e event.Event)) error {

	return e.AddEventListener("dblclick", handler)
}

func (e Element) OnError(handler func(e event.Event)) error {

	return e.AddEventListener("error", handler)
}

func (e Element) OnFocus(handler func(e event.Event)) error {

	return e.AddEventListener("focus", handler)
}

func (e Element) OnFocusIn(handler func(e event.Event)) error {

	return e.AddEventListener("focusin", handler)
}

func (e Element) OnFocusOut(handler func(e event.Event)) error {

	return e.AddEventListener("focusout", handler)
}

func (e Element) OnFullScreenChange(handler func(e event.Event)) error {

	return e.AddEventListener("fullscreenchange", handler)
}

func (e Element) OnFullScreenError(handler func(e event.Event)) error {

	return e.AddEventListener("fullscreenerror", handler)
}

func (e Element) OnKeyDown(handler func(e event.Event)) error {

	return e.AddEventListener("keydown", handler)
}

func (e Element) OnKeyUp(handler func(e event.Event)) error {

	return e.AddEventListener("keyup", handler)
}

func (e Element) OnMouseDown(handler func(e event.Event)) error {

	return e.AddEventListener("mousedown", handler)
}

func (e Element) OnMouseEnter(handler func(e event.Event)) error {

	return e.AddEventListener("mouseenter", handler)
}

func (e Element) OnMouseLeave(handler func(e event.Event)) error {

	return e.AddEventListener("mouseleave", handler)
}

func (e Element) OnMouseMove(handler func(e event.Event)) error {

	return e.AddEventListener("mousemove", handler)
}

func (e Element) OnMouseOut(handler func(e event.Event)) error {

	return e.AddEventListener("mouseout", handler)
}

func (e Element) OnMouseOver(handler func(e event.Event)) error {

	return e.AddEventListener("mouseover", handler)
}

func (e Element) OnMouseUp(handler func(e event.Event)) error {

	return e.AddEventListener("mouseup", handler)
}

func (e Element) OnPaste(handler func(e event.Event)) error {

	return e.AddEventListener("paste", handler)
}

func (e Element) OnScroll(handler func(e event.Event)) error {

	return e.AddEventListener("scroll", handler)
}

func (e Element) OnSelect(handler func(e event.Event)) error {

	return e.AddEventListener("select", handler)
}

func (e Element) OnTouchCancel(handler func(e event.Event)) error {

	return e.AddEventListener("touchcancel", handler)
}

func (e Element) OnTouchEnd(handler func(e event.Event)) error {

	return e.AddEventListener("touchend", handler)
}

func (e Element) OnTouchMove(handler func(e event.Event)) error {

	return e.AddEventListener("touchmove", handler)
}

func (e Element) OnTouchStart(handler func(e event.Event)) error {

	return e.AddEventListener("touchstart", handler)
}

func (e Element) OnTouchWheel(handler func(e event.Event)) error {

	return e.AddEventListener("wheel", handler)
}
