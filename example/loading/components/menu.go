package components

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/documentfragment"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlanchorelement"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmllielement"
	"github.com/realPy/hogosuru/htmlstyleelement"
	htmltemplatelement "github.com/realPy/hogosuru/htmltemplateelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/promise"
)

type Menu struct {
	parentNode   node.Node
	itemTemplate htmltemplatelement.HtmlTemplateElement
	div          htmldivelement.HtmlDivElement
	Items        map[string]string
	selected     htmlanchorelement.HtmlAnchorElement
}

func (m *Menu) RefreshMenu(d document.Document) {

	if df, err := documentfragment.New(); hogosuru.AssertErr(err) {

		df.AppendChild(m.div.Node)
		ulmenu, _ := m.div.QuerySelector("#menu")

		if n, err := df.QuerySelector("#menuitem"); hogosuru.AssertErr(err) {
			for r, err := ulmenu.FirstChild(); err == nil; r, err = ulmenu.FirstChild() {
				ulmenu.RemoveChild(r)
			}

			if p, err := n.Discover(); hogosuru.AssertErr(err) {
				if t, ok := p.(htmltemplatelement.HtmlTemplateElement); ok {
					/*
						for r, err := t.NextSibling(); err == nil; r, err = t.NextSibling() {
							ulmenu.RemoveChild(r)


						}*/

					if fragment, err := t.Content(); hogosuru.AssertErr(err) {
						for key, value := range m.Items {

							if clone, err := d.ImportNode(fragment.Node, true); hogosuru.AssertErr(err) {
								if clonefragment, ok := clone.(documentfragment.DocumentFragment); ok {
									if n, err := clonefragment.QuerySelector("#menudata"); hogosuru.AssertErr(err) {
										if discovernode, err := n.Discover(); hogosuru.AssertErr(err) {
											if a, ok := discovernode.(htmlanchorelement.HtmlAnchorElement); ok {
												a.SetHref(key)

												a.OnClick(func(e event.Event) {
													if liobj, err := m.selected.ParentNode(); hogosuru.AssertErr(err) {
														if discoverli, err := liobj.Discover(); hogosuru.AssertErr(err) {
															if li, ok := discoverli.(htmllielement.HtmlLIElement); ok {
																li.SetClassName("")

															}
														}

													}

													if liobj, err := a.ParentNode(); hogosuru.AssertErr(err) {
														if discoverli, err := liobj.Discover(); hogosuru.AssertErr(err) {
															if li, ok := discoverli.(htmllielement.HtmlLIElement); ok {
																li.SetClassName("selected")

															}
														}

													}

													e.PreventDefault()
													m.selected = a
												})

												a.SetTextContent(value)
												ulmenu.AppendChild(clonefragment.Node)
											}
										}

									}

								}

							}

						}

					}

				}

			}
			/*
				if p, err := n.ParentNode(); hogosuru.AssertErr(err) {
					p.RemoveChild(n)
				}*/
		}

	}

}

func (m *Menu) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	m.parentNode = n

	htmltemplatelement.GetInterface()
	htmlanchorelement.GetInterface()
	htmllielement.GetInterface()

	if style, err := htmlstyleelement.New(d); hogosuru.AssertErr(err) {

		if head, err := d.Head(); hogosuru.AssertErr(err) {
			style.SetTextContent(menucss)
			head.AppendChild(style.Node)

		}

	}

	if menu, err := htmldivelement.New(d); hogosuru.AssertErr(err) {

		menu.SetID("menu-tab")
		menu.SetClassName("tab")
		m.div = menu
		m.div.SetInnerHTML(menuhtml)

	}
	m.RefreshMenu(d)
	return nil, nil
}

func (m *Menu) Node(r hogosuru.Rendering) node.Node {

	return m.div.Node
}

func (m *Menu) OnEndChildRendering(r hogosuru.Rendering) {

}

func (m *Menu) OnEndChildsRendering() {
	m.parentNode.AppendChild(m.div.Node)

}

func (l *Menu) OnUnload() {

}
