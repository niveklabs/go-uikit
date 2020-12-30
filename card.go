package uikit

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

//TODO - implement https://getuikit.com/docs/card

// UICard is a component
type UICard interface {
	app.UI

	// Class adds a CSS class to the card.
	Class(v string) UICard

	// Content sets the main content.
	Content(elems ...app.UI) UICard

	// Default visually styled box.
	Default() UICard

	// Primary emphasize with a primary color.
	Primary() UICard

	// Secondary emphasize with a secondary background color.
	Secondary() UICard

	// Small card.
	Small() UICard

	// Large card.
	Large() UICard

	// Hover effect on the card.
	Hover() UICard
}

type card struct {
	app.Compo

	Iclass   string
	Icontent []app.UI
}

func (c *card) Class(v string) UICard {
	if c.Iclass != "" {
		c.Iclass += " "
	}

	c.Iclass += v
	return c
}

func (c *card) Content(elems ...app.UI) UICard {
	c.Icontent = app.FilterUIElems(elems...)
	return c
}

// Card returns a card component.
func Card() UICard {
	return &card{
		Iclass: "uk-card",
	}
}

func (c *card) Small() UICard {
	return c.Class("uk-card-small")
}

func (c *card) Large() UICard {
	return c.Class("uk-card-large")
}

func (c *card) Hover() UICard {
	return c.Class("uk-card-hover")
}

func (c *card) Default() UICard {
	return c.Class("uk-card-default")
}

func (c *card) Primary() UICard {
	return c.Class("uk-card-primary")
}

func (c *card) Secondary() UICard {
	return c.Class("uk-card-secondary")
}

func (c *card) Render() app.UI {
	return app.Div().
		Class(c.Iclass).
		Body(c.Icontent...).
		Class("uk-card-body")
}
