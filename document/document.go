package document

import (
	"sync"

	"syscall/js"

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

func New() (Document, error) {

	var d Document
	var err error
	if di := GetJSInterface(); di != nil {
		d.BaseObject = d.SetObject(di.objectInterface)

	} else {

		err = ErrNotImplemented
	}

	return d, err
}
