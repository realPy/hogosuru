package document

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var docinterface js.Value

type Document struct {
	node.Node
}

func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if docinterface, err = js.Global().GetWithErr("document"); err != nil {
			docinterface = js.Null()
		}
	})

	baseobject.Register(docinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return docinterface
}

func New() (Document, error) {

	var d Document
	var err error
	if di := GetInterface(); !di.IsNull() {
		d.BaseObject = d.SetObject(di)

	} else {

		err = ErrNotImplemented
	}

	return d, err
}

func NewFromJSObject(obj js.Value) (Document, error) {
	var d Document

	if dci := GetInterface(); !dci.IsNull() {
		if obj.InstanceOf(dci) {

			d.BaseObject = d.SetObject(obj)
			return d, nil
		}
	}
	return d, ErrNotADocument
}
