package loadingview

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/example/loading/components"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type RedSquare struct {
	parentNode node.Node
	div        htmldivelement.HtmlDivElement
}

func (rs *RedSquare) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	components.StartLoader("RedSquare loading")
	rs.parentNode = n
	if div, err := htmldivelement.New(d); hogosuru.AssertErr(err) {
		div.SetID("redsquare")
		div.Style_().SetProperty("background-color", "red")
		div.Style_().SetProperty("position", "fixed")
		div.Style_().SetProperty("top", "0")
		div.Style_().SetProperty("right", "0")
		div.Style_().SetProperty("bottom", "0")
		div.Style_().SetProperty("left", "0")

		rs.div = div
	}

	var p promise.Promise

	p, _ = promise.SetTimeout(func() (interface{}, error) {
		return nil, nil
	}, 3000)

	return &p, nil
}

func (rs *RedSquare) OnEndChildRendering(r hogosuru.Rendering) {

}

func (rs *RedSquare) OnEndChildsRendering(tree node.Node) {
	rs.parentNode.AppendChild(tree)
	components.StopLoader()
}

func (rs *RedSquare) Node() node.Node {

	return rs.div.Node
}

func (rs *RedSquare) OnUnload() {

}
