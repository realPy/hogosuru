package testingutils

import (
	"errors"
	"reflect"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
)

func InvokeCheck(t *testing.T, object interface{}, expectDesc map[string]interface{}) {

	t.Run(expectDesc["method"].(string), func(t *testing.T) {
		var argsReflect []reflect.Value

		if argsMethodsExist, ok := expectDesc["args"]; ok {
			if argsMethods, ok := argsMethodsExist.([]interface{}); ok {

				for _, i := range argsMethods {
					argsReflect = append(argsReflect, reflect.ValueOf(i))
				}

			}

		}

		if method := reflect.ValueOf(object).MethodByName(expectDesc["method"].(string)); method != (reflect.Value{}) {

			val := method.Call(argsReflect)
			var typechecking string

			if valtype, ok := expectDesc["type"]; ok {
				typechecking = valtype.(string)
			}

			if err, ok := val[len(val)-1].Interface().(error); ok {
				if typechecking == "error" {
					if !errors.Is(expectDesc["resultattempt"].(error), err) {
						AssertErr(t, err)
					}
				} else {
					AssertErr(t, err)
				}

			} else {

				if gettermethod, ok := expectDesc["gettermethod"]; ok {

					var getterArgsReflect []reflect.Value

					if getterArgsMethodsExist, ok := expectDesc["getterargs"]; ok {
						if argsMethods, ok := getterArgsMethodsExist.([]interface{}); ok {

							for _, i := range argsMethods {
								getterArgsReflect = append(getterArgsReflect, reflect.ValueOf(i))
							}

						}

					}

					val2 := reflect.ValueOf(object).MethodByName(gettermethod.(string)).Call(getterArgsReflect)
					if err, ok := val2[1].Interface().(error); ok {
						AssertErr(t, err)
					} else {
						val = val2
					}

				}

				switch typechecking {
				case "constructnamechecking":
					if objfrom, ok := val[0].Interface().(baseobject.ObjectFrom); ok {
						bobj := objfrom.BaseObject_()

						if settermethod := reflect.ValueOf(bobj).MethodByName("ConstructName_"); settermethod != (reflect.Value{}) {
							valconstruct := settermethod.Call([]reflect.Value{})

							AssertExpect(t, expectDesc["resultattempt"], valconstruct[0].Interface())

						} else {
							t.Errorf("Method %s not found in %T", "ConstructName_", bobj)
						}

					} else {
						t.Errorf("constructnamechecking need a baseobject")
					}
				case "tostringchecking":
					if objfrom, ok := val[0].Interface().(baseobject.ObjectFrom); ok {
						bobj := objfrom.BaseObject_()

						if settermethod := reflect.ValueOf(bobj).MethodByName("ToString_"); settermethod != (reflect.Value{}) {
							valconstruct := settermethod.Call([]reflect.Value{})

							AssertExpect(t, expectDesc["resultattempt"], valconstruct[0].Interface())

						} else {
							t.Errorf("Method %s not found in %T", "ToString_", bobj)
						}

					} else {
						t.Errorf("tostringchecking need a baseobject")
					}

				case "contains":
					AssertStringContains(t, expectDesc["resultattempt"], val[0].Interface())
				default:
					AssertExpect(t, expectDesc["resultattempt"], val[0].Interface())
				}

			}
		} else {
			t.Errorf("Method %s not found in %T", expectDesc["method"].(string), object)
		}

	})

}
