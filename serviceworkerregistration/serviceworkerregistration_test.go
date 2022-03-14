package serviceworkerregistration

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`swregistercontainer=navigator.serviceWorker;
	promise=swregistercontainer.register('testserviceworkerregistration.js');

		p=promise.then(function(registration) {
			reg=registration;
		  });


	  
	
	`)

	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := promise.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objreg, err := baseobject.Get(js.Global(), "reg"); testingutils.AssertErr(t, err) {

					if servicew, err := NewFromJSObject(objreg); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, "ServiceWorkerRegistration", servicew.ConstructName_())

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
				t.Errorf("ServiceWorker request timeout")

			}

		}

	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Installing", "type": "constructnamechecking", "resultattempt": "ServiceWorker"},
	{"method": "Active", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "Scope", "type": "contains", "resultattempt": "localhost"},
	{"method": "Waiting", "type": "error", "resultattempt": baseobject.ErrUndefinedValue},
	{"method": "NavigationPreload", "type": "constructnamechecking", "resultattempt": "NavigationPreloadManager"},
	{"method": "PushManager", "type": "constructnamechecking", "resultattempt": "PushManager"},
	{"method": "GetNotifications", "args": []interface{}{"test"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "ShowNotification", "args": []interface{}{"test"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Unregister", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Update", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "UpdateViaCache", "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := promise.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objreg, err := baseobject.Get(js.Global(), "reg"); testingutils.AssertErr(t, err) {

					if servicew, err := NewFromJSObject(objreg); testingutils.AssertErr(t, err) {

						for _, result := range methodsAttempt {
							testingutils.InvokeCheck(t, servicew, result)
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
				t.Errorf("ServiceWorker request timeout")

			}

		}

	}

}
