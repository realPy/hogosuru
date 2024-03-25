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

	if s, err := NewReadableStream(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object ReadableStream]", s.ToString_())

	}

	if s, err := NewWritableStream(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object WritableStream]", s.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("r=new ReadableStream();w=new WritableStream();")

	if obj, err := baseobject.Get(js.Global(), "r"); testingutils.AssertErr(t, err) {
		if d, err := NewReadableStreamFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object ReadableStream]", d.ToString_())

		}
	}
	if obj, err := baseobject.Get(js.Global(), "w"); testingutils.AssertErr(t, err) {
		if d, err := NewWriteableStreamFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object WritableStream]", d.ToString_())

		}
	}

}

func TestLocked(t *testing.T) {
	if s, err := NewReadableStream(); testingutils.AssertErr(t, err) {

		if locked, err := s.Locked(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, false, locked)
		}

	}

	if s, err := NewWritableStream(); testingutils.AssertErr(t, err) {

		if locked, err := s.Locked(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, false, locked)
		}

	}
}

func TestCancelReadable(t *testing.T) {
	if s, err := NewReadableStream(); testingutils.AssertErr(t, err) {

		if pcancel, err := s.Cancel(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Promise]", pcancel.ToString_())
		}

	}
}

func TestAbortWritable(t *testing.T) {
	if w, err := NewWritableStream(); testingutils.AssertErr(t, err) {

		if pabort, err := w.Abort("i want"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Promise]", pabort.ToString_())
		}

	}
}

func TestCloseWritable(t *testing.T) {
	if w, err := NewWritableStream(); testingutils.AssertErr(t, err) {

		if pabort, err := w.Close(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Promise]", pabort.ToString_())
		}

	}
}

func TestGetReader(t *testing.T) {
	if s, err := NewReadableStream(); testingutils.AssertErr(t, err) {

		if reader, err := s.GetReader(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object ReadableStreamDefaultReader]", reader.ToString_())
		}
		//doc say that the stream is locked when get reader so check it
		if locked, err := s.Locked(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, true, locked)
		}
	}
}

func TestTeeReadable(t *testing.T) {
	if s, err := NewReadableStream(); testingutils.AssertErr(t, err) {

		if a, err := s.Tee(); testingutils.AssertErr(t, err) {

			if testingutils.AssertExpect(t, 2, len(a)) {

				for i := 0; i < 2; i++ {
					testingutils.AssertExpect(t, "[object ReadableStream]", a[i].ToString_())
				}

			}

		}
	}
}
