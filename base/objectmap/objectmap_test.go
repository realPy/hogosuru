package objectmap

import (
	"testing"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestNew(t *testing.T) {

	a := array.New_(array.New_("a", "b"), array.New_("c"))

	if m, err := New(a); testingutils.AssertErr(t, err) {

		if s, err := m.Size(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 2, s)

		}
	}
}

func TestClear(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); testingutils.AssertErr(t, err) {
		if err := m.Clear(); testingutils.AssertErr(t, err) {
			if s, err := m.Size(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, 0, s)

			}
		}

	}

}

func TestHas(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); testingutils.AssertErr(t, err) {
		if ok, err := m.Has("a"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, true, ok)

		}

		if ok, err := m.Has("c"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, true, ok)

		}

		if ok, err := m.Has("d"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, false, ok)
		}
	}

}

func TestDelete(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); testingutils.AssertErr(t, err) {

		if ok, err := m.Delete("a"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, true, ok)

		}

		if ok, err := m.Has("a"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, false, ok)

		}

	}

}

func TestEntries(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); testingutils.AssertErr(t, err) {

		if it, err := m.Entries(); testingutils.AssertErr(t, err) {

			var i int

			for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {
				aobj1, _ := a.GetValue(i)
				if a1, ok := aobj1.(array.ArrayFrom); ok {
					indexobj1, _ := a1.Array_().GetValue(0)
					valueobj1, _ := a1.Array_().GetValue(1)
					if index.(string) != indexobj1.(string) {
						t.Errorf("index get %s have %s", indexobj1.(string), index.(string))
					}

					if value.(string) != valueobj1.(string) {
						t.Errorf("value get %s have %s", valueobj1.(string), value.(string))
					}

				}

				i++

			}
		}
	}

}

func TestForEach(t *testing.T) {

	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); testingutils.AssertErr(t, err) {

		var i int
		if err := m.ForEach(func(value, index interface{}) {

			aobj1, _ := a.GetValue(i)
			if a1, ok := aobj1.(array.ArrayFrom); ok {
				indexobj1, _ := a1.Array_().GetValue(0)
				valueobj1, _ := a1.Array_().GetValue(1)

				if index.(string) != indexobj1.(string) {
					t.Errorf("index get %s have %s", indexobj1.(string), index.(string))
				}

				if value.(string) != valueobj1.(string) {
					t.Errorf("value get %s have %s", valueobj1.(string), value.(string))
				}

			}
			i++
		}); err != nil {
			t.Error(err.Error())
		}

		if it, err := m.Entries(); testingutils.AssertErr(t, err) {

			var i int

			for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {
				aobj1, _ := a.GetValue(i)
				if a1, ok := aobj1.(array.ArrayFrom); ok {
					indexobj1, _ := a1.Array_().GetValue(0)
					valueobj1, _ := a1.Array_().GetValue(1)
					if index.(string) != indexobj1.(string) {
						t.Errorf("index get %s have %s", indexobj1.(string), index.(string))
					}

					if value.(string) != valueobj1.(string) {
						t.Errorf("value get %s have %s", valueobj1.(string), value.(string))
					}

				}

				i++

			}
		}
	}

}

func TestKeys(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); testingutils.AssertErr(t, err) {

		if it, err := m.Keys(); testingutils.AssertErr(t, err) {

			var i int

			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {
				aobj1, _ := a.GetValue(i)
				if a1, ok := aobj1.(array.ArrayFrom); ok {
					indexobj1, _ := a1.Array_().GetValue(0)

					if value.(string) != indexobj1.(string) {
						t.Errorf("value get %s have %s", indexobj1.(string), value.(string))
					}

				}

				i++

			}
		}
	}
}

func TestValues(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); testingutils.AssertErr(t, err) {

		if it, err := m.Values(); testingutils.AssertErr(t, err) {

			var i int

			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {
				aobj1, _ := a.GetValue(i)
				if a1, ok := aobj1.(array.ArrayFrom); ok {
					valueobj1, _ := a1.Array_().GetValue(1)

					if value.(string) != valueobj1.(string) {
						t.Errorf("value get %s have %s", valueobj1.(string), value.(string))
					}

				}

				i++

			}
		}
	}

}

func TestGet(t *testing.T) {

	a := array.New_(array.New_("a", "b"), array.New_("c"))

	if m, err := New(a); testingutils.AssertErr(t, err) {

		if value, err := m.Get("a"); testingutils.AssertErr(t, err) {

			if value.(string) != "b" {
				t.Errorf("Must get b have %s", value.(string))
			}

		}
	}

}
func TestSet(t *testing.T) {

	a := array.New_(array.New_("a", "b"), array.New_("c"))

	if m, err := New(a); testingutils.AssertErr(t, err) {

		if err := m.Set("hello", "World"); testingutils.AssertErr(t, err) {

			if s, err := m.Size(); err == nil {

				if s != 3 {
					t.Errorf("Size must be 3 have %d", s)
				}

			} else {
				t.Error(err.Error())
			}

			if value, err := m.Get("hello"); err == nil {

				if value.(string) != "World" {
					t.Errorf("Must get World have %s", value.(string))
				}

			} else {
				t.Error(err.Error())
			}
		}
	}

}
func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
