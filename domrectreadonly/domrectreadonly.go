package domrectreadonly

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var domrectreadonlyinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var domrectreadonlyinstance JSInterface
		var err error
		if domrectreadonlyinstance.objectInterface, err = js.Global().GetWithErr("DOMRectReadOnly"); err == nil {
			domrectreadonlyinterface = &domrectreadonlyinstance
		}
	})

	return domrectreadonlyinterface
}

//DOMRectReadOnly struct
type DOMRectReadOnly struct {
	baseobject.BaseObject
}

func New() (DOMRectReadOnly, error) {

	var d DOMRectReadOnly

	if di := GetJSInterface(); di != nil {

		d.BaseObject = d.SetObject(di.objectInterface.New())
		return d, nil
	}
	return d, ErrNotImplemented
}

func (d DOMRectReadOnly) getAttributeDouble(attribute string) (float64, error) {

	var err error
	var obj js.Value
	var result float64

	if obj, err = d.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeNumber {
			result = obj.Float()
		} else {
			err = baseobject.ErrObjectNotNumber
		}
	}

	return result, err
}

func (d DOMRectReadOnly) Bottom() (float64, error) {
	return d.getAttributeDouble("bottom")
}

func (d DOMRectReadOnly) Height() (float64, error) {
	return d.getAttributeDouble("height")
}

func (d DOMRectReadOnly) Left() (float64, error) {
	return d.getAttributeDouble("left")
}
func (d DOMRectReadOnly) Right() (float64, error) {
	return d.getAttributeDouble("right")
}
func (d DOMRectReadOnly) Top() (float64, error) {
	return d.getAttributeDouble("top")
}
func (d DOMRectReadOnly) Width() (float64, error) {
	return d.getAttributeDouble("width")
}

func (d DOMRectReadOnly) X() (float64, error) {
	return d.getAttributeDouble("x")
}

func (d DOMRectReadOnly) Y() (float64, error) {
	return d.getAttributeDouble("y")
}

func (d DOMRectReadOnly) FromRect() {
	//TODO IMPLEMENT
}

func (d DOMRectReadOnly) ToJSON() (string, error) {
	var err error
	var result string
	var obj js.Value

	if obj, err = d.JSObject().CallWithErr("toJSON"); err == nil {
		if obj.Type() == js.TypeString {
			result = obj.String()
		} else {
			err = baseobject.ErrObjectNotString
		}
	}
	return result, err
}
