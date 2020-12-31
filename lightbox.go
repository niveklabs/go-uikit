// Code generated by go generate; DO NOT EDIT.

package uikit

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// UILightbox is a component
type UILightbox interface {
	app.UI

	// Class adds a CSS class to the lightbox.
	Class(v string) UILightbox

	// Content sets the main content.
	Content(elems ...app.UI) UILightbox

	// Option sets a component option.
	Option(k string, v interface{}) UILightbox
}

type lightbox struct {
	app.Compo

	Iclass   string
	Icontent []app.UI

	Ioptions map[string]interface{}
}

func (l *lightbox) Class(v string) UILightbox {
	if l.Iclass != "" {
		l.Iclass += " "
	}

	l.Iclass += v
	return l
}

func (l *lightbox) Content(elems ...app.UI) UILightbox {
	l.Icontent = app.FilterUIElems(elems...)
	return l
}

func (l *lightbox) Option(k string, v interface{}) UILightbox {
	if l.Ioptions == nil {
		l.Ioptions = make(map[string]interface{}, 0)
	}
	l.Ioptions[k] = v
	return l
}

func (l *lightbox) Render() app.UI {

	opts, _ := JSONString(l.Ioptions)

	return app.Div().
		DataSet("uk-lightbox", opts).
		Class(l.Iclass).
		Body(l.Icontent...)
}
