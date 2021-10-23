package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlanchorelement"
	"github.com/realPy/hogosuru/htmlbrelement"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/htmldatalistelement"
	"github.com/realPy/hogosuru/htmldetailselement"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmldlistelement"
	"github.com/realPy/hogosuru/htmlembedelement"
	"github.com/realPy/hogosuru/htmlfieldsetelement"
	"github.com/realPy/hogosuru/htmlformelement"
	"github.com/realPy/hogosuru/htmlheadingelement"
	"github.com/realPy/hogosuru/htmliframelement"
	"github.com/realPy/hogosuru/htmlimageelement"
	"github.com/realPy/hogosuru/htmlinputelement"
	"github.com/realPy/hogosuru/htmllabelelement"
	"github.com/realPy/hogosuru/htmllegendelement"
	"github.com/realPy/hogosuru/htmlmeterelement"
	"github.com/realPy/hogosuru/htmlprogresselement"
	"github.com/realPy/hogosuru/htmlscriptelement"
	"github.com/realPy/hogosuru/htmlselectelement"
	"github.com/realPy/hogosuru/htmltableelement"
	htmltemplatelement "github.com/realPy/hogosuru/htmltemplateelement"
	"github.com/realPy/hogosuru/htmltextareaelement"
	"github.com/realPy/hogosuru/promise"
)

func main() {
	hogosuru.Init()
	d := document.New_()

	nod := d.Body_()

	if text, err := nod.TextContent(); err == nil {
		nod.Export("pou")
		println("<--" + text + "-->")
	}

	if elem, err := d.CreateElement("b"); err == nil {

		if t, err := d.CreateTextNode("Hello"); err == nil {

			elem.AppendChild(t)
			elem.Export("manu")
		} else {
			println(err.Error())
		}

		nod.AppendChild(elem.Node)
	} else {
		println(err.Error())
	}

	if elem, err := d.CreateElement("p"); err == nil {

		elem.SetInnerHTML("<b>World</b>")
		nod.AppendChild(elem.Node)
	} else {
		println(err.Error())
	}

	nodelist := d.QuerySelectorAll_(".pictureContainer")
	println("Found", nodelist.Length(), "elements")
	nodelist.Item_(0).Export("node1")
	/*
		d.AddEventListener("mousemove", func(e event.Event) {
			println("mouse move", e.JSObject().Get("clientX").String(), e.JSObject().Get("clientY").String())
		})
	*/

	if clickbutton, err := d.GetElementById("clickme"); err == nil {

		clickbutton.OnClick(func(e event.Event) {

			if testinput, err := d.GetElementById("test"); err == nil {
				attributes, _ := testinput.Attributes()

				if attr, err := attributes.GetNamedItem("type"); err == nil {
					if str, err := attr.Value(); err == nil {
						println("type->" + str)
					}

				} else {
					println("erreur" + err.Error())
				}

				//easy method

				if str, err := testinput.GetAttribute("type"); err == nil {
					println("Second method type->" + str)
				} else {
					println("erreur" + err.Error())
				}

			}

		})
	}

	p, _ := d.CreateElement("input")
	p.SetAttribute("type", "checkbox")

	h, _ := htmlinputelement.NewFromElement(p)
	h.SetChecked(true)
	nod.AppendChild(h.Node)
	h.Focus()
	h.SetDataset("toto", "value")

	v, _ := h.Dataset("toto")
	println(v.(string))

	input, _ := htmlinputelement.New(d)
	input.SetAttribute("type", "checkbox")
	nod.AppendChild(input.Node)
	//h.SetHidden(true)
	h.Export("mat")

	progress, _ := htmlprogresselement.New(d)
	progress.SetMax(100)
	progress.SetValue(50)

	nod.AppendChild(progress.Node)

	if anchor, err := htmlanchorelement.New(d); err == nil {
		anchor.SetHref("https://google.fr")

		anchor.SetText("Cliquez ici")
		anchor.SetAttribute("info", "color:green")
		anchor.Export("poo")
		anchor.Style_().SetProperty("color", "green")
		anchor.Style_().SetProperty("font-weight", "bold")

		nod.AppendChild(anchor.Node)
	} else {
		println("erreur", err.Error())
	}

	if br, err := htmlbrelement.New(d); err == nil {

		br.SetDataset("test", "test")
		nod.AppendChild(br.Node)
	} else {
		println("erreur", err.Error())
	}

	if form, err := htmlformelement.New(d); err == nil {
		form.SetID("pouet")
		nod.AppendChild(form.Node)

		if b1, err := htmlbuttonelement.New(d); err == nil {
			b1.SetName("submit")
			b1.SetTextContent("submit")
			form.AppendChild(b1.Node)
		} else {
			println("erreur", err.Error())
		}

	} else {
		println("erreur", err.Error())
	}

	inputlist, _ := htmlinputelement.New(d)
	inputlist.SetAttribute("list", "ice-cream-flavors")
	inputlist.SetTextContent("ddd")
	nod.AppendChild(inputlist.Node)
	if datalist, err := htmldatalistelement.New(d); err == nil {

		datalist.SetID("ice-cream-flavors")
		opt, _ := d.CreateElement("option")
		opt.SetAttribute("value", "chocolate")
		datalist.AppendChild(opt.Node)
		opt, _ = d.CreateElement("option")
		opt.SetAttribute("value", "coconut")
		datalist.AppendChild(opt.Node)
		opt, _ = d.CreateElement("option")
		opt.SetAttribute("value", "vanilla")
		datalist.AppendChild(opt.Node)

		nod.AppendChild(datalist.Node)
	} else {
		println("erreur", err.Error())
	}

	if details, err := htmldetailselement.New(d); err == nil {
		details.SetTextContent("A keyboard.")
		sum, _ := d.CreateElement("summary")
		sum.SetTextContent("I have keys but no doors. I have space but no room. You can enter but can’t leave. What am I?")
		details.AppendChild(sum.Node)
		nod.AppendChild(details.Node)
	} else {
		println("erreur", err.Error())
	}

	if dlist, err := htmldlistelement.New(d); err == nil {
		nod.AppendChild(dlist.Node)

	} else {
		println("erreur", err.Error())
	}

	if embed, err := htmlembedelement.New(d); err == nil {

		embed.SetType("video/webm")
		embed.SetWidth("250")
		embed.SetHeight("200")
		embed.SetSrc("https://www.youtube.com/embed/tgbNymZ7vqY")
		nod.AppendChild(embed.Node)

	} else {
		println("erreur", err.Error())
	}

	if formelem, err := htmlformelement.New(d); err == nil {

		if fieldset, err := htmlfieldsetelement.New(d); err == nil {

			l1, _ := htmllegendelement.New(d)
			l1.SetTextContent("Choose your favorite monster")
			fieldset.AppendChild(l1.Node)
			i1, _ := htmlinputelement.New(d)
			i1.SetType("radio")
			i1.SetName("monster")
			i1.SetID("kraken")
			fieldset.AppendChild(i1.Node)
			label1, _ := htmllabelelement.New(d)
			label1.SetHtmlFor("kraken")
			label1.SetTextContent("Kraken")
			fieldset.AppendChild(label1.Node)
			br1, _ := htmlbrelement.New(d)
			fieldset.AppendChild(br1.Node)

			i2, _ := htmlinputelement.New(d)
			i2.SetType("radio")
			i2.SetName("monster")
			i2.SetID("sasquatch")
			fieldset.AppendChild(i2.Node)
			label2, _ := htmllabelelement.New(d)
			label2.SetHtmlFor("sasquatch")
			label2.SetTextContent("Sasquatch")
			fieldset.AppendChild(label2.Node)
			br2, _ := htmlbrelement.New(d)
			fieldset.AppendChild(br2.Node)

			i3, _ := htmlinputelement.New(d)
			i3.SetType("radio")
			i3.SetName("monster")
			i3.SetID("mothman")
			fieldset.AppendChild(i3.Node)
			label3, _ := htmllabelelement.New(d)
			label3.SetHtmlFor("mothman")
			label3.SetTextContent("Mothman")
			fieldset.AppendChild(label3.Node)
			br3, _ := htmlbrelement.New(d)
			fieldset.AppendChild(br3.Node)

			formelem.AppendChild(fieldset.Node)

		} else {
			println("erreur", err.Error())
		}

		nod.AppendChild(formelem.Node)

	} else {
		println("erreur", err.Error())
	}

	if h1, err := htmlheadingelement.NewH1(d); err == nil {
		h1.SetTextContent("The title")
		nod.AppendChild(h1.Node)

	} else {
		println("erreur", err.Error())
	}

	template, _ := htmltemplatelement.New(d)
	nod.AppendChild(template.Node)

	if img, err := htmlimageelement.New(d); err == nil {

		img.SetSrc("https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/500px-Image_created_with_a_mobile_phone.png")
		img.SetHidden(true)
		if p, err := img.Decode(); err == nil {
			p.Then(func(obj interface{}) *promise.Promise {

				//nod.AppendChild(img.Node)
				p, _ := template.ParentNode()
				p.ReplaceChild(img.Node, template.Node)

				img.SetHidden(false)
				return nil

			}, func(e error) {
				println("erreur", e.Error())
			})
		} else {
			println("erreur", err.Error())
		}

	} else {
		println("erreur", err.Error())
	}

	br, _ := htmlbrelement.New(d)
	nod.AppendChild(br.Node)
	if iframe, err := htmliframelement.New(d); err == nil {
		iframe.SetSrcdoc("<!DOCTYPE html><p>Hello World!</p>")
		iframe.SetWidth("25%")
		iframe.SetHeight("50")

		nod.AppendChild(iframe.Node)

	} else {
		println("erreur", err.Error())
	}
	br, _ = htmlbrelement.New(d)
	nod.AppendChild(br.Node)
	if meter, err := htmlmeterelement.New(d); err == nil {
		meter.SetMin(0)
		meter.SetMax(100)
		meter.SetLow(33)
		meter.SetHigh(75)
		meter.SetValue(50)

		nod.AppendChild(meter.Node)

	} else {
		println("erreur", err.Error())
	}

	buttonLoad, _ := htmlbuttonelement.New(d)
	buttonLoad.SetTextContent("click here to load script")
	buttonLoad.OnClick(func(e event.Event) {

		script, _ := htmlscriptelement.New(d)
		script.SetAsync(true)
		script.SetText("window.alert(\"Hello !\");")

		if element, err := d.Head(); err == nil {

			element.AppendChild(script.Node)
		} else {
			println("erreur", err.Error())
		}

	})

	nod.AppendChild(buttonLoad.Node)

	div, _ := htmldivelement.New(d)
	content := `<label for="pet-select">Choose a pet:</label>
	<select name="pets" id="pet-select">
	<option value="">--Please choose an option--</option>
	<option value="dog">Dog</option>
	<option value="cat">Cat</option>
	<option value="hamster">Hamster</option>
	<option value="parrot">Parrot</option>
	<option value="spider">Spider</option>
	<option value="goldfish">Goldfish</option>
</select>`

	nod.AppendChild(div.Node)
	div.SetOuterHTML(content)

	selectelem, _ := d.GetElementById("pet-select")

	selecth, _ := htmlselectelement.NewFromElement(selectelem)

	selecth.OnInput(func(e event.Event) {

		println("selected!!")
	})

	table, _ := htmltableelement.New(d)

	caption, _ := table.CreateCaption()
	tfoot, _ := table.CreateTFoot()
	tfoot.SetTextContent("foot")
	caption.SetTextContent("hello")
	row, _ := table.InsertRow()
	cell, _ := row.InsertCell()
	cell.SetTextContent("pouet")

	nod.AppendChild(table.Node)

	textarea, _ := htmltextareaelement.New(d)
	nod.AppendChild(textarea.Node)

	ch := make(chan struct{})
	<-ch

}
