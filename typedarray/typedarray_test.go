package typedarray

import (
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNewUInt8Array(t *testing.T) {
	if a, err := NewUint8Array(10); testingutils.AssertErr(t, err) {
		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {
			if l != 10 {
				t.Errorf("Size must be 10 have %d", l)
			}

			if b, err := a.Subarray(2, 8); testingutils.AssertErr(t, err) {
				if bb, ok := b.(Uint8ArrayFrom); ok {
					if l1, err := bb.Uint8Array_().ByteLength(); testingutils.AssertErr(t, err) {
						if l1 != 6 {
							t.Error("Must be equal 6")
						}
					}

				} else {
					t.Error("Not a Uint8array")
				}

			}

		}

	}

}

func TestNewInt8Array(t *testing.T) {
	if a, err := NewInt8Array(10); testingutils.AssertErr(t, err) {
		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {
			if l != 10 {
				t.Errorf("Size must be 10 have %d", l)
			}

		}

		a.Fill(200)
		if v, err := a.GetValue(0); testingutils.AssertErr(t, err) {

			if v.(int) > 0 {
				t.Error("Must be a negative value")
			}
		}

	}
}

func TestNewUInt8ClampedArray(t *testing.T) {
	if a, err := NewUint8ClampedArray(10); testingutils.AssertErr(t, err) {
		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {
			if l != 10 {
				t.Errorf("Size must be 10 have %d", l)
			}

		}

		a.SetValue(0, 3.6)
		if v, err := a.GetValue(0); testingutils.AssertErr(t, err) {

			if v.(int) != 4 {
				t.Error("Must be a egal 4")
			}
		}

	}

}

func TestNewUInt16Array(t *testing.T) {
	if a, err := NewUint16Array(10); testingutils.AssertErr(t, err) {
		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {
			if l != 20 {
				t.Errorf("Size must be 20 have %d", l)
			}

		}

		a.SetValue(0, 1000)
		if v, err := a.GetValue(0); testingutils.AssertErr(t, err) {

			if v.(int) != 1000 {
				t.Error("Must be a egal 1000")
			}
		}

		if n, err := a.BYTES_PER_ELEMENT(); testingutils.AssertErr(t, err) {
			if n != 2 {
				t.Errorf("BYTES_PER_ELEMENT must be 2 have %d", n)
			}
		}

	}

}

func TestNewInt16Array(t *testing.T) {
	if a, err := NewInt16Array(10); testingutils.AssertErr(t, err) {
		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {
			if l != 20 {
				t.Errorf("Size must be 20 have %d", l)
			}

		}

		a.SetValue(0, 45000)
		if v, err := a.GetValue(0); testingutils.AssertErr(t, err) {

			if v.(int) > 0 {
				t.Error("Must be <0")
			}
		}

		if n, err := a.BYTES_PER_ELEMENT(); testingutils.AssertErr(t, err) {
			if n != 2 {
				t.Errorf("BYTES_PER_ELEMENT must be 2 have %d", n)
			}
		}

	}

}

func TestNewUint32Array(t *testing.T) {
	if a, err := NewUint32Array(10); testingutils.AssertErr(t, err) {
		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {
			if l != 40 {
				t.Errorf("Size must be 20 have %d", l)
			}

		}

		if n, err := a.BYTES_PER_ELEMENT(); testingutils.AssertErr(t, err) {
			if n != 4 {
				t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
			}
		}

	}

}

func TestNewInt32Array(t *testing.T) {
	if a, err := NewInt32Array(10); testingutils.AssertErr(t, err) {
		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {
			if l != 40 {
				t.Errorf("Size must be 20 have %d", l)
			}

		}

		a.SetValue(0, 0xEFFFFFFF)

		if v, err := a.GetValue(0); testingutils.AssertErr(t, err) {

			if v.(int) > 0 {
				t.Error("Must be <0")
			}
		}

		if n, err := a.BYTES_PER_ELEMENT(); testingutils.AssertErr(t, err) {
			if n != 4 {
				t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
			}
		}

	}

}

func TestNewFloat32Array(t *testing.T) {
	if a, err := NewFloat32Array(10); testingutils.AssertErr(t, err) {
		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {
			if l != 40 {
				t.Errorf("Size must be 20 have %d", l)
			}

		}
		a.SetValue(0, float64(3.14))

		if v, err := a.GetValue(0); testingutils.AssertErr(t, err) {

			if float32(v.(float64)) != float32(3.14) {
				t.Errorf("Must be 3.14 have %f ", v.(float64))
			}
		}

		if n, err := a.BYTES_PER_ELEMENT(); testingutils.AssertErr(t, err) {
			if n != 4 {
				t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
			}
		}

	}

}
func TestNewFloat64Array(t *testing.T) {
	if a, err := NewFloat64Array(10); testingutils.AssertErr(t, err) {
		if l, err := a.ByteLength(); testingutils.AssertErr(t, err) {
			if l != 80 {
				t.Errorf("Size must be 20 have %d", l)
			}

		}

		a.SetValue(0, float64(3.14))

		if v, err := a.GetValue(0); testingutils.AssertErr(t, err) {

			if v.(float64) != float64(3.14) {
				t.Errorf("Must be 3.14 have %f ", v.(float64))
			}
		}

		if n, err := a.BYTES_PER_ELEMENT(); testingutils.AssertErr(t, err) {
			if n != 8 {
				t.Errorf("BYTES_PER_ELEMENT must be 4 have %d", n)
			}
		}

	}

}
