package json

// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/JSON

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/object"
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

	if ji := GetJSInterface(); ji != nil {

		j.BaseObject = j.SetObject(obj)
		return j, nil

	}

	return j, ErrNotAJson

}

func extractJsonFromObject(jsobj js.Value) interface{} {
	var retvalue interface{}

	if obj, err := object.NewFromJSObject(jsobj); err == nil {

		if ok, err := array.IsArray(obj.BaseObject); ok && err == nil {
			var array []interface{}
			keys, _ := object.Values(obj)
			itkeys, _ := keys.Values()

			for _, vkey, err := itkeys.Next(); err == nil; _, vkey, err = itkeys.Next() {

				if obj1, ok := vkey.(baseobject.BaseObject); !ok {
					array = append(array, vkey)

				} else {
					array = append(array, extractJsonFromObject(obj1.JSObject()))
				}

			}
			retvalue = array

		} else {
			json := make(map[string]interface{})
			keys, _ := object.Keys(obj)

			itkeys, _ := keys.Values()

			for _, vkey, err := itkeys.Next(); err == nil; _, vkey, err = itkeys.Next() {

				if key, ok := vkey.(string); ok {
					if value, err := jsobj.GetWithErr(key); err == nil {
						i := baseobject.GoValue(value)
						if obj1, ok := i.(baseobject.BaseObject); !ok {
							json[key] = i
						} else {
							json[key] = extractJsonFromObject(obj1.JSObject())

						}

					}

				}

			}
			retvalue = json
		}

	}

	return retvalue

}

func (j Json) Map() interface{} {

	return extractJsonFromObject(j.JSObject())

}

/*
func (j Json) GetItem(item string) interface{} {
baseobject.Eval()
}
*./
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

			return stringObject.String(), nil
		}

	} else {
		err = ErrNotImplemented
	}

	return "", err

}
