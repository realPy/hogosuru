package documentfragment

import (
	"errors"
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()`)

	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if o, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object DocumentFragment]", o.ToString_())

		}
	}

}

func TestNew(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()`)
	if d, err := New(); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, "[object DocumentFragment]", d.ToString_())

	}
}

func TestChildElementCount(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	df.append(div)
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if n, err := d.ChildElementCount(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, 1, n)
			}

		}
	}

}

func TestChildren(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if c, err := d.Children(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLCollection]", c.ToString_())
			}

		}
	}

}

func TestFirstElementChild(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if c, err := d.FirstElementChild(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLDivElement]", c.ToString_())
			}

		}
	}

}

func TestLastElementChild(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if c, err := d.LastElementChild(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLSpanElement]", c.ToString_())
			}

		}
	}

}

func TestQueryAppend(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	b=document.createElement("b")
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if divobj, err := baseobject.Get(js.Global(), "div"); testingutils.AssertErr(t, err) {

				if e, err := element.NewFromJSObject(divobj); testingutils.AssertErr(t, err) {
					testingutils.AssertErr(t, d.Append(e))

					if count, err := d.ChildElementCount(); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, 1, count)
					}

					span, _ := baseobject.Get(js.Global(), "span")
					b, _ := baseobject.Get(js.Global(), "b")
					es, _ := element.NewFromJSObject(span)
					eb, _ := element.NewFromJSObject(b)

					testingutils.AssertErr(t, d.Append(es, eb))

					if count, err := d.ChildElementCount(); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, 3, count)
					}

					if c, err := d.FirstElementChild(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "[object HTMLDivElement]", c.ToString_())
					}
					if c, err := d.LastElementChild(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "[object HTMLElement]", c.ToString_())
					}

				}

			}

		}
	}

}

func TestQueryPrepend(t *testing.T) {

	baseobject.Eval(`
	df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	b=document.createElement("b")
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if divobj, err := baseobject.Get(js.Global(), "div"); testingutils.AssertErr(t, err) {

				if e, err := element.NewFromJSObject(divobj); testingutils.AssertErr(t, err) {
					testingutils.AssertErr(t, d.Prepend(e))

					if count, err := d.ChildElementCount(); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, 1, count)
					}

					span, _ := baseobject.Get(js.Global(), "span")
					b, _ := baseobject.Get(js.Global(), "b")
					es, _ := element.NewFromJSObject(span)
					eb, _ := element.NewFromJSObject(b)

					testingutils.AssertErr(t, d.Prepend(es, eb))

					if count, err := d.ChildElementCount(); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, 3, count)
					}

					if c, err := d.FirstElementChild(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "[object HTMLSpanElement]", c.ToString_())
					}
					if c, err := d.LastElementChild(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "[object HTMLDivElement]", c.ToString_())
					}
				}

			}

		}
	}

}

func TestQuerySelector(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if c, err := d.QuerySelector("div"); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLDivElement]", c.ToString_())
			}

			_, err := d.QuerySelector("button")
			testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))

		}
	}

}

func TestQuerySelectorAll(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if c, err := d.QuerySelectorAll("div"); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object NodeList]", c.ToString_())
			}
		}
	}

}

func TestReplaceChild(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	df.append(div)
	df.append(span)
	b=document.createElement("b")
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			b, _ := baseobject.Get(js.Global(), "b")

			eb, _ := element.NewFromJSObject(b)

			if c, err := d.FirstElementChild(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLDivElement]", c.ToString_())

				if old, err := d.ReplaceChild(eb.Node, c.Node); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "[object HTMLDivElement]", old.ToString_())
					if c2, err := d.FirstElementChild(); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "[object HTMLElement]", c2.ToString_())

					}
				}

			}

		}
	}

}
func TestGetElementById(t *testing.T) {

	baseobject.Eval(`df=new DocumentFragment()
	div=document.createElement("div")
	span=document.createElement("span")
	div.id="test"
	df.append(div)
	span.id="test2"
	df.append(span)
	
	`)
	if obj, err := baseobject.Get(js.Global(), "df"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			if c, err := d.GetElementById("test"); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLDivElement]", c.ToString_())
			}
			if c, err := d.GetElementById("test2"); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLSpanElement]", c.ToString_())
			}
			_, err := d.GetElementById("unknown")
			testingutils.AssertExpect(t, true, errors.Is(err, baseobject.ErrUndefinedValue))
		}
	}

}
