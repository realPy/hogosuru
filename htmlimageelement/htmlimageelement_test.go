package htmlimageelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`i= document.createElement("img")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLImageElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "i"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLImageElement", b.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Alt", "resultattempt": ""},
	{"method": "Complete", "resultattempt": true},
	{"method": "CrossOrigin", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "CurrentSrc", "resultattempt": ""},
	{"method": "Decoding", "resultattempt": "auto"},
	{"method": "Height", "resultattempt": 0},
	{"method": "IsMap", "resultattempt": false},
	{"method": "Loading", "resultattempt": "auto"},
	{"method": "NaturalHeight", "resultattempt": 0},
	{"method": "NaturalWidth", "resultattempt": 0},
	{"method": "Src", "resultattempt": ""},
	{"method": "Width", "resultattempt": 0},
	{"method": "X", "resultattempt": 0},
	{"method": "Y", "resultattempt": 0},
	{"method": "SetAlt", "args": []interface{}{"n2"}, "gettermethod": "Alt", "resultattempt": "n2"},
	{"method": "SetCrossOrigin", "args": []interface{}{"anonymous"}, "gettermethod": "CrossOrigin", "resultattempt": "anonymous"},
	{"method": "SetDecoding", "args": []interface{}{"sync"}, "gettermethod": "Decoding", "resultattempt": "sync"},
	{"method": "SetHeight", "args": []interface{}{200}, "gettermethod": "Height", "resultattempt": 200},
	{"method": "SetLoading", "args": []interface{}{"eager"}, "gettermethod": "Loading", "resultattempt": "eager"},
	{"method": "SetSrc", "args": []interface{}{"https://github.com/realPy/hogosuru/blob/main/ressources/virtualRendering.png?raw=true"}, "gettermethod": "Src", "resultattempt": "https://github.com/realPy/hogosuru/blob/main/ressources/virtualRendering.png?raw=true"},
	{"method": "SetWidth", "args": []interface{}{200}, "gettermethod": "Width", "resultattempt": 200},
	{"method": "Decode", "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "i"); testingutils.AssertErr(t, err) {

		if image, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, image, result)
			}

		}

	}
}
