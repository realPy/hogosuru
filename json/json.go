package json

import (
	"fmt"

	"github.com/realPy/jswasm/js"
)

type Json struct {
	object js.Value
}

func JsonParse(jsonstr string) (Json, error) {

	if jsonObject, err := js.Global().GetWithErr("JSON"); err == nil {
		if json, err := jsonObject.CallWithErr("parse", jsonstr); err != nil {
			return Json{}, err
		} else {
			return Json{object: json}, nil
		}

	}
	return Json{}, fmt.Errorf("Unable to get JSON interface")
}

func (j Json) Get(key string) js.Value {
	return j.object.Get(key)
}

func recurseUmarshalJson(object js.Value) map[string]interface{} {

	var json map[string]interface{} = make(map[string]interface{})
	if Object, err := js.Global().GetWithErr("Object"); err == nil {

		if value, err := Object.CallWithErr("keys", object); err == nil {

			for i := 0; i < value.Length(); i++ {
				key := value.Index(i).String()
				jsvalue := object.Get(key)
				switch jsvalue.Type() {
				case js.TypeNumber:
					json[key] = jsvalue.Float()
				case js.TypeString:
					json[key] = jsvalue.String()
				case js.TypeBoolean:
					json[key] = jsvalue.Bool()
				case js.TypeObject:
					json[key] = recurseUmarshalJson(jsvalue)

				}
			}

		}
	}

	return json
}
func (j Json) GoJson() map[string]interface{} {

	return recurseUmarshalJson(j.object)
}
