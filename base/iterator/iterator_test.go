package iterator

import (
	"errors"
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`str = "hp";
	it=str[Symbol.iterator]()
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "it"); testingutils.AssertErr(t, err) {

		if it, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "Iterator", it.ConstructName_())
		}

	}

}

func TestNext(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "it"); testingutils.AssertErr(t, err) {

		if it, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			_, val, err := it.Next()
			testingutils.AssertExpect(t, "h", val)
			testingutils.AssertExpect(t, true, errors.Is(err, nil))
			_, val, err = it.Next()
			testingutils.AssertExpect(t, "p", val)
			testingutils.AssertExpect(t, true, errors.Is(err, nil))
			_, val, err = it.Next()
			testingutils.AssertExpect(t, true, errors.Is(err, EOI))
		}

	}

}
