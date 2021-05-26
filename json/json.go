package json

// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/JSON

import (
	"sync"

	"syscall/js"
)

var singleton sync.Once

var jsoninterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Json  struct
type Json struct {
	baseobject.BaseObject
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var jsoninstance JSInterface
		var err error
		if jsoninstance.objectInterface, err = js.Global().GetWithErr("JSON"); err == nil {
			jsoninterface = &jsoninstance
		}
	})

	return jsoninterface
}

func Parse(data string) (Json, error) {

	var jsonObject js.Value
	var err error
	if jsoni := GetJSInterface(); jsoni != nil {

		if jsonObject, err = jsoni.objectInterface.CallWithErr("parse", data); err != nil {
			return Json{}, err
		} else {
			return NewFromJSObject(jsonObject)
		}

	} else {
		err = ErrNotImplemented
	}

	return Json{}, err
}

func NewFromJSObject(obj js.Value) (Json, error) {
	var j Json

	if ji := GetJSInterface(); ji != nil {

		j.BaseObject = j.SetObject(obj)
		return j, nil

	}

	return j, ErrNotAJson

}

/*
func (j Json) Get(key string) js.Value {
	return j.JSObject().Get(key)
}

func (j Json) GoJson() jsbaseobject.GOMap {
	return jsbaseobject.Map(j.JSObject())

}
*/

func (j Json) Stringify() (string, error) {

	var stringObject js.Value
	var err error
	if jsoni := GetJSInterface(); jsoni != nil {

		if stringObject, err = jsoni.objectInterface.CallWithErr("stringify", j.JSObject()); err != nil {
			return "", err
		} else {

			return stringbaseobject.String(), nil
		}

	} else {
		err = ErrNotImplemented
	}

	return "", err

}
