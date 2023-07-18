package hogosuru

import (
	"reflect"
	"strings"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/nodelist"
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
	"HtmlTableSectionElement": "t", //thead tfoot
	"HtmlTemplateElement":     "template",
	"HtmlTextAreaElement":     "textarea",
	"HtmlTimeElement":         "time",
	"HtmlTitleElement":        "title",
}

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
					taghtml = tag
				}

				//TODO: Separate with : for use h1|h2|h2|h3|h4|h5|h6, td|th, tfoot|theader|tbody

				nodelist, err := q.QuerySelectorAll(taghtml)
				if err != nil {
					return err
				}
				ln := nodelist.Length()
				array := reflect.MakeSlice(f.Type(), 0, ln)
				for i := 0; i < ln; i++ {

					if objjs, err := nodelist.GetIndex(i); err != nil {
						if r, err := baseobject.Discover(objjs); err != nil {
							array = reflect.Append(array, reflect.ValueOf(r))
						}

					}
				}
				if v.Kind() == reflect.Struct {
					f.Set(array)
				}

			}

		} else {

			elem, err := q.QuerySelector(tag)
			if err != nil {
				return err
			}
			if delem, err := elem.Discover(); err != nil {

				if v.Kind() == reflect.Struct {
					f.Set(reflect.ValueOf(delem))
				}
			}
		}
	}

	return nil
}
