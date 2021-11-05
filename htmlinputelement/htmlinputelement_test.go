package htmlinputelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`f=document.createElement("form")
	i= document.createElement("input")
	f.appendChild(i)`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLInputElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "i"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLInputElement", b.ConstructName_())
		}

	}
}

var methodsCommon []map[string]interface{} = []map[string]interface{}{
	{"method": "Form", "type": "constructnamechecking", "resultattempt": "HTMLFormElement"},
	{"method": "FormAction", "type": "contains", "resultattempt": "http://localhost"},
	{"method": "FormEncType", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "FormMethod", "resultattempt": ""},
	{"method": "FormNoValidate", "resultattempt": false},
	{"method": "FormTarget", "resultattempt": ""},
}

func TestCommon(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "i"); testingutils.AssertErr(t, err) {

		if input, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsCommon {
				testingutils.InvokeCheck(t, input, result)
			}

		}

	}
}

//apply on txt
var methodsTxt []map[string]interface{} = []map[string]interface{}{
	{"method": "Name", "resultattempt": ""},
	{"method": "Type", "resultattempt": "text"},
	{"method": "Disabled", "resultattempt": false},
	{"method": "Autofocus", "resultattempt": false},
	{"method": "Required", "resultattempt": false},
	{"method": "Value", "resultattempt": ""},
	{"method": "Validity", "type": "constructnamechecking", "resultattempt": "ValidityState"},
	{"method": "ValidationMessage", "resultattempt": ""},
	{"method": "WillValidate", "resultattempt": true},
	{"method": "SetFormAction", "args": []interface{}{"sync"}, "gettermethod": "FormAction", "type": "contains", "resultattempt": "/sync"},
	{"method": "SetFormEncType", "args": []interface{}{"multipart/form-data"}, "gettermethod": "FormEncType", "resultattempt": "multipart/form-data"},
	{"method": "SetFormMethod", "args": []interface{}{"post"}, "gettermethod": "FormMethod", "resultattempt": "post"},
	{"method": "SetFormNoValidate", "args": []interface{}{true}, "gettermethod": "FormNoValidate", "resultattempt": true},
	{"method": "SetFormTarget", "args": []interface{}{"_self"}, "gettermethod": "FormTarget", "resultattempt": "_self"},
	{"method": "SetName", "args": []interface{}{"hello"}, "gettermethod": "Name", "resultattempt": "hello"},
	{"method": "SetType", "args": []interface{}{"reset"}, "gettermethod": "Type", "resultattempt": "reset"},
	{"method": "SetDisabled", "args": []interface{}{true}, "gettermethod": "Disabled", "resultattempt": true},
	{"method": "SetAutofocus", "args": []interface{}{true}, "gettermethod": "Autofocus", "resultattempt": true},
	{"method": "SetRequired", "args": []interface{}{true}, "gettermethod": "Required", "resultattempt": true},
	{"method": "SetValue", "args": []interface{}{"value"}, "gettermethod": "Value", "resultattempt": "value"},
}

func TestInputTxt(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "i"); testingutils.AssertErr(t, err) {

		if input, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsTxt {
				testingutils.InvokeCheck(t, input, result)
			}

		}

	}
}

var methodsCheckbox []map[string]interface{} = []map[string]interface{}{

	{"method": "Checked", "resultattempt": false},
	{"method": "DefaultChecked", "resultattempt": false},
	{"method": "Indeterminate", "resultattempt": false},

	{"method": "SetChecked", "args": []interface{}{true}, "gettermethod": "Checked", "resultattempt": true},
	{"method": "SetDefaultChecked", "args": []interface{}{true}, "gettermethod": "DefaultChecked", "resultattempt": true},
	{"method": "SetIndeterminate", "args": []interface{}{true}, "gettermethod": "Indeterminate", "resultattempt": true},
}

func TestInputCheckbox(t *testing.T) {

	baseobject.Eval(`ick= document.createElement("input")
	ick.type="checkbox"`)

	if obj, err := baseobject.Get(js.Global(), "ick"); testingutils.AssertErr(t, err) {

		if input, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsCheckbox {
				testingutils.InvokeCheck(t, input, result)
			}

		}

	}
}

var methodsImage []map[string]interface{} = []map[string]interface{}{
	{"method": "Alt", "resultattempt": ""},
	{"method": "Height", "resultattempt": 0},
	{"method": "Src", "resultattempt": ""},
	{"method": "Width", "resultattempt": 0},
	{"method": "SetAlt", "args": []interface{}{"n2"}, "gettermethod": "Alt", "resultattempt": "n2"},
	{"method": "SetHeight", "args": []interface{}{250}, "gettermethod": "Height", "resultattempt": 250},
	{"method": "SetSrc", "args": []interface{}{"https://github.com/realPy/hogosuru/blob/main/ressources/virtualRendering.png?raw=true"}, "gettermethod": "Src", "resultattempt": "https://github.com/realPy/hogosuru/blob/main/ressources/virtualRendering.png?raw=true"},
	{"method": "SetWidth", "args": []interface{}{200}, "gettermethod": "Width", "resultattempt": 200},
}

func TestInputImage(t *testing.T) {

	baseobject.Eval(`iimg= document.createElement("input")
	iimg.type="image"`)

	if obj, err := baseobject.Get(js.Global(), "iimg"); testingutils.AssertErr(t, err) {

		if input, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsImage {
				testingutils.InvokeCheck(t, input, result)
			}

		}

	}
}

var methodsFile []map[string]interface{} = []map[string]interface{}{
	{"method": "Files", "type": "constructnamechecking", "resultattempt": "FileList"},
	{"method": "DirName", "resultattempt": ""},
	{"method": "SetDirName", "args": []interface{}{"5"}, "gettermethod": "DirName", "resultattempt": "5"},
	{"method": "Multiple", "resultattempt": false},
	{"method": "SetMultiple", "args": []interface{}{true}, "gettermethod": "Multiple", "resultattempt": true},
}

func TestFile(t *testing.T) {

	baseobject.Eval(`ifile= document.createElement("input")
	ifile.type="file"`)

	if obj, err := baseobject.Get(js.Global(), "ifile"); testingutils.AssertErr(t, err) {

		if input, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsFile {
				testingutils.InvokeCheck(t, input, result)
			}

		}

	}
}

var methodsNumber []map[string]interface{} = []map[string]interface{}{
	{"method": "Autocomplete", "resultattempt": ""},
	{"method": "SetAutocomplete", "args": []interface{}{"off"}, "gettermethod": "Autocomplete", "resultattempt": "off"},
	{"method": "Max", "resultattempt": ""},
	{"method": "SetMax", "args": []interface{}{"10"}, "gettermethod": "Max", "resultattempt": "10"},
	{"method": "MaxLength", "resultattempt": -1},
	{"method": "SetMaxLength", "args": []interface{}{10}, "gettermethod": "MaxLength", "resultattempt": 10},
	{"method": "Min", "resultattempt": ""},
	{"method": "SetMin", "args": []interface{}{"5"}, "gettermethod": "Min", "resultattempt": "5"},
	{"method": "MinLength", "resultattempt": -1},
	{"method": "SetMinLength", "args": []interface{}{5}, "gettermethod": "MinLength", "resultattempt": 5},
	{"method": "Pattern", "resultattempt": ""},
	{"method": "SetPattern", "args": []interface{}{"####"}, "gettermethod": "Pattern", "resultattempt": "####"},
	{"method": "Placeholder", "resultattempt": ""},
	{"method": "SetPlaceholder", "args": []interface{}{"####"}, "gettermethod": "Placeholder", "resultattempt": "####"},
	{"method": "ReadOnly", "resultattempt": false},
	{"method": "SetReadOnly", "args": []interface{}{true}, "gettermethod": "ReadOnly", "resultattempt": true},
	{"method": "Step", "resultattempt": ""},
	{"method": "SetStep", "args": []interface{}{"5"}, "gettermethod": "Step", "resultattempt": "5"},
}

func TestNumber(t *testing.T) {

	baseobject.Eval(`inumber= document.createElement("input")
	inumber.type="number"`)

	if obj, err := baseobject.Get(js.Global(), "inumber"); testingutils.AssertErr(t, err) {

		if input, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsNumber {
				testingutils.InvokeCheck(t, input, result)
			}

		}

	}
}

var methodsText []map[string]interface{} = []map[string]interface{}{
	{"method": "SelectionStart", "resultattempt": 20},
	{"method": "SetSelectionStart", "args": []interface{}{1}, "gettermethod": "SelectionStart", "resultattempt": 1},
	{"method": "SelectionEnd", "resultattempt": 20},
	{"method": "SetSelectionEnd", "args": []interface{}{8}, "gettermethod": "SelectionEnd", "resultattempt": 8},
	//My chrome return none, CI return default "forward"... so ensure no errror
	{"method": "SelectionDirection", "type": "error", "resultattempt": nil},
	{"method": "SetSelectionDirection", "args": []interface{}{"forward"}, "gettermethod": "SelectionDirection", "resultattempt": "forward"},
	{"method": "Size", "resultattempt": 20},
	{"method": "SetSize", "args": []interface{}{8}, "gettermethod": "Size", "resultattempt": 8},
	{"method": "Labels", "type": "constructnamechecking", "resultattempt": "NodeList"},
}

func TestText(t *testing.T) {

	baseobject.Eval(`
	l= document.createElement("label")
	itext= document.createElement("input")
	itext.type="text"
	itext.value="osdiosidosidosidosdi"
	l.appendChild(itext)
	`)

	if obj, err := baseobject.Get(js.Global(), "itext"); testingutils.AssertErr(t, err) {

		if input, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsText {
				testingutils.InvokeCheck(t, input, result)
			}

		}

	}
}
