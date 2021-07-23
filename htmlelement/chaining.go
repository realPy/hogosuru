package htmlelement

import "github.com/realPy/hogosuru/cssstyledeclaration"

func (h HtmlElement) Style_() cssstyledeclaration.CSSStyleDeclaration {
	style, _ := h.Style()
	return style
}
