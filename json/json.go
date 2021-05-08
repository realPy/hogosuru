package json

import (
	"sync"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
	jsobject "github.com/realPy/jswasm/object"
)

var singleton sync.Once

var jsoninterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Json  struct
type Json struct {
	object.Object
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
	j.Object = j.SetObject(obj)
	return j, nil

}

func (j Json) Get(key string) js.Value {
	return j.JSObject().Get(key)
}

func (j Json) GoJson() jsobject.GOMap {
	return jsobject.Map(j.JSObject())

}

func (j Json) Stringify() (string, error) {

	var stringObject js.Value
	var err error
	if jsoni := GetJSInterface(); jsoni != nil {

		if stringObject, err = jsoni.objectInterface.CallWithErr("stringify", j.JSObject()); err != nil {
			return "", err
		} else {

			return stringObject.String(), nil
		}

	} else {
		err = ErrNotImplemented
	}

	return "", err

}
