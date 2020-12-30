package uikit

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// TODO - implement https://getuikit.com/docs/grid

// UIGrid is a component
type UIGrid interface {
	app.UI

	// Class adds a CSS class to the grid.
	Class(v string) UIGrid

	// Content sets the main content.
	Content(elems ...app.UI) UIGrid

	// Option sets a single component option.
	Option(k string, v interface{}) UIGrid

	// Divider modifier to separate grid cells with lines
	Divider() UIGrid
}

type grid struct {
	app.Compo

	Iclass   string
	Icontent []app.UI
	Ioptions map[string]interface{}
}

func (g *grid) Class(v string) UIGrid {
	if g.Iclass != "" {
		g.Iclass += " "
	}

	g.Iclass += v
	return g
}

func (g *grid) Content(elems ...app.UI) UIGrid {
	g.Icontent = app.FilterUIElems(elems...)
	return g
}

func (g *grid) Option(k string, v interface{}) UIGrid {
	if g.Ioptions == nil {
		g.Ioptions = make(map[string]interface{}, 0)
	}
	g.Ioptions[k] = v
	return g
}

func (g *grid) Divider() UIGrid {
	g.Class("uk-grid-divider")
	return g
}

// Grid returns a grid component.
func Grid() UIGrid {
	return &grid{}
}

func (g *grid) Render() app.UI {

	opts, _ := JSONString(g.Ioptions)

	return app.Div().
		DataSet("uk-grid", opts).
		Class(g.Iclass).
		Body(g.Icontent...)
}
