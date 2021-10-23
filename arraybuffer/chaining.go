package arraybuffer

func (a ArrayBuffer) ByteLength_() int64 {
	length, err := a.ByteLength()

	if err != nil {
		a.Debug(err.Error())
	}
	return length
}
