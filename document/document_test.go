package document

import (
	"strings"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestDomain(t *testing.T) {

	if d, err := New(); err == nil {

		if str, err := d.Domain(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "localhost", str)
		}

		if err := d.SetDomain("testing.com"); err == nil {
			t.Error("Must return error")
		} else {

			testingutils.AssertExpect(t, "Failed to set the 'domain' property on 'Document': 'testing.com' is not a suffix of 'localhost'.", err.Error())

		}

	}

}

func TestTitle(t *testing.T) {

	if d, err := New(); err == nil {

		if str, err := d.Title(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "Go wasm", str)

		}

		testingutils.AssertErr(t, d.SetTitle("Hello"))

		if str, err := d.Title(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "Hello", str)

		}

	}

}

func TestCookie(t *testing.T) {

	if d, err := New(); err == nil {

		if str, err := d.Cookie(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "", str)

		}
		testingutils.AssertErr(t, d.SetCookie("hello world"))
		if str, err := d.Cookie(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "hello world", str)
		}

	}

}

func TestGetBody(t *testing.T) {

	baseobject.Eval("document.body.id=\"hello\"")

	if d, err := New(); testingutils.AssertErr(t, err) {
		if b, err := d.Body(); testingutils.AssertErr(t, err) {

			if id, err := b.ID(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "hello", id)

			}
		}

	}

}

func TestActiveElement(t *testing.T) {
	if d, err := New(); testingutils.AssertErr(t, err) {
		if el, err := d.ActiveElement(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLBodyElement]", el.ToString_())
		}
	}
}

func TestCharacterSet(t *testing.T) {
	if d, err := New(); testingutils.AssertErr(t, err) {
		if s, err := d.CharacterSet(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "UTF-8", s)
		}

	}
}

func TestChildElementCount(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		if c, err := d.ChildElementCount(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, c)
		}

	}
}

func TestChildren(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		if collections, err := d.Children(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLHtmlElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}

	}
}

func TestCompatMode(t *testing.T) {

	//var expectElementString []string = []string{"", ""}
	if d, err := New(); testingutils.AssertErr(t, err) {
		if compatmode, err := d.CompatMode(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "CSS1Compat", compatmode)
		}

	}
}

func TestContentType(t *testing.T) {

	if d, err := New(); err == nil {

		if str, err := d.ContentType(); testingutils.AssertErr(t, err) {
			if str != "text/html" {
				t.Errorf("Content Type must be text/html, have %s", str)
			}
		}

	}

}
func TestDocumentElement(t *testing.T) {

	if d, err := New(); err == nil {

		if elem, err := d.DocumentElement(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}

func TestDocumentURI(t *testing.T) {

	if d, err := New(); err == nil {

		if uri, err := d.DocumentURI(); testingutils.AssertErr(t, err) {
			var expect string = "http://localhost"
			if !strings.Contains(uri, expect) {
				t.Errorf("Must contain %s have %s", expect, uri)
			}

		}

	}
}

func TestEmbeds(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("document.body.appendChild(document.createElement(\"embed\"))")

		if collections, err := d.Embeds(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLEmbedElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}

	}
}

func TestFirstElementChild(t *testing.T) {

	if d, err := New(); err == nil {

		if elem, err := d.FirstElementChild(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}
func TestForms(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("document.body.appendChild(document.createElement(\"form\"))")

		if collections, err := d.Forms(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLFormElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}

	}
}
