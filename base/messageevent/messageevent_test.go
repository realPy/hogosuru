package messageevent

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`mevent=new MessageEvent("test",{data:Object(),origin:"or",lastEventId:"123"})
	`)

	m.Run()
}

func TestNew(t *testing.T) {

	if k, err := New("test"); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object MessageEvent]", k.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "mevent"); testingutils.AssertErr(t, err) {
		if mevent, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object MessageEvent]", mevent.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Data", "type": "constructnamechecking", "resultattempt": "Object"},
	{"method": "Source", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "Origin", "resultattempt": "or"},
	{"method": "LastEventId", "resultattempt": "123"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "mevent"); testingutils.AssertErr(t, err) {

		if mevent, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, mevent, result)
			}

		}

	}
}
