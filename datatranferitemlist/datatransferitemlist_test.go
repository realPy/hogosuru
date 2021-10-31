package datatranferitemlist

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/file"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`dt=new DataTransfer()
	itemlist=dt.items`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj, err = baseobject.Get(js.Global(), "itemlist"); testingutils.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object DataTransferItemList]", d.ToString_())

		}
	}

}

func TestAdd(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj, err = baseobject.Get(js.Global(), "itemlist"); testingutils.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if f, err := file.New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); testingutils.AssertErr(t, err) {

				testingutils.AssertErr(t, d.Add(f))
				if l, err := d.Length(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, 1, l)
				}

			}

		}
	}

}

func TestClear(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj, err = baseobject.Get(js.Global(), "itemlist"); testingutils.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if f, err := file.New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); testingutils.AssertErr(t, err) {

				testingutils.AssertErr(t, d.Add(f))
				testingutils.AssertErr(t, d.Clear())
				if l, err := d.Length(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, 0, l)
				}

			}

		}
	}

}

func TestRemove(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj, err = baseobject.Get(js.Global(), "itemlist"); testingutils.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			d.Clear()

			if f, err := file.New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); testingutils.AssertErr(t, err) {

				testingutils.AssertErr(t, d.Add(f))

				testingutils.AssertErr(t, d.Remove(0))
				if l, err := d.Length(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, 0, l)
				}

			}

		}
	}

}

func TestDataTransferItem(t *testing.T) {

	var err error
	var obj js.Value
	var d DataTransferItemList

	if obj, err = baseobject.Get(js.Global(), "itemlist"); testingutils.AssertErr(t, err) {

		if d, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			d.Clear()

			if f, err := file.New(array.From_("(⌐□_□)"), "chucknorris.png", map[string]interface{}{"type": "image/png"}); testingutils.AssertErr(t, err) {
				testingutils.AssertErr(t, d.Add(f))

				if item, err := d.DataTransferItem(0); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "[object DataTransferItem]", item.ToString_())
				}

				_, err := d.DataTransferItem(1)

				testingutils.AssertExpect(t, err, baseobject.ErrUndefinedValue)
			}

		}
	}

}
