package clipboard

import (
	"fmt"
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/blob"
	"github.com/realPy/hogosuru/base/clipboarditem"
	"github.com/realPy/hogosuru/testingutils"
)

var clipitem clipboarditem.ClipboardItem

var methodsAttempt []map[string]interface{}

func TestMain(m *testing.M) {
	var err error
	var b blob.Blob
	baseobject.SetSyscall()

	if b, err = blob.New("{\"hello\"}"); err != nil {
		fmt.Printf("error %s\n", err.Error())
	}
	blobitem := map[string]blob.Blob{"application/json": b}

	if clipitem, err = clipboarditem.New(blobitem); err != nil {
		fmt.Printf("error %s\n", err.Error())
	}

	methodsAttempt = []map[string]interface{}{
		{"method": "Read", "type": "constructnamechecking", "resultattempt": "Promise"},
		{"method": "ReadText", "type": "constructnamechecking", "resultattempt": "Promise"},
		{"method": "Write", "args": []interface{}{[]clipboarditem.ClipboardItem{clipitem}}, "type": "constructnamechecking", "resultattempt": "Promise"},
		{"method": "WriteText", "args": []interface{}{"hello"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	}

	baseobject.Eval(`clip=window.navigator.clipboard`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "clip"); testingutils.AssertErr(t, err) {
		if nav, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "Clipboard", nav.ConstructName_())

		}
	}

}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "clip"); testingutils.AssertErr(t, err) {

		if clip, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, clip, result)
			}

		}

	}
}
