package date

func New_(values ...interface{}) Date {
	d, err := New(values...)

	if err != nil {
		d.Debug(err.Error())
	}
	return d
}
