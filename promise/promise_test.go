package promise

import (
	"errors"
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/object"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {
	var io chan bool = make(chan bool)

	//Start promise and wait result

	if p, err := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return "finished", nil
	}); err == nil {

		p.Then(func(i interface{}) *Promise {

			if i.(string) != "finished" {
				t.Errorf("Invalid receive data")
			}
			io <- true
			return nil
		}, func(e error) {

			t.Errorf(e.Error())
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

	//testing chaining promise

	if p, err := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return "finished", nil
	}); err == nil {

		p1, _ := p.Then(func(i interface{}) *Promise {

			if i.(string) != "finished" {
				t.Errorf("Invalid receive data")
			}

			psub, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

				return "psub finished", nil
			})

			return &psub
		}, func(e error) {

			t.Errorf(e.Error())
		})

		p1.Then(func(i interface{}) *Promise {

			if i.(string) != "psub finished" {
				t.Errorf("Invalid receive data %s", i.(string))
			}
			io <- true
			return nil
		}, func(e error) {
			t.Errorf(e.Error())
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

	if p, err := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return nil, errors.New("Problem")
	}); err == nil {

		p.Then(func(i interface{}) *Promise {
			t.Errorf("Must not enter here ")
			return nil
		}, func(e error) {
			io <- true
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}

func TestCatch(t *testing.T) {
	var io chan bool = make(chan bool)

	//Start promise and wait result

	if p, err := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return nil, errors.New("Problem")
	}); err == nil {

		p.Catch(func(e error) {
			io <- true

		})
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}

func TestAll(t *testing.T) {

	var io chan bool = make(chan bool)
	array.GetInterface()

	p1, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return "Hello", nil
	})

	p2, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return "World", nil
	})

	if allp, err := All(p1, p2); err == nil {

		allp.Then(func(i interface{}) *Promise {

			if a, ok := i.(array.ArrayFrom); ok {
				if str, err := a.Array_().ToString(); err == nil {
					if str == "Hello,World" {
						io <- true
					} else {
						t.Errorf("Mistmatch %s", str)
					}
				} else {
					t.Errorf(err.Error())
				}
			} else {
				t.Error("Must be an array")
			}

			return nil
		}, func(e error) {

			t.Error(err.Error())
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

	p3, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return "Hello", nil
	})

	p4, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return nil, errors.New("error")
	})

	if allp, err := All(p3, p4); err == nil {

		allp.Then(func(i interface{}) *Promise {

			t.Error("Must return error")

			return nil
		}, func(e error) {
			io <- true
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}
}

func TestAllSettled(t *testing.T) {

	var io chan bool = make(chan bool)
	array.GetInterface()
	object.GetInterface()
	p1, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return "Hello", nil
	})

	p2, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return nil, errors.New("unknown error")
	})

	if allp, err := AllSettled(p1, p2); err == nil {

		allp.Then(func(i interface{}) *Promise {

			if a, ok := i.(array.ArrayFrom); ok {

				if it, err := a.Array_().Entries(); err == nil {

					for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {

						if obj, ok := value.(object.ObjectFrom); ok {
							if i, ok := index.(int); ok {
								switch i {
								case 0:
									if obj.Object_().GetAttributeString_("status") != "fulfilled" {
										t.Errorf("Status invalid %s", obj.Object_().GetAttributeString_("status"))
									}
									if obj.Object_().GetAttributeString_("value") != "Hello" {
										t.Errorf("Value invalid %s", obj.Object_().GetAttributeString_("value"))
									}
								case 1:
									if obj.Object_().GetAttributeString_("status") != "rejected" {
										t.Errorf("Status invalid %s", obj.Object_().GetAttributeString_("status"))
									}
									if obj.Object_().GetAttributeString_("reason") != "unknown error" {
										t.Errorf("Reason invalid %s", obj.Object_().GetAttributeString_("reason"))
									}
								}

							}
						}
					}

					io <- true
				} else {
					t.Error(err.Error())
				}

			} else {
				t.Error("Must be an array")
			}

			return nil
		}, func(e error) {
			io <- true
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}
}

func TestAny(t *testing.T) {

	var io chan bool = make(chan bool)

	wait, _ := SetTimeout(1000)

	p1w, _ := wait.Then(func(ix interface{}) *Promise {
		p1, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

			return "Hello", nil
		})
		return &p1
	}, nil)

	p2, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

		return "World", nil
	})

	if anyp, err := Any(p1w, p2); err == nil {

		anyp.Then(func(i interface{}) *Promise {

			if i.(string) != "World" {
				t.Error("Must match World")
			}
			io <- true
			return nil
		}, func(e error) {
			t.Error(err.Error())
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}

func TestRace(t *testing.T) {

	var io chan bool = make(chan bool)

	wait500, _ := SetTimeout(500)
	wait100, _ := SetTimeout(100)

	p1w, _ := wait500.Then(func(ix interface{}) *Promise {
		p1, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

			return "one", nil
		})
		return &p1
	}, nil)

	p2w, _ := wait100.Then(func(ix interface{}) *Promise {
		p2, _ := New(func(resolvefunc, errfunc js.Value) (interface{}, error) {

			return "two", nil
		})

		return &p2
	}, nil)

	if anyp, err := Race(p1w, p2w); err == nil {

		anyp.Then(func(i interface{}) *Promise {

			if i.(string) != "two" {
				t.Error("Must match World")
			}
			io <- true
			return nil
		}, func(e error) {
			t.Error(err.Error())
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(200) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}

func TestFinally(t *testing.T) {

	var io chan bool = make(chan bool)

	wait500, _ := SetTimeout(500)
	wait100, _ := SetTimeout(100)

	if fin, err := All(wait100, wait500); err == nil {

		fin.Finally(func() {

			io <- true
		})

	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(500) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}

func TestReject(t *testing.T) {

	var io chan bool = make(chan bool)

	if p, err := Reject(errors.New("Failed")); err == nil {

		p.Then(nil, func(e error) {

			if e.Error() != "Error: Failed" {
				t.Errorf("Must receive Failed receive %s", e.Error())
			}
			io <- true
		})
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}
func TestResolve(t *testing.T) {

	var io chan bool = make(chan bool)

	if p, err := Resolve(100); err == nil {

		p.Then(func(i interface{}) *Promise {

			if i.(int) != 100 {
				t.Errorf("Must equal to 100 get %d", i.(int))
			}
			io <- true
			return nil

		}, nil)
	} else {
		t.Error(err.Error())
	}

	select {
	case <-io:
	case <-time.After(time.Duration(100) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}
