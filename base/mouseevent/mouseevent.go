package mouseevent

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/eventtarget"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var mouseeventinterface js.Value

// MouseEvent MouseEvent struct
type MouseEvent struct {
	//Must be herited from mouseevent
	event.Event
}

type MouseEventFrom interface {
	MouseEvent_() MouseEvent
}

func (m MouseEvent) MouseEvent_() MouseEvent {
	return m
}

// GetInterface get MouseEvent interface
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if mouseeventinterface, err = baseobject.Get(js.Global(), "MouseEvent"); err != nil {
			mouseeventinterface = js.Undefined()
		}
		baseobject.Register(mouseeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		eventtarget.GetInterface()
	})

	return mouseeventinterface
}

func New(typeevent string, init ...map[string]interface{}) (MouseEvent, error) {

	var m MouseEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if mei := GetInterface(); !mei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(init) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(init[0]))
		}
		if obj, err = baseobject.New(mei, arrayJS...); err == nil {
			m.BaseObject = m.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return m, err
}

func NewFromJSObject(obj js.Value) (MouseEvent, error) {
	var m MouseEvent
	var err error
	if mei := GetInterface(); !mei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(mei) {
				m.BaseObject = m.SetObject(obj)

			} else {
				err = ErrNotAMouseEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return m, err
}

func (m MouseEvent) AltKey() (bool, error) {

	return m.GetAttributeBool("altKey")
}

func (m MouseEvent) Button() (int, error) {

	return m.GetAttributeInt("button")
}

func (m MouseEvent) Buttons() (int, error) {

	return m.GetAttributeInt("buttons")
}

func (m MouseEvent) ClientX() (float64, error) {

	return m.GetAttributeDouble("clientX")
}

func (m MouseEvent) ClientY() (float64, error) {

	return m.GetAttributeDouble("clientY")
}

func (m MouseEvent) CtrlKey() (bool, error) {

	return m.GetAttributeBool("ctrlKey")
}

func (m MouseEvent) MetaKey() (bool, error) {

	return m.GetAttributeBool("metaKey")
}

func (m MouseEvent) MovementX() (int, error) {

	return m.GetAttributeInt("movementX")
}

func (m MouseEvent) MovementY() (int, error) {

	return m.GetAttributeInt("movementY")
}

func (m MouseEvent) OffsetX() (float64, error) {

	return m.GetAttributeDouble("offsetX")
}

func (m MouseEvent) OffsetY() (float64, error) {

	return m.GetAttributeDouble("offsetY")
}

func (m MouseEvent) PageX() (int, error) {

	return m.GetAttributeInt("pageX")
}

func (m MouseEvent) PageY() (int, error) {

	return m.GetAttributeInt("pageY")
}

func (m MouseEvent) Region() (string, error) {

	return m.GetAttributeString("region")
}

func (m MouseEvent) RelatedTarget() (eventtarget.EventTarget, error) {
	var err error
	var obj interface{}

	var e eventtarget.EventTarget

	if obj, err = m.GetAttributeGlobal("relatedTarget"); err == nil {

		if obj != nil {
			if efrom, ok := obj.(eventtarget.EventTargetFrom); ok {
				e = efrom.EventTarget_()
			}
		} else {
			err = baseobject.ErrUndefinedValue
		}

	}
	return e, err
}

func (m MouseEvent) ScreenX() (float64, error) {

	return m.GetAttributeDouble("screenX")
}

func (m MouseEvent) ScreenY() (float64, error) {

	return m.GetAttributeDouble("screenY")
}

func (m MouseEvent) ShiftKey() (bool, error) {

	return m.GetAttributeBool("shiftKey")
}

func (m MouseEvent) X() (float64, error) {

	return m.GetAttributeDouble("x")
}

func (m MouseEvent) Y() (float64, error) {

	return m.GetAttributeDouble("y")
}

func (m MouseEvent) GetModifierState(args string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = m.Call("getModifierState", js.ValueOf(args)); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}
