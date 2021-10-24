package abortcontroller

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestAbort(t *testing.T) {

	var isAborted bool = false
	if a, err := New(); testingutils.AssertErr(t, err) {

		if as, err := a.Signal(); testingutils.AssertErr(t, err) {

			as.OnAbort(func(e event.Event) {
				isAborted = true
			})
			a.Abort()

			testingutils.AssertExpect(t, true, isAborted)

		}
	}
}
