package navigationpreloadmanager

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
	promise=swregistercontainer.register('testnavigationpreloadmanager.js');

		p=promise.then(function(registration) {
			np=registration.navigationPreload;

		  });


	  
	
	`)

	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := promise.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objpm, err := baseobject.Get(js.Global(), "np"); testingutils.AssertErr(t, err) {

					if pm, err := NewFromJSObject(objpm); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, "NavigationPreloadManager", pm.ConstructName_())

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
	{"method": "Enable", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "Disable", "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "SetHeaderValue", "args": []interface{}{"test"}, "type": "constructnamechecking", "resultattempt": "Promise"},
	{"method": "GetState", "type": "constructnamechecking", "resultattempt": "Promise"},
}

func TestMethods(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := promise.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objpm, err := baseobject.Get(js.Global(), "np"); testingutils.AssertErr(t, err) {

					if pm, err := NewFromJSObject(objpm); testingutils.AssertErr(t, err) {

						for _, result := range methodsAttempt {
							testingutils.InvokeCheck(t, pm, result)
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
