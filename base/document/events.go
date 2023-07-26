package document

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/dragevent"
	"github.com/realPy/hogosuru/base/event"
)

func (d Document) OnCopy(handler func(e event.Event)) (js.Func, error) {

	return d.AddEventListener("copy", handler)
}

func (d Document) OnCut(handler func(e event.Event)) (js.Func, error) {

	return d.AddEventListener("cut", handler)
}

func (d Document) OnPaste(handler func(e event.Event)) (js.Func, error) {

	return d.AddEventListener("paste", handler)
}

func (d Document) OnDrag(handler func(e dragevent.DragEvent)) (js.Func, error) {

	return d.AddEventListener("drag", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(dragevent.DragEventFrom); ok {
				handler(m.DragEvent_())
			}
		}
	})
}

func (d Document) OnDragStart(handler func(e dragevent.DragEvent)) (js.Func, error) {

	return d.AddEventListener("dragstart", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(dragevent.DragEventFrom); ok {
				handler(m.DragEvent_())
			}
		}
	})
}

func (d Document) OnDragEnd(handler func(e dragevent.DragEvent)) (js.Func, error) {

	return d.AddEventListener("dragend", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(dragevent.DragEventFrom); ok {
				handler(m.DragEvent_())
			}
		}
	})
}

func (d Document) OnDragOver(handler func(e dragevent.DragEvent)) (js.Func, error) {

	return d.AddEventListener("dragover", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(dragevent.DragEventFrom); ok {
				handler(m.DragEvent_())
			}
		}
	})
}

func (d Document) OnDragEnter(handler func(e dragevent.DragEvent)) (js.Func, error) {

	return d.AddEventListener("dragenter", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(dragevent.DragEventFrom); ok {
				handler(m.DragEvent_())
			}
		}
	})
}

func (d Document) OnDragLeave(handler func(e dragevent.DragEvent)) (js.Func, error) {

	return d.AddEventListener("dragleave", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(dragevent.DragEventFrom); ok {
				handler(m.DragEvent_())
			}
		}
	})
}

func (d Document) OnDrop(handler func(e dragevent.DragEvent)) (js.Func, error) {

	return d.AddEventListener("drop", func(e event.Event) {

		if globalObj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := globalObj.(dragevent.DragEventFrom); ok {
				handler(m.DragEvent_())
			}
		}
	})
}

func (d Document) OnTouchCancel(handler func(e event.Event)) (js.Func, error) {

	return d.AddEventListener("touchcancel", handler)
}

func (d Document) OnTouchEnd(handler func(e event.Event)) (js.Func, error) {

	return d.AddEventListener("touchend", handler)
}

func (d Document) OnTouchMove(handler func(e event.Event)) (js.Func, error) {

	return d.AddEventListener("touchmove", handler)
}

func (d Document) OnTouchStart(handler func(e event.Event)) (js.Func, error) {

	return d.AddEventListener("touchstart", handler)
}

func (d Document) OnScroll(handler func(e event.Event)) (js.Func, error) {

	return d.AddEventListener("scroll", handler)
}
