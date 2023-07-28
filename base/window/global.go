package window

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
)

func Alert(message string) error {
	w, err := self()
	if err != nil {
		return err
	}
	_, err = w.Call("alert", js.ValueOf(message))
	return err
}

func Confirm(message string) (bool, error) {
	var ret = false
	w, err := self()
	if err != nil {
		return ret, err
	}
	if obj, err := w.Call("confirm", js.ValueOf(message)); err == nil {

		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return ret, err
}
func Prompt(message, input string) (*string, error) {
	var ret *string = nil
	w, err := self()
	if err != nil {
		return ret, err
	}
	if obj, err := w.Call("prompt", js.ValueOf(message), js.ValueOf(input)); err == nil {

		if obj.Type() == js.TypeString {
			v := obj.String()
			ret = &v
		}
	}
	return ret, err
}

func Atob(encoded string) (string, error) {

	var err error
	var result string
	var obj js.Value

	w, err := self()
	if err != nil {
		return "", err
	}
	if obj, err = w.Call("atob", js.ValueOf(encoded)); err == nil {
		if obj.Type() == js.TypeString {
			result = obj.String()
		} else {
			err = baseobject.ErrObjectNotString
		}
	}
	return result, err
}

func Btoa(message string) (string, error) {

	var err error
	var result string
	var obj js.Value

	w, err := self()
	if err != nil {
		return "", err
	}
	if obj, err = w.Call("btoa", js.ValueOf(message)); err == nil {
		if obj.Type() == js.TypeString {
			result = obj.String()
		} else {
			err = baseobject.ErrObjectNotString
		}
	}
	return result, err
}

func Open(url string, opts ...interface{}) (Window, error) {
	w, err := self()
	if err != nil {
		return Window{}, err
	}
	var arrayJS []interface{}

	arrayJS = append(arrayJS, js.ValueOf(url))
	for _, opt := range opts {
		arrayJS = append(arrayJS, js.ValueOf(opt))
	}

	if obj, err := w.Call("open", arrayJS...); err == nil {
		var w Window
		w.BaseObject = w.SetObject(obj)
		return w, nil
	} else {
		return Window{}, err
	}

}
