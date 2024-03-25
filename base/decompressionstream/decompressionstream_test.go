package decompressionstream

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`dec=new DecompressionStream("gzip")`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "dec"); testingutils.AssertErr(t, err) {
		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "DecompressionStream", nav.ConstructName_())

		}
	}

	decstream, err := New("gzip")
	testingutils.AssertExpect(t, nil, err)
	testingutils.AssertExpect(t, "DecompressionStream", decstream.ConstructName_())
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Readable", "type": "constructnamechecking", "resultattempt": "ReadableStream"},
	{"method": "Writable", "type": "constructnamechecking", "resultattempt": "WritableStream"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "dec"); testingutils.AssertErr(t, err) {

		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, nav, result)
			}

		}

	}
}
