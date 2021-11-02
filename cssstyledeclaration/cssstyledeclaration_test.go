package cssstyledeclaration

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

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "style"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object CSSStyleDeclaration]", o.ToString_())

		}
	}

}

func TestParentRule(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "style"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if p, err := o.ParentRule(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object CSSStyleRule]", p.ToString_())

			}

		}
	}

}

func TestGetPropertyValue(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "style"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := o.GetPropertyValue("color"); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "red", v)
			}

		}
	}

}

func TestSetProperty(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "style"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if err := o.SetProperty("color", "blue"); testingutils.AssertErr(t, err) {

				if v, err := o.GetPropertyValue("color"); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, "blue", v)
				}

			}

		}
	}

}

func TestItem(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "style"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := o.Item(0); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "color", v)

			}

		}
	}

}

func TestGetPropertyPriority(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red!important; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "style"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := o.GetPropertyPriority("color"); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "important", v)

			}

		}
	}

}
func TestRemoveProperty(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	style=document.styleSheets[0].rules[0].style
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "style"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := o.RemoveProperty("color"); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "red", v)
				if v2, err := o.GetPropertyValue("color"); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, "", v2)
				}

			}

		}
	}

}
