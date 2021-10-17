package fetch

import (
	"testing"
	"time"

	"github.com/realPy/hogosuru/json"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/response"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestNew(t *testing.T) {
	var io chan bool = make(chan bool)

	//Start promise and wait result
	if f, err := New("https://httpbin.org/get"); err == nil {
		f.Then(func(r response.Response) *promise.Promise {

			if status, err := r.Status(); err == nil {
				if status != 200 {
					t.Errorf("Status must be 200 , give %d", status)
				}
				io <- true
			} else {
				t.Error(err.Error())
			}
			return nil
		}, func(e error) {

			t.Error(e.Error())
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(2000) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

	var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/json",
		"XCustomValue": "Test"}

	var fetchOpts map[string]interface{} = map[string]interface{}{"method": "GET", "headers": headers}

	//Start promise and wait result
	if f, err := New("https://httpbin.org/get", fetchOpts); err == nil {
		textpromise, _ := f.Then(func(r response.Response) *promise.Promise {
			if status, err := r.Status(); err == nil {
				if status != 200 {
					t.Errorf("Status must be 200 , give %d", status)
				} else {

					if promise, err := r.Text(); err == nil {
						return &promise
					} else {
						t.Error(err.Error())
					}

				}
			} else {
				t.Error(err.Error())
			}
			return nil
		}, func(e error) {

			t.Error(e.Error())
		})

		textpromise.Then(func(i interface{}) *promise.Promise {

			if j, err := json.Parse(i.(string)); err == nil {
				goValue := j.Map()

				headers := goValue.(map[string]interface{})["headers"]

				if headers != nil {
					customValue := headers.(map[string]interface{})["Xcustomvalue"]

					if customValue != nil {
						if customValue.(string) == "Test" {
							io <- true
						} else {
							t.Errorf("Xcustomvalue not match %s", customValue.(string))
						}
					} else {
						t.Error("No Xcustomvalue headers present")
					}

				} else {
					t.Error("No headers present")
				}

			} else {
				t.Error(err.Error())
			}

			return nil
		}, func(e error) {
			t.Error(e.Error())
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(4000) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}
