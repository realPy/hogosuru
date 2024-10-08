package serviceworkerregistration

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/event"
)

func (s ServiceWorkerRegistration) OnUpdateFound(handler func(e event.Event)) (js.Func, error) {

	return s.AddEventListener("updatefound", handler)
}
