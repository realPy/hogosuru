package htmlstruct

import (
	"reflect"
	"strings"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/nodelist"
)

var tagtoelem = map[string]string{

	"HtmlAnchorElement":       "a",
	"HtmlAreaElement":         "area",
	"HtmlBodyElement":         "body",
	"HtmlBRElement":           "br",
	"HtmlButtonElement":       "button",
	"HtmlDataElement":         "data",
	"HtmlDataListElement":     "datalist",
	"HtmlDetailsElement":      "details",
	"HtmlDivElement":          "div",
	"HtmlDListElement":        "dl",
	"HtmlEmbedElement":        "embed",
	"HtmlFieldSetElement":     "fieldset",
	"HtmlFormElement":         "form",
	"HtmlHeadElement":         "head",
	"HtmlHeadingElement":      "h", //h1,h2...
	"HtmlHrElement":           "hr",
	"HtmlIFrameElement":       "iframe",
	"HtmlImageElement":        "image",
	"HtmlInputElement":        "input",
	"HtmlLabelElement":        "label",
	"HtmlLegendElement":       "legend",
	"HtmlLIElement":           "li",
	"HtmlLinkElement":         "link",
	"HtmlMapElement":          "map",
	"HtmlMetaElement":         "meta",
	"HtmlMeterElement":        "meter",
	"HtmlOptionElement":       "option",
	"HtmlParagraphElement":    "p",
	"HtmlPreElement":          "pre",
	"HtmlProgressElement":     "progress",
	"HtmlQuoteElement":        "quote",
	"HtmlScriptElement":       "script",
	"HtmlSelectElement":       "select",
	"HtmlSourceElement":       "source",
	"HtmlSpanElement":         "span",
	"HtmlStyleElement":        "style",
	"HtmlTableCaptionElement": "caption",
	"HtmlTableCellElement":    "t", //th or td
	"HtmlTableElement":        "table",
	"HtmlTableRowElement":     "tr",
	"HtmlTableSectionElement": "t", //thead tfoot tbody
	"HtmlTemplateElement":     "template",
	"HtmlTextAreaElement":     "textarea",
	"HtmlTimeElement":         "time",
	"HtmlTitleElement":        "title",
}

/* Syntax for hogosuru struct tag
    _ htmlbodyelement.HtmlBodyElement                   `hogosuru:"body:nth-of-type(1)"` attach the first body element
	_ htmlbuttonelement.HtmlButtonElement               `hogosuru:"button.innerBox"` attach the button for class innerBox
	_ htmldivelement.HtmlDivElement                     `hogosuru:"#divid"` attach the div with id="divid"
	_ []htmldivelement.HtmlDivElement                   `hogosuru:"[]"` Get all divs
	_ []htmldivelement.HtmlDivElement                   `hogosuru:"[]div.toto"` Get all divs with class toto
    _ []htmlheadingelement.HtmlHeadingElement           `hogosuru:"[]:1"` Get all h1
	_ []htmlheadingelement.HtmlHeadingElement           `hogosuru:"[]:2"` Get all h2
	_ []htmltablesectionelement.HtmlTableSectionElement `hogosuru:"[]:head"` Get all thead
*/

func getTagByType(t string) string {
	v, _ := tagtoelem[t]
	return v
}

type QuerySelector interface {
	QuerySelectorAll(selector string) (nodelist.NodeList, error)
	QuerySelector(selector string) (element.Element, error)
}

func Unmarshal(q QuerySelector, s interface{}) error {
	o := reflect.ValueOf(s)
	v := reflect.Indirect(o)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("hogosuru")
		f := v.FieldByName(field.Name)
		if strings.HasPrefix(tag, "[]") {
			if f.Kind() == reflect.Slice {
				tag = strings.TrimPrefix(tag, "[]")

				taghtml := getTagByType(f.Type().Elem().Name())

				if tag != "" {
					stag := strings.Split(tag, ":")
					if len(stag) > 1 {
						taghtml = taghtml + stag[1]
					} else {
						taghtml = tag
					}

				}

				nodelist, err := q.QuerySelectorAll(taghtml)
				if err != nil {
					return err
				}
				ln := nodelist.Length()
				array := reflect.MakeSlice(f.Type(), 0, ln)
				for i := 0; i < ln; i++ {

					if objjs, err := nodelist.GetIndex(i); err == nil {
						if r, err := baseobject.Discover(objjs); err == nil {
							array = reflect.Append(array, reflect.ValueOf(r))
						}

					}
				}
				if v.Kind() == reflect.Struct {
					if f.CanSet() {
						f.Set(array)
					}
				}

			}

		} else {
			if tag != "" {
				elem, err := q.QuerySelector(tag)
				if err != nil {
					return err
				}
				if delem, err := elem.Discover(); err == nil {

					if v.Kind() == reflect.Struct {
						if f.CanSet() {
							f.Set(reflect.ValueOf(delem))
						}

					}
				}
			}

		}
	}

	return nil
}
