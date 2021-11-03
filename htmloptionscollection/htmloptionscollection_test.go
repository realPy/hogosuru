package htmloptionscollection

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`s=document.createElement("select")
	options=s.options`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "options"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLOptionsCollection", b.ConstructName_())
		}

	}
}

var getterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Length", "resultattempt": 0},
}

func TestGetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "options"); testingutils.AssertErr(t, err) {

		if area, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range getterAttempt {
				testingutils.InvokeCheck(t, area, result)
			}

		}

	}
}
