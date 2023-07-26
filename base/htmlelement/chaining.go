package htmlelement

import "github.com/realPy/hogosuru/base/cssstyledeclaration"

func (h HtmlElement) Style_() cssstyledeclaration.CSSStyleDeclaration {
	style, _ := h.Style()
	return style
}
