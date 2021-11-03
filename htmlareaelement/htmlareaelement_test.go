package htmlareaelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`a= document.createElement("area")
	a.accessKey="o"
	a.download="tdownload"
	a.href="https://myuser:mypass@www.test.com:444?q=123#tag"
	a.rel="nofollow"
	a.tabIndex=123
	a.text="textanchor"
	a.type="customtype"
	a.target="thistarget"
	a.alt="altvalue"
	a.shape="rect"
	a.coords="0,0,253,27"
	`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if a, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLAreaElement", a.ConstructName_())
		}

	}

}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "a"); testingutils.AssertErr(t, err) {

		if a, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLAreaElement", a.ConstructName_())
		}

	}
}

var getterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "AccessKey", "resultattempt": "o"},
	{"method": "Alt", "resultattempt": "altvalue"},
	{"method": "Coords", "resultattempt": "0,0,253,27"},
	{"method": "Download", "resultattempt": "tdownload"},
	{"method": "Hash", "resultattempt": "#tag"},
	{"method": "Host", "resultattempt": "www.test.com:444"},
	{"method": "Hostname", "resultattempt": "www.test.com"},
	{"method": "Href", "resultattempt": "https://myuser:mypass@www.test.com:444/?q=123#tag"},
	{"method": "Origin", "resultattempt": "https://www.test.com:444"},
	{"method": "Password", "resultattempt": "mypass"},
	{"method": "Pathname", "resultattempt": "/"},
	{"method": "Port", "resultattempt": "444"},
	{"method": "Protocol", "resultattempt": "https:"},
	{"method": "ReferrerPolicy", "resultattempt": ""},
	{"method": "Rel", "resultattempt": "nofollow"},
	{"method": "RelList", "type": "constructnamechecking", "resultattempt": "DOMTokenList"},
	{"method": "Search", "resultattempt": "?q=123"},
	{"method": "Shape", "resultattempt": "rect"},
	{"method": "TabIndex", "resultattempt": 123},
	{"method": "Target", "resultattempt": "thistarget"},
	{"method": "Username", "resultattempt": "myuser"},
}

func TestGetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "a"); testingutils.AssertErr(t, err) {

		if area, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range getterAttempt {
				testingutils.InvokeCheck(t, area, result)
			}

		}

	}
}

var setterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "SetAccessKey", "args": []interface{}{"i"}, "gettermethod": "AccessKey", "resultattempt": "i"},
	{"method": "SetAlt", "args": []interface{}{"n2"}, "gettermethod": "Alt", "resultattempt": "n2"},
	{"method": "SetCoords", "args": []interface{}{"99,99,66,27"}, "gettermethod": "Coords", "resultattempt": "99,99,66,27"},
	{"method": "SetDownload", "args": []interface{}{"testvalue"}, "gettermethod": "Download", "resultattempt": "testvalue"},
	{"method": "SetPort", "args": []interface{}{"445"}, "gettermethod": "Port", "resultattempt": "445"},
	{"method": "SetHost", "args": []interface{}{"testhost"}, "gettermethod": "Host", "resultattempt": "testhost:445"},
	{"method": "SetHostname", "args": []interface{}{"testhostname"}, "gettermethod": "Hostname", "resultattempt": "testhostname"},
	{"method": "SetHref", "args": []interface{}{"http://pp:ss@www.noone.com:444?q=456#nosecure"}, "gettermethod": "Href", "resultattempt": "http://pp:ss@www.noone.com:444/?q=456#nosecure"},
	{"method": "SetPassword", "args": []interface{}{"nopass"}, "gettermethod": "Password", "resultattempt": "nopass"},
	{"method": "SetUsername", "args": []interface{}{"toto"}, "gettermethod": "Username", "resultattempt": "toto"},
	{"method": "SetPathname", "args": []interface{}{"/toto/"}, "gettermethod": "Pathname", "resultattempt": "/toto/"},
	{"method": "SetProtocol", "args": []interface{}{"http"}, "gettermethod": "Protocol", "resultattempt": "http:"},
	{"method": "SetReferrerPolicy", "args": []interface{}{"no-referrer"}, "gettermethod": "ReferrerPolicy", "resultattempt": "no-referrer"},
	{"method": "SetRel", "args": []interface{}{"alternate"}, "gettermethod": "Rel", "resultattempt": "alternate"},
	{"method": "SetSearch", "args": []interface{}{"?p=098"}, "gettermethod": "Search", "resultattempt": "?p=098"},
	{"method": "SetShape", "args": []interface{}{"circle"}, "gettermethod": "Shape", "resultattempt": "circle"},
	{"method": "SetTabIndex", "args": []interface{}{777}, "gettermethod": "TabIndex", "resultattempt": 777},
	{"method": "SetTarget", "args": []interface{}{"yes"}, "gettermethod": "Target", "resultattempt": "yes"},
}

func TestSetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "a"); testingutils.AssertErr(t, err) {

		if area, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range setterAttempt {

				testingutils.InvokeCheck(t, area, result)

			}

		}

	}

}
