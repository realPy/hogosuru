package domtokenlist

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`l=document.createElement("link")
	l.rel="a b c"
	list=l.relList
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "list"); testingutils.AssertErr(t, err) {
		if list, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "DOMTokenList", list.ConstructName_())

		}
	}

}

func TestItem(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "list"); testingutils.AssertErr(t, err) {
		if list, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if i, err := list.Item(0); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "a", i)
			}

		}
	}

}

func TestContains(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "list"); testingutils.AssertErr(t, err) {
		if list, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if b, err := list.Contains("b"); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, true, b)
			}

			if b, err := list.Contains("d"); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, false, b)
			}

		}
	}

}

func TestAdd(t *testing.T) {

	baseobject.Eval(`ladd=document.createElement("link")
	ladd.rel="a b c"
	listadd=ladd.relList
	`)

	if obj, err := baseobject.Get(js.Global(), "listadd"); testingutils.AssertErr(t, err) {
		if list, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, list.Add("d", "e"))
			testingutils.AssertExpect(t, "a b c d e", list.ToString_())

		}
	}

}

func TestRemove(t *testing.T) {

	baseobject.Eval(`lr=document.createElement("link")
	lr.rel="a b c"
	listr=lr.relList
	`)

	if obj, err := baseobject.Get(js.Global(), "listr"); testingutils.AssertErr(t, err) {
		if list, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, list.Remove("b", "c"))
			testingutils.AssertExpect(t, "a", list.ToString_())

		}
	}

}

func TestReplace(t *testing.T) {

	baseobject.Eval(`lre=document.createElement("link")
	lre.rel="a b c"
	listre=lre.relList
	`)

	if obj, err := baseobject.Get(js.Global(), "listre"); testingutils.AssertErr(t, err) {
		if list, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, list.Replace("b", "k"))
			testingutils.AssertExpect(t, "a k c", list.ToString_())

		}
	}

}

func TestToggle(t *testing.T) {

	baseobject.Eval(`lt=document.createElement("link")
	lt.rel="a b c"
	listt=lt.relList
	`)

	if obj, err := baseobject.Get(js.Global(), "listt"); testingutils.AssertErr(t, err) {
		if list, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if b, err := list.Toggle("b"); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, false, b)
				testingutils.AssertExpect(t, "a c", list.ToString_())

			}

		}
	}

}

func TestSupports(t *testing.T) {

	baseobject.Eval(`ls=document.createElement("link")
	ls.rel="a b c"
	lists=ls.relList
	`)

	if obj, err := baseobject.Get(js.Global(), "lists"); testingutils.AssertErr(t, err) {
		if list, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if b, err := list.Supports(""); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, false, b)

			}

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Entries", "type": "tostringchecking", "resultattempt": "[object Array Iterator]"},
	{"method": "Keys", "type": "tostringchecking", "resultattempt": "[object Array Iterator]"},
	{"method": "Values", "type": "tostringchecking", "resultattempt": "[object Array Iterator]"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "list"); testingutils.AssertErr(t, err) {

		if maphtml, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, maphtml, result)
			}

		}

	}
}

func TestForEach(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "list"); testingutils.AssertErr(t, err) {
		if list, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			var i int
			list.ForEach(func(s string) {

				i++

			})
			testingutils.AssertExpect(t, 3, i)
		}
	}

}
