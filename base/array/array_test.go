package array

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestNewEmpty(t *testing.T) {

	var err error

	var a Array
	var len int

	if a, err = NewEmpty(6); testingutils.AssertErr(t, err) {
		if len, err = a.Length(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 6, len)

		}
	}

}

func TestFrom(t *testing.T) {

	var err error

	var a Array
	t.Run("From string", func(t *testing.T) {
		if a, err = From("test"); testingutils.AssertErr(t, err) {
			var str string
			if str, err = a.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "t,e,s,t", str)
			}
		}
	})
	t.Run("From array", func(t *testing.T) {
		if a, err = From(New_(1, 2, 3, 4), func(i interface{}) interface{} {
			if vi, ok := i.(int); ok {

				return vi * 3
			}
			return i
		}); testingutils.AssertErr(t, err) {
			var str string
			if str, err = a.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "3,6,9,12", str)

			}
		}

	})

}

func TestNewFromJSObject(t *testing.T) {
	var err error
	var obj js.Value
	var a Array

	baseobject.Eval("customarray=new Array(1,2,5)")
	if obj, err = baseobject.Get(js.Global(), "customarray"); testingutils.AssertErr(t, err) {
		if a, err = NewFromJSObject(obj); testingutils.AssertErr(t, err) {
			var str string
			if str, err = a.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "1,2,5", str)
			}
		}
	}
}

func TestConcat(t *testing.T) {

	var a Array
	var err error

	if a, err = New(1, 2, 3); testingutils.AssertErr(t, err) {
		if c, err := a.Concat(New_(6, 7, 8)); testingutils.AssertErr(t, err) {
			var str string
			if str, err = c.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "1,2,3,6,7,8", str)
			}
		}
	}
}

func TestCopyWithin(t *testing.T) {

	var a Array
	var err error

	if a, err = New("a", "b", "c", "d", "e"); testingutils.AssertErr(t, err) {
		if c, err := a.CopyWithin(0, 3, 4); testingutils.AssertErr(t, err) {
			var str string
			if str, err = c.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "d,b,c,d,e", str)

			}
		}
	}

}

func TestEntries(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{"a", "b", "c", "d", "e"}

	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if it, err := a.Entries(); testingutils.AssertErr(t, err) {
			var loop int
			for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {

				if str, ok := value.(string); ok {
					if i, ok := index.(int); ok {

						if str != goArray[i] {
							testingutils.AssertExpect(t, goArray[i], str)
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

	if a, err = NewEmpty(5); testingutils.AssertErr(t, err) {

		if err := a.Fill(7); testingutils.AssertErr(t, err) {

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

		}
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
		}); testingutils.AssertErr(t, err) {
			if str, err := b.ToString(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "exuberant,destruction", str)

			}
		}

	}

}

func TestFind(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{5, 8, 12, 130, 44}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {
		if found, err := a.Find(func(i interface{}) bool {

			if i.(int) > 10 {

				return true
			}

			return false

		}); testingutils.AssertErr(t, err) {

			if found != nil {
				testingutils.AssertExpect(t, 12, found)
			} else {
				t.Errorf("no element found")
			}

		}

	}

}

func TestFindIndex(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{5, 8, 12, 130, 44}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {
		if found, err := a.FindIndex(func(i interface{}) bool {

			if i.(int) == 12 {

				return true
			}

			return false

		}); testingutils.AssertErr(t, err) {
			if found >= 0 {
				testingutils.AssertExpect(t, 2, found)

			} else {
				t.Errorf("no element found")
			}

		}

	}

}

func TestFlat(t *testing.T) {

	var a Array
	var err error

	var goArray []interface{} = []interface{}{1, 2, []interface{}{3, 4}}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if b, err := a.Flat(); testingutils.AssertErr(t, err) {

			if str, err := b.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "1,2,3,4", str)

			}
		}

	}

}

func TestFlatMap(t *testing.T) {

	var goArray []interface{} = []interface{}{1, 2, 3, 4}

	if a, err := New(goArray...); testingutils.AssertErr(t, err) {

		if b, err := a.FlatMap(func(i1 interface{}, i2 int) interface{} {

			b1 := Of_(i1.(int) * 2)
			return b1.JSObject()

		}); testingutils.AssertErr(t, err) {

			if str, err := b.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "2,4,6,8", str)
			}
		}

	}

}

func TestForEach(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		var count int = 0
		a.ForEach(func(i interface{}) {

			testingutils.AssertExpect(t, goArray[count], i)
			count++

		})

		if count != len(goArray) {
			t.Errorf("Bad number of element")
		}

	}

}

func TestIncludes(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if ok, err := a.Includes("limit"); testingutils.AssertErr(t, err) {
			if !ok {

				t.Errorf("Must include limit")

			}
		}

		if ok, err := a.Includes("limit2"); testingutils.AssertErr(t, err) {
			if ok {

				t.Errorf("Must not include limit")

			}
		}

	}

}

func TestIndexOf(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		obj := a.JSObject().Index(2)
		b, _ := baseobject.NewFromJSObject(obj)

		if index, err := a.IndexOf(b); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 2, index)

			if index != 2 {
				t.Errorf("index must be 2 have %d when searching %s", index, obj.String())
			}

		}

		if index, err := a.IndexOf("elite"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 2, index)

		}

	}

}

func TestIsArray(t *testing.T) {
	var a Array
	var err error

	if a, err = NewEmpty(3); testingutils.AssertErr(t, err) {
		if ok, err := IsArray(a.BaseObject); testingutils.AssertErr(t, err) {
			if !ok {
				t.Errorf("Must be an array")
			}
		}

		if ok, err := IsArray(baseobject.BaseObject{}); testingutils.AssertErr(t, err) {
			if ok {
				t.Errorf("Must not be an array")
			}
		}

	}

}

func TestJoin(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Hello", "World", "elite"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {
		if str, err := a.Join("|"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "Hello|World|elite", str)

		}

	}

}

func TestKeys(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Hello", "World", "elite"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {
		var i int = 0
		if it, err := a.Keys(); testingutils.AssertErr(t, err) {
			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {

				testingutils.AssertExpect(t, value, i)

				i++

			}
		}

	}
}

func TestLastIndexOf(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"spray", "limit", "elite", "exuberant", "destruction", "present", "limit"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		obj := a.JSObject().Index(6)
		b, _ := baseobject.NewFromJSObject(obj)

		if index, err := a.LastIndexOf(b); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 6, index)
		}

		if index, err := a.LastIndexOf("limit"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 6, index)

		}

	}

}

func TestMap(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{1, 2, 3, 4}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {
		if b, err := a.Map(func(i interface{}) interface{} {
			if vi, ok := i.(int); ok {

				return vi * 3
			}
			return i
		}); testingutils.AssertErr(t, err) {
			if str, err := b.ToString(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "3,6,9,12", str)

			}
		}

	}
}

func TestPop(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"hello"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if err := a.Pop(); testingutils.AssertErr(t, err) {
			if l, _ := a.Length(); l != 0 {

				t.Errorf("Array must be empty now")

			}
		}
	}

}

func TestPush(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"hello"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if length, err := a.Push("world"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, length, 2)

			if str, err := a.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "hello,world", str)

			}

		}
	}

}

func TestReduce(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{1, 2, 3, 4}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if value, err := a.Reduce(func(accumulateur, value interface{}, opts ...interface{}) interface{} {
			val := accumulateur.(int) + value.(int)

			return val
		}); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 10, value)

		}

	}
}

func TestReduceRight(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{9, 6, 8, 40}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if value, err := a.ReduceRight(func(accumulateur, value interface{}, opts ...interface{}) interface{} {
			val := accumulateur.(int) - value.(int)

			return val
		}); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 17, value)

		}

	}
}

func TestReverse(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{9, 6, 8, 40}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if err := a.Reverse(); testingutils.AssertErr(t, err) {
			if str, err := a.ToString(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "40,8,6,9", str)

			}

		}
	}
}

func TestShift(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{9, 6, 8, 40}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if v, err := a.Shift(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 9, v)

			if str, err := a.ToString(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "6,8,40", str)

			}
		}

	}
}

func TestSlice(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{"ant", "bison", "camel", "duck", "elephant"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if v, err := a.Slice(2); testingutils.AssertErr(t, err) {

			if str, err := v.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "camel,duck,elephant", str)
			}

		}

		if v, err := a.Slice(2, 4); testingutils.AssertErr(t, err) {

			if str, err := v.ToString(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "camel,duck", str)
			}

		}
		if v, err := a.Slice(-2); testingutils.AssertErr(t, err) {

			if str, err := v.ToString(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "duck,elephant", str)

			}

		}

	}
}

func TestSome(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{9, 6, 8, 40}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if ok, err := a.Some(func(i interface{}) bool {

			if i.(int) == 40 {
				return true
			}

			return false
		}); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, true, ok)

		}

		if ok, err := a.Some(func(i interface{}) bool {

			if i.(int) == 42 {
				return true
			}

			return false
		}); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, false, ok)

		}

	}
}

func TestSort(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{"March", "Jan", "Feb", "Dec"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if err := a.Sort(); testingutils.AssertErr(t, err) {
			if str, err := a.ToString(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "Dec,Feb,Jan,March", str)

			}
		}
	}
}

func TestSplice(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Jan", "March", "April", "June"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if err := a.Splice(1, 0, "Feb"); testingutils.AssertErr(t, err) {
			if str, err := a.ToString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "Jan,Feb,March,April,June", str)
			}
		}

		if err := a.Splice(4, 1, "May"); testingutils.AssertErr(t, err) {
			if str, err := a.ToString(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "Jan,Feb,March,April,May", str)

			}
		}
	}
}

func TestUnshift(t *testing.T) {

	var a Array
	var err error
	var goArray []interface{} = []interface{}{1, 2, 3}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {

		if l, err := a.Unshift(4, 5); testingutils.AssertErr(t, err) {

			if str, err := a.ToString(); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "4,5,1,2,3", str)

			}
			testingutils.AssertExpect(t, 5, l)

		}

	}
}

func TestValues(t *testing.T) {
	var a Array
	var err error
	var goArray []interface{} = []interface{}{"Hello", "World", "elite"}
	if a, err = New(goArray...); testingutils.AssertErr(t, err) {
		var i int = 0
		if it, err := a.Values(); testingutils.AssertErr(t, err) {
			for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {
				testingutils.AssertExpect(t, goArray[i], value)
				i++

			}
		}

	}

}

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}
