package arraybuffer

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func TestNew(t *testing.T) {

	if a, err := New(8); err == nil {

		if l, err := a.ByteLength(); err == nil {
			if l != 8 {
				t.Error("Size must be 8")
			}

		} else {
			t.Error(err.Error())
		}
	} else {
		t.Error(err.Error())
	}
}

func TestSlice(t *testing.T) {

	if a, err := New(32); err == nil {

		if b, err := a.Slice(10); err == nil {
			if b.ByteLength_() != 22 {
				t.Errorf("Must be size to 5 have %d", b.ByteLength_())
			}

		} else {
			t.Error(err.Error())
		}

		if b, err := a.Slice(10, 16); err == nil {
			if b.ByteLength_() != 6 {
				t.Errorf("Must be size to 5 have %d", b.ByteLength_())
			}

		} else {
			t.Error(err.Error())
		}

	} else {
		t.Error(err.Error())
	}
}
