package json

// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/JSON

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/gomap"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/object"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var jsoninterface js.Value

// Json  struct
type Json struct {
	object.Object
}

type JsonFrom interface {
	Json_() Json
}

func (i Json) Json_() Json {
	return i
}

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if jsoninterface, err = baseobject.Get(js.Global(), "JSON"); err != nil {
			jsoninterface = js.Undefined()
		}
		baseobject.Register(jsoninterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return jsoninterface
}

func Parse(data string) (Json, error) {

	var jsonObject js.Value
	var err error
	if jsoni := GetInterface(); !jsoni.IsUndefined() {

		if jsonObject, err = baseobject.Call(jsoni, "parse", data); err != nil {
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

	if jsoni := GetInterface(); !jsoni.IsUndefined() {

		j.BaseObject = j.SetObject(obj)
		return j, nil

	}

	return j, ErrNotAJson

}

func (j Json) Map() interface{} {

	return gomap.MapFromJSObject(j.JSObject())

}

func Stringify(opts ...interface{}) (string, error) {

	var arrayJS []interface{}
	var err error
	var stringObject js.Value

	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}
	if jsoni := GetInterface(); !jsoni.IsUndefined() {

		if stringObject, err = baseobject.Call(jsoni, "stringify", arrayJS); err != nil {
			return "", err
		} else {

			return stringObject.String(), nil
		}

	} else {
		err = ErrNotImplemented
	}

	return "", err

}

func StringifyObject(object interface{}) (string, error) {

	var err error
	var stringObject js.Value

	if jsoni := GetInterface(); !jsoni.IsUndefined() {

		if stringObject, err = baseobject.Call(jsoni, "stringify", js.ValueOf(object)); err != nil {
			return "", err
		} else {

			return stringObject.String(), nil
		}

	} else {
		err = ErrNotImplemented
	}

	return "", err

}
