package objectmap

func New_(values ...interface{}) ObjectMap {
	m, _ := New(values...)
	return m
}

func (o ObjectMap) Has_(key interface{}) bool {
	has, err := o.Has(key)

	if err != nil {
		o.Debug(err.Error())
	}
	return has
}

func (o ObjectMap) Get_(key interface{}) interface{} {
	get, err := o.Get(key)

	if err != nil {
		o.Debug(err.Error())
	}
	return get
}
