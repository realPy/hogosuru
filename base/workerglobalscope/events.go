package workerglobalscope

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/event"
)

func (w WorkerGlobalScope) OnError(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("error", handler)
}

func (w WorkerGlobalScope) OnLanguageChange(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("languagechange", handler)
}

func (w WorkerGlobalScope) OnOffline(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("offline", handler)
}

func (w WorkerGlobalScope) OnOnline(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("online", handler)
}
