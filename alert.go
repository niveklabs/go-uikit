package uikit

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// UIAlert is a component
type UIAlert interface {
	app.UI

	// Class adds a CSS class to the section.
	Class(c string) UIAlert

	// Content sets the main content.
	Content(elems ...app.UI) UIAlert

	// Option sets a single component option.
	Option(k string, v interface{}) UIAlert

	// Primary gives the message a prominent styling.
	Primary() UIAlert

	// Success indicates success or a positive message.
	Success() UIAlert

	// Warning indicates a message containing a warning.
	Warning() UIAlert

	// Danger indicates an important or error message.
	Danger() UIAlert

	// Close adds close functionality.
	Close(v bool) UIAlert
}

type alert struct {
	app.Compo

	Iclass   string
	Icontent []app.UI
	Ioptions map[string]interface{}
	Iclose   bool
}

func (a *alert) Class(c string) UIAlert {
	if a.Iclass != "" {
		a.Iclass += " "
	}

	a.Iclass += c
	return a
}

func (a *alert) Content(elems ...app.UI) UIAlert {
	a.Icontent = app.FilterUIElems(elems...)
	return a
}

func (a *alert) Option(k string, v interface{}) UIAlert {
	if a.Ioptions == nil {
		a.Ioptions = make(map[string]interface{}, 0)
	}
	a.Ioptions[k] = v
	return a
}

func (a *alert) Primary() UIAlert {
	a.Class("uk-alert-primary")
	return a
}

func (a *alert) Success() UIAlert {
	a.Class("uk-alert-success")
	return a
}

func (a *alert) Warning() UIAlert {
	a.Class("uk-alert-warning")
	return a
}

func (a *alert) Danger() UIAlert {
	a.Class("uk-alert-danger")
	return a
}

func (a *alert) Close(v bool) UIAlert {
	a.Iclose = v
	return a
}

// Alert returns a alert component.
func Alert() UIAlert {
	return &alert{}
}

func (a *alert) Render() app.UI {

	opts, _ := JSONString(a.Ioptions)

	if a.Iclose {
		contents := []app.UI{
			app.A().
				Href("#").
				DataSet("uk-close", "").
				Class("uk-alert-close"),
		}

		a.Icontent = append(contents, a.Icontent...)
	}

	return app.Div().
		DataSet("uk-alert", opts).
		Class(a.Iclass).
		Body(a.Icontent...)
}
