package document

import (
	"syscall/js"

	"github.com/realPy/hogosuru/base/nodelist"
)

// Close Closer interface
func (d Document) Close() error {

	_, err := d.Call("close")
	return err

}

func (d Document) GetElementsByName(name string) (nodelist.NodeList, error) {

	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj, err = d.Call("getElementsByName", js.ValueOf(name)); err == nil {

		nlist, err = nodelist.NewFromJSObject(obj)
	}
	return nlist, err
}

func (d Document) getSelection() {
	//TO IMPLEMENT
}

func (d Document) HasFocus() (bool, error) {

	return d.GetAttributeBool("hasFocus")
}

// Close Closer interface
func (d Document) Open() error {

	_, err := d.Call("open")
	return err
}

/* TO IMPLEMENTED
document.queryCommandValue
document.queryCommandSupported
document.queryCommandState
document.queryCommandIndeterm
document.queryCommandEnabled
*/

// Write Writer interface
func (d Document) Write(p []byte) (n int, err error) {
	n = len(p)
	_, err = d.Call("write", js.ValueOf(string(p)))
	return
}

func (d Document) Writeln(text string) error {

	_, err := d.Call("writeln", js.ValueOf(text))
	return err
}
