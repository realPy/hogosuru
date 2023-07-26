package history

import (
	"errors"
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/objectmap"
	"github.com/realPy/hogosuru/testingutils"
)

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval(`h= window.history
	`)

	if obj, err := baseobject.Get(js.Global(), "h"); testingutils.AssertErr(t, err) {

		if h, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object History]", h.ToString_())
		}

	}
}

func TestState(t *testing.T) {

	baseobject.Eval(`h= window.history
	`)

	if obj, err := baseobject.Get(js.Global(), "h"); testingutils.AssertErr(t, err) {

		if h, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			_, err := h.State()

			testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))
			o, _ := objectmap.New(array.New_(array.New_("title", "teststate")))
			h.PushState(o, "hello")
			state, err := h.State()
			testingutils.AssertExpect(t, "[object Map]", state.(objectmap.ObjectMap).ObjectMap_().ToString_())
		}

	}

}

func TestLength(t *testing.T) {

	baseobject.Eval(`h= window.history
	`)

	if obj, err := baseobject.Get(js.Global(), "h"); testingutils.AssertErr(t, err) {

		if h, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if l, err := h.Length(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, true, l >= 1)
				o, _ := objectmap.New(array.New_(array.New_("title", "testLength1")))
				h.PushState(o, "hello")
				if l2, err := h.Length(); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, 1, l2-l)
				}
			}
		}

	}

}

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
