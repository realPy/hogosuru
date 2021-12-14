package url

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`u=new URL('http://user:pass@mydomain.com/test?arg=3')`)
	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "u"); testingutils.AssertErr(t, err) {
		if event, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "URL", event.ConstructName_())

		}
	}

}

/*
var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Hash", "resultattempt": ""},
	{"method": "Host", "type": "contains", "resultattempt": "localhost"},
	{"method": "Hostname", "resultattempt": "localhost"},
	{"method": "Href", "type": "contains", "resultattempt": "localhost"},
	{"method": "Origin", "type": "contains", "resultattempt": "localhost"},
	{"method": "Pathname", "resultattempt": "/"},
	{"method": "Protocol", "resultattempt": "http:"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "l"); testingutils.AssertErr(t, err) {

		if location, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, location, result)
			}

		}

	}
}

func TestPort(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "l"); testingutils.AssertErr(t, err) {

		if location, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if str, err := location.Port(); testingutils.AssertErr(t, err) {
				if i, err := strconv.Atoi(str); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, true, i > 0)

				}

			}

		}
	}
}
*/
