package navigator

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`nav=window.navigator`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "nav"); testingutils.AssertErr(t, err) {
		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "Navigator", nav.ConstructName_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "CookieEnabled", "resultattempt": true},
	//	{"method": "DeviceMemory", "resultattempt": float64(8.0)},
	{"method": "UserAgent", "type": "contains", "resultattempt": "HeadlessChrome"},
	//	{"method": "Language", "resultattempt": "fr"},
	{"method": "Vendor", "resultattempt": "Google Inc."},
	{"method": "JavaEnabled", "resultattempt": false},
	{"method": "Permissions", "type": "constructnamechecking", "resultattempt": "Permissions"},
	{"method": "Clipboard", "type": "constructnamechecking", "resultattempt": "Clipboard"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "nav"); testingutils.AssertErr(t, err) {

		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, nav, result)
			}

		}

	}
}
