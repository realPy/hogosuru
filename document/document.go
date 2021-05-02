package document

import "github.com/realPy/jswasm/js"

//Root Get the root obj document
func Root() js.Value {

	//We panic if document not exist
	root := js.Global().Get("document")
	return root
}
