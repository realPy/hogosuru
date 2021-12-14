package urlsearchparams

import (
	"errors"
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestNew(t *testing.T) {

	if h, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "URLSearchParams", h.ConstructName_())
	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval(`u= new URLSearchParams()
	`)

	if obj, err := baseobject.Get(js.Global(), "u"); testingutils.AssertErr(t, err) {

		if h, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "URLSearchParams", h.ConstructName_())
		}

	}
}

func TestAppend(t *testing.T) {

	if u, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, u.Append("TestKey", "1234"))
		if v, err := u.Get("TestKey"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "1234", v)
		}
	}

}

func TestDelete(t *testing.T) {

	if u, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, u.Append("TestKey", "1234"))

		testingutils.AssertErr(t, u.Delete("TestKey"))

		_, err := u.Get("TestKey")

		testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))
	}

}

func TestHas(t *testing.T) {

	if u, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, u.Append("TestKey", "1234"))

		if b, err := u.Has("TestKey"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, true, b)
		}
		testingutils.AssertErr(t, u.Delete("TestKey"))

		if b, err := u.Has("TestKey"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, false, b)
		}

	}

}

func TestKeys(t *testing.T) {

	if u, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, u.Append("TestKey", "1234"))

		if it, err := u.Keys(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Iterator]", it.ToString_())
		}
	}
}

func TestValues(t *testing.T) {

	if u, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, u.Append("TestKey", "1234"))

		if it, err := u.Keys(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Iterator]", it.ToString_())
		}
	}
}

func TestSet(t *testing.T) {

	if u, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, u.Append("TestKey", "1234"))

		testingutils.AssertErr(t, u.Set("TestKey", "4567"))
		if v, err := u.Get("TestKey"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "4567", v)
		}

	}
}

func TestSort(t *testing.T) {

	if u, err := New("c=4&a=2&b=3&a=1"); testingutils.AssertErr(t, err) {

		if err := u.Sort(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "a=2&a=1&b=3&c=4", u.ToString_())
		}

	}
}

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
