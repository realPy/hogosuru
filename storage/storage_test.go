package storage

import (
	"testing"
	"time"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/window"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	hogosuru.Init()

	m.Run()
}

type People struct {
	Email   string `indexeddb:"store=people,version=1,keypath=email"`
	Name    string `indexeddb:"name"`
	Surname string `indexeddb:"surname"`
}

func TestPutPeople(t *testing.T) {

	var io chan bool = make(chan bool)
	w, err := window.New()
	testingutils.AssertExpect(t, nil, err)
	p, err := Register(w, People{}, "put")
	testingutils.AssertExpect(t, nil, err)
	p.Then(func(i interface{}) *promise.Promise {
		//waiting ready to use
		people := People{Email: "hello@world.com", Name: "Hello", Surname: "World"}

		people1, _ := Add(people, true)

		people1.Then(func(i interface{}) *promise.Promise {
			people.Name = "\\o/"
			put, err := Put(people, true)
			testingutils.AssertExpect(t, nil, err)
			put.Then(func(i interface{}) *promise.Promise {
				peoples := []People{}
				peopleget, err := GetAll(&peoples)
				if err != nil {
					panic(err)
				}

				peopleget.Then(func(i interface{}) *promise.Promise {
					testingutils.AssertExpect(t, 1, len(peoples))
					expected := []People{People{Email: "hello@world.com", Name: "\\o/", Surname: "World"}}
					testingutils.AssertExpect(t, expected, peoples)

					io <- true
					return nil
				}, nil)

				return &peopleget
			}, nil)
			return &put
		}, nil)

		return nil
	}, func(err error) {
		panic(err)
	})

	select {
	case <-io:
	case <-time.After(time.Duration(45000) * time.Millisecond):
		t.Errorf("No message channel receive")
	}

}

func TestDeletePeople(t *testing.T) {

	t.Run("Register and insert 2 people delete 1", func(t *testing.T) {
		var io chan bool = make(chan bool)
		w, err := window.New()
		testingutils.AssertExpect(t, nil, err)
		p, err := Register(w, People{}, "del1")
		testingutils.AssertExpect(t, nil, err)
		p.Then(func(i interface{}) *promise.Promise {
			//waiting ready to use
			people1 := People{Email: "hello@world.com", Name: "Hello", Surname: "World"}
			p1, _ := Add(people1, true)
			people2 := People{Email: "hello2@world.com", Name: "Hello", Surname: "World"}
			p2, _ := Add(people2, true)

			allp, err := promise.All(p1, p2)

			if err != nil {
				panic(err)
			}

			allp.Then(func(i interface{}) *promise.Promise {

				peoples := []People{}

				pdelete, err := Delete(people2, false)
				if err != nil {
					panic(err)
				}

				pdelete.Then(func(i interface{}) *promise.Promise {
					pe, err := GetAll(&peoples)
					if err != nil {
						panic(err)
					}

					pe.Then(func(i interface{}) *promise.Promise {
						testingutils.AssertExpect(t, 1, len(peoples))
						expected := []People{People{Email: "hello@world.com", Name: "Hello", Surname: "World"}}
						testingutils.AssertExpect(t, expected, peoples)

						io <- true
						return nil
					}, nil)
					return &pe
				}, nil)

				return &pdelete
			}, nil)

			return nil
		}, func(err error) {
			panic(err)
		})

		select {
		case <-io:
		case <-time.After(time.Duration(45000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

	t.Run("Register and insert 2 people in add2 with same key and rollback on", func(t *testing.T) {
		var io chan bool = make(chan bool)
		w, err := window.New()
		testingutils.AssertExpect(t, nil, err)
		p, err := Register(w, People{}, "add2")
		testingutils.AssertExpect(t, nil, err)
		p.Then(func(i interface{}) *promise.Promise {
			//waiting ready to use

			p1, _ := Add(People{Email: "hello@world.com", Name: "Hello", Surname: "World"}, true)

			p2, _ := Add(People{Email: "hello@world.com", Name: "Hello", Surname: "World"}, true)

			allp, err := promise.All(p1, p2)

			if err != nil {
				panic(err)
			}

			allp.Then(func(i interface{}) *promise.Promise {

				return nil
			}, func(err error) {
				testingutils.AssertExpect(t, "Key already exists in the object store.", err.Error())
				//transaction is aborted
				io <- true
			})

			return nil
		}, func(err error) {
			panic(err)
		})

		select {
		case <-io:
		case <-time.After(time.Duration(45000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

	t.Run("Register and insert 2 people in add3 with same key and rollback off", func(t *testing.T) {
		var io chan bool = make(chan bool)
		w, err := window.New()
		testingutils.AssertExpect(t, nil, err)
		p, err := Register(w, People{}, "add3")
		testingutils.AssertExpect(t, nil, err)
		p.Then(func(i interface{}) *promise.Promise {
			//waiting ready to use

			p1, _ := Add(People{Email: "hello@world.com", Name: "Hello", Surname: "World"}, false)

			p2, _ := Add(People{Email: "hello@world.com", Name: "Hello", Surname: "World"}, false)

			allp, err := promise.All(p1, p2)

			if err != nil {
				panic(err)
			}

			allp.Then(func(i interface{}) *promise.Promise {

				return nil
			}, func(err error) {
				testingutils.AssertExpect(t, "Key already exists in the object store.", err.Error())
				//transaction is aborted
				peoples := []People{}
				pe, err := GetAll(&peoples)
				if err != nil {
					panic(err)
				}

				pe.Then(func(i interface{}) *promise.Promise {
					testingutils.AssertExpect(t, 1, len(peoples))
					expected := []People{People{Email: "hello@world.com", Name: "Hello", Surname: "World"}}
					testingutils.AssertExpect(t, expected, peoples)

					io <- true
					return nil
				}, nil)

			})

			return nil
		}, func(err error) {
			panic(err)
		})

		select {
		case <-io:
		case <-time.After(time.Duration(45000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}

	})

}
