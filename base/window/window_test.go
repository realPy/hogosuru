package window

import (
	"strings"
	"testing"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	hogosuru.Init()
	m.Run()
}

func TestWindow(t *testing.T) {

	if w, err := New(); testingutils.AssertErr(t, err) {

		if d, err := w.Document(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLDocument]", d.ToString_())
		}

	}

}

func TestHistory(t *testing.T) {

	if w, err := New(); testingutils.AssertErr(t, err) {

		if h, err := w.History(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object History]", h.ToString_())
		}

	}

}
func TestLocation(t *testing.T) {

	if w, err := New(); testingutils.AssertErr(t, err) {

		if l, err := w.Location(); testingutils.AssertErr(t, err) {
			var expect string = "http://localhost"
			if !strings.Contains(l.ToString_(), expect) {
				t.Errorf("Must contain %s have %s", expect, l.ToString_())
			}
		}

	}

}

func TestLocalStorage(t *testing.T) {

	if w, err := New(); testingutils.AssertErr(t, err) {

		if l, err := w.LocalStorage(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Storage]", l.ToString_())
		}

	}

}

func TestSessionStorage(t *testing.T) {

	if w, err := New(); testingutils.AssertErr(t, err) {

		if l, err := w.SessionStorage(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Storage]", l.ToString_())
		}

	}

}

func TestIndexdedDB(t *testing.T) {

	if w, err := New(); testingutils.AssertErr(t, err) {

		if i, err := w.IndexdedDB(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object IDBFactory]", i.ToString_())
		}

	}

}

func TestNavigator(t *testing.T) {

	if w, err := New(); testingutils.AssertErr(t, err) {

		if i, err := w.Navigator(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Navigator]", i.ToString_())
		}

	}

}

func TestAtob(t *testing.T) {

	v, err := Atob("SGVsbG93b3JsZA==")
	testingutils.AssertExpect(t, "Helloworld", v)
	testingutils.AssertExpect(t, nil, err)

}

func TestBtoa(t *testing.T) {

	v, err := Btoa("Helloworld")
	testingutils.AssertExpect(t, "SGVsbG93b3JsZA==", v)
	testingutils.AssertExpect(t, nil, err)

}
