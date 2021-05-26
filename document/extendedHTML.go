package document

import (
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/nodelist"
)

//Close Closer interface
func (d Document) Close() error {

	_, err := d.JSObject().CallWithErr("close")
	return err

}

func (d Document) GetElementsByName(name string) (nodelist.NodeList, error) {

	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj, err = d.JSObject().CallWithErr("getElementsByName", js.ValueOf(name)); err == nil {

		nlist, err = nodelist.NewFromJSObject(obj)
	}
	return nlist, err
}

func (d Document) getSelection() {
	//TO IMPLEMENT
}

func (d Document) HasFocus() (bool, error) {

	var err error
	var result bool
	var obj js.Value

	if obj, err = d.JSObject().GetWithErr("hasFocus"); err == nil {
		if obj.Type() == js.TypeBoolean {
			result = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}
	return result, err
}

//Close Closer interface
func (d Document) Open() error {

	_, err := d.JSObject().CallWithErr("open")
	return err
}

/* TO IMPLEMENTED
document.queryCommandValue
document.queryCommandSupported
document.queryCommandState
document.queryCommandIndeterm
document.queryCommandEnabled
*/

//Write Writer interface
func (d Document) Write(p []byte) (n int, err error) {
	n = len(p)
	_, err = d.JSObject().CallWithErr("write", js.ValueOf(string(p)))
	return
}

func (d Document) Writeln(text string) error {

	_, err := d.JSObject().CallWithErr("writeln", js.ValueOf(text))
	return err
}
