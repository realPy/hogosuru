package typedarray

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestNewUInt8Array(t *testing.T) {
	if a, err := NewUint8Array(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 10 {
				t.Errorf("Size must be 10 have %d", l)
			}

		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}

func TestNewInt8Array(t *testing.T) {
	if a, err := NewInt8Array(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 10 {
				t.Errorf("Size must be 10 have %d", l)
			}

		} else {
			t.Error(err.Error())
		}

		a.Fill(200)
		if v, err := a.GetValue(0); err == nil {

			if v.(int) > 0 {
				t.Error("Must be a negative value")
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}

func TestNewUInt16Array(t *testing.T) {
	if a, err := NewUint16Array(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 20 {
				t.Errorf("Size must be 20 have %d", l)
			}

		} else {
			t.Error(err.Error())
		}

		a.SetValue(0, 1000)
		if v, err := a.GetValue(0); err == nil {

			if v.(int) != 1000 {
				t.Error("Must be a egal 1000")
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}

func TestNewUInt8ClampedArray(t *testing.T) {
	if a, err := NewUint8ClampedArray(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 10 {
				t.Errorf("Size must be 10 have %d", l)
			}

		} else {
			t.Error(err.Error())
		}

		a.SetValue(0, 3.6)
		if v, err := a.GetValue(0); err == nil {

			if v.(int) != 4 {
				t.Error("Must be a egal 4")
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}
