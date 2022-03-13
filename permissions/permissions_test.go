package permissions

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {

	baseobject.SetSyscall()

	baseobject.Eval(`perm=navigator.permissions`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "perm"); testingutils.AssertErr(t, err) {
		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "Permissions", nav.ConstructName_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Query", "args": []interface{}{map[string]interface{}{"name": "clipboard-read"}}, "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "perm"); testingutils.AssertErr(t, err) {

		if clip, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, clip, result)
			}

		}

	}
}
