package permissionstatus

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {

	baseobject.SetSyscall()

	baseobject.Eval(`
	p=navigator.permissions.query({name:'clipboard-read'}).then(function(permissionStatus) {
		permstatus=permissionStatus;
		});
		`)
	m.Run()

}

func TestNewFromJSObject(t *testing.T) {
	var wchan chan bool = make(chan bool)
	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := promise.NewFromJSObject(obj); hogosuru.AssertErr(err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objperm, err := baseobject.Get(js.Global(), "permstatus"); testingutils.AssertErr(t, err) {

					if permstatus, err := NewFromJSObject(objperm); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, "PermissionStatus", permstatus.ConstructName_())

					}

				}
				wchan <- true

				return nil
			}, func(e error) {

				t.Errorf(e.Error())
			})

			select {
			case <-wchan:
			case <-time.After(time.Duration(200) * time.Millisecond):
				t.Errorf("Permission request timeout")

			}

		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Name", "resultattempt": "clipboard_read"},
	{"method": "State", "resultattempt": "prompt"},
}

func TestMethods(t *testing.T) {

	var wchan chan bool = make(chan bool)
	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := promise.NewFromJSObject(obj); hogosuru.AssertErr(err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objperm, err := baseobject.Get(js.Global(), "permstatus"); testingutils.AssertErr(t, err) {

					if permstatus, err := NewFromJSObject(objperm); testingutils.AssertErr(t, err) {

						for _, result := range methodsAttempt {
							testingutils.InvokeCheck(t, permstatus, result)
						}

					}

				}
				wchan <- true

				return nil
			}, func(e error) {

				t.Errorf(e.Error())
			})

			select {
			case <-wchan:
			case <-time.After(time.Duration(200) * time.Millisecond):
				t.Errorf("Permission request timeout")

			}

		}

	}

}
