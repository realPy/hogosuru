package arraybuffer

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

	if a, err := New(8); testingutils.AssertErr(t, err) {

		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, int64(8), l)

		}
	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("ab = new ArrayBuffer()")

	if obj, err := baseobject.Get(js.Global(), "ab"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object ArrayBuffer]", d.ToString_())

		}
	}

}

func TestSlice(t *testing.T) {

	if a, err := New(32); testingutils.AssertErr(t, err) {

		if b, err := a.Slice(10); testingutils.AssertErr(t, err) {

			if l, err := b.ByteLength(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, int64(22), l)

			}

		}

		if b, err := a.Slice(10, 16); testingutils.AssertErr(t, err) {

			if l, err := b.ByteLength(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, int64(6), l)

			}

		}

	}
}

func TestIsView(t *testing.T) {

	baseobject.Eval("customuint16=new Uint16Array()")
	if obj, err := baseobject.Get(js.Global(), "customuint16"); testingutils.AssertErr(t, err) {
		if a, err := baseobject.NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if ok, err := IsView(a); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, true, ok)

			}
		}

	}
	baseobject.Eval("customuint16=\"string\"")
	if obj, err := baseobject.Get(js.Global(), "customuint16"); testingutils.AssertErr(t, err) {
		if a, err := baseobject.NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if ok, err := IsView(a); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, false, ok)
			}
		}

	}
}
