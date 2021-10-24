package testingutils

import "testing"

func AssertErr(t *testing.T, err error) bool {

	if err != nil {
		t.Errorf(err.Error())
		return false
	}

	return true
}

func AssertExpect(t *testing.T, exp interface{}, get interface{}) bool {

	switch expval := exp.(type) {
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

	}

	return false
}
