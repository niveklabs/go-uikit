// Code generated by go generate; DO NOT EDIT.

package uikit

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// UICard is a component
type UICard interface {
	app.UI

	// Class adds a CSS class to the card.
	Class(v string) UICard

	// Content sets the main content.
	Content(elems ...app.UI) UICard

	// Default to create a visually styled box.
	Default() UICard

	// Hover to create a hover effect on the card.
	Hover() UICard

	// Large to apply a larger padding.
	Large() UICard

	// Primary to modify the card and emphasize it with a primary color.
	Primary() UICard

	// Secondary to modify the card and give it a secondary background color.
	Secondary() UICard

	// Small to apply a smaller padding.
	Small() UICard
}

type card struct {
	app.Compo

	Iclass   string
	Icontent []app.UI
}

// Card returns a card component.
func Card() UICard {
	return &card{
		Iclass: "uk-card",
	}
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

func (c *card) Default() UICard {
	c.Class("uk-card-default")
	return c
}

func (c *card) Hover() UICard {
	c.Class("uk-card-hover")
	return c
}

func (c *card) Large() UICard {
	c.Class("uk-card-large")
	return c
}

func (c *card) Primary() UICard {
	c.Class("uk-card-primary")
	return c
}

func (c *card) Secondary() UICard {
	c.Class("uk-card-secondary")
	return c
}

func (c *card) Small() UICard {
	c.Class("uk-card-small")
	return c
}

func (c *card) Render() app.UI {
	return app.Div().
		Class(c.Iclass).
		Body(c.Icontent...)
}
