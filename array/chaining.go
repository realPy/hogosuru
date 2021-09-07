package array

import "syscall/js"

func New_(values ...interface{}) Array {
	a, err := New(values...)

	if err != nil {
		a.Debug(err.Error())
	}
	return a
}

func Of_(values ...interface{}) Array {
	return New_(values...)
}

func NewEmpty_(size int) Array {
	a, err := NewEmpty(size)
	if err != nil {
		a.Debug(err.Error())
	}
	return a
}
func From_(iterable interface{}, f ...func(interface{}) interface{}) Array {

	a, err := From(iterable, f...)
	if err != nil {
		a.Debug(err.Error())
	}
	return a
}

func NewFromJSObject_(obj js.Value) Array {
	a, err := NewFromJSObject(obj)
	if err != nil {
		a.Debug(err.Error())
	}
	return a
}
func (a Array) Concat_(a2 Array) Array {

	arr, err := a.Concat(a2)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) CopyWithin_(cible int, opts ...int) Array {
	arr, err := a.CopyWithin(cible, opts...)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr

}

func (a Array) Filter_(f func(interface{}) bool) Array {
	arr, err := a.Filter(f)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) Flat_(opts ...int) Array {
	arr, err := a.Flat(opts...)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) FlatMap_(f func(interface{}, int) interface{}) Array {
	arr, err := a.FlatMap(f)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) Map_(f func(interface{}) interface{}) Array {
	arr, err := a.Map(f)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}

func (a Array) Slice_(opts ...int) Array {
	arr, err := a.Slice(opts...)
	if err != nil {
		a.Debug(err.Error())
	}
	return arr
}
