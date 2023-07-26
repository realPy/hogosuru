package validitystate

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`b=document.createElement("button")
	validity=b.validity
	`)

	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "validity"); testingutils.AssertErr(t, err) {
		if mevent, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object ValidityState]", mevent.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "BadInput", "resultattempt": false},
	{"method": "CustomError", "resultattempt": false},
	{"method": "PatternMismatch", "resultattempt": false},
	{"method": "RangeOverflow", "resultattempt": false},
	{"method": "RangeUnderflow", "resultattempt": false},
	{"method": "StepMismatch", "resultattempt": false},
	{"method": "TooLong", "resultattempt": false},
	{"method": "TooShort", "resultattempt": false},
	{"method": "TypeMismatch", "resultattempt": false},
	{"method": "Valid", "resultattempt": true},
	{"method": "ValueMissing", "resultattempt": false},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "validity"); testingutils.AssertErr(t, err) {

		if mevent, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, mevent, result)
			}

		}

	}
}
