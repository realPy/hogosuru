package document

import (
	"testing"
)

func AssertErr(t *testing.T, err error) bool {

	if err != nil {
		t.Errorf(err.Error())
		return false
	}

	return true
}

func TestMain(m *testing.M) {
	/*
		baseobject.Eval(`go.importObject.env["syscall/js.valueSetErr"] = 	(ret_addr, v_addr, p_ptr, p_len, x_addr) => {
				const v = loadValue(v_addr);
				const p = loadString(p_ptr, p_len);
				const x = loadValue(x_addr);
				try {
					Reflect.set(v, p, x);
					mem().setUint8(ret_addr + 8, 1);
				} catch (err) {
					storeValue(ret_addr, err);
					mem().setUint8(ret_addr+ 8, 0);
				}

			}`)
	*/

	m.Run()
}

func TestGetBody(t *testing.T) {

	if d, err := New(); err == nil {

		if str, err := d.ContentType(); AssertErr(t, err) {
			if str != "text/html" {
				t.Errorf("Content Type must be text/html, have %s", str)
			}
		}

		if str, err := d.Domain(); AssertErr(t, err) {
			if str != "localhost" {
				t.Errorf("Domain must be text/html, have %s", str)
			}
		}

		//d.SetDomain("toto.com")
		AssertErr(t, d.SetDomain("toto.com"))

		if str, err := d.Title(); AssertErr(t, err) {
			if str != "Go wasm" {
				t.Errorf("Title must be Go wasm, have %s", str)
			}

		}

		AssertErr(t, d.SetTitle("Hello"))

		if str, err := d.Title(); AssertErr(t, err) {
			if str != "Hello" {
				t.Errorf("Title must be Hello, have %s", str)
			}
		}

		if str, err := d.Cookie(); AssertErr(t, err) {
			if str != "" {
				t.Errorf("Cookie must be hello world, have %s", str)
			}
		}
		AssertErr(t, d.SetCookie("hello world"))
		if str, err := d.Cookie(); AssertErr(t, err) {
			if str != "hello world" {
				t.Errorf("Cookie must be hello world, have %s", str)
			}
		}

	}

}
