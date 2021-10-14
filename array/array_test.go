package array

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
)

func TestNewEmpty(t *testing.T) {

	var err error

	var a Array
	var len int

	if a, err = NewEmpty(6); err == nil {
		if len, err = a.Length(); err == nil {
			if len != 6 {
				t.Errorf("Size mismatch")
			}
		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf(err.Error())
	}

}

func TestFrom(t *testing.T) {

	var err error

	var a Array

	if a, err = From("test"); err == nil {
		var str string
		if str, err = a.ToString(); err == nil {
			if str != "t,e,s,t" {
				t.Errorf("not match %s", str)
			}

		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf(err.Error())
	}
	//with mapping
	if a, err = From(New_(1, 2, 3, 4), func(i interface{}) interface{} {
		if vi, ok := i.(int); ok {

			return vi * 3
		}
		return i
	}); err == nil {
		var str string
		if str, err = a.ToString(); err == nil {
			if str != "3,6,9,12" {
				t.Errorf("not match %s", str)
			}

		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf(err.Error())
	}

}

func TestNewFromJSObject(t *testing.T) {
	var err error
	var obj js.Value
	var a Array

	baseobject.Eval("customarray=new Array(1,2,5)")
	if obj, err = js.Global().GetWithErr("customarray"); err == nil {

		if a, err = NewFromJSObject(obj); err == nil {
			var str string
			if str, err = a.ToString(); err == nil {
				if str != "1,2,5" {
					t.Errorf("not match %s", str)
				}

			} else {
				t.Errorf(err.Error())
			}
		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf(err.Error())
	}
}

func TestConcat(t *testing.T) {

	var a Array
	var err error

	if a, err = New(1, 2, 3); err == nil {
		if c, err := a.Concat(New_(6, 7, 8)); err == nil {
			var str string
			if str, err = c.ToString(); err == nil {
				if str != "1,2,3,6,7,8" {
					t.Errorf("not match %s", str)
				}

			} else {
				t.Errorf(err.Error())
			}
		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}
}

func TestCopyWithin(t *testing.T) {

	var a Array
	var err error

	if a, err = New("a", "b", "c", "d", "e"); err == nil {
		if c, err := a.CopyWithin(0, 3, 4); err == nil {
			var str string
			if str, err = c.ToString(); err == nil {
				if str != "d,b,c,d,e" {
					t.Errorf("not match %s", str)
				}

			} else {
				t.Errorf(err.Error())
			}
		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf(err.Error())
	}

}

func TestEntries(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{"a", "b", "c", "d", "e"}

	if a, err = New(goArray...); err == nil {

		if it, err := a.Entries(); err == nil {
			var loop int
			for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {

				if str, ok := value.(string); ok {
					if i, ok := index.(int); ok {

						if str != goArray[i] {
							t.Errorf("content not match %s", str)
						}
					} else {
						t.Errorf("Index is not int")
					}

				} else {
					t.Errorf("Value is not string")
				}

				loop++

			}
			if loop != len(goArray) {
				t.Errorf("Loop entries not match")
			}

		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestEvery(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{1, 2, 3, 4, 5}

	if a, err = New(goArray...); err == nil {

		if b, _ := a.Every(func(i interface{}) bool {

			if i.(int) < 13 {
				return true
			}

			return false
		}); !b {
			t.Errorf("number must be < 13")
		}

		if b, _ := a.Every(func(i interface{}) bool {

			if i.(int) < 2 {
				return true
			}

			return false
		}); b {
			t.Errorf("some number must be >  2")
		}
	}

}

func TestFill(t *testing.T) {
	var a Array
	var err error

	if a, err = NewEmpty(5); err == nil {

		if err := a.Fill(7); err == nil {

			if ok, _ := a.Every(func(i interface{}) bool {

				if i.(int) == 7 {
					return true
				}

				return false
			}); !ok {
				t.Errorf("must be fill with 7")
			}

			if ok, _ := a.Every(func(i interface{}) bool {

				if i.(int) != 7 {
					return true
				}

				return false
			}); ok {
				t.Errorf("must be fill with 7")
			}

		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}
}

func TestFilter(t *testing.T) {
	var a Array
	var err error

	if a, err = New("spray", "limit", "elite", "exuberant", "destruction", "present"); err == nil {

		//select all len word > 7
		if b, err := a.Filter(func(i interface{}) bool {
			if len(i.(string)) > 7 {
				return true
			}
			return false
		}); err == nil {
			if str, err := b.ToString(); err == nil {
				if str != "exuberant,destruction" {
					t.Errorf("Mistmatch")
				}
			} else {
				t.Errorf(err.Error())
			}
		} else {
			t.Errorf(err.Error())
		}

	}

}

func TestFind(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{5, 8, 12, 130, 44}
	if a, err = New(goArray...); err == nil {
		if found, err := a.Find(func(i interface{}) bool {

			if i.(int) > 10 {

				return true
			}

			return false

		}); err == nil {

			if found != nil {
				if found.(int) != 12 {
					t.Errorf("Value mismatch")
				}
			} else {
				t.Errorf("no element found")
			}

		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestFindIndex(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{5, 8, 12, 130, 44}
	if a, err = New(goArray...); err == nil {
		if found, err := a.FindIndex(func(i interface{}) bool {

			if i.(int) == 12 {

				return true
			}

			return false

		}); err == nil {
			if found >= 0 {
				if found != 2 {
					t.Errorf("Value mismatch %d", found)
				}
			} else {
				t.Errorf("no element found")
			}

		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestFlat(t *testing.T) {

	var a Array
	var err error

	var goArray []interface{} = []interface{}{1, 2, []interface{}{3, 4}}
	if a, err = New(goArray...); err == nil {

		if b, err := a.Flat(); err == nil {

			if str, err := b.ToString(); err == nil {
				if str != "1,2,3,4" {
					t.Errorf("Mistmatch %s", str)
				}
			} else {
				t.Errorf(err.Error())
			}
		}

	}

}

func TestFlatMap(t *testing.T) {

	var goArray []interface{} = []interface{}{1, 2, 3, 4}

	if a, err := New(goArray...); err == nil {

		if b, err := a.FlatMap(func(i1 interface{}, i2 int) interface{} {

			b1 := Of_(i1.(int) * 2)
			return b1.JSObject()

		}); err == nil {

			if str, err := b.ToString(); err == nil {
				if str != "2,4,6,8" {
					t.Errorf("Mistmatch %s", str)
				}
			} else {
				t.Errorf(err.Error())
			}
		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestForEach(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	if a, err = New(goArray...); err == nil {

		var count int = 0
		a.ForEach(func(i interface{}) {

			if i.(string) != goArray[count].(string) {
				t.Errorf("Mistmatch value %s", goArray[count])
			}
			count++

		})

		if count != len(goArray) {
			t.Errorf("Bad number of element")
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestIncludes(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	if a, err = New(goArray...); err == nil {

		if ok, err := a.Includes("limit"); err == nil {
			if !ok {

				t.Errorf("Must include limit")

			}
		} else {
			t.Errorf(err.Error())
		}

		if ok, err := a.Includes("limit2"); err == nil {
			if ok {

				t.Errorf("Must not include limit")

			}
		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestIndexOf(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	if a, err = New(goArray...); err == nil {

		obj := a.JSObject().Index(2)
		b, _ := baseobject.NewFromJSObject(obj)

		if index, err := a.IndexOf(b); err == nil {

			if index != 2 {
				t.Errorf("index must be 2 have %d when searching %s", index, obj.String())
			}

		} else {
			t.Errorf(err.Error())
		}

		if index, err := a.IndexOf("elite"); err == nil {

			if index != 2 {
				t.Errorf("index must be 2 have %d", index)
			}

		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestIsArray(t *testing.T) {
	var a Array
	var err error

	if a, err = NewEmpty(3); err == nil {
		if ok, err := IsArray(a.BaseObject); err == nil {
			if !ok {
				t.Errorf("Must be an array")
			}
		} else {
			t.Errorf(err.Error())
		}

		if ok, err := IsArray(baseobject.BaseObject{}); err == nil {
			if ok {
				t.Errorf("Must not be an array")
			}
		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestJoin(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Hello", "World", "elite"}
	if a, err = New(goArray...); err == nil {
		if str, err := a.Join("|"); err == nil {
			if str != "Hello|World|elite" {
				t.Errorf("Mistmatch %s", str)
			}

		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestKeys(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Hello", "World", "elite"}
	if a, err = New(goArray...); err == nil {
		var i int = 0
		if it, err := a.Keys(); err == nil {
			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

				if value.(int) != i {
					t.Errorf("not match index %d", value.(int))
				}
				i++

			}
		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestLastIndexOf(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present", "limit"}
	if a, err = New(goArray...); err == nil {

		obj := a.JSObject().Index(6)
		b, _ := baseobject.NewFromJSObject(obj)

		if index, err := a.LastIndexOf(b); err == nil {

			if index != 6 {
				t.Errorf("index must be 6 have %d when searching %s", index, obj.String())
			}

		} else {
			t.Errorf(err.Error())
		}

		if index, err := a.LastIndexOf("limit"); err == nil {

			if index != 6 {
				t.Errorf("index must be 6 have %d", index)
			}

		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}

}

func TestMap(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{1, 2, 3, 4}
	if a, err = New(goArray...); err == nil {
		if b, err := a.Map(func(i interface{}) interface{} {
			if vi, ok := i.(int); ok {

				return vi * 3
			}
			return i
		}); err == nil {
			if str, err := b.ToString(); err == nil {
				if str != "3,6,9,12" {
					t.Errorf("Mistmatch %s", str)
				}
			} else {
				t.Errorf(err.Error())
			}
		} else {
			t.Errorf(err.Error())
		}

	} else {
		t.Errorf(err.Error())
	}
}
func TestMain(m *testing.M) {
	m.Run()
}
