package serviceworkerglobalscope

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/event"
)

func (w ServiceWorkerGlobalScope) OnActivate(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("activate", handler)
}

func (w ServiceWorkerGlobalScope) OnContentDelete(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("contentdelete", handler)
}

func (w ServiceWorkerGlobalScope) OnFetch(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("fetch", handler)
}

func (w ServiceWorkerGlobalScope) OnInstall(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("install", handler)
}

func (w ServiceWorkerGlobalScope) OnMessage(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("message", handler)
}

func (w ServiceWorkerGlobalScope) OnPeriodicSync(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("periodicsync", handler)
}

func (w ServiceWorkerGlobalScope) OnPush(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("push", handler)
}

func (w ServiceWorkerGlobalScope) OnPushSubscriptionChange(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("onpushsubscriptionchange", handler)
}
