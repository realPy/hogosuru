package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/array"
)

//eval code

func main() {
	hogosuru.Init()
	arr := array.Of_(1, 2, 3, 4)
	arr2 := array.New_(5, 6, 7, 8)
	arrx := array.New_(5)
	arrx.Export("manux")
	arr = arr.Concat_(arr2)

	it, _ := arr.Entries()
	arr3 := array.New_(arr, arr2)
	arr3.Export("arr3")
	arr4 := arr3.Flat_()
	arr4.Export("arr4")

	for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {
		println("<", index, ":", value)

	}
	var compare int = 13
	if b, _ := arr.Every(func(i interface{}) bool {

		if i.(int) < compare {
			return true
		}

		return false
	}); b {
		println("All elements is < ", compare)
	} else {
		println("Some elements are  > ", compare)
	}

	compare = 5
	if b, _ := arr.Every(func(i interface{}) bool {

		if i.(int) < compare {
			return true
		}

		return false
	}); b {
		println("All elements is < ", compare)
	} else {
		println("Some elements are  > ", compare)
	}

	value, _ := arr.Find(func(i interface{}) bool {
		if i.(int) == 31 {
			return true
		}
		return false
	})
	println(value)
	arr.Fill("k")
	arr.Export("arr")

	afrom := array.From_("hello")
	afrom.Export("hello")

	if ok, _ := afrom.Includes("h"); ok {
		println("include h")
	}

	if index, err := afrom.IndexOf("e"); err == nil {
		println("position e:", index)
	}

	afrom.Reverse()

	ato := afrom.Slice_(2, 4)
	ato.Export("hello2")

	it, _ = ato.Values()

	for _, value, err := it.Next(); err == nil; _, value, err = it.Next() {
		println("<", ":", value)
	}

	ch := make(chan struct{})
	<-ch

}
