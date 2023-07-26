package htmlstruct

import (
	"testing"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/htmlbuttonelement"
	"github.com/realPy/hogosuru/base/htmldivelement"
	"github.com/realPy/hogosuru/base/htmlheadingelement"
	"github.com/realPy/hogosuru/base/htmltemplateelement"
	"github.com/realPy/hogosuru/testingutils"
)

const htmlContent = `Hello world
<div id="div1" class="toto">div1classtoto</div>
<div id="searchbyid">divs</div>
<div class="toto"></div>
<button id="press">Press</button>
<h1>H1 1</h1>
<h1>H1 2</h1>
<h2>H2 1</h2>
<button class="inner">Button by class</button>`

type MyMainWindow struct {
	FirstDiv         htmldivelement.HtmlDivElement           `hogosuru:"div:nth-of-type(1)"`
	DivSearchbyid    htmldivelement.HtmlDivElement           `hogosuru:"#searchbyid"`
	Divs             []htmldivelement.HtmlDivElement         `hogosuru:"[]"`
	H1s              []htmlheadingelement.HtmlHeadingElement `hogosuru:"[]:1"`
	H2s              []htmlheadingelement.HtmlHeadingElement `hogosuru:"[]:2"`
	Button           htmlbuttonelement.HtmlButtonElement     `hogosuru:"#press"`
	ButtonClassInner htmlbuttonelement.HtmlButtonElement     `hogosuru:"button.inner"`
	Divtoto          []htmldivelement.HtmlDivElement         `hogosuru:"[]div.toto"`
}

func TestMain(m *testing.M) {

	baseobject.SetSyscall()
	hogosuru.Init()
	m.Run()
}

func TestUnmarshal(t *testing.T) {
	var w MyMainWindow

	doc, err := document.New()

	if err != nil {
		panic(err)
	}
	// using fragment instead document is a workaround for unit test only
	// wasmbrowsertest stop to work if the main document is internally modify
	tpl, _ := htmltemplateelement.New(doc)
	tpl.SetInnerHTML(htmlContent)
	fragment, _ := tpl.Content()
	Unmarshal(fragment, &w)
	testingutils.AssertExpect(t, "div1", w.FirstDiv.ID_())
	testingutils.AssertExpect(t, "searchbyid", w.DivSearchbyid.ID_())
	testingutils.AssertExpect(t, 3, len(w.Divs))
	testingutils.AssertExpect(t, "H1 1", w.H1s[0].TextContent_())
	testingutils.AssertExpect(t, "H2 1", w.H2s[0].TextContent_())
	testingutils.AssertExpect(t, "press", w.Button.ID_())
	testingutils.AssertExpect(t, "Button by class", w.ButtonClassInner.TextContent_())
	testingutils.AssertExpect(t, 2, len(w.Divtoto))
}
