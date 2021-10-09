package htmlbodyelement

import (
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

func (h HtmlBodyElement) OnAfterPrint(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onafterprint", handler)
}

func (h HtmlBodyElement) OnBeforePrint(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onbeforeprint", handler)
}

func (h HtmlBodyElement) OnBeforeUnload(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onbeforeunload", handler)
}

func (h HtmlBodyElement) OnHashChange(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onhashchange", handler)
}

func (h HtmlBodyElement) OnMessage(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onmessage", handler)
}

func (h HtmlBodyElement) OnMessageError(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onmessageerror", handler)
}

func (h HtmlBodyElement) OnOffline(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onoffline", handler)
}

func (h HtmlBodyElement) OnOnline(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("ononline", handler)
}

func (h HtmlBodyElement) OnPageHide(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onpagehide", handler)
}

func (h HtmlBodyElement) OnPageShow(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onpageshow", handler)
}

func (h HtmlBodyElement) OnPopState(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onpopstate", handler)
}

func (h HtmlBodyElement) OnRejectionHandled(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onrejectionhandled", handler)
}

func (h HtmlBodyElement) OnResize(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onresize", handler)
}

func (h HtmlBodyElement) OnStorage(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onstorage", handler)
}

func (h HtmlBodyElement) OnUnhandledRejection(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onunhandledrejection", handler)
}

func (h HtmlBodyElement) OnUnload(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("onunload", handler)
}
