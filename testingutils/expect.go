package testingutils

import (
	"testing"
)

func AssertErr(t *testing.T, err error) bool {

	if err != nil {
		t.Errorf(err.Error())
		return false
	}

	return true
}

func AssertExpect(t *testing.T, exp interface{}, get interface{}) bool {

	switch expval := exp.(type) {
	case error:
		if getval, ok := get.(error); ok {
			if expval.Error() != getval.Error() {
				t.Errorf("Expect %s have %s", expval.Error(), getval.Error())
			}
		} else {
			t.Errorf("Expect type var not match")
		}
	case int:
		if getval, ok := get.(int); ok {
			if expval != getval {
				t.Errorf("Expect %d have %d", expval, getval)
			} else {
				return true
			}

		} else {
			t.Errorf("Expect type var not match")
		}

		return false
	case int64:
		if getval, ok := get.(int64); ok {
			if expval != getval {
				t.Errorf("Expect %d have %d", expval, getval)
			} else {
				return true
			}

		} else {
			t.Errorf("Expect type var not match")
		}

		return false

	case string:
		if getval, ok := get.(string); ok {
			if expval != getval {
				t.Errorf("Expect %s have %s", expval, getval)
			} else {
				return true
			}

		} else {
			t.Errorf("Expect type var not match")
		}
		return false
	case bool:
		if getval, ok := get.(bool); ok {
			if expval != getval {
				t.Errorf("Expect %t have %t", expval, getval)
			} else {
				return true
			}

		} else {
			t.Errorf("Expect type var not match")
		}
		return false
	default:

		t.Errorf("Undefined expect type")
	}

	return false
}
