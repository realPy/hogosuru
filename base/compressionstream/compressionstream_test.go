package compressionstream

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`enc=new CompressionStream("gzip")`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "enc"); testingutils.AssertErr(t, err) {
		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "CompressionStream", nav.ConstructName_())

		}
	}

	encstream, err := New("gzip")
	testingutils.AssertExpect(t, nil, err)
	testingutils.AssertExpect(t, "CompressionStream", encstream.ConstructName_())
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Readable", "type": "constructnamechecking", "resultattempt": "ReadableStream"},
	{"method": "Writable", "type": "constructnamechecking", "resultattempt": "WritableStream"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "enc"); testingutils.AssertErr(t, err) {

		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, nav, result)
			}

		}

	}
}
