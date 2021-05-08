package blob

import (
	"sync"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

var singleton sync.Once

var blobinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var blobinstance JSInterface
		var err error
		if blobinstance.objectInterface, err = js.Global().GetWithErr("Blob"); err == nil {
			blobinterface = &blobinstance
		}
	})

	return blobinterface
}

type Blob struct {
	object.Object
}

/**** new need more information implemented later
func New() (Blob, error) {

	var b Blob
	if bi := GetJSInterface(); bi != nil {
		b.Object = b.SetObject(bi.objectInterface.New())
		return b, nil
	}
	return b, ErrNotImplemented
}
*/

func NewFromJSObject(obj js.Value) (Blob, error) {
	var b Blob

	if object.String(obj) == "[object Blob]" {
		b.Object = b.SetObject(obj)
		return b, nil
	}

	return b, ErrNotABlob
}

func (b Blob) IsClosed() (bool, error) {
	var err error
	var obj js.Value

	if obj, err = b.JSObject().GetWithErr("isClosed"); err == nil {

		return obj.Bool(), nil
	}
	return true, err
}

func (b Blob) Size() (int, error) {
	var err error
	var obj js.Value
	if obj, err = b.JSObject().GetWithErr("size"); err == nil {

		return obj.Int(), nil
	}
	return 0, err
}
func (b Blob) Type() (string, error) {
	var err error
	var obj js.Value

	if obj, err = b.JSObject().GetWithErr("type"); err == nil {

		return obj.String(), nil
	}
	return "", err
}

func (b Blob) Close() error {
	_, err := b.JSObject().CallWithErr("close")
	return err
}

func (b Blob) Slice(begin, end int) (Blob, error) {
	var blob js.Value
	var err error
	if blob, err = b.JSObject().CallWithErr("slice", js.ValueOf(begin), js.ValueOf(end)); err == nil {
		var newblob Blob
		object := newblob.SetObject(blob)
		newblob.Object = object
		return newblob, nil
	}
	return Blob{}, err
}
