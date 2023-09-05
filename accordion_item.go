package uikit

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// UIAccordionItem is a component
type UIAccordionItem interface {
	app.UI

	// Class adds a CSS class to the section.
	Class(c string) UIAccordionItem

	// Title defines and styles the toggle for accordion item
	Title(v string) UIAccordionItem

	// Content sets the main content for accordion item.
	Content(elems ...app.UI) UIAccordionItem
}

type accordionItem struct {
	app.Compo

	Iclass   string
	Ititle   string
	Icontent []app.UI
}

func (a *accordionItem) Class(c string) UIAccordionItem {
	if a.Iclass != "" {
		a.Iclass += " "
	}

	a.Iclass += c
	return a
}

func (a *accordionItem) Title(v string) UIAccordionItem {
	a.Ititle = v
	return a
}

func (a *accordionItem) Content(elems ...app.UI) UIAccordionItem {
	a.Icontent = app.FilterUIElems(elems...)
	return a
}

// AccordionItem returns a UI Accordion Item
func AccordionItem() UIAccordionItem {
	return &accordionItem{}
}

func (a *accordionItem) Render() app.UI {
	return app.Li().
		Class(a.Iclass).
		Body(
			app.A().
				Class("uk-accordion-title").
				Href("#"). // TODO fix this
				Text(a.Ititle),
			app.Div().
				Class("uk-accordion-content").
				Body(
					a.Icontent...,
				),
		)
}
