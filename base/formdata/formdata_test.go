package formdata

import (
	"errors"
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/htmlformelement"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()

	m.Run()
}

func TestNew(t *testing.T) {

	if f, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "[object FormData]", f.ToString_())
	}
	baseobject.Eval(`f= document.createElement("form")
	intext=document.createElement("input")
	intext.type="text"
	intext.name="hello"
	intext.value="world"
	f.appendChild(intext)
	`)

	if obj, err := baseobject.Get(js.Global(), "f"); testingutils.AssertErr(t, err) {

		if form, err := htmlformelement.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if f, err := New(form); testingutils.AssertErr(t, err) {

				if v, err := f.Get("hello"); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, "world", v)
				}

				_, err := f.Get("hell")
				testingutils.AssertExpect(t, true, errors.Is(ErrNotAFormValueNotFound, err))

			}
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval(`f= new FormData()
	`)

	if obj, err := baseobject.Get(js.Global(), "f"); testingutils.AssertErr(t, err) {

		if f, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object FormData]", f.ToString_())
		}

	}
}

func TestGet(t *testing.T) {

	baseobject.Eval(`f= new FormData()
	f.append("hello","world")
	`)

	if obj, err := baseobject.Get(js.Global(), "f"); testingutils.AssertErr(t, err) {

		if f, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := f.Get("hello"); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "world", v)
			}

			_, err := f.Get("hell")
			testingutils.AssertExpect(t, true, errors.Is(ErrNotAFormValueNotFound, err))

		}

	}

}

func TestAppend(t *testing.T) {

	if f, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, f.Append("data", "test"))
		if v, err := f.Get("data"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "test", v)
		}

	}
}

func TestDelete(t *testing.T) {

	if f, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, f.Append("data", "test"))

		testingutils.AssertErr(t, f.Delete("data"))

		_, err := f.Get("data")
		testingutils.AssertExpect(t, true, errors.Is(ErrNotAFormValueNotFound, err))

	}
}

func TestEntries(t *testing.T) {
	if f, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, f.Append("data", "test"))
		testingutils.AssertErr(t, f.Append("hello", "world"))
		if it, err := f.Entries(); testingutils.AssertErr(t, err) {

			var expecdata map[string]string = map[string]string{"data": "test", "hello": "world"}

			for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {

				_, ok := expecdata[index.(string)]
				testingutils.AssertExpect(t, true, ok)

				testingutils.AssertExpect(t, expecdata[index.(string)], value.(string))

			}
		}
	}

}

func TestHas(t *testing.T) {

	if f, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, f.Append("data", "test"))

		if ok, err := f.Has("data"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, true, ok)

		}

		if ok, err := f.Has("c"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, false, ok)

		}

	}

}

func TestKeys(t *testing.T) {
	if f, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, f.Append("data", "test"))
		testingutils.AssertErr(t, f.Append("hello", "world"))
		if it, err := f.Keys(); testingutils.AssertErr(t, err) {

			var expecdata map[string]int = map[string]int{"data": 0, "hello": 1}
			var i int
			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

				v, ok := expecdata[value.(string)]
				testingutils.AssertExpect(t, true, ok)
				testingutils.AssertExpect(t, i, v)
				i++
			}
			testingutils.AssertExpect(t, 2, i)

		}
	}

}

func TestSet(t *testing.T) {
	if f, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, f.Append("data", "test"))
		testingutils.AssertErr(t, f.Append("hello", "world"))
		testingutils.AssertErr(t, f.Set("hello", "you"))

		if v, err := f.Get("hello"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "you", v)
		}

	}

}

func TestValues(t *testing.T) {
	if f, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertErr(t, f.Append("data", "test"))
		testingutils.AssertErr(t, f.Append("hello", "world"))
		if it, err := f.Values(); testingutils.AssertErr(t, err) {

			var expecdata map[string]int = map[string]int{"test": 0, "world": 1}
			var i int
			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

				v, ok := expecdata[value.(string)]
				testingutils.AssertExpect(t, true, ok)
				testingutils.AssertExpect(t, i, v)
				i++
			}
			testingutils.AssertExpect(t, 2, i)

		}
	}

}
