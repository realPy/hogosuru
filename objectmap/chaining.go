package objectmap

func New_(values ...interface{}) ObjectMap {
	m, _ := New(values...)
	return m
}
