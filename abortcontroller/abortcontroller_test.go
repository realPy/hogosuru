package abortcontroller

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestAbort(t *testing.T) {

	var isAborted bool = false
	if a, err := New(); err == nil {

		if as, err := a.Signal(); err == nil {

			as.OnAbort(func(e event.Event) {
				isAborted = true
			})
			a.Abort()

			if isAborted == false {
				t.Error("Must be aborted")

			}
		} else {
			t.Error(err.Error())
		}
	} else {
		t.Error(err.Error())
	}
}
