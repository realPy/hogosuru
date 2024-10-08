package abortsignal

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/event"
)

func (a AbortSignal) OnAbort(handler func(e event.Event)) (js.Func, error) {

	return a.AddEventListener("abort", handler)
}
