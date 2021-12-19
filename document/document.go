package document

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/documentfragment"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var docinterface js.Value

type Document struct {
	node.Node
}

type DocumentFrom interface {
	Document_() Document
}

func (d Document) Document_() Document {
	return d
}

func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if docinterface, err = baseobject.Get(js.Global(), "Document"); err != nil {
			docinterface = js.Undefined()
		}
		baseobject.Register(docinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
		node.GetInterface()
		documentfragment.GetInterface()

	})

	return docinterface
}

func New() (Document, error) {
	var d Document
	var di, dobj js.Value
	var err error
	if di = GetInterface(); di.IsUndefined() {
		return d, ErrNotImplemented
	}
	if dobj, err = baseobject.Get(js.Global(), "document"); err != nil {
		return d, err
	}
	d.BaseObject = d.SetObject(dobj)
	return d, nil
}

func NewFromJSObject(obj js.Value) (Document, error) {
	var d Document
	var dci js.Value
	if dci = GetInterface(); dci.IsUndefined() {
		return d, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return d, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(dci) {
		return d, ErrNotADocument
	}
	d.BaseObject = d.SetObject(obj)
	return d, nil
}
