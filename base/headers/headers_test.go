package headers

import (
	"errors"
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestNew(t *testing.T) {

	if h, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "Headers", h.ConstructName_())
	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval(`h= new Headers()
	`)

	if obj, err := baseobject.Get(js.Global(), "h"); testingutils.AssertErr(t, err) {

		if h, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "Headers", h.ConstructName_())
		}

	}
}

func TestAppend(t *testing.T) {

	if h, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, h.Append("X-custom", "1234"))
		if v, err := h.Get("X-custom"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "1234", v)
		}
	}

}

func TestDelete(t *testing.T) {

	if h, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, h.Append("X-custom", "1234"))

		testingutils.AssertErr(t, h.Delete("X-custom"))

		_, err := h.Get("X-custom")

		testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))
	}

}

func TestHas(t *testing.T) {

	if h, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, h.Append("X-custom", "1234"))

		if b, err := h.Has("X-custom"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, true, b)
		}
		testingutils.AssertErr(t, h.Delete("X-custom"))

		if b, err := h.Has("X-custom"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, false, b)
		}

	}

}

func TestKeys(t *testing.T) {

	if h, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, h.Append("X-custom", "1234"))

		if it, err := h.Keys(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Headers Iterator]", it.ToString_())
		}
	}
}

func TestValues(t *testing.T) {

	if h, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, h.Append("X-custom", "1234"))

		if it, err := h.Keys(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Headers Iterator]", it.ToString_())
		}
	}
}

func TestSet(t *testing.T) {

	if h, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, h.Append("X-custom", "1234"))

		testingutils.AssertErr(t, h.Set("X-custom", "4567"))
		if v, err := h.Get("X-custom"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "4567", v)
		}

	}
}

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
