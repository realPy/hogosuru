package abortcontroller

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if a, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object AbortController]", a.ToString_())

	}
}
func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("abortctrl=new AbortController()")

	if obj, err := baseobject.Get(js.Global(), "abortctrl"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object AbortController]", d.ToString_())

		}
	}

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
