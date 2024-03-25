package stream

import "github.com/realPy/hogosuru/base/baseobject"

type TransformStream struct {
	baseobject.BaseObject
}

func TransfertToTransformStream(b baseobject.BaseObject) TransformStream {
	var t TransformStream
	t.BaseObject = t.SetObject(b.JSObject())
	return t
}
