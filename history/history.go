package history

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
)

// https://developer.mozilla.org/fr/docs/Web/API/History_API

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

// var historyinterface *JSInterface

var historyinterface js.Value

//JSInterface JSInterface struct
// type JSInterface struct {
// 	objectInterface js.Value
// }

//History struct
type History struct {
	baseobject.BaseObject
}

type HistoryFrom interface {
	History_() History
}

func (h History) History_() History {
	return h
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if historyinterface, err = baseobject.Get(js.Global(), "History"); err != nil {
			historyinterface = js.Undefined()
		}
		baseobject.Register(historyinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return historyinterface
}

func NewFromJSObject(obj js.Value) (History, error) {
	var h History
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHistory
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h History) Forward() error {
	var err error
	_, err = h.Call("forward")
	return err
}

func (h History) Back() error {
	var err error
	_, err = h.Call("back")
	return err
}

func (h History) Go(position int) error {
	var err error
	_, err = h.Call("go", js.ValueOf(position))

	return err
}

func (h History) Length() (int, error) {
	var err error

	obj, err := h.Get("length")

	return obj.Int(), err
}

func (h History) PushState(obj interface{}, title string, url ...string) error {
	var err error
	var arrayJS []interface{}

	if objGo, ok := obj.(baseobject.ObjectFrom); ok {
		arrayJS = append(arrayJS, objGo.JSObject())
	} else {
		arrayJS = append(arrayJS, js.ValueOf(obj))
	}

	arrayJS = append(arrayJS, js.ValueOf(title))

	if len(url) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(url[0]))
	}

	_, err = h.Call("pushState", arrayJS...)

	return err
}

func (h History) ReplaceState(obj interface{}, title string, url ...string) error {
	var err error

	var arrayJS []interface{}

	if objGo, ok := obj.(baseobject.ObjectFrom); ok {
		arrayJS = append(arrayJS, objGo.JSObject())
	} else {
		arrayJS = append(arrayJS, js.ValueOf(obj))
	}

	arrayJS = append(arrayJS, js.ValueOf(title))

	if len(url) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(url[0]))
	}

	_, err = h.Call("replaceState", arrayJS...)

	return err
}

func (h History) State() (interface{}, error) {
	var err error
	var obj js.Value
	var ret interface{}

	if obj, err = h.Get("state"); err == nil {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			ret, err = baseobject.Discover(obj)
		}

	}

	return ret, err
}
