package serviceworker

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`swregistercontainer=navigator.serviceWorker;
	promise=swregistercontainer.register('testserviceworker.js');
	p=promise.then(function(registration) {
		serviceworker=registration.installing

	  });
	`)

	m.Run()

}

func TestNewFromJSObject(t *testing.T) {

	var wchan chan bool = make(chan bool)

	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := promise.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objsw, err := baseobject.Get(js.Global(), "serviceworker"); testingutils.AssertErr(t, err) {

					if servicew, err := NewFromJSObject(objsw); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, "ServiceWorker", servicew.ConstructName_())

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
	{"method": "ScriptURL", "type": "contains", "resultattempt": "/testserviceworker.js"},
	{"method": "State", "resultattempt": "installing"},
}

func TestMethods(t *testing.T) {

	var wchan chan bool = make(chan bool)
	if obj, err := baseobject.Get(js.Global(), "p"); testingutils.AssertErr(t, err) {

		if p, err := promise.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			p.Then(func(i interface{}) *promise.Promise {

				if objsw, err := baseobject.Get(js.Global(), "serviceworker"); testingutils.AssertErr(t, err) {

					if servicew, err := NewFromJSObject(objsw); testingutils.AssertErr(t, err) {

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
