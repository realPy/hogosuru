package mouseevent

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`mevent=new MouseEvent("test",{altKey:true,button:2,buttons:16,clientX:200,clientY:453,ctrlKey:true,metaKey:true,movementX:100,movementY:300,screenX:89,screenY:43,shiftKey:true})

	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if k, err := New("test"); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object MouseEvent]", k.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "mevent"); testingutils.AssertErr(t, err) {
		if mevent, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object MouseEvent]", mevent.ToString_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "AltKey", "resultattempt": true},
	{"method": "Button", "resultattempt": 2},
	{"method": "Buttons", "resultattempt": 16},
	{"method": "ClientX", "resultattempt": float64(200)},
	{"method": "ClientY", "resultattempt": float64(453)},
	{"method": "CtrlKey", "resultattempt": true},
	{"method": "MetaKey", "resultattempt": true},
	{"method": "MovementX", "resultattempt": 100},
	{"method": "MovementY", "resultattempt": 300},
	{"method": "OffsetX", "resultattempt": float64(200)},
	{"method": "OffsetY", "resultattempt": float64(453)},
	{"method": "PageX", "resultattempt": 200},
	{"method": "PageY", "resultattempt": 453},
	{"method": "RelatedTarget", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "ScreenX", "resultattempt": float64(89)},
	{"method": "ScreenY", "resultattempt": float64(43)},
	{"method": "ShiftKey", "resultattempt": true},
	{"method": "X", "resultattempt": float64(200)},
	{"method": "Y", "resultattempt": float64(453)},
	{"method": "GetModifierState", "args": []interface{}{"Accel"}, "resultattempt": true},
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
