package domrectlist

import "github.com/realPy/hogosuru/base/domrect"

func (d DOMRectList) Item_(index int) domrect.DOMRect {
	domrect, _ := d.Item(index)
	return domrect
}
