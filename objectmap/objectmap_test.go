package objectmap

import (
	"testing"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
)

func TestNew(t *testing.T) {

	a := array.New_(array.New_("a", "b"), array.New_("c"))

	if m, err := New(a); err == nil {

		if s, err := m.Size(); err == nil {

			if s != 2 {
				t.Errorf("Size must be 2 have %d", s)
			}

		} else {
			t.Error(err.Error())
		}
	} else {
		t.Error(err.Error())
	}
}

func TestClear(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); err == nil {
		if err := m.Clear(); err == nil {
			if s, err := m.Size(); err == nil {

				if s != 0 {
					t.Errorf("Size must be 0 have %d", s)
				}

			} else {
				t.Error(err.Error())
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}

func TestHas(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); err == nil {
		if ok, err := m.Has("a"); err == nil {
			if !ok {
				t.Error("Must have a key")

			}
		} else {
			t.Error(err.Error())
		}

		if ok, err := m.Has("c"); err == nil {
			if !ok {
				t.Error("Must have c key")

			}
		} else {
			t.Error(err.Error())
		}

		if ok, err := m.Has("d"); err == nil {
			if ok {
				t.Error("Must not have d key")

			}
		} else {
			t.Error(err.Error())
		}
	} else {
		t.Error(err.Error())
	}

}

func TestDelete(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); err == nil {

		if ok, err := m.Delete("a"); err == nil {

			if !ok {
				t.Error("Must delete a key")
			}

		} else {
			t.Error(err.Error())
		}

		if ok, err := m.Has("a"); err == nil {
			if ok {
				t.Error("A must be delete")
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}

func TestEntries(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); err == nil {

		if it, err := m.Entries(); err == nil {

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
		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Error(err.Error())
	}

}

func TestForEach(t *testing.T) {

	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); err == nil {

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

		if it, err := m.Entries(); err == nil {

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
		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Error(err.Error())
	}

}

func TestKeys(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); err == nil {

		if it, err := m.Keys(); err == nil {

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
		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Error(err.Error())
	}

}

func TestValues(t *testing.T) {
	a := array.New_(array.New_("a", "b"), array.New_("c", "d"))

	if m, err := New(a); err == nil {

		if it, err := m.Values(); err == nil {

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
		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Error(err.Error())
	}

}

func TestGet(t *testing.T) {

	a := array.New_(array.New_("a", "b"), array.New_("c"))

	if m, err := New(a); err == nil {

		if value, err := m.Get("a"); err == nil {

			if value.(string) != "b" {
				t.Errorf("Must get b have %s", value.(string))
			}

		} else {
			t.Error(err.Error())
		}
	} else {
		t.Error(err.Error())
	}

}
func TestSet(t *testing.T) {

	a := array.New_(array.New_("a", "b"), array.New_("c"))

	if m, err := New(a); err == nil {

		if err := m.Set("hello", "World"); err == nil {

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
		} else {
			t.Error(err.Error())
		}
	} else {
		t.Error(err.Error())
	}

}
func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
