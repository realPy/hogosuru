package gomap

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/object"
)

func MapFromJSObject(jsobj js.Value) interface{} {
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
							arrayret = append(arrayret, MapFromJSObject(obj1.JSObject()))
						}

					}

				}

			}

			retvalue = arrayret

		} else {

			vmap := make(map[string]interface{})
			keys, _ := obj.Keys()

			itkeys, _ := keys.Values()

			for _, vkey, err := itkeys.Next(); err == nil; _, vkey, err = itkeys.Next() {

				if key, ok := vkey.(string); ok {

					if value, err := baseobject.Get(jsobj, key); err == nil {
						i := baseobject.GoValue_(value)
						if obj1, ok := i.(baseobject.ObjectFrom); !ok {
							vmap[key] = i
						} else {
							vmap[key] = MapFromJSObject(obj1.JSObject())

						}

					}

				}

			}
			retvalue = vmap
		}

	}

	return retvalue

}
