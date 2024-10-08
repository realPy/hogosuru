package document

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/documentfragment"
	"github.com/realPy/hogosuru/base/dragevent"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/node"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

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
		dragevent.GetInterface()

	})

	return docinterface
}

func New() (Document, error) {

	var d Document
	var err error
	if di := GetInterface(); !di.IsUndefined() {

		if dobj, err := baseobject.Get(js.Global(), "document"); err == nil {

			d.BaseObject = d.SetObject(dobj)
		}

	} else {

		err = ErrNotImplemented
	}

	return d, err
}

func NewFromJSObject(obj js.Value) (Document, error) {
	var d Document
	var err error
	if dci := GetInterface(); !dci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(dci) {

				d.BaseObject = d.SetObject(obj)

			} else {
				err = ErrNotADocument
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}
