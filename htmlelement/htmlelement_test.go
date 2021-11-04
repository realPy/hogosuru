package htmlelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`e=document.createElement("unk")
	`)
	m.Run()
}
func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "e"); testingutils.AssertErr(t, err) {

		if a, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLUnknownElement", a.ConstructName_())
		}

	}
}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "AccessKey", "resultattempt": ""},
	{"method": "AccessKeyLabel", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "ClassName", "resultattempt": ""},
	{"method": "ContentEditable", "resultattempt": "inherit"},
	{"method": "IsContentEditable", "resultattempt": false},
	{"method": "Dataset", "args": []interface{}{"test"}, "type": "error", "resultattempt": ErrDatasetNotFound},
	{"method": "Dir", "resultattempt": ""},
	{"method": "Hidden", "resultattempt": false},
	{"method": "Lang", "resultattempt": ""},
	{"method": "OffsetHeight", "resultattempt": 0},
	{"method": "OffsetLeft", "resultattempt": 0},
	{"method": "OffsetTop", "resultattempt": 0},
	{"method": "OffsetWidth", "resultattempt": 0},
	{"method": "OffsetParent", "type": "error", "resultattempt": ErrParentNotFound},
	{"method": "Title", "resultattempt": ""},
	{"method": "Style", "type": "constructnamechecking", "resultattempt": "CSSStyleDeclaration"},
	{"method": "Blur", "type": "error", "resultattempt": nil},
	{"method": "Click", "type": "error", "resultattempt": nil},
	{"method": "Focus", "type": "error", "resultattempt": nil},
	{"method": "SetAccessKey", "args": []interface{}{"i"}, "gettermethod": "AccessKey", "resultattempt": "i"},
	{"method": "SetAccessKeyLabel", "args": []interface{}{"i"}, "gettermethod": "AccessKeyLabel", "resultattempt": "i"},
	{"method": "SetClassName", "args": []interface{}{"test"}, "gettermethod": "ClassName", "resultattempt": "test"},
	{"method": "SetContentEditable", "args": []interface{}{"true"}, "gettermethod": "ContentEditable", "resultattempt": "true"},
	{"method": "SetDataset", "args": []interface{}{"test", "myvalue"}, "gettermethod": "Dataset", "getterargs": []interface{}{"test"}, "resultattempt": "myvalue"},
	{"method": "SetDir", "args": []interface{}{"auto"}, "gettermethod": "Dir", "resultattempt": "auto"},
	{"method": "SetHidden", "args": []interface{}{true}, "gettermethod": "Hidden", "resultattempt": true},
	{"method": "SetLang", "args": []interface{}{"fr"}, "gettermethod": "Lang", "resultattempt": "fr"},
	{"method": "SetTitle", "args": []interface{}{"mytitle"}, "gettermethod": "Title", "resultattempt": "mytitle"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "e"); testingutils.AssertErr(t, err) {

		if area, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, area, result)
			}

		}

	}
}
