package cssstyledeclaration

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/cssrule"
)

var singleton sync.Once

var cssstyledeclarationinterface js.Value

//CSSStyleDeclaration struct
type CSSStyleDeclaration struct {
	baseobject.BaseObject
}

type CSSStyleDeclarationFrom interface {
	CSSStyleDeclaration_() CSSStyleDeclaration
}

func (c CSSStyleDeclaration) CSSStyleDeclaration_() CSSStyleDeclaration {
	return c
}

func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if cssstyledeclarationinterface, err = baseobject.Get(js.Global(), "CSSStyleDeclaration"); err != nil {
			cssstyledeclarationinterface = js.Undefined()
		}
		baseobject.Register(cssstyledeclarationinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return cssstyledeclarationinterface
}

func NewFromJSObject(obj js.Value) (CSSStyleDeclaration, error) {
	var c CSSStyleDeclaration
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				c.BaseObject = c.SetObject(obj)

			} else {
				err = ErrNotAnCSSStyleDeclaration
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return c, err
}

func (c CSSStyleDeclaration) ParentRule() (cssrule.CSSRule, error) {
	var err error
	var obj js.Value
	var cr cssrule.CSSRule
	if obj, err = c.Get("parentRule"); err == nil {

		if obj.IsUndefined() {
			err = baseobject.ErrNotAnObject

		} else {
			cr, err = cssrule.NewFromJSObject(obj)
		}
	}
	return cr, err
}

func (c CSSStyleDeclaration) SetProperty(propertyName string, opts ...string) error {
	var err error
	var arrayJS []interface{}

	arrayJS = append(arrayJS, js.ValueOf(propertyName))

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}
	_, err = c.Call("setProperty", arrayJS...)
	return err

}

func (c CSSStyleDeclaration) Item(index int) (string, error) {
	var err error
	var obj js.Value
	var ret string

	if obj, err = c.Call("item", js.ValueOf(index)); err == nil {
		ret = obj.String()
	}
	return ret, err
}

func (c CSSStyleDeclaration) GetPropertyPriority(property string) (string, error) {
	var err error
	var obj js.Value
	var ret string

	if obj, err = c.Call("getPropertyPriority", js.ValueOf(property)); err == nil {
		ret = obj.String()
	}
	return ret, err
}

func (c CSSStyleDeclaration) GetPropertyValue(property string) (string, error) {
	var err error
	var obj js.Value
	var ret string

	if obj, err = c.Call("getPropertyValue", js.ValueOf(property)); err == nil {
		ret = obj.String()
	}
	return ret, err
}

func (c CSSStyleDeclaration) RemoveProperty(property string) (string, error) {
	var err error
	var obj js.Value
	var ret string

	if obj, err = c.Call("removeProperty", js.ValueOf(property)); err == nil {
		ret = obj.String()
	}
	return ret, err
}
