package customevent

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

func TestNew(t *testing.T) {

	if d, err := New("hello", "world"); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object CustomEvent]", d.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("customevent=new CustomEvent(\"hello\")")

	if obj, err := baseobject.Get(js.Global(), "customevent"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object CustomEvent]", d.ToString_())

		}
	}

}

func TestDetail(t *testing.T) {

	if d, err := New("hello", "world"); testingutils.AssertErr(t, err) {

		if detail, err := d.Detail(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "world", detail)
		}

	}
}
