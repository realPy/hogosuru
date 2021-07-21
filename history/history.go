package history

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

// https://developer.mozilla.org/fr/docs/Web/API/History_API

var singleton sync.Once

// var historyinterface *JSInterface

var historyinterface js.Value

//JSInterface JSInterface struct
// type JSInterface struct {
// 	objectInterface js.Value
// }

//HTMLCollection struct
type History struct {
	baseobject.BaseObject
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var window js.Value
		var err error

		if window, err = js.Global().GetWithErr("window"); err == nil {
			if historyinterface, err = window.GetWithErr("history"); err != nil {
				historyinterface = js.Null()
			}
		} else {
			historyinterface = js.Null()
		}
	})

	baseobject.Register(historyinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return historyinterface
}

func NewFromJSObject(obj js.Value) (History, error) {
	var h History

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrCantImplementedHistory
}

func GetHistory() (History, error) {
	return NewFromJSObject(GetInterface())
}

func (h History) Forward() error {
	var err error
	_, err = h.JSObject().CallWithErr("forward")
	return err
}

func (h History) Back() error {
	var err error
	_, err = h.JSObject().CallWithErr("back")
	return err
}

func (h History) Go(position int) error {
	var err error
	_, err = h.JSObject().CallWithErr("go", js.ValueOf(position))

	return err
}

/*
func (h History) Length(position int) error {
	var err error
	if hist := GetJSInterface(); hist != nil {
		_, err = hist.objectInterface.CallWithErr("length", js.ValueOf(position))
	}

	return err
}

func PushState(obj baseobject.BaseObject, name string, page string) error {
	var err error
	if hist := GetJSInterface(); hist != nil {
		_, err = hist.objectInterface.CallWithErr("pushState", js.ValueOf(obj), js.ValueOf(name), js.ValueOf(page))
	}

	return err
}

func ReplaceState(obj baseobject.BaseObject, name string, page string) error {
	var err error
	if hist := GetJSInterface(); hist != nil {
		_, err = hist.objectInterface.CallWithErr("replaceState", js.ValueOf(obj), js.ValueOf(name), js.ValueOf(page))
	}

	return err
}

func State() (js.Value, error) {
	var err error
	var obj js.Value
	if hist := GetJSInterface(); hist != nil {
		obj, err := hist.objectInterface.CallWithErr("state")
		return obj, err
	}

	return obj, err
}
*/
