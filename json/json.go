package json

// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/JSON

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
	"github.com/realPy/hogosuru/object"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var jsoninterface js.Value

//Json  struct
type Json struct {
	object.Object
}

type JsonFrom interface {
	Json_() Json
}

func (i Json) Json_() Json {
	return i
}

//GetInterface get the JS interface
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

func extractJsonFromObject(jsobj js.Value) interface{} {
	var retvalue interface{}

	if obj, err := object.NewFromJSObject(jsobj); err == nil {

		if ok, err := array.IsArray(obj.BaseObject); ok && err == nil {

			var arrayret []interface{}

			if a, err := array.NewFromJSObject(obj.JSObject()); err == nil {
				if it, err := a.Entries(); err == nil {
					for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

						if obj1, ok := value.(baseobject.ObjectFrom); !ok {
							arrayret = append(arrayret, value)

						} else {
							arrayret = append(arrayret, extractJsonFromObject(obj1.JSObject()))
						}

					}

				}

			}

			retvalue = arrayret

		} else {

			json := make(map[string]interface{})
			keys, _ := obj.Keys()

			itkeys, _ := keys.Values()

			for _, vkey, err := itkeys.Next(); err == nil; _, vkey, err = itkeys.Next() {

				if key, ok := vkey.(string); ok {

					if value, err := baseobject.Get(jsobj, key); err == nil {
						i := baseobject.GoValue_(value)
						if obj1, ok := i.(baseobject.ObjectFrom); !ok {
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
