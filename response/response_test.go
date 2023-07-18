package response

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	var io chan bool = make(chan bool)

	js.Global().Set("waiting", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		io <- true
		return nil
	}))

	baseobject.Eval(`resp = fetch("http://localhost/get")
	fetch('http://localhost/get')
	.then(function(response) {

		resp=response
		waiting()
	})
`)
	<-io
	m.Run()
}

func TestNew(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object Response]", d.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("r=new Response()")

	if obj, err := baseobject.Get(js.Global(), "r"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Response]", d.ToString_())

		}
	}

}

func TestStatus(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())
			if s, err := response.Status(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, 200, s)
			}
		}
	}

}

func TestStatusText(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if s, err := response.StatusText(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "", s)
			}
		}
	}

}

func TestRedirected(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if b, err := response.Redirected(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, false, b)
			}

		}
	}

}

func TestOk(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if b, err := response.Ok(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, true, b)
			}

		}
	}

}

func TestType(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if typef, err := response.Type(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "cors", typef)
			}

		}
	}

}

func TestUrl(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if url, err := response.Url(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "http://localhost/get", url)
			}

		}
	}

}

func TestBody(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if stream, err := response.Body(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object ReadableStream]", stream.ToString_())
			}

		}
	}

}

func TestBodyUsed(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if b, err := response.BodyUsed(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, false, b)
			}

		}
	}

}

func TestClone(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if clone, err := response.Clone(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object Response]", clone.ToString_())
			}

		}
	}

}

func TestError(t *testing.T) {

	if response, err := Error(); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "[object Response]", response.ToString_())

		if resptype, err := response.Type(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "error", resptype)
		}

	}

}

func TestArrayBuffer(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if p, err := response.ArrayBuffer(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object Promise]", p.ToString_())
			}

		}
	}

}

func TestBlob(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if p, err := response.Blob(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object Promise]", p.ToString_())
			}

		}
	}

}

func TestJson(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if p, err := response.Json(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object Promise]", p.ToString_())
			}

		}
	}

}

func TestText(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "resp"); testingutils.AssertErr(t, err) {

		if response, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Response]", response.ToString_())

			if p, err := response.Text(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object Promise]", p.ToString_())
			}

		}
	}

}
