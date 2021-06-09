package document

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var docinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if docinterface, err = js.Global().GetWithErr("document"); err != nil {
			docinterface = js.Null()
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
	if di := GetInterface(); !di.IsNull() {
		d.BaseObject = d.SetObject(di)

	} else {

		err = ErrNotImplemented
	}

	return d, err
}
