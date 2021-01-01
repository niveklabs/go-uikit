// Code generated by go generate; DO NOT EDIT.

package uikit

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// UIMarker is a component that creates a marker icon that can be displayed on top of images
type UIMarker interface {
	app.UI

	// Class adds a CSS class to the marker.
	Class(v string) UIMarker

	// Content sets the main content.
	Content(elems ...app.UI) UIMarker

	// Option sets a component option.
	Option(k string, v interface{}) UIMarker
}

type marker struct {
	app.Compo

	Iclass   string
	Icontent []app.UI

	Ioptions map[string]interface{}
}

// Marker returns a marker component.
func Marker() UIMarker {
	return &marker{}
}

func (m *marker) Class(v string) UIMarker {
	if m.Iclass != "" {
		m.Iclass += " "
	}

	m.Iclass += v
	return m
}

func (m *marker) Content(elems ...app.UI) UIMarker {
	m.Icontent = app.FilterUIElems(elems...)
	return m
}

func (m *marker) Option(k string, v interface{}) UIMarker {
	if m.Ioptions == nil {
		m.Ioptions = make(map[string]interface{}, 0)
	}
	m.Ioptions[k] = v
	return m
}

func (m *marker) Render() app.UI {
	opts, _ := JSONString(m.Ioptions)

	return app.A().
		DataSet("uk-marker", opts).
		Class(m.Iclass).
		Body(m.Icontent...)
}
