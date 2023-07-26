package stylesheet

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/node"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var stylesheetinterface js.Value

// StyleSheet struct
type StyleSheet struct {
	baseobject.BaseObject
}

type StyleSheetFrom interface {
	StyleSheet_() StyleSheet
}

func (s StyleSheet) StyleSheet_() StyleSheet {
	return s
}

func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if stylesheetinterface, err = baseobject.Get(js.Global(), "StyleSheet"); err != nil {
			stylesheetinterface = js.Undefined()
		}
		baseobject.Register(stylesheetinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return stylesheetinterface
}

func NewFromJSObject(obj js.Value) (StyleSheet, error) {
	var s StyleSheet
	var err error
	if dli := GetInterface(); !dli.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dli) {
				s.BaseObject = s.SetObject(obj)

			} else {
				err = ErrNotAnStyleSheet
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return s, err
}

func (s StyleSheet) Disabled() (bool, error) {
	return s.GetAttributeBool("disabled")
}

func (s StyleSheet) SetDisabled(value bool) error {
	return s.SetAttributeBool("disabled", value)
}

func (s StyleSheet) Href() (string, error) {
	return s.GetAttributeString("href")
}

func (s StyleSheet) OwnerNode() (node.Node, error) {
	var err error
	var obj js.Value
	var n node.Node
	if obj, err = s.Get("ownerNode"); err == nil {

		if obj.IsUndefined() {
			err = baseobject.ErrNotAnObject

		} else {
			n, err = node.NewFromJSObject(obj)
		}
	}
	return n, err
}

func (s StyleSheet) ParentStyleSheet() (StyleSheet, error) {
	var err error
	var obj js.Value
	var ps StyleSheet
	if obj, err = s.Get("parentStyleSheet"); err == nil {

		if obj.IsUndefined() {
			err = baseobject.ErrNotAnObject

		} else {
			ps, err = NewFromJSObject(obj)
		}
	}
	return ps, err
}

/*
func (s StyleSheet) Media() {
//TODO IMPLEMENT
}*/

func (s StyleSheet) Title() (string, error) {
	return s.GetAttributeString("title")
}

func (s StyleSheet) Type() (string, error) {
	return s.GetAttributeString("type")
}
