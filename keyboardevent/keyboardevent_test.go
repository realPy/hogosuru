package keyboardevent

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()

	m.Run()
}

func TestNew(t *testing.T) {

	if k, err := New("keypress", map[string]interface{}{"code": "KeyA"}); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object KeyboardEvent]", k.ToString_())
		if v, err := k.Code(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "KeyA", v)
		}
	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\")")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object KeyboardEvent]", k.ToString_())

		}
	}

}

func TestAltKey(t *testing.T) {

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\",{altKey:true})")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := k.AltKey(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, true, v)
			}

		}
	}

}

func TestCtrlKey(t *testing.T) {

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\",{ctrlKey:true})")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := k.CtrlKey(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, true, v)
			}

		}
	}

}

func TestKey(t *testing.T) {

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\",{key:\"Enter\"})")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := k.Key(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "Enter", v)
			}

		}
	}

}

func TestCode(t *testing.T) {
  
	//independent position (FR position) == (press q obtain KeyA, press a obtain KeyQ )

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\",{code:\"KeyA\"})")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := k.Code(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "KeyA", v)
			}

		}
	}

}

func TestIsComposing(t *testing.T) {

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\",{isComposing:true})")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := k.IsComposing(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, true, v)
			}

		}
	}

}

func TestLocation(t *testing.T) {

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\",{location:1})")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := k.Location(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, int64(1), v)
			}

		}
	}

}

func TestMetaKey(t *testing.T) {

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\",{metaKey:true})")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := k.MetaKey(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, true, v)
			}

		}
	}

}

func TestRepeat(t *testing.T) {

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\",{repeat:true})")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := k.Repeat(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, true, v)
			}

		}
	}

}

func TestShiftKey(t *testing.T) {

	baseobject.Eval("keypress=new KeyboardEvent(\"keyup\",{shiftKey:true})")

	if obj, err := baseobject.Get(js.Global(), "keypress"); testingutils.AssertErr(t, err) {
		if k, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := k.ShiftKey(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, true, v)
			}

		}
	}

}
