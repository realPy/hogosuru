package document

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/object"
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

func (d Document) CharacterSet() (string, error) {

	var err error
	var obj js.Value
	var cs string = ""
	if obj, err = d.JSObject().GetWithErr("characterSet"); err == nil {

		cs = obj.String()
	}
	return cs, err
}

func (d Document) ChildElementCount() int {
	var err error
	var obj js.Value

	if obj, err = d.JSObject().GetWithErr("childElementCount"); err == nil {
		if obj.Type() == js.TypeNumber {
			return obj.Int()
		}
	}

	return 0
}

func (d Document) Children() (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = d.JSObject().GetWithErr("children"); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) CompatMode() (string, error) {

	var err error
	var obj js.Value
	var mode string = "s"

	if obj, err = d.JSObject().GetWithErr("compatMode"); err == nil {

		mode = obj.String()
	}
	return mode, err
}

func (d *Document) ContentType() (string, error) {

	var err error
	var obj js.Value
	var ct string = ""
	if obj, err = d.JSObject().GetWithErr("contentType"); err == nil {

		ct = obj.String()
	}
	return ct, err
}

func (d *Document) Doctype() {

}

func (d Document) DocumentElement() element.Element {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.JSObject().GetWithErr("documentElement"); err == nil {
		elem = element.NewFromJSObject(obj)
	} else {
		elem.Error = &err
	}
	return elem
}

func (d *Document) DocumentURI() (string, error) {

	var err error
	var obj js.Value
	var uri string = ""

	if obj, err = d.JSObject().GetWithErr("documentURI"); err == nil {

		uri = obj.String()
	}
	return uri, err
}

func (d Document) Embeds() (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = d.JSObject().GetWithErr("embeds"); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) FirstElementChild() element.Element {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.JSObject().GetWithErr("firstElementChild"); err == nil {
		elem = element.NewFromJSObject(obj)
	} else {
		elem.Error = &err
	}
	return elem
}

func (d Document) Fonts() {

}

func (d Document) Forms() (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = d.JSObject().GetWithErr("forms"); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) FullscreenElement() element.Element {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.JSObject().GetWithErr("fullscreenElement"); err == nil {
		elem = element.NewFromJSObject(obj)
	} else {
		elem.Error = &err
	}
	return elem
}

func (d Document) Head() (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = d.JSObject().GetWithErr("head"); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) Hidden() (bool, error) {

	var err error
	var obj js.Value
	var ret bool

	if obj, err = d.JSObject().GetWithErr("hidden"); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return ret, err
}

func (d Document) Images() (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = d.JSObject().GetWithErr("images"); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}
func (d Document) Implementation() {
	//not implemented
}

func (d Document) LastElementChild() element.Element {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.JSObject().GetWithErr("lastElementChild"); err == nil {
		elem = element.NewFromJSObject(obj)
	} else {
		elem.Error = &err
	}
	return elem
}

func (d Document) Links() (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = d.JSObject().GetWithErr("links"); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) PictureInPictureElement() element.Element {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.JSObject().GetWithErr("pictureInPictureElement"); err == nil {
		elem = element.NewFromJSObject(obj)
	} else {
		elem.Error = &err
	}
	return elem
}

func (d Document) PictureInPictureEnabled() (bool, error) {

	var err error
	var obj js.Value
	var ret bool

	if obj, err = d.JSObject().GetWithErr("pictureInPictureEnabled"); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = object.ErrObjectNotBool
		}
	}

	return ret, err
}

func (d Document) Plugins() (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = d.JSObject().GetWithErr("plugins"); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) PointerLockElement() element.Element {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.JSObject().GetWithErr("pointerLockElement"); err == nil {
		elem = element.NewFromJSObject(obj)
	} else {
		elem.Error = &err
	}
	return elem
}

func (d Document) Scripts() (htmlcollection.HTMLCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HTMLCollection

	if obj, err = d.JSObject().GetWithErr("scripts"); err == nil {
		collection, err = htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (d Document) ScrollingElement() element.Element {
	var err error
	var obj js.Value
	var elem element.Element

	if obj, err = d.JSObject().GetWithErr("scrollingElement"); err == nil {
		elem = element.NewFromJSObject(obj)
	} else {
		elem.Error = &err
	}
	return elem
}

func (d Document) VisibilityState() (string, error) {

	var err error
	var obj js.Value
	var vis string = ""

	if obj, err = d.JSObject().GetWithErr("visibilityState"); err == nil {

		vis = obj.String()
	}
	return vis, err
}

func (d Document) QuerySelector(selector string) (js.Value, error) {

	return d.JSObject().CallWithErr("querySelector", js.ValueOf(selector))
}

func (d Document) Cookie() (string, error) {
	var err error
	var obj js.Value
	var cookie string = ""
	if obj, err = d.JSObject().GetWithErr("cookie"); err == nil {

		cookie = obj.String()
	}
	return cookie, err
}

func (d Document) SetCookie(cookie string) error {

	return d.JSObject().SetWithErr("cookie", js.ValueOf(cookie))
}
