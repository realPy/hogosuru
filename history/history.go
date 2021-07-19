package history

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

// https://developer.mozilla.org/fr/docs/Web/API/History_API

var singleton sync.Once

var historyinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//HTMLCollection struct
type Hystory struct {
	baseobject.BaseObject
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var historyinstance JSInterface
		var windowinstance js.Value
		var err error

		if windowinstance, err = js.Global().GetWithErr("window"); err == nil {
			if historyinstance.objectInterface, err = windowinstance.GetWithErr("history"); err == nil {
				historyinterface = &historyinstance
			} else {
				println("%s", err)
			}
		} else {
			println("%s", err)
		}
	})

	return historyinterface
}

func Forward() error {
	var err error
	if hist := GetJSInterface(); hist != nil {
		_, err = hist.objectInterface.CallWithErr("forward")
	}

	return err
}

func Back() error {
	var err error
	if hist := GetJSInterface(); hist != nil {
		_, err = hist.objectInterface.CallWithErr("back")
	}

	return err
}

func Go(position int) error {
	var err error
	if hist := GetJSInterface(); hist != nil {
		_, err = hist.objectInterface.CallWithErr("go", js.ValueOf(position))
	}

	return err
}

func Length(position int) error {
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
