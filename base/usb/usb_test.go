package usb

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/clipboarditem"
	"github.com/realPy/hogosuru/testingutils"
)

var clipitem clipboarditem.ClipboardItem

var methodsAttempt []map[string]interface{}

func TestMain(m *testing.M) {

	baseobject.SetSyscall()

	methodsAttempt = []map[string]interface{}{
		{"method": "GetDevices", "type": "constructnamechecking", "resultattempt": "Promise"},
		{"method": "RequestDevices", "args": []interface{}{map[string]int{"vendorId": 0x11}}, "type": "constructnamechecking", "resultattempt": "Promise"},
	}

	baseobject.Eval(`usbobj=navigator.usb`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "usbobj"); testingutils.AssertErr(t, err) {
		if usb, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "USB", usb.ConstructName_())

		}
	}
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "usbobj"); testingutils.AssertErr(t, err) {

		if clip, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, clip, result)
			}

		}

	}
}
