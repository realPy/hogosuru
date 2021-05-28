package domrect

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/domrectreadonly"
)

var singleton sync.Once

var domrectinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var domrectinstance JSInterface
		var err error
		if domrectinstance.objectInterface, err = js.Global().GetWithErr("DOMRect"); err == nil {
			domrectinterface = &domrectinstance
		}
	})

	return domrectinterface
}

//DOMRectReadOnly struct
type DOMRect struct {
	domrectreadonly.DOMRectReadOnly
}

func New() (DOMRect, error) {

	var d DOMRect

	if di := GetJSInterface(); di != nil {

		d.BaseObject = d.SetObject(di.objectInterface.New())
		return d, nil
	}
	return d, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (DOMRect, error) {
	var d DOMRect
	var err error
	if di := GetJSInterface(); di != nil {
		if obj.InstanceOf(di.objectInterface) {
			d.BaseObject = d.SetObject(obj)

		} else {
			err = ErrNotAnDOMRect
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func (d DOMRect) SetBottom(value float64) error {

	return d.JSObject().SetWithErr("bottom", js.ValueOf(value))
}
func (d DOMRect) SetHeight(value float64) error {

	return d.JSObject().SetWithErr("height", js.ValueOf(value))
}
func (d DOMRect) SetLeft(value float64) error {

	return d.JSObject().SetWithErr("left", js.ValueOf(value))
}
func (d DOMRect) SetRight(value float64) error {

	return d.JSObject().SetWithErr("right", js.ValueOf(value))
}

func (d DOMRect) SetTop(value float64) error {

	return d.JSObject().SetWithErr("top", js.ValueOf(value))
}

func (d DOMRect) SetWidth(value float64) error {

	return d.JSObject().SetWithErr("width", js.ValueOf(value))
}

func (d DOMRect) SetX(value float64) error {

	return d.JSObject().SetWithErr("x", js.ValueOf(value))
}

func (d DOMRect) SetY(value float64) error {

	return d.JSObject().SetWithErr("y", js.ValueOf(value))
}
