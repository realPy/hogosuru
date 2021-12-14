package url

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/file"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`u=new URL('http://user:pass@www.mydomain.com:8888/test?arg=3#tag')`)

	m.Run()
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "u"); testingutils.AssertErr(t, err) {
		if event, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "URL", event.ConstructName_())

		}
	}

}

var methodsAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Hash", "resultattempt": "#tag"},
	{"method": "Host", "type": "contains", "resultattempt": "mydomain.com"},
	{"method": "Hostname", "resultattempt": "www.mydomain.com"},
	{"method": "Href", "type": "contains", "resultattempt": "www.mydomain.com"},
	{"method": "Origin", "type": "contains", "resultattempt": "www.mydomain.com"},
	{"method": "Port", "resultattempt": "8888"},
	{"method": "Pathname", "resultattempt": "/test"},
	{"method": "Protocol", "resultattempt": "http:"},
	{"method": "Username", "resultattempt": "user"},
	{"method": "Password", "resultattempt": "pass"},
	{"method": "Search", "resultattempt": "?arg=3"},
	{"method": "ToJSON", "resultattempt": "http://user:pass@www.mydomain.com:8888/test?arg=3#tag"},
	{"method": "SearchParams", "type": "tostringchecking", "resultattempt": "arg=3"},
}

func TestMethods(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "u"); testingutils.AssertErr(t, err) {

		if urlObj, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range methodsAttempt {
				testingutils.InvokeCheck(t, urlObj, result)
			}

		}

	}
}

var methodsSetterAttempt []map[string]interface{} = []map[string]interface{}{

	{"method": "SetHash", "args": []interface{}{"pouet"}, "gettermethod": "Hash", "resultattempt": "#pouet"},
	{"method": "SetHost", "args": []interface{}{"anotherdomain.com"}, "gettermethod": "Host", "resultattempt": "anotherdomain.com"},
	{"method": "SetHostname", "args": []interface{}{"www.anotheordomain.com"}, "gettermethod": "Hostname", "resultattempt": "www.anotheordomain.com"},
	{"method": "SetPathname", "args": []interface{}{"test"}, "gettermethod": "Pathname", "resultattempt": "/test"},
	{"method": "SetPort", "args": []interface{}{"3333"}, "gettermethod": "Port", "resultattempt": "3333"},
	{"method": "SetProtocol", "args": []interface{}{"https"}, "gettermethod": "Protocol", "resultattempt": "https:"},
	{"method": "SetUsername", "args": []interface{}{"user"}, "gettermethod": "Username", "resultattempt": "user"},
	{"method": "SetPassword", "args": []interface{}{"pass"}, "gettermethod": "Password", "resultattempt": "pass"},
	{"method": "SetSearch", "args": []interface{}{"p=yes"}, "gettermethod": "Search", "resultattempt": "?p=yes"},

	{"method": "SetHref", "args": []interface{}{"https://developer.mozilla.org/en-US/docs/Web/API/URL/href"}, "gettermethod": "Href", "resultattempt": "https://developer.mozilla.org/en-US/docs/Web/API/URL/href"},
}

func TestSetter(t *testing.T) {

	if urlObj, err := New("http://localhost"); testingutils.AssertErr(t, err) {

		for _, result := range methodsSetterAttempt {
			testingutils.InvokeCheck(t, urlObj, result)
		}

	}
}

func TestCreateObjectURL(t *testing.T) {

	baseobject.Eval("file = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })")

	if obj, err := baseobject.Get(js.Global(), "file"); testingutils.AssertErr(t, err) {

		if f, err := file.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if s, err := CreateObjectURL(f); testingutils.AssertErr(t, err) {

				testingutils.AssertStringContains(t, "blob:http://", s)
			}

		}
	}

}

func TestRevokeObjectURL(t *testing.T) {

	baseobject.Eval("file2 = new File(['(⌐□_□)'], 'chucknorris.png', { type: 'image/png' })")

	if obj, err := baseobject.Get(js.Global(), "file2"); testingutils.AssertErr(t, err) {

		if f, err := file.NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			if s, err := CreateObjectURL(f); testingutils.AssertErr(t, err) {

				testingutils.AssertErr(t, RevokeObjectURL(s))

			}

		}
	}

}
