package domrect

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/domrectreadonly"
)

var singleton sync.Once

var domrectinterface js.Value

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if domrectinterface, err = baseobject.Get(js.Global(), "DOMRect"); err != nil {
			domrectinterface = js.Undefined()
		}
		baseobject.Register(domrectinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return domrectinterface
}

//DOMRectReadOnly struct
type DOMRect struct {
	domrectreadonly.DOMRectReadOnly
}

type DOMRectFrom interface {
	DOMRect_() DOMRect
}

func (d DOMRect) DOMRect_() DOMRect {
	return d
}

func New() (DOMRect, error) {

	var d DOMRect
	var obj js.Value
	var err error
	if di := GetInterface(); !di.IsUndefined() {

		if obj, err = baseobject.New(di); err == nil {
			d.BaseObject = d.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func NewFromJSObject(obj js.Value) (DOMRect, error) {
	var d DOMRect
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				d.BaseObject = d.SetObject(obj)

			} else {
				err = ErrNotAnDOMRect
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func (d DOMRect) SetBottom(value float64) error {

	return d.Set("bottom", js.ValueOf(value))
}
func (d DOMRect) SetHeight(value float64) error {

	return d.Set("height", js.ValueOf(value))
}
func (d DOMRect) SetLeft(value float64) error {

	return d.Set("left", js.ValueOf(value))
}
func (d DOMRect) SetRight(value float64) error {

	return d.Set("right", js.ValueOf(value))
}

func (d DOMRect) SetTop(value float64) error {

	return d.Set("top", js.ValueOf(value))
}

func (d DOMRect) SetWidth(value float64) error {

	return d.Set("width", js.ValueOf(value))
}

func (d DOMRect) SetX(value float64) error {

	return d.Set("x", js.ValueOf(value))
}

func (d DOMRect) SetY(value float64) error {

	return d.Set("y", js.ValueOf(value))
}
