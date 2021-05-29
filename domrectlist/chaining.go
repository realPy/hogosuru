package domrectlist

import "github.com/realPy/hogosuru/domrect"

func (d DOMRectList) Item_(index int) domrect.DOMRect {
	domrect, _ := d.Item(index)
	return domrect
}
