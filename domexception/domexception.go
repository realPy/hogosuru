package domexception

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var domexceptioninterface js.Value

//DomException DomException struct
type DomException struct {
	baseobject.BaseObject
}

type DomExceptionFrom interface {
	DomException_() DomException
}

func (d DomException) DomException_() DomException {
	return d
}

//GetJSInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if domexceptioninterface, err = baseobject.Get(js.Global(), "DOMException"); err != nil {
			domexceptioninterface = js.Undefined()
		}

		baseobject.Register(domexceptioninterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})
	return domexceptioninterface
}

func New(opts ...string) (DomException, error) {

	var e DomException
	var arrayJS []interface{}
	var obj js.Value
	var err error
	if len(opts) < 3 {
		for _, opt := range opts {
			arrayJS = append(arrayJS, js.ValueOf(opt))
		}
	}

	if ei := GetInterface(); !ei.IsUndefined() {

		if obj, err = baseobject.New(ei, arrayJS...); err == nil {
			e.BaseObject = e.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func NewFromJSObject(obj js.Value) (DomException, error) {
	var d DomException
	var err error
	if di := GetInterface(); !di.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(di) {
				d.BaseObject = d.SetObject(obj)
			} else {
				err = ErrNotADOMException
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func (d DomException) Message() (string, error) {
	return d.GetAttributeString("message")
}

func (d DomException) Name() (string, error) {
	return d.GetAttributeString("name")
}
