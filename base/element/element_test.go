package element

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/attr"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/htmlcollection"
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
	div.id="pouet"
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

	{"method": "QuerySelector", "args": []interface{}{"#pouet"}, "type": "constructnamechecking", "resultattempt": "HTMLDivElement"},
	{"method": "QuerySelectorAll", "args": []interface{}{"div"}, "type": "constructnamechecking", "resultattempt": "NodeList"},

	{"method": "SetID", "args": []interface{}{"test"}, "gettermethod": "ID", "resultattempt": "test"},
	{"method": "InnerHTML", "resultattempt": "<div id=\"pouet\"></div>"},
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
	{"method": "Matches", "args": []interface{}{"#test"}, "resultattempt": true},

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

func TestPrepend(t *testing.T) {

	baseobject.Eval(`
	elementspanap= document.createElement("span")
	elementdivap= document.createElement("div")

	`)

	if objspan, err := baseobject.Get(js.Global(), "elementspanap"); testingutils.AssertErr(t, err) {

		if span, err := NewFromJSObject(objspan); testingutils.AssertErr(t, err) {

			if objdiv, err := baseobject.Get(js.Global(), "elementdivap"); testingutils.AssertErr(t, err) {

				if div, err := NewFromJSObject(objdiv); testingutils.AssertErr(t, err) {

					testingutils.AssertErr(t, span.Prepend(div))

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

func TestInsertAdjacentElement(t *testing.T) {

	baseobject.Eval(`
	elementp= document.createElement("p")
	elementp.textContent="hello"
	elementbri= document.createElement("br")

	`)

	if objp, err := baseobject.Get(js.Global(), "elementp"); testingutils.AssertErr(t, err) {

		if p, err := NewFromJSObject(objp); testingutils.AssertErr(t, err) {

			if objbr, err := baseobject.Get(js.Global(), "elementbri"); testingutils.AssertErr(t, err) {

				if br, err := NewFromJSObject(objbr); testingutils.AssertErr(t, err) {

					if elem, err := p.InsertAdjacentElement("afterbegin", br); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, "[object HTMLBRElement]", elem.ToString_())
						if val, err := p.OuterHTML(); testingutils.AssertErr(t, err) {
							testingutils.AssertExpect(t, "<p><br>hello</p>", val)

						}

					}

				}
			}

		}
	}

}

func TestInsertAdjacentHTML(t *testing.T) {

	baseobject.Eval(`
	elementp= document.createElement("p")
	elementp.textContent="hello"
	elementbri= document.createElement("br")

	`)

	if objp, err := baseobject.Get(js.Global(), "elementp"); testingutils.AssertErr(t, err) {

		if p, err := NewFromJSObject(objp); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, p.InsertAdjacentHTML("afterbegin", "<div>test</div>"))

			if val, err := p.OuterHTML(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "<p><div>test</div>hello</p>", val)

			}

		}
	}

}

func TestInsertAdjacentText(t *testing.T) {

	baseobject.Eval(`
	elementp= document.createElement("p")
	elementp.textContent="hello"
	elementbri= document.createElement("br")

	`)

	if objp, err := baseobject.Get(js.Global(), "elementp"); testingutils.AssertErr(t, err) {

		if p, err := NewFromJSObject(objp); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, p.InsertAdjacentText("beforeend", "this is text <br>"))

			if val, err := p.OuterHTML(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "<p>hellothis is text &lt;br&gt;</p>", val)

			}

		}
	}

}

func TestRemove(t *testing.T) {

	baseobject.Eval(`
	p= document.createElement("p")
	div= document.createElement("div")
	p.appendChild(div)

	`)

	if objdiv, err := baseobject.Get(js.Global(), "div"); testingutils.AssertErr(t, err) {

		if div, err := NewFromJSObject(objdiv); testingutils.AssertErr(t, err) {

			if objp, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

				if p, err := NewFromJSObject(objp); testingutils.AssertErr(t, err) {

					testingutils.AssertErr(t, div.Remove())
					if val, err := p.OuterHTML(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "<p></p>", val)

					}

				}
			}

		}
	}

}

func TestRemoveAttribute(t *testing.T) {

	baseobject.Eval(`
	p= document.createElement("p")
	p.setAttribute("hello","world")
	p.setAttribute("hello1","world1")
	`)

	if objp, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := NewFromJSObject(objp); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, p.RemoveAttribute("hello"))
			if val, err := p.OuterHTML(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "<p hello1=\"world1\"></p>", val)

			}

		}
	}

}

func TestRemoveAttributeNS(t *testing.T) {

	baseobject.Eval(`
	p= document.createElement("p")
	p.setAttribute("hello","world")
	p.setAttribute("hello1","world1")
	p.setAttributeNS("space","hello","world")
	`)

	if objp, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := NewFromJSObject(objp); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, p.RemoveAttributeNS("space", "hello"))
			if val, err := p.OuterHTML(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "<p hello=\"world\" hello1=\"world1\"></p>", val)

			}

		}
	}

}

func TestReplaceChildren(t *testing.T) {
	baseobject.Eval(`
	div= document.createElement("div")
	pTemp = document.createElement("p")
	pTemp.innerText = "remove me"
	div.append(pTemp)
	span = document.createElement("span")
	span.innerText = "done"
	`)

	if objdiv, err := baseobject.Get(js.Global(), "div"); testingutils.AssertErr(t, err) {
		if div, err := NewFromJSObject(objdiv); testingutils.AssertErr(t, err) {
			if objspan, err := baseobject.Get(js.Global(), "span"); testingutils.AssertErr(t, err) {
				if span, err := NewFromJSObject(objspan); testingutils.AssertErr(t, err) {
					testingutils.AssertErr(t, div.ReplaceChildren("well ", span.Node))
					if val, err := div.OuterHTML(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "<div>well <span>done</span></div>", val)
					}
				}
			}
		}
	}
}

func TestSetAttribute(t *testing.T) {

	baseobject.Eval(`
	p= document.createElement("p")

	`)

	if objp, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := NewFromJSObject(objp); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, p.SetAttribute("space", "hello"))
			if val, err := p.OuterHTML(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "<p space=\"hello\"></p>", val)

			}

		}
	}

}

func TestSetAttributeNS(t *testing.T) {

	baseobject.Eval(`
	p= document.createElement("p")

	`)

	if objp, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := NewFromJSObject(objp); testingutils.AssertErr(t, err) {

			testingutils.AssertErr(t, p.SetAttributeNS("space", "hello", "world"))
			if val, err := p.OuterHTML(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "<p hello=\"world\"></p>", val)

			}

		}
	}

}

func TestToggleAttribute(t *testing.T) {

	baseobject.Eval(`
	p= document.createElement("p")
	`)

	if objp, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := NewFromJSObject(objp); testingutils.AssertErr(t, err) {

			if b, err := p.ToggleAttribute("disabled", true); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, true, b)
				if val, err := p.OuterHTML(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "<p disabled=\"\"></p>", val)

				}
			}

		}
	}

}

var methodsAttemptAccessibility []map[string]interface{} = []map[string]interface{}{

	{"method": "AriaAtomic", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaAtomic", "args": []interface{}{"true"}, "gettermethod": "AriaAtomic", "resultattempt": "true"},

	{"method": "AriaAutoComplete", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaAutoComplete", "args": []interface{}{"inline"}, "gettermethod": "AriaAutoComplete", "resultattempt": "inline"},

	{"method": "AriaBusy", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaBusy", "args": []interface{}{"true"}, "gettermethod": "AriaBusy", "resultattempt": "true"},

	{"method": "AriaChecked", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaChecked", "args": []interface{}{"true"}, "gettermethod": "AriaChecked", "resultattempt": "true"},

	{"method": "AriaColCount", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaColCount", "args": []interface{}{"2"}, "gettermethod": "AriaColCount", "resultattempt": "2"},

	{"method": "AriaColIndex", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaColIndex", "args": []interface{}{"1"}, "gettermethod": "AriaColIndex", "resultattempt": "1"},

	{"method": "AriaColIndexText", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaColIndexText", "args": []interface{}{"11"}, "gettermethod": "AriaColIndexText", "resultattempt": "11"},

	{"method": "AriaColSpan", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaColSpan", "args": []interface{}{"1"}, "gettermethod": "AriaColSpan", "resultattempt": "1"},

	{"method": "AriaCurrent", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaCurrent", "args": []interface{}{"page"}, "gettermethod": "AriaCurrent", "resultattempt": "page"},

	{"method": "AriaDescription", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaDescription", "args": []interface{}{"test"}, "gettermethod": "AriaDescription", "resultattempt": "test"},

	{"method": "AriaDisabled", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaDisabled", "args": []interface{}{"true"}, "gettermethod": "AriaDisabled", "resultattempt": "true"},

	{"method": "AriaExpanded", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaExpanded", "args": []interface{}{"true"}, "gettermethod": "AriaExpanded", "resultattempt": "true"},

	{"method": "AriaHasPopup", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaHasPopup", "args": []interface{}{"true"}, "gettermethod": "AriaHasPopup", "resultattempt": "true"},

	{"method": "AriaHidden", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaHidden", "args": []interface{}{"true"}, "gettermethod": "AriaHidden", "resultattempt": "true"},

	{"method": "AriaKeyShortcuts", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaKeyShortcuts", "args": []interface{}{"true"}, "gettermethod": "AriaKeyShortcuts", "resultattempt": "true"},

	{"method": "AriaLabel", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaLabel", "args": []interface{}{"true"}, "gettermethod": "AriaLabel", "resultattempt": "true"},

	{"method": "AriaLevel", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaLevel", "args": []interface{}{"1"}, "gettermethod": "AriaLevel", "resultattempt": "1"},

	{"method": "AriaLive", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaLive", "args": []interface{}{"assertive"}, "gettermethod": "AriaLive", "resultattempt": "assertive"},

	{"method": "AriaModal", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaModal", "args": []interface{}{"true"}, "gettermethod": "AriaModal", "resultattempt": "true"},

	{"method": "AriaMultiline", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaMultiline", "args": []interface{}{"true"}, "gettermethod": "AriaMultiline", "resultattempt": "true"},

	{"method": "AriaMultiSelectable", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaMultiSelectable", "args": []interface{}{"true"}, "gettermethod": "AriaMultiSelectable", "resultattempt": "true"},

	{"method": "AriaOrientation", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaOrientation", "args": []interface{}{"horizontal"}, "gettermethod": "AriaOrientation", "resultattempt": "horizontal"},

	{"method": "AriaPlaceholder", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaPlaceholder", "args": []interface{}{"true"}, "gettermethod": "AriaPlaceholder", "resultattempt": "true"},

	{"method": "AriaPosInSet", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaPosInSet", "args": []interface{}{"1"}, "gettermethod": "AriaPosInSet", "resultattempt": "1"},

	{"method": "AriaPressed", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaPressed", "args": []interface{}{"true"}, "gettermethod": "AriaPressed", "resultattempt": "true"},

	{"method": "AriaReadOnly", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaReadOnly", "args": []interface{}{"true"}, "gettermethod": "AriaReadOnly", "resultattempt": "true"},

	{"method": "AriaRelevant", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaRelevant", "args": []interface{}{"text"}, "gettermethod": "AriaRelevant", "resultattempt": "text"},

	{"method": "AriaRequired", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaRequired", "args": []interface{}{"true"}, "gettermethod": "AriaRequired", "resultattempt": "true"},

	{"method": "AriaRoleDescription", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaRoleDescription", "args": []interface{}{"test"}, "gettermethod": "AriaRoleDescription", "resultattempt": "test"},

	{"method": "AriaRowCount", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaRowCount", "args": []interface{}{"1"}, "gettermethod": "AriaRowCount", "resultattempt": "1"},

	{"method": "AriaRowIndex", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaRowIndex", "args": []interface{}{"1"}, "gettermethod": "AriaRowIndex", "resultattempt": "1"},

	{"method": "AriaRowIndexText", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaRowIndexText", "args": []interface{}{"1"}, "gettermethod": "AriaRowIndexText", "resultattempt": "1"},

	{"method": "AriaRowSpan", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaRowSpan", "args": []interface{}{"1"}, "gettermethod": "AriaRowSpan", "resultattempt": "1"},

	{"method": "AriaSelected", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaSelected", "args": []interface{}{"true"}, "gettermethod": "AriaSelected", "resultattempt": "true"},

	{"method": "AriaSetSize", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaSetSize", "args": []interface{}{"true"}, "gettermethod": "AriaSetSize", "resultattempt": "true"},

	{"method": "AriaSort", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaSort", "args": []interface{}{"ascending"}, "gettermethod": "AriaSort", "resultattempt": "ascending"},

	{"method": "AriaValueMax", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaValueMax", "args": []interface{}{"9"}, "gettermethod": "AriaValueMax", "resultattempt": "9"},

	{"method": "AriaValueMin", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaValueMin", "args": []interface{}{"2"}, "gettermethod": "AriaValueMin", "resultattempt": "2"},

	{"method": "AriaValueNow", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaValueNow", "args": []interface{}{"2"}, "gettermethod": "AriaValueNow", "resultattempt": "2"},

	{"method": "AriaValueText", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "SetAriaValueText", "args": []interface{}{"2"}, "gettermethod": "AriaValueText", "resultattempt": "2"},
}

func TestMethodsAccessibility(t *testing.T) {

	baseobject.Eval(`
	
	elementaccessibility= document.createElement("title")
	`)

	if obj, err := baseobject.Get(js.Global(), "elementaccessibility"); testingutils.AssertErr(t, err) {

		if elem, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttemptAccessibility {
				testingutils.InvokeCheck(t, elem, result)
			}

		}

	}
}
