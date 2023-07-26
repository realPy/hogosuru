package blob

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/arraybuffer"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/typedarray"
	"github.com/realPy/hogosuru/testingutils"
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

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("blob1=new Blob()")

	if obj, err := baseobject.Get(js.Global(), "blob1"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Blob]", d.ToString_())

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

func TestNewWith2ArrayBuffer(t *testing.T) {

	if a, err := arraybuffer.New(8); testingutils.AssertErr(t, err) {
		if viewuint8, err := typedarray.NewInt8Array(a); testingutils.AssertErr(t, err) {
			viewuint8.Fill(7)

			astring := array.From_("Hello World")
			if struint8, err := typedarray.NewUint8ArrayFrom(astring); testingutils.AssertErr(t, err) {

				if appendblob, err := New(viewuint8, struint8); testingutils.AssertErr(t, err) {

					if s, err := appendblob.Size(); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, int64(19), s)

					}

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

		if b, err := struint8.Buffer(); testingutils.AssertErr(t, err) {

			if ab, err := NewWithArrayBuffer(b); testingutils.AssertErr(t, err) {

				if blob2, err := ab.Slice(0, 6); testingutils.AssertErr(t, err) {
					if s, err := blob2.Size(); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, int64(6), s)

					}
				}

			}
		}

	}

}

func TestStream(t *testing.T) {
	astring := array.From_("Hello World")

	if struint8, err := typedarray.NewUint8ArrayFrom(astring); testingutils.AssertErr(t, err) {

		if b, err := struint8.Buffer(); testingutils.AssertErr(t, err) {

			if ab, err := NewWithArrayBuffer(b); testingutils.AssertErr(t, err) {

				if stream, err := ab.Stream(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "[object ReadableStream]", stream.ToString_())
				}

			}
		}

	}

}

func TestArrayBuffer(t *testing.T) {
	astring := array.From_("Hello World")

	if struint8, err := typedarray.NewUint8ArrayFrom(astring); testingutils.AssertErr(t, err) {

		if b, err := struint8.Buffer(); testingutils.AssertErr(t, err) {

			if ab, err := NewWithArrayBuffer(b); testingutils.AssertErr(t, err) {

				if blobbuffer, err := ab.ArrayBuffer(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "[object Promise]", blobbuffer.ToString_())
				}

			}
		}

	}

}
func TestText(t *testing.T) {
	astring := array.From_("Hello World")

	if struint8, err := typedarray.NewUint8ArrayFrom(astring); testingutils.AssertErr(t, err) {

		if b, err := struint8.Buffer(); testingutils.AssertErr(t, err) {

			if ab, err := NewWithArrayBuffer(b); testingutils.AssertErr(t, err) {

				if blobtext, err := ab.Text(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "[object Promise]", blobtext.ToString_())
				}

			}
		}

	}

}
