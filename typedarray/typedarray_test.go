package typedarray

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNewUInt8Array(t *testing.T) {
	if a, err := NewUint8Array(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 10 {
				t.Errorf("Size must be 10 have %d", l)
			}

			if b, err := a.Subarray(2, 8); err == nil {
				if bb, ok := b.(Uint8ArrayFrom); ok {
					if l1, err := bb.Uint8Array_().ByteLength(); err == nil {
						if l1 != 6 {
							t.Error("Must be equal 6")
						}
					} else {
						t.Error("Not a Uint8array")
					}

				} else {
					t.Error("Not a Uint8array")
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

		if n, err := a.BYTES_PER_ELEMENT(); err == nil {
			if n != 2 {
				t.Errorf("BYTES_PER_ELEMENT must be 2 have %d", n)
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}

func TestNewInt16Array(t *testing.T) {
	if a, err := NewInt16Array(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 20 {
				t.Errorf("Size must be 20 have %d", l)
			}

		} else {
			t.Error(err.Error())
		}

		a.SetValue(0, 45000)
		if v, err := a.GetValue(0); err == nil {

			if v.(int) > 0 {
				t.Error("Must be <0")
			}
		} else {
			t.Error(err.Error())
		}

		if n, err := a.BYTES_PER_ELEMENT(); err == nil {
			if n != 2 {
				t.Errorf("BYTES_PER_ELEMENT must be 2 have %d", n)
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}

func TestNewUint32Array(t *testing.T) {
	if a, err := NewUint32Array(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 40 {
				t.Errorf("Size must be 20 have %d", l)
			}

		} else {
			t.Error(err.Error())
		}

		if n, err := a.BYTES_PER_ELEMENT(); err == nil {
			if n != 4 {
				t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}

func TestNewInt32Array(t *testing.T) {
	if a, err := NewInt32Array(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 40 {
				t.Errorf("Size must be 20 have %d", l)
			}

		} else {
			t.Error(err.Error())
		}

		a.SetValue(0, 0xEFFFFFFF)

		if v, err := a.GetValue(0); err == nil {

			if v.(int) > 0 {
				t.Error("Must be <0")
			}
		} else {
			t.Error(err.Error())
		}

		if n, err := a.BYTES_PER_ELEMENT(); err == nil {
			if n != 4 {
				t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}

func TestNewFloat32Array(t *testing.T) {
	if a, err := NewFloat32Array(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 40 {
				t.Errorf("Size must be 20 have %d", l)
			}

		} else {
			t.Error(err.Error())
		}

		a.SetValue(0, float64(3.14))

		if v, err := a.GetValue(0); err == nil {

			if float32(v.(float64)) != float32(3.14) {
				t.Errorf("Must be 3.14 have %f ", v.(float64))
			}
		} else {
			t.Error(err.Error())
		}

		if n, err := a.BYTES_PER_ELEMENT(); err == nil {
			if n != 4 {
				t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}
func TestNewFloat64Array(t *testing.T) {
	if a, err := NewFloat64Array(10); err == nil {
		if l, err := a.ByteLength(); err == nil {
			if l != 80 {
				t.Errorf("Size must be 20 have %d", l)
			}

		} else {
			t.Error(err.Error())
		}

		a.SetValue(0, float64(3.14))

		if v, err := a.GetValue(0); err == nil {

			if v.(float64) != float64(3.14) {
				t.Errorf("Must be 3.14 have %f ", v.(float64))
			}
		} else {
			t.Error(err.Error())
		}

		if n, err := a.BYTES_PER_ELEMENT(); err == nil {
			if n != 8 {
				t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
			}
		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}

}
