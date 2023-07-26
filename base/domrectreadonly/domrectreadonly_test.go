package domrectreadonly

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`r=new DOMRect()
	r.x=3
	r.y=8
	r.width=10
	r.height=13
	ro=DOMRectReadOnly.fromRect(r)
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "ro"); testingutils.AssertErr(t, err) {
		if rect, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object DOMRectReadOnly]", rect.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "X", "resultattempt": float64(3)},
	{"method": "Width", "resultattempt": float64(10)},
	{"method": "Right", "resultattempt": float64(13)},
	{"method": "Left", "resultattempt": float64(3)},
	{"method": "Y", "resultattempt": float64(8)},
	{"method": "Height", "resultattempt": float64(13)},
	{"method": "Top", "resultattempt": float64(8)},
	{"method": "Bottom", "resultattempt": float64(21)},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "ro"); testingutils.AssertErr(t, err) {

		if meta, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, meta, result)
			}

		}

	}
}
