package element

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/attr"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`
	elementspan= document.createElement("span")
	document.body.appendChild(elementspan)
	element= document.createElement("title")
	element.setAttribute("hello","world")
	listattr=element.attributes
	attr=listattr.item(0)
	div=document.createElement("div")
	element.appendChild(div)
	collection=element.children
	document.body.appendChild(element)
	element2= document.createElement("br")
	document.body.appendChild(element2)
	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "element"); testingutils.AssertErr(t, err) {

		if e, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTitleElement", e.ConstructName_())
		}

	}

}

func TestItemFromHTMLCollection(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "collection"); testingutils.AssertErr(t, err) {

		if c, err := htmlcollection.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if e, err := ItemFromHTMLCollection(c, 0); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "HTMLDivElement", e.ConstructName_())

			}
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "Attributes", "type": "constructnamechecking", "resultattempt": "NamedNodeMap"},
	{"method": "ChildElementCount", "resultattempt": 1},
	{"method": "Children", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "ClassList", "type": "constructnamechecking", "resultattempt": "DOMTokenList"},
	{"method": "ClassName", "resultattempt": ""},
	{"method": "SetClassName", "args": []interface{}{"n2"}, "gettermethod": "ClassName", "resultattempt": "n2"},
	{"method": "ClientHeight", "resultattempt": 0},
	{"method": "ClientLeft", "resultattempt": 0},
	{"method": "ClientTop", "resultattempt": 0},
	{"method": "ClientWidth", "resultattempt": 0},
	{"method": "ID", "resultattempt": ""},
	{"method": "SetID", "args": []interface{}{"test"}, "gettermethod": "ID", "resultattempt": "test"},
	{"method": "InnerHTML", "resultattempt": "<div></div>"},
	{"method": "SetInnerHTML", "args": []interface{}{"test"}, "gettermethod": "InnerHTML", "resultattempt": "test"},
	{"method": "LocalName", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "NamespaceURI", "resultattempt": "http://www.w3.org/1999/xhtml"},
	{"method": "NextElementSibling", "type": "constructnamechecking", "resultattempt": "HTMLBRElement"},
	{"method": "PreviousElementSibling", "type": "constructnamechecking", "resultattempt": "HTMLSpanElement"},
	{"method": "TagName", "resultattempt": "TITLE"},
	{"method": "Prefix", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "ScrollHeight", "resultattempt": 0},
	{"method": "SetScrollHeight", "args": []interface{}{0}, "gettermethod": "ScrollHeight", "resultattempt": 0},
	{"method": "ScrollLeft", "resultattempt": 0},
	{"method": "SetScrollLeft", "args": []interface{}{0}, "gettermethod": "ScrollLeft", "resultattempt": 0},
	{"method": "ScrollTop", "resultattempt": 0},
	{"method": "SetScrollTop", "args": []interface{}{0}, "gettermethod": "ScrollTop", "resultattempt": 0},
	{"method": "ScrollWidth", "resultattempt": 0},
	{"method": "SetScrollWidth", "args": []interface{}{0}, "gettermethod": "ScrollWidth", "resultattempt": 0},
	{"method": "Closest", "args": []interface{}{"body"}, "type": "constructnamechecking", "resultattempt": "HTMLBodyElement"},
	{"method": "GetAttribute", "args": []interface{}{"hello"}, "resultattempt": "world"},
	{"method": "GetAttributeNS", "args": []interface{}{"name", "hello"}, "type": "error", "resultattempt": ErrAttributeEmpty},
	{"method": "GetAttributeNames", "type": "constructnamechecking", "resultattempt": "Array"},
	{"method": "GetBoundingClientRect", "type": "constructnamechecking", "resultattempt": "DOMRect"},
	{"method": "GetClientRects", "type": "constructnamechecking", "resultattempt": "DOMRectList"},
	{"method": "GetElementsByClassName", "args": []interface{}{"div"}, "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "GetElementsByTagName", "args": []interface{}{"div"}, "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "GetElementsByTagNameNS", "args": []interface{}{"namespace", "div"}, "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "HasAttribute", "args": []interface{}{"hello"}, "resultattempt": true},
	{"method": "HasPointerCapture", "args": []interface{}{0}, "resultattempt": false},

	{"method": "OuterHTML", "resultattempt": "<title hello=\"world\" class=\"n2\" id=\"test\">test</title>"},
	{"method": "SetOuterHTML", "args": []interface{}{"<title helloZ=\"world\" class=\"n2\" id=\"test\">test</title>"}, "gettermethod": "OuterHTML", "resultattempt": "<title hello=\"world\" class=\"n2\" id=\"test\">test</title>"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "element"); testingutils.AssertErr(t, err) {

		if elem, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, elem, result)
			}

		}

	}
}
func TestOwnerElementForAttr(t *testing.T) {

	if objattr, err := baseobject.Get(js.Global(), "attr"); testingutils.AssertErr(t, err) {

		if attr, err := attr.NewFromJSObject(objattr); testingutils.AssertErr(t, err) {

			if elem, err := OwnerElementForAttr(attr); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "HTMLTitleElement", elem.ConstructName_())

			}
		}
	}

}

func TestAfter(t *testing.T) {

	baseobject.Eval(`
	elementspanaf= document.createElement("span")
	elementdivaf= document.createElement("div")
	elementspanaf.appendChild(elementdivaf)
	elementbraf= document.createElement("br")

	`)

	if objspan, err := baseobject.Get(js.Global(), "elementspanaf"); testingutils.AssertErr(t, err) {

		if span, err := NewFromJSObject(objspan); testingutils.AssertErr(t, err) {

			if objdiv, err := baseobject.Get(js.Global(), "elementdivaf"); testingutils.AssertErr(t, err) {

				if div, err := NewFromJSObject(objdiv); testingutils.AssertErr(t, err) {
					if objbr, err := baseobject.Get(js.Global(), "elementbraf"); testingutils.AssertErr(t, err) {

						if br, err := NewFromJSObject(objbr); testingutils.AssertErr(t, err) {

							testingutils.AssertErr(t, div.After(br))

							if val, err := span.OuterHTML(); testingutils.AssertErr(t, err) {
								testingutils.AssertExpect(t, "<span><div></div><br></span>", val)

							}

						}
					}

				}
			}

		}
	}

}

func TestAppend(t *testing.T) {

	baseobject.Eval(`
	elementspanap= document.createElement("span")
	elementdivap= document.createElement("div")

	`)

	if objspan, err := baseobject.Get(js.Global(), "elementspanap"); testingutils.AssertErr(t, err) {

		if span, err := NewFromJSObject(objspan); testingutils.AssertErr(t, err) {

			if objdiv, err := baseobject.Get(js.Global(), "elementdivap"); testingutils.AssertErr(t, err) {

				if div, err := NewFromJSObject(objdiv); testingutils.AssertErr(t, err) {

					testingutils.AssertErr(t, span.Append(div))

					if er, err := span.FirstChild(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "HTMLDivElement", er.ConstructName_())

					}

				}
			}
		}
	}

}

func TestBefore(t *testing.T) {

	baseobject.Eval(`
	elementspanbf= document.createElement("span")
	elementdivbf= document.createElement("div")
	elementspanbf.appendChild(elementdivbf)
	elementbrbf= document.createElement("br")

	`)

	if objspan, err := baseobject.Get(js.Global(), "elementspanbf"); testingutils.AssertErr(t, err) {

		if span, err := NewFromJSObject(objspan); testingutils.AssertErr(t, err) {

			if objdiv, err := baseobject.Get(js.Global(), "elementdivbf"); testingutils.AssertErr(t, err) {

				if div, err := NewFromJSObject(objdiv); testingutils.AssertErr(t, err) {
					if objbr, err := baseobject.Get(js.Global(), "elementbrbf"); testingutils.AssertErr(t, err) {

						if br, err := NewFromJSObject(objbr); testingutils.AssertErr(t, err) {

							testingutils.AssertErr(t, div.Before(br))

							if val, err := span.OuterHTML(); testingutils.AssertErr(t, err) {
								testingutils.AssertExpect(t, "<span><br><div></div></span>", val)

							}

						}
					}

				}
			}

		}
	}

}
