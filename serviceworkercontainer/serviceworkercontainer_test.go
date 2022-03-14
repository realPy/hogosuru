package serviceworkercontainer

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`sw=window.navigator.serviceWorker`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "sw"); testingutils.AssertErr(t, err) {
		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "ServiceWorkerContainer", nav.ConstructName_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Controller", "type": "error", "resultattempt": ErrControllerNotDefined},
	{"method": "Ready", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "GetRegistration", "args": []interface{}{"url"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Register", "args": []interface{}{"url"}, "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "sw"); testingutils.AssertErr(t, err) {

		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, nav, result)
			}

		}

	}
}
