package iterator

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
)

// Iterator iterator
type Iterator struct {
	baseobject.BaseObject
}

type IteratorFrom interface {
	Iterator_() Iterator
}

func (i Iterator) Iterator_() Iterator {
	return i
}

func NewFromJSObject(obj js.Value) (Iterator, error) {
	var i Iterator
	var err error
	var objf js.Value

	symbol := js.Global().Get("Symbol")
	it := symbol.Get("iterator")

	if objf, err = baseobject.Get(obj, it); err == nil {
		if objf.Type() == js.TypeFunction {
			i.BaseObject = i.SetObject(obj)
		} else {
			err = NotAnIterator
		}
	}

	return i, err
}

func pairValues(obj js.Value) (interface{}, interface{}) {

	var value interface{}
	var index interface{}

	if obj.Type() == js.TypeObject {
		if obj.Length() == 2 {

			index = baseobject.GoValue_(obj.Index(0))

			value = baseobject.GoValue_(obj.Index(1))

		}

	}
	return index, value
}

/* Parse using

for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {

}
*/

func (i Iterator) Next() (interface{}, interface{}, error) {

	var err error
	var done bool = true
	var obj, doneobj, valueobj js.Value
	var index interface{}
	var value interface{}

	if obj, err = i.Call("next"); err == nil {

		if doneobj, err = baseobject.Get(obj, "done"); err == nil {
			if doneobj.Type() == js.TypeBoolean {
				done = doneobj.Bool()
			} else {
				err = baseobject.ErrObjectNotBool
			}
		}
		if done {
			err = EOI

		} else {

			if valueobj, err = baseobject.Get(obj, "value"); err == nil {
				if valueobj.Type() == js.TypeObject {
					index, value = pairValues(valueobj)
				} else {
					value, err = baseobject.GoValue(valueobj)
				}

			}
		}

	}
	return index, value, err
}
