package number

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
func TestBaseObjectString(t *testing.T) {

	var obj js.Value
	var err error

	t.Run("1 is Int", func(t *testing.T) {
		baseobject.Eval("intvalue=1")

		if obj, err = baseobject.Get(js.Global(), "intvalue"); testingutils.AssertErr(t, err) {

			if b, err := IsInteger(obj); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, true, b)
			}

		}

	})

	t.Run("1.3 is Float", func(t *testing.T) {
		baseobject.Eval("intvalue=1.3")

		if obj, err = baseobject.Get(js.Global(), "intvalue"); testingutils.AssertErr(t, err) {

			if b, err := IsInteger(obj); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, false, b)
			}

		}

	})

	t.Run("str is not int", func(t *testing.T) {
		baseobject.Eval("intvalue='hello'")

		if obj, err = baseobject.Get(js.Global(), "intvalue"); testingutils.AssertErr(t, err) {

			if b, err := IsInteger(obj); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, false, b)
			}

		}

	})

}
