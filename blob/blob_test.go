package blob

import (
	"testing"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
	"github.com/realPy/hogosuru/typedarray"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if a, err := New(); testingutils.AssertErr(t, err) {

		if s, err := a.Size(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, int64(0), s)

		}
	}
}

func TestNewWithArrayBuffer(t *testing.T) {

	if a, err := arraybuffer.New(8); testingutils.AssertErr(t, err) {
		if viewuint8, err := typedarray.NewInt8Array(a); testingutils.AssertErr(t, err) {
			viewuint8.Fill(7)

			if ab, err := NewWithArrayBuffer(a); testingutils.AssertErr(t, err) {

				if s, err := ab.Size(); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, int64(8), s)

				}

			}

		}
	}

}

func TestIsClosed(t *testing.T) {

	if a, err := New(); testingutils.AssertErr(t, err) {

		_, err := a.IsClosed()
		testingutils.AssertExpect(t, baseobject.ErrNotImplementedFunc, err)
	}
}

func TestClosed(t *testing.T) {

	if a, err := New(); testingutils.AssertErr(t, err) {

		err := a.Close()
		testingutils.AssertExpect(t, baseobject.ErrNotImplementedFunc, err)
	}
}

func TestSlice(t *testing.T) {
	astring := array.From_("Hello World")

	if struint8, err := typedarray.NewUint8ArrayFrom(astring); testingutils.AssertErr(t, err) {
		if b, err := struint8.(typedarray.Uint8Array).Buffer(); testingutils.AssertErr(t, err) {

			if ab, err := NewWithArrayBuffer(b); testingutils.AssertErr(t, err) {

				if s, err := ab.Size(); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, int64(11), s)

				}

			}
		}

	}

}
