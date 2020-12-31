package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
)

type component struct {
	Name      string
	Doc       string
	Modifiers []modifier
	Parts     []part
	Elem      string
	Option    bool
}

type modifier struct {
	Name  string
	Class string
	Doc   string
}

type part struct {
	Name string
}

var components = []component{
	{
		Name:   "accordion",
		Doc:    "create a list of items that can be shown individually by clicking an item's header.",
		Elem:   "Ul",
		Option: true,
	},
	{
		Name:      "alert",
		Doc:       "display success, warning and error messages",
		Modifiers: modifierByPrefix("uk-alert"),
		Elem:      "Div",
		Option:    true,
	},
	{
		Name:      "card",
		Doc:       "create layout boxes with different styles",
		Modifiers: modifierByPrefix("uk-card"),
		Elem:      "Div",
	},
	{
		Name:      "container",
		Doc:       "allows you to align and center your page content",
		Modifiers: modifierByPrefix("uk-container"),
		Elem:      "Div",
	},
	{
		Name:      "grid",
		Doc:       "create a fully responsive, fluid and nestable grid layout",
		Modifiers: modifierByPrefix("uk-grid"),
		Elem:      "Div",
		Option:    true,
	},
	{
		Name:      "leader",
		Doc:       "create dot leaders for pricing menus or tables of contents",
		Modifiers: modifierByPrefix("uk-leader"),
		Elem:      "Div",
		Option:    true,
	},
	{
		Name:      "lightbox",
		Doc:       "create a responsive lightbox gallery with images and videos",
		Modifiers: modifierByPrefix("uk-lightbox"),
		Elem:      "Div",
		Option:    true,
	},
	{
		Name:      "marker",
		Doc:       "create a marker icon that can be displayed on top of images",
		Modifiers: modifierByPrefix("uk-marker"),
		Elem:      "A",
		Option:    true,
	},
	{
		Name:      "section",
		Doc:       "create horizontal layout sections with different background colors and styles",
		Modifiers: modifierByPrefix("uk-section", "uk-padding-remove-vertical"),
		Elem:      "Div",
	},
}

var modifiers = map[string]modifier{
	// Alert
	"uk-alert-primary": {
		Name:  "Primary",
		Class: "uk-alert-primary",
		Doc:   "gives the message a prominent styling.",
	},
	"uk-alert-success": {
		Name:  "Success",
		Class: "uk-alert-success",
		Doc:   "indicates success or a positive message.",
	},
	"uk-alert-warning": {
		Name:  "Warning",
		Class: "uk-alert-warning",
		Doc:   "indicates a message containing a warning.",
	},
	"uk-alert-danger": {
		Name:  "Danger",
		Class: "uk-alert-danger",
		Doc:   "indicates an important or error message.",
	},
	// Card
	"uk-card-default": {
		Name:  "Default",
		Class: "uk-card-default",
		Doc:   "to create a visually styled box.",
	},
	"uk-card-primary": {
		Name:  "Primary",
		Class: "uk-card-primary",
		Doc:   "to modify the card and emphasize it with a primary color.",
	},
	"uk-card-secondary": {
		Name:  "Secondary",
		Class: "uk-card-secondary",
		Doc:   "to modify the card and give it a secondary background color.",
	},
	"uk-card-hover": {
		Name:  "Hover",
		Class: "uk-card-hover",
		Doc:   "to create a hover effect on the card.",
	},
	"uk-card-small": {
		Name:  "Small",
		Class: "uk-card-small",
		Doc:   "to apply a smaller padding.",
	},
	"uk-card-large": {
		Name:  "Large",
		Class: "uk-card-large",
		Doc:   "to apply a larger padding.",
	},
	// Container
	"uk-container-xsmall": {
		Name:  "XSmall",
		Class: "uk-container-xsmall",
		Doc:   "for a xsmall container.",
	},
	"uk-container-small": {
		Name:  "Small",
		Class: "uk-container-small",
		Doc:   "for a small container.",
	},
	"uk-container-large": {
		Name:  "Large",
		Class: "uk-container-large",
		Doc:   "for a large container.",
	},
	"uk-container-xlarge": {
		Name:  "XLarge",
		Class: "uk-container-xlarge",
		Doc:   "for a xlarge container.",
	},
	"uk-container-expand": {
		Name:  "Expand",
		Class: "uk-container-expand",
		Doc:   "if you do not want to limit the container width but still want the dynamic horizontal padding.",
	},
	// Grid
	"uk-grid-small": {
		Name:  "Small",
		Class: "uk-grid-small",
		Doc:   "to apply a small gap.",
	},
	"uk-grid-medium": {
		Name:  "Medium",
		Class: "uk-grid-medium",
		Doc:   "to apply a medium gap like the default one, but without a breakpoint.",
	},
	"uk-grid-large": {
		Name:  "Large",
		Class: "uk-grid-large",
		Doc:   "to apply a large gap with breakpoints.",
	},
	"uk-grid-collapse": {
		Name:  "Collapse",
		Class: "uk-grid-collapse",
		Doc:   "to remove the grid gap entirely.",
	},
	"uk-grid-divider": {
		Name:  "Divider",
		Class: "uk-grid-divider",
		Doc:   "to separate grid cells with lines.",
	},
	"uk-grid-match": {
		Name:  "MatchHeight",
		Class: "uk-grid-match",
		Doc:   "to match the height of the direct child of each cell.",
	},
	// Section
	"uk-section-default": {
		Name:  "Default",
		Class: "uk-section-default",
		Doc:   "adds the default background color of your site.",
	},
	"uk-section-muted": {
		Name:  "Muted",
		Class: "uk-section-muted",
		Doc:   "adds a muted background color.",
	},
	"uk-section-primary": {
		Name:  "Primary",
		Class: "uk-section-primary",
		Doc:   "adds a primary background color.",
	},
	"uk-section-secondary": {
		Name:  "Secondary",
		Class: "uk-section-secondary",
		Doc:   "adds a secondary background color.",
	},
	"uk-section-xsmall": {
		Name:  "XSmall",
		Class: "uk-section-xsmall",
		Doc:   "to decrease a section's padding to a minimum.",
	},
	"uk-section-small": {
		Name:  "Small",
		Class: "uk-section-small",
		Doc:   "to decrease a section's padding.",
	},
	"uk-section-large": {
		Name:  "Large",
		Class: "uk-section-large",
		Doc:   "to increase a section's padding.",
	},
	"uk-section-xlarge": {
		Name:  "XLarge",
		Class: "uk-section-xlarge",
		Doc:   "to further increase a section's padding.",
	},
	// Padding
	"uk-padding-remove-vertical": {
		Name:  "RemoveVerticalPadding",
		Class: "uk-padding-remove-vertical",
		Doc:   "removes top and bottom padding from an element.",
	},
}

var uikitTmpl = `
// Code generated by go generate; DO NOT EDIT.

package uikit

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// UI{{title .Name}} is a component
type UI{{title .Name}} interface {
	app.UI

	// Class adds a CSS class to the {{.Name}}.
	Class(v string) UI{{title .Name}}

	// Content sets the main content.
	Content(elems ...app.UI) UI{{title .Name}}
{{if .Option}}
	// Option sets a component option.
	Option(k string, v interface{}) UI{{title .Name}}
{{end}}
{{range $value := .Modifiers}}
	// {{title $value.Name}} {{$value.Doc}}
	{{title $value.Name}}() UI{{title $.Name}}
{{end}}
}

type {{.Name}} struct {
	app.Compo

	Iclass   string
	Icontent []app.UI
{{if .Option}}
	Ioptions map[string]interface{}
{{end}}
}

// {{title .Name}} returns a {{.Name}} component.
func {{title .Name}}() UI{{title .Name}} {
{{- if .Option}}
	return &{{.Name}}{}
{{else}}
	return &{{.Name}}{
		Iclass: "uk-{{.Name}}",
	}
{{end -}}
}

func ({{id .Name}} *{{.Name}}) Class(v string) UI{{title .Name}} {
	if {{id .Name}}.Iclass != "" {
		{{id .Name}}.Iclass += " "
	}

	{{id .Name}}.Iclass += v
	return {{id .Name}}
}

func ({{id .Name}} *{{.Name}}) Content(elems ...app.UI) UI{{title .Name}} {
	{{id .Name}}.Icontent = app.FilterUIElems(elems...)
	return {{id .Name}}
}

{{if .Option}}
func ({{id .Name}} *{{.Name}}) Option(k string, v interface{}) UI{{title .Name}} {
	if {{id .Name}}.Ioptions == nil {
		{{id .Name}}.Ioptions = make(map[string]interface{}, 0)
	}
	{{id .Name}}.Ioptions[k] = v
	return {{id .Name}}
}
{{end}}

{{range $value := .Modifiers}}
	func ({{id $.Name}} *{{$.Name}}) {{title $value.Name}}() UI{{title $.Name}} {
		{{id $.Name}}.Class("{{$value.Class}}")
		return {{id $.Name}}
	}
{{end}}


func ({{id .Name}} *{{.Name}}) Render() app.UI {
{{- if .Option}}
	opts, _ := JSONString({{id .Name}}.Ioptions)
{{end}}
	return app.{{.Elem}}().
{{if .Option}}
		DataSet("uk-{{lower .Name}}", opts).
{{end}}
		Class({{id .Name}}.Iclass).
		Body({{id .Name}}.Icontent...)
}
`

func main() {

	generateUIkitGo()

}

func generateUIkitGo() {

	funcMap := template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
		"id": func(name string) string {
			return strings.ToLower(string(name[0]))
		},
	}

	for _, c := range components {

		f, err := os.Create(fmt.Sprintf("%s.go", strings.ToLower(c.Name)))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		tmpl, err := template.New("component").Funcs(funcMap).Parse(uikitTmpl)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(f, c)
		if err != nil {
			panic(err)
		}
	}
}

func modifierByNames(names ...string) []modifier {
	res := make([]modifier, 0, len(names))
	for _, n := range names {
		mod, ok := modifiers[n]
		if !ok {
			panic("unknown modifier: " + n)
		}
		res = append(res, mod)
	}

	sort.Slice(res, func(i, j int) bool {
		return strings.Compare(res[i].Name, res[j].Name) <= 0
	})

	return res
}

func modifierByPrefix(prefixes ...string) []modifier {
	res := make([]modifier, 0, len(prefixes))
	for _, prefix := range prefixes {

		for k, mod := range modifiers {
			if strings.HasPrefix(k, prefix) {
				res = append(res, mod)
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return strings.Compare(res[i].Name, res[j].Name) <= 0
	})

	return res
}
