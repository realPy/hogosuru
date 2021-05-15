package document

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var docinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var docinstance JSInterface
		var err error
		if docinstance.objectInterface, err = js.Global().GetWithErr("document"); err == nil {
			docinterface = &docinstance
		}
	})

	return docinterface
}

type Document struct {
	node.Node
}

func New() Document {

	var d Document
	if di := GetJSInterface(); di != nil {
		d.Object = d.SetObject(di.objectInterface)
		return d
	}
	d.Error = &ErrNotImplemented

	return d
}

func (d Document) ActiveElement() element.Element {
	var elem element.Element
	var elemObject js.Value
	var err error

	elem.Error = d.Error
	if d.Error == nil {
		if elemObject, err = d.JSObject().GetWithErr("activeElement"); err == nil {

			elem = element.NewFromJSObject(elemObject)

		} else {
			elem.Error = &err
		}

	}

	return elem
}

func (d Document) Body() node.Node {
	var body node.Node
	var bodyObject js.Value
	var err error

	body.Error = d.Error
	if d.Error == nil {
		if bodyObject, err = d.JSObject().GetWithErr("body"); err == nil {

			body = node.NewFromJSObject(bodyObject)

		} else {
			body.Error = &err
		}

	}

	return body
}

func (d Document) QuerySelector(selector string) (js.Value, error) {

	return d.JSObject().CallWithErr("querySelector", js.ValueOf(selector))
}

func (d Document) Cookie() string {
	var err error
	var obj js.Value

	if obj, err = d.JSObject().GetWithErr("cookie"); err == nil {

		return obj.String()
	} else {
		d.Error = &err
	}
	return ""
}

func (d *Document) SetCookie(cookie string) {
	var err error
	if d.Error == nil {
		if err = d.JSObject().SetWithErr("cookie", js.ValueOf(cookie)); err != nil {

			d.Error = &err
		}
	}
}
