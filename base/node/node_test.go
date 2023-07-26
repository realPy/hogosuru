package node

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`node= document.createElement("title")
	node.text="hello"
	div=document.createElement("div")
	node.appendChild(div)
	span=document.createElement("span")
	span.appendChild(node)
	br=document.createElement("br")
	span.appendChild(br)
	document.body.appendChild(span)


	`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "node"); testingutils.AssertErr(t, err) {

		if ti, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLTitleElement", ti.ConstructName_())
		}

	}

}

func TestNodeValue(t *testing.T) {
	if obj, err := baseobject.Get(js.Global(), "node"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if c, err := node.FirstChild(); testingutils.AssertErr(t, err) {
				if v, err := c.NodeValue(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "hello", v)
				}

			}
		}

	}

}

func TestSetNodeValue(t *testing.T) {

	baseobject.Eval(`nodeset= document.createElement("title")
	nodeset.text="hello"
	`)
	if obj, err := baseobject.Get(js.Global(), "nodeset"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if c, err := node.FirstChild(); testingutils.AssertErr(t, err) {

				testingutils.AssertErr(t, c.SetNodeValue("world"))

				if v, err := c.NodeValue(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "world", v)
				}

			}
		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "BaseURI", "type": "contains", "resultattempt": "http://localhost"},
	{"method": "FirstChild", "type": "constructnamechecking", "resultattempt": "Text"},
	{"method": "IsConnected", "resultattempt": true},
	{"method": "LastChild", "type": "constructnamechecking", "resultattempt": "HTMLDivElement"},
	{"method": "NextSibling", "type": "constructnamechecking", "resultattempt": "HTMLBRElement"},
	{"method": "NodeName", "resultattempt": "TITLE"},
	{"method": "NodeType", "resultattempt": 1},
	{"method": "OwnerDocument", "type": "constructnamechecking", "resultattempt": "HTMLDocument"},
	{"method": "ParentNode", "type": "constructnamechecking", "resultattempt": "HTMLSpanElement"},
	{"method": "ParentElement", "type": "constructnamechecking", "resultattempt": "HTMLSpanElement"},
	{"method": "TextContent", "resultattempt": "hello"},
	{"method": "SetTextContent", "args": []interface{}{"mytitle"}, "gettermethod": "TextContent", "resultattempt": "mytitle"},
	{"method": "GetRootNode", "type": "constructnamechecking", "resultattempt": "HTMLDocument"},
	{"method": "IsDefaultNamespace", "args": []interface{}{"none"}, "resultattempt": false},
	{"method": "LookupPrefix", "args": []interface{}{"none"}, "resultattempt": ""},
	{"method": "LookupNamespaceURI", "args": []interface{}{"none"}, "resultattempt": nil},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "node"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, node, result)
			}

		}

	}
}

func TestPreviousSibling(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "br"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if c, err := node.PreviousSibling(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "HTMLTitleElement", c.ConstructName_())

			}
		}

	}

}

func TestAppendChild(t *testing.T) {

	baseobject.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
`)
	if obj, err := baseobject.Get(js.Global(), "appendnode"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if divobj, err := baseobject.Get(js.Global(), "div"); testingutils.AssertErr(t, err) {

				if div, err := NewFromJSObject(divobj); testingutils.AssertErr(t, err) {

					testingutils.AssertErr(t, node.AppendChild(div))

					if c, err := node.FirstChild(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "HTMLDivElement", c.ConstructName_())

					}

				}
			}

		}

	}

}

func TestCloneNode(t *testing.T) {

	baseobject.Eval(`clonenode= document.createElement("title")
	clonenode.text="hello"
`)
	if obj, err := baseobject.Get(js.Global(), "clonenode"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if clone, err := node.CloneNode(true); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "HTMLTitleElement", clone.ConstructName_())
				testingutils.AssertErr(t, clone.SetTextContent("world"))

				v1, _ := clone.TextContent()
				v2, _ := node.TextContent()
				testingutils.AssertExpect(t, true, v1 != v2)

			}

		}

	}

}

func TestCompareDocumentPosition(t *testing.T) {

	baseobject.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
	appendnode.appendChild(div)

`)
	if obj, err := baseobject.Get(js.Global(), "appendnode"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if divobj, err := baseobject.Get(js.Global(), "div"); testingutils.AssertErr(t, err) {

				if div, err := NewFromJSObject(divobj); testingutils.AssertErr(t, err) {

					if n, err := node.CompareDocumentPosition(div); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, 20, n)

					}

				}
			}

		}

	}

}

func TestContains(t *testing.T) {

	baseobject.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
	appendnode.appendChild(div)

`)
	if obj, err := baseobject.Get(js.Global(), "appendnode"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if divobj, err := baseobject.Get(js.Global(), "div"); testingutils.AssertErr(t, err) {

				if div, err := NewFromJSObject(divobj); testingutils.AssertErr(t, err) {

					if b, err := node.Contains(div); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, true, b)

					}

					node.RemoveChild(div)

					if b, err := node.Contains(div); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, false, b)

					}
				}
			}

		}

	}

}

func TestHasChildNodes(t *testing.T) {

	baseobject.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
	appendnode.appendChild(div)

`)
	if obj, err := baseobject.Get(js.Global(), "appendnode"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if divobj, err := baseobject.Get(js.Global(), "div"); testingutils.AssertErr(t, err) {

				if div, err := NewFromJSObject(divobj); testingutils.AssertErr(t, err) {

					if b, err := node.HasChildNodes(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, true, b)

					}
					node.RemoveChild(div)

					if b, err := node.HasChildNodes(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, false, b)

					}
				}
			}

		}

	}

}

func TestInsertBefore(t *testing.T) {

	baseobject.Eval(`appendnode= document.createElement("title")
	div=document.createElement("div")
	span=document.createElement("span")
	appendnode.appendChild(div)

`)
	if obj, err := baseobject.Get(js.Global(), "appendnode"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if divobj, err := baseobject.Get(js.Global(), "div"); testingutils.AssertErr(t, err) {

				if div, err := NewFromJSObject(divobj); testingutils.AssertErr(t, err) {

					if spanobj, err := baseobject.Get(js.Global(), "span"); testingutils.AssertErr(t, err) {

						if span, err := NewFromJSObject(spanobj); testingutils.AssertErr(t, err) {

							if n, err := node.InsertBefore(span, div); testingutils.AssertErr(t, err) {

								testingutils.AssertExpect(t, "HTMLSpanElement", n.ConstructName_())

								if next, err := node.FirstChild(); testingutils.AssertErr(t, err) {

									testingutils.AssertExpect(t, "HTMLSpanElement", next.ConstructName_())

								}

							}

						}
					}

				}
			}

		}

	}

}

func TestIsEqualNode(t *testing.T) {

	baseobject.Eval(`clonenode= document.createElement("title")
	clonenode.text="hello"

`)
	if obj, err := baseobject.Get(js.Global(), "clonenode"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if clone, err := node.CloneNode(true); testingutils.AssertErr(t, err) {

				if b, err := clone.IsEqualNode(node); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, true, b)

				}
				clone.SetTextContent("world")
				if b, err := clone.IsEqualNode(node); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, false, b)

				}

			}

		}

	}

}

func TestIsSameNode(t *testing.T) {

	baseobject.Eval(`clonenode= document.createElement("title")
	clonenode.text="hello"

`)
	if obj, err := baseobject.Get(js.Global(), "clonenode"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if clone, err := node.CloneNode(true); testingutils.AssertErr(t, err) {

				if b, err := clone.IsSameNode(node); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, false, b)

				}
				if b, err := node.IsSameNode(node); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, true, b)

				}

			}

		}

	}

}

func TestNormalize(t *testing.T) {
	baseobject.Eval(`node= document.createElement("title")
	text1=  document.createTextNode("-Partie 1 ")
	text2=  document.createTextNode("Partie 2 -")
`)
	if obj, err := baseobject.Get(js.Global(), "node"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if text2obj, err := baseobject.Get(js.Global(), "text1"); testingutils.AssertErr(t, err) {

				if text1, err := NewFromJSObject(text2obj); testingutils.AssertErr(t, err) {

					if text2obj, err := baseobject.Get(js.Global(), "text2"); testingutils.AssertErr(t, err) {

						if text2, err := NewFromJSObject(text2obj); testingutils.AssertErr(t, err) {

							node.AppendChild(text1)
							node.AppendChild(text2)

							testingutils.AssertErr(t, node.Normalize())

							if textn, err := node.FirstChild(); testingutils.AssertErr(t, err) {

								testingutils.AssertExpect(t, "-Partie 1 Partie 2 -", textn.TextContent_())

							}

						}
					}

				}
			}

		}
	}

}

func TestRemoveChild(t *testing.T) {
	baseobject.Eval(`node= document.createElement("title")
	text1=  document.createTextNode("-Partie 1 ")
	text2=  document.createTextNode("Partie 2 -")
`)
	if obj, err := baseobject.Get(js.Global(), "node"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if text2obj, err := baseobject.Get(js.Global(), "text1"); testingutils.AssertErr(t, err) {

				if text1, err := NewFromJSObject(text2obj); testingutils.AssertErr(t, err) {

					if text2obj, err := baseobject.Get(js.Global(), "text2"); testingutils.AssertErr(t, err) {

						if text2, err := NewFromJSObject(text2obj); testingutils.AssertErr(t, err) {

							node.AppendChild(text1)
							node.AppendChild(text2)

							if removetext, err := node.RemoveChild(text1); testingutils.AssertErr(t, err) {

								testingutils.AssertExpect(t, "Text", removetext.ConstructName_())

								if n, err := node.FirstChild(); testingutils.AssertErr(t, err) {

									testingutils.AssertExpect(t, "Partie 2 -", n.TextContent_())

								}

							}

						}
					}

				}
			}

		}
	}

}

func TestReplaceChild(t *testing.T) {
	baseobject.Eval(`node= document.createElement("title")
	text1=  document.createTextNode("-Partie 1 ")
	text2=  document.createTextNode("Partie 2 -")
	text3=  document.createTextNode("Partie 3 -")
`)
	if obj, err := baseobject.Get(js.Global(), "node"); testingutils.AssertErr(t, err) {

		if node, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if text2obj, err := baseobject.Get(js.Global(), "text1"); testingutils.AssertErr(t, err) {

				if text1, err := NewFromJSObject(text2obj); testingutils.AssertErr(t, err) {

					if text2obj, err := baseobject.Get(js.Global(), "text2"); testingutils.AssertErr(t, err) {

						if text2, err := NewFromJSObject(text2obj); testingutils.AssertErr(t, err) {

							node.AppendChild(text1)
							node.AppendChild(text2)
							if text3obj, err := baseobject.Get(js.Global(), "text3"); testingutils.AssertErr(t, err) {

								if text3, err := NewFromJSObject(text3obj); testingutils.AssertErr(t, err) {
									if replace, err := node.ReplaceChild(text3, text1); testingutils.AssertErr(t, err) {

										testingutils.AssertExpect(t, "-Partie 1 ", replace.TextContent_())
										if n, err := node.FirstChild(); testingutils.AssertErr(t, err) {

											testingutils.AssertExpect(t, "Partie 3 -", n.TextContent_())

										}

									}

								}
							}

						}
					}

				}
			}

		}
	}

}
