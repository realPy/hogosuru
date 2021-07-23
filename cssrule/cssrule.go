package cssrule

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/stylesheet"
)

var singleton sync.Once

var cssruleinterface js.Value

//CSSRule struct
type CSSRule struct {
	baseobject.BaseObject
}

func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if cssruleinterface, err = js.Global().GetWithErr("CSSRule"); err != nil {
			cssruleinterface = js.Null()
		}
	})

	baseobject.Register(cssruleinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return cssruleinterface
}

func NewFromJSObject(obj js.Value) (CSSRule, error) {
	var c CSSRule
	var err error
	if dli := GetInterface(); !dli.IsNull() {
		if obj.InstanceOf(dli) {
			c.BaseObject = c.SetObject(obj)

		} else {
			err = ErrNotAnCSSRule
		}
	} else {
		err = ErrNotImplemented
	}
	return c, err
}

func (c CSSRule) CssText() (string, error) {

	return c.GetAttributeString("cssText")
}

func (c CSSRule) SetCssText(value string) error {
	return c.SetAttributeString("cssText", value)
}

func (c CSSRule) ParentRule() (CSSRule, error) {
	var err error
	var obj js.Value
	var cr CSSRule
	if obj, err = c.JSObject().GetWithErr("parentRule"); err == nil {

		if obj.IsNull() {
			err = baseobject.ErrNotAnObject

		} else {
			cr, err = NewFromJSObject(obj)
		}
	}
	return cr, err
}

func (c CSSRule) ParentStyleSheet() (stylesheet.StyleSheet, error) {
	var err error
	var obj js.Value
	var s stylesheet.StyleSheet
	if obj, err = c.JSObject().GetWithErr("parentStyleSheet"); err == nil {

		if obj.IsNull() {
			err = baseobject.ErrNotAnObject

		} else {
			s, err = stylesheet.NewFromJSObject(obj)
		}
	}
	return s, err
}
