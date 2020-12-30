package uikit

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// UIContainer is a component
type UIContainer interface {
	app.UI

	// Class adds a CSS class to the container.
	Class(v string) UIContainer

	// Content sets the main content.
	Content(elems ...app.UI) UIContainer

	// XSmall container.
	XSmall() UIContainer

	// Small container.
	Small() UIContainer

	// Large container.
	Large() UIContainer

	// XLarge container.
	XLarge() UIContainer

	// Expand if you do not want to limit the container width but still want the dynamic horizontal padding.
	Expand() UIContainer
}

type container struct {
	app.Compo

	Iclass   string
	Icontent []app.UI
}

func (c *container) Class(v string) UIContainer {
	if c.Iclass != "" {
		c.Iclass += " "
	}

	c.Iclass += v
	return c
}

func (c *container) Content(elems ...app.UI) UIContainer {
	c.Icontent = app.FilterUIElems(elems...)
	return c
}

// Container returns a container component.
func Container() UIContainer {
	return &container{
		Iclass: "uk-container",
	}
}

func (c *container) XSmall() UIContainer {
	return c.Class("uk-container-xsmall")
}

func (c *container) Small() UIContainer {
	return c.Class("uk-container-small")
}

func (c *container) Large() UIContainer {
	return c.Class("uk-container-large")
}

func (c *container) XLarge() UIContainer {
	return c.Class("uk-container-xlarge")
}

func (c *container) Expand() UIContainer {
	return c.Class("uk-container-expand")
}

func (c *container) Render() app.UI {
	return app.Div().
		Class(c.Iclass).
		Body(c.Icontent...)
}
