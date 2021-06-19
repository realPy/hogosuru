package array

import "syscall/js"

func New_(values ...interface{}) Array {
	a, _ := New(values...)
	return a
}

func Of_(values ...interface{}) Array {
	return New_(values...)
}

func NewEmpty_(size int) Array {
	a, _ := NewEmpty(size)
	return a
}
func From_(iterable interface{}, f ...func(interface{}) interface{}) Array {
	a, _ := From(iterable, f...)
	return a
}

func NewFromJSObject_(obj js.Value) Array {
	a, _ := NewFromJSObject(obj)
	return a
}
func (a Array) Concat_(a2 Array) Array {
	arr, _ := a.Concat(a2)
	return arr
}

func (a Array) CopyWithin_(cible int, opts ...int) Array {
	arr, _ := a.CopyWithin(cible, opts...)
	return arr

}

func (a Array) Filter_(f func(interface{}) bool) Array {
	arr, _ := a.Filter(f)
	return arr
}

func (a Array) Flat_(opts ...int) Array {
	arr, _ := a.Flat(opts...)
	return arr
}

func (a Array) FlatMap_(f func(interface{}, int) interface{}) Array {
	arr, _ := a.FlatMap(f)
	return arr
}

func (a Array) Map_(f func(interface{}) interface{}) Array {
	arr, _ := a.Map(f)
	return arr
}

func (a Array) Slice_(opts ...int) Array {
	arr, _ := a.Slice(opts...)
	return arr
}
