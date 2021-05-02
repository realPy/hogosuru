package json

import (
	"sync"

	"github.com/realPy/jswasm/js"
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
	object js.Value
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

//NewJsonFromString Parse a json str
func NewJsonFromString(jsonstr string) (Json, error) {
	var json Json

	if jsoni := GetJSInterface(); jsoni != nil {

		if json, err := jsoni.objectInterface.CallWithErr("parse", jsonstr); err != nil {
			return Json{}, err
		} else {
			return Json{object: json}, nil
		}

	}

	return json, ErrNotImplemented

}

func (j Json) Get(key string) js.Value {
	return j.object.Get(key)
}

func recurseUmarshalJson(object js.Value) jsobject.GOMap {

	return jsobject.Map(object)

	return jsobject.GOMap{}
}
func (j Json) GoJson() jsobject.GOMap {

	return recurseUmarshalJson(j.object)
}
