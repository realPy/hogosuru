package components

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/customevent"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmlprogresselement"
	"github.com/realPy/hogosuru/htmlstyleelement"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/object"
	"github.com/realPy/hogosuru/promise"
)

type Loader struct {
	parentNode node.Node
	div        htmldivelement.HtmlDivElement
	divText    htmldivelement.HtmlDivElement
	progress   htmlprogresselement.HtmlProgressElement
}

func (l *Loader) Hide(hide bool) {
	if hide {

		l.div.Style_().SetProperty("display", "none")
	} else {
		l.div.Style_().SetProperty("display", "")
	}

}

func (l *Loader) SetProgressValue(value float64) {
	l.progress.SetValue(value)

}

func (l *Loader) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []hogosuru.Rendering) {

	l.parentNode = n

	//use autodiscover

	htmlprogresselement.GetInterface()
	htmldivelement.GetInterface()
	customevent.GetInterface()
	object.GetInterface()

	if style, err := htmlstyleelement.New(d); hogosuru.AssertErr(err) {

		if head, err := d.Head(); hogosuru.AssertErr(err) {
			style.SetTextContent(loadercss)
			//style.SetInnerHTML(loadercss)
			head.AppendChild(style.Node)

		}

	}

	if loader, err := htmldivelement.New(d); hogosuru.AssertErr(err) {
		loader.SetID("loader-container")
		loader.SetClassName("loader")

		loader.SetInnerHTML(loaderhtml)
		l.div = loader
	}

	return nil, nil
}

func (l *Loader) Node(r hogosuru.Rendering) node.Node {

	return l.div.Node
}

func (l *Loader) OnEndChildRendering(r hogosuru.Rendering) {

}

func (l *Loader) OnEndChildsRendering() {
	l.parentNode.AppendChild(l.div.Node)
	if n, err := l.div.QuerySelector("#loadprogress"); hogosuru.AssertErr(err) {
		if p, err := n.Discover(); hogosuru.AssertErr(err) {

			if e, ok := p.(htmlprogresselement.HtmlProgressElement); ok {
				l.progress = e
			}

		}

	}

	if n, err := l.div.QuerySelector("#loadtext"); hogosuru.AssertErr(err) {
		if p, err := n.Discover(); hogosuru.AssertErr(err) {

			if e, ok := p.(htmldivelement.HtmlDivElement); ok {
				l.divText = e
			}

		}

	}

	if d, err := document.New(); hogosuru.AssertErr(err) {

		d.AddEventListener("hogoloader", func(e event.Event) {

			if obj, err := baseobject.Discover(e.JSObject()); err == nil {
				if c, ok := obj.(customevent.CustomEventFrom); ok {
					if detail, err := c.CustomEvent_().Detail(); hogosuru.AssertErr(err) {

						if objdetail, ok := detail.(object.Object); ok {

							if mapObject, err := objdetail.Map(); hogosuru.AssertErr(err) {

								if mapObject.Has_("type") {

									switch typ := mapObject.Get_("type").(string); typ {
									case "start":
										l.SetProgressValue(100)

										l.divText.SetTextContent(mapObject.Get_("message").(string))
										l.Hide(false)
									case "stop":
										l.Hide(true)
									}

								}
							}

						}

					}

				}
			}

		})

	}

}

func (l *Loader) OnUnload() {

}

func StartLoader(message string) {

	if d, err := document.New(); hogosuru.AssertErr(err) {

		if e, err := customevent.New("hogoloader", map[string]interface{}{"type": "start", "message": message}); hogosuru.AssertErr(err) {
			d.DispatchEvent(e.Event)
		}

	}

}

func StopLoader() {

	if d, err := document.New(); hogosuru.AssertErr(err) {

		if e, err := customevent.New("hogoloader", map[string]interface{}{"type": "stop"}); hogosuru.AssertErr(err) {
			d.DispatchEvent(e.Event)
		}

	}

}
