package progressevent

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`pevent=new ProgressEvent("test",{lengthComputable:true,loaded:12345,total:1234567})
	`)

	m.Run()
}

func TestNew(t *testing.T) {

	if k, err := New("test"); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object ProgressEvent]", k.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "pevent"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object ProgressEvent]", k.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "LengthComputable", "resultattempt": true},
	{"method": "Loaded", "resultattempt": 12345},
	{"method": "Total", "resultattempt": 1234567},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "pevent"); testingutils.AssertErr(t, err) {

		if pevent, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, pevent, result)
			}

		}

	}
}
