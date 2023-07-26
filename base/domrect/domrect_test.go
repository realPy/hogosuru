package domrect

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`r=new DOMRect()`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "r"); testingutils.AssertErr(t, err) {
		if rect, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object DOMRect]", rect.ToString_())

		}
	}

}

func TestNew(t *testing.T) {

	if rect, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "[object DOMRect]", rect.ToString_())

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "X", "resultattempt": float64(0)},
	{"method": "SetX", "args": []interface{}{float64(10.0)}, "gettermethod": "X", "resultattempt": float64(10.0)},
	{"method": "Width", "resultattempt": float64(0)},
	{"method": "SetWidth", "args": []interface{}{float64(55.0)}, "gettermethod": "Width", "resultattempt": float64(55.0)},
	{"method": "Right", "resultattempt": float64(65)},
	{"method": "Left", "resultattempt": float64(10)},

	{"method": "Y", "resultattempt": float64(0)},
	{"method": "SetY", "args": []interface{}{float64(3.0)}, "gettermethod": "Y", "resultattempt": float64(3.0)},
	{"method": "Height", "resultattempt": float64(0)},
	{"method": "SetHeight", "args": []interface{}{float64(45.0)}, "gettermethod": "Height", "resultattempt": float64(45.0)},
	{"method": "Top", "resultattempt": float64(3)},
	{"method": "Bottom", "resultattempt": float64(48)},
	{"method": "RectReadOnly", "type": "constructnamechecking", "resultattempt": "DOMRectReadOnly"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "r"); testingutils.AssertErr(t, err) {

		if meta, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, meta, result)
			}

		}

	}
}
