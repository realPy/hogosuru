package fetch

import (
	"testing"
	"time"

	"github.com/realPy/hogosuru/abortcontroller"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/json"
	"github.com/realPy/hogosuru/object"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/response"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {
	var io chan bool = make(chan bool)

	//Start promise and wait result
	t.Run("Get ", func(t *testing.T) {
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
	})
	t.Run("Get with custom headers", func(t *testing.T) {

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
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

	t.Run("Post with custom headers", func(t *testing.T) {

		var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/json",
			"XCustomValue": "Test"}

		var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers}

		//Start promise and wait result
		if f, err := New("https://httpbin.org/post", fetchOpts); err == nil {
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
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

	object.GetInterface()

	t.Run("Post with custom headers and json response and form data ", func(t *testing.T) {

		var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded",
			"XCustomValue": "Test"}

		var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers, "body": "data=test"}

		//Start promise and wait result
		if f, err := New("https://httpbin.org/post", fetchOpts); err == nil {
			jsonpromise, _ := f.Then(func(r response.Response) *promise.Promise {
				if status, err := r.Status(); err == nil {
					if status != 200 {
						t.Errorf("Status must be 200 , give %d", status)
					} else {

						if promise, err := r.Json(); err == nil {
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

			jsonpromise.Then(func(i interface{}) *promise.Promise {

				if obj, ok := i.(object.ObjectFrom); ok {

					j, _ := json.NewFromJSObject(obj.Object_().JSObject())
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
					t.Error("No a json")
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

	})

}

func TestNewCancelable(t *testing.T) {
	var io chan bool = make(chan bool)

	t.Run("Post with custom headers and json response and form data ", func(t *testing.T) {

		var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded",
			"XCustomValue": "Test"}

		var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers, "body": "data=test", "mode": "no-cors"}

		if f, err := NewCancelable("http://httpbin.org/post", fetchOpts); err == nil {
			f.Then(func(r response.Response) *promise.Promise {

				t.Error("Must not get response")
				return nil
			}, func(e error) {
				if e.Error() != "The user aborted a request." {
					t.Error("Error mismatch")
				}
				io <- true
			})

			f.Abort()
		} else {
			t.Error(err.Error())
		}

		select {
		case <-io:
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

	t.Run("Post with custom headers and json response and custom signal ", func(t *testing.T) {

		var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded",
			"XCustomValue": "Test"}

		abortctrl, _ := abortcontroller.New()

		s, _ := abortctrl.Signal()

		var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers, "body": "data=test", "mode": "no-cors", "signal": s.JSObject()}

		if f, err := NewCancelable("http://httpbin.org/post", fetchOpts); err == nil {
			f.Then(func(r response.Response) *promise.Promise {

				t.Error("Must not get response")
				return nil
			}, func(e error) {
				if e.Error() != "The user aborted a request." {
					t.Error("Error mismatch")
				}
				io <- true
			})

			if err := f.Abort(); err != ErrSignalNotManaged {

				t.Error("Must throw an error")
			}
			abortctrl.Abort()
		} else {
			t.Error(err.Error())
		}

		select {
		case <-io:
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

}
