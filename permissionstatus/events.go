package permissionstatus

import (
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

func (p PermissionStatus) OnChange(handler func(e event.Event)) (js.Func, error) {

	return p.AddEventListener("change", handler)
}
