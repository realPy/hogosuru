package event

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`event=new Event("test",{"bubbles":true, "cancelable":true, "composed":true})
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if e, err := New("test", map[string]interface{}{"bubbles": true, "cancelable": true, "composed": true}); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object Event]", e.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "event"); testingutils.AssertErr(t, err) {
		if event, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Event]", event.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Bubbles", "resultattempt": true},
	{"method": "Cancelable", "resultattempt": true},
	{"method": "Composed", "resultattempt": true},
	{"method": "EventPhase", "resultattempt": 0},
	{"method": "Type", "resultattempt": "test"},
	{"method": "IsTrusted", "resultattempt": false},
	{"method": "Target", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "CurrentTarget", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "PreventDefault", "type": "error", "resultattempt": nil},
	{"method": "StopImmediatePropagation", "type": "error", "resultattempt": nil},
	{"method": "StopPropagation", "type": "error", "resultattempt": nil},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "event"); testingutils.AssertErr(t, err) {

		if event, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, event, result)
			}

		}

	}
}
