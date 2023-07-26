package baseobject

import (
	"syscall/js"
	"testing"
)

func TestIsInteger(t *testing.T) {

	var obj js.Value
	var err error

	t.Run("1 is Int", func(t *testing.T) {
		Eval("intvalue=1")

		if obj, err = Get(js.Global(), "intvalue"); err == nil {

			if b, err := IsInteger(obj); err == nil {
				if b != true {
					t.Errorf("Must return true")
				}
			} else {
				t.Errorf(err.Error())
			}

		} else {
			t.Errorf(err.Error())
		}

	})

	t.Run("1.3 is Float", func(t *testing.T) {
		Eval("intvalue=1.3")

		if obj, err = Get(js.Global(), "intvalue"); err == nil {

			if b, err := IsInteger(obj); err == nil {
				if b != false {
					t.Errorf("Must return false")
				}
			} else {
				t.Errorf(err.Error())
			}

		} else {
			t.Errorf(err.Error())
		}

	})

	t.Run("str is not int", func(t *testing.T) {
		Eval("intvalue='hello'")

		if obj, err = Get(js.Global(), "intvalue"); err == nil {

			if b, err := IsInteger(obj); err == nil {
				if b != false {
					t.Errorf("Must return false")
				}
			} else {
				t.Errorf(err.Error())
			}

		} else {
			t.Errorf(err.Error())
		}

	})

}
