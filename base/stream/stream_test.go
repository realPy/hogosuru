package stream

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

	if s, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object ReadableStream]", s.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("r=new ReadableStream()")

	if obj, err := baseobject.Get(js.Global(), "r"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object ReadableStream]", d.ToString_())

		}
	}

}

func TestLocked(t *testing.T) {
	if s, err := New(); testingutils.AssertErr(t, err) {

		if locked, err := s.Locked(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, false, locked)
		}

	}
}

func TestCancel(t *testing.T) {
	if s, err := New(); testingutils.AssertErr(t, err) {

		if pcancel, err := s.Cancel(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Promise]", pcancel.ToString_())
		}

	}
}

func TestGetReader(t *testing.T) {
	if s, err := New(); testingutils.AssertErr(t, err) {

		if reader, err := s.GetReader(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object ReadableStreamDefaultReader]", reader.ToString_())
		}
		//doc say that the stream is locked when get reader so check it
		if locked, err := s.Locked(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, true, locked)
		}
	}
}

func TestTee(t *testing.T) {
	if s, err := New(); testingutils.AssertErr(t, err) {

		if a, err := s.Tee(); testingutils.AssertErr(t, err) {

			if testingutils.AssertExpect(t, 2, len(a)) {

				for i := 0; i < 2; i++ {
					testingutils.AssertExpect(t, "[object ReadableStream]", a[i].ToString_())
				}

			}

		}
	}
}
