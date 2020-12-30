package uikit

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// UIAccordion is a component
type UIAccordion interface {
	app.UI

	// Class adds a CSS class to the accordion.
	Class(c string) UIAccordion

	// Content sets the main content.
	Content(elems ...app.UI) UIAccordion

	// Option sets a single component option.
	Option(k string, v interface{}) UIAccordion
}

type accordion struct {
	app.Compo

	Iclass   string
	Icontent []app.UI
	Ioptions map[string]interface{}
}

func (a *accordion) Class(c string) UIAccordion {
	if a.Iclass != "" {
		a.Iclass += " "
	}

	a.Iclass += c
	return a
}

func (a *accordion) Content(elems ...app.UI) UIAccordion {
	a.Icontent = app.FilterUIElems(elems...)
	return a
}

func (a *accordion) Option(k string, v interface{}) UIAccordion {
	if a.Ioptions == nil {
		a.Ioptions = make(map[string]interface{}, 0)
	}
	a.Ioptions[k] = v
	return a
}

// Accordion returns a UI Accordion
func Accordion() UIAccordion {
	return &accordion{}
}

func (a *accordion) Render() app.UI {

	opts, _ := JSONString(a.Ioptions)

	return app.Ul().
		DataSet("uk-accordion", opts).
		Class(a.Iclass).
		Body(a.Icontent...)
}
