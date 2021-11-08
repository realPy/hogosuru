package dragevent

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`devent=new DragEvent("test",{dataTransfer:new DataTransfer()})
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if k, err := New("test"); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object DragEvent]", k.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "devent"); testingutils.AssertErr(t, err) {
		if mevent, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object DragEvent]", mevent.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "DataTransfer", "type": "constructnamechecking", "resultattempt": "DataTransfer"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "devent"); testingutils.AssertErr(t, err) {

		if mevent, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, mevent, result)
			}

		}

	}
}
