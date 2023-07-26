package cssrule

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
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
	prule=document.styleSheets[0].rules[0].style.parentRule
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "prule"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object CSSStyleRule]", o.ToString_())

		}
	}

}

func TestCssText(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "h1 { color: red; font-size: 50px; }"
	document.head.appendChild(s)
	prule=document.styleSheets[0].rules[0].style.parentRule
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "prule"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if v, err := o.CssText(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "h1 { color: red; font-size: 50px; }", v)
			}

		}
	}

}

func TestParentRule(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "@supports (display: flex) { @media screen and (min-width: 900px) { article { display: flex; } } }"
	document.head.appendChild(s)
	rule=document.styleSheets[0].rules[0].cssRules[0]
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "rule"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if p, err := o.ParentRule(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object CSSSupportsRule]", p.ToString_())

			}

		}
	}

}

func TestParentStyleSheet(t *testing.T) {

	baseobject.Eval(`s=document.createElement("style")
	s.textContent= "@supports (display: flex) { @media screen and (min-width: 900px) { article { display: flex; } } }"
	document.head.appendChild(s)
	rule=document.styleSheets[0].rules[0].cssRules[0]
	document.head.removeChild(s)
	`)

	//CSSSyleRule is derivated from CSSRule
	if obj, err := baseobject.Get(js.Global(), "rule"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if p, err := o.ParentStyleSheet(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object CSSStyleSheet]", p.ToString_())

			}

		}
	}

}
