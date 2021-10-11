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
	node       node.Node
	div        htmldivelement.HtmlDivElement
}

func (rs *RedSquare) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	components.StartLoader("RedSquare loading")
	rs.parentNode = n

	var p promise.Promise

	p, _ = promise.SetTimeout(1000)

	if rs.div.Empty() {
		if div, err := htmldivelement.New(d); hogosuru.AssertErr(err) {
			div.SetID("redsquare")
			div.Style_().SetProperty("background-color", "white")
			div.Style_().SetProperty("position", "fixed")
			div.Style_().SetProperty("top", "0")
			div.Style_().SetProperty("right", "0")
			div.Style_().SetProperty("bottom", "0")
			div.Style_().SetProperty("left", "0")

			rs.div = div
		}
		return &p, []hogosuru.Rendering{&components.Menu{Items: map[string]string{"#1": "menu 1", "#2": "menu 2", "#3": "Menu 3"}}}
	}

	return nil, nil
}

func (rs *RedSquare) OnEndChildRendering(r hogosuru.Rendering) {

}

func (rs *RedSquare) OnEndChildsRendering() {

	rs.parentNode.AppendChild(rs.div.Node)
	rs.node = rs.div.Node
	components.StopLoader()
}

func (rs *RedSquare) Node(r hogosuru.Rendering) node.Node {

	return rs.div.Node
}

func (rs *RedSquare) OnUnload() {

	p, _ := rs.node.ParentNode()

	p.RemoveChild(rs.node)

}
