package objectmap

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/iterator"
)

var singleton sync.Once

var mapinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if mapinterface, err = js.Global().GetWithErr("Map"); err != nil {
			mapinterface = js.Null()
		}

	})

	return mapinterface
}

//ObjectMap
type ObjectMap struct {
	baseobject.BaseObject
}

func New(values ...interface{}) (ObjectMap, error) {

	var o ObjectMap
	var err error

	var arrayJS []interface{}

	for _, value := range values {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			arrayJS = append(arrayJS, objGo.JSObject())
		} else {
			arrayJS = append(arrayJS, js.ValueOf(value))
		}

	}

	if omi := GetInterface(); !omi.IsNull() {

		o.BaseObject = o.SetObject(omi.New(arrayJS...))

	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func NewFromGo(values map[string]interface{}) (ObjectMap, error) {

	var arrayJS []interface{}

	for key, value := range values {
		arrayJS = append(arrayJS, array.New_(key, value))
	}

	return New(array.New_(arrayJS...))

}

func (o ObjectMap) Clear() error {
	var err error
	_, err = o.JSObject().CallWithErr("clear")
	return err
}

func (o ObjectMap) Delete(b baseobject.BaseObject) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj, err = o.JSObject().CallWithErr("delete", b.JSObject()); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (o ObjectMap) Entries() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = o.JSObject().CallWithErr("entries"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) ForEach(f func(ObjectMap, interface{}, interface{})) error {
	var err error

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(o, baseobject.GoValue(args[0]), baseobject.GoValue(args[0]))
		return nil
	})

	_, err = o.JSObject().CallWithErr("forEach", jsfunc)

	return err
}

func (o ObjectMap) Get(b baseobject.BaseObject) (interface{}, error) {

	var err error
	var obj js.Value
	var result interface{}

	if obj, err = o.JSObject().CallWithErr("get", b.JSObject()); err == nil {
		result = baseobject.GoValue(obj)
	}
	return result, err
}

func (o ObjectMap) Has(b baseobject.BaseObject) (bool, error) {
	var err error
	var obj js.Value
	var result bool
	if obj, err = o.JSObject().CallWithErr("has", b.JSObject()); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return result, err
}

func (o ObjectMap) Keys() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = o.JSObject().CallWithErr("keys"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}

func (o ObjectMap) Set(b baseobject.BaseObject) error {
	var err error
	_, err = o.JSObject().CallWithErr("set", b.JSObject())
	return err
}

func (o ObjectMap) Values() (iterator.Iterator, error) {
	var err error
	var obj js.Value
	var iter iterator.Iterator

	if obj, err = o.JSObject().CallWithErr("values"); err == nil {
		iter = iterator.NewFromJSObject(obj)
	}

	return iter, err
}
