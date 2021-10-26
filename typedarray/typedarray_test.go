package typedarray

import (
	"testing"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNewUint8Array(t *testing.T) {
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

func TestNewUint8ArrayFrom(t *testing.T) {

	if struint8, err := NewUint8ArrayFrom(array.From_("Hello World")); testingutils.AssertErr(t, err) {

		if l, err := struint8.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(11), l)
		}
	}
}

func TestNewUint8ArrayOf(t *testing.T) {

	if ofuint8, err := NewUint8ArrayOf(3, 8, 4); testingutils.AssertErr(t, err) {

		if l, err := ofuint8.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(3), l)
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

func TestNewInt8ArrayFrom(t *testing.T) {

	if strint8, err := NewInt8ArrayFrom(array.From_("Hello World")); testingutils.AssertErr(t, err) {

		if l, err := strint8.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(11), l)
		}
	}
}

func TestNewInt8ArrayOf(t *testing.T) {

	if offint8, err := NewInt8ArrayOf(3, 8, 4); testingutils.AssertErr(t, err) {

		if l, err := offint8.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(3), l)
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

func TestNewUint8ClampedArrayFrom(t *testing.T) {

	if struintclamp8, err := NewUint8ClampedArrayFrom(array.From_("Hello World")); testingutils.AssertErr(t, err) {

		if l, err := struintclamp8.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(11), l)
		}
	}
}

func TestNewUint8ClampedArrayOf(t *testing.T) {

	if ofuintclamped8, err := NewUint8ClampedArrayOf(3, 8, 4); testingutils.AssertErr(t, err) {

		if l, err := ofuintclamped8.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(3), l)
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
func TestNewUint16ArrayFrom(t *testing.T) {

	if struint16, err := NewUint16ArrayFrom(array.From_("Hello World")); testingutils.AssertErr(t, err) {

		if l, err := struint16.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(22), l)
		}
	}
}

func TestNewUint16ArrayOf(t *testing.T) {

	if ofuint16, err := NewUint16ArrayOf(3, 8, 4); testingutils.AssertErr(t, err) {

		if l, err := ofuint16.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(6), l)
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

func TestNewInt16ArrayFrom(t *testing.T) {

	if strint16, err := NewInt16ArrayFrom(array.From_("Hello World")); testingutils.AssertErr(t, err) {

		if l, err := strint16.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(22), l)
		}
	}
}

func TestNewInt16ArrayOf(t *testing.T) {

	if offint16, err := NewInt16ArrayOf(3, 8, 4); testingutils.AssertErr(t, err) {

		if l, err := offint16.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(6), l)
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

func TestNewUint32ArrayFrom(t *testing.T) {

	if struint32, err := NewUint32ArrayFrom(array.From_("Hello World")); testingutils.AssertErr(t, err) {

		if l, err := struint32.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(44), l)
		}
	}
}

func TestNewUint32ArrayOf(t *testing.T) {

	if offuint32, err := NewUint32ArrayOf(3, 8, 4); testingutils.AssertErr(t, err) {

		if l, err := offuint32.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(12), l)
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

func TestNewInt32ArrayFrom(t *testing.T) {

	if strint32, err := NewInt32ArrayFrom(array.From_("Hello World")); testingutils.AssertErr(t, err) {

		if l, err := strint32.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(44), l)
		}
	}
}

func TestNewInt32ArrayOf(t *testing.T) {

	if offint32, err := NewInt32ArrayOf(3, 8, 4); testingutils.AssertErr(t, err) {

		if l, err := offint32.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(12), l)
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

func TestNewFloat32ArrayFrom(t *testing.T) {

	if strfloat32, err := NewFloat32ArrayFrom(array.From_("Hello World")); testingutils.AssertErr(t, err) {

		if l, err := strfloat32.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(44), l)
		}
	}
}

func TestNewFloat32ArrayOf(t *testing.T) {

	if offloat32, err := NewFloat32ArrayOf(3.2, 8.3, 4.3); testingutils.AssertErr(t, err) {

		if l, err := offloat32.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(12), l)
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

func TestNewFloat64ArrayFrom(t *testing.T) {

	if strfloat64, err := NewFloat64ArrayFrom(array.From_("Hello World")); testingutils.AssertErr(t, err) {

		if l, err := strfloat64.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(88), l)
		}
	}
}

func TestNewFloat64ArrayOf(t *testing.T) {

	if offloat64, err := NewFloat64ArrayOf(3.2, 8.3, 4.3); testingutils.AssertErr(t, err) {

		if l, err := offloat64.ByteLength(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, int64(24), l)
		}

	}
}
