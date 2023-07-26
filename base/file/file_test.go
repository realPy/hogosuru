package file

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

/*
file = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })
*/

func TestNew(t *testing.T) {

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object File]", f.ToString_())

	}

}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("file = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })")

	if obj, err := baseobject.Get(js.Global(), "file"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object File]", d.ToString_())

		}
	}

}

func TestName(t *testing.T) {

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); testingutils.AssertErr(t, err) {

		if name, err := f.Name(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "chucknorris.png", name)
		}

	}

}

func TestType(t *testing.T) {

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); testingutils.AssertErr(t, err) {

		if typefile, err := f.Type(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "image/png", typefile)
		}

	}

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png"); testingutils.AssertErr(t, err) {

		if typefile, err := f.Type(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "", typefile)
		}

	}

}

func TestLastModifiedDate(t *testing.T) {

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); testingutils.AssertErr(t, err) {

		if lastmodified, err := f.LastModifiedDate(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "Date", lastmodified.ConstructName_())
		}

	}

}

func TestLastModified(t *testing.T) {

	if f, err := New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); testingutils.AssertErr(t, err) {

		if lastmodified, err := f.LastModified(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, true, lastmodified > 0)
		}

	}

}
