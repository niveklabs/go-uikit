package uikit

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// UISection is a component
type UISection interface {
	app.UI

	// Class adds a CSS class to the section.
	Class(c string) UISection

	// Content sets the main content.
	Content(elems ...app.UI) UISection

	// Default adds the default background color of your site.
	Default() UISection

	// Muted adds a muted background color.
	Muted() UISection

	// Primary adds a primary background color.
	Primary() UISection

	// Secondary adds a secondary background color.
	Secondary() UISection

	// PreserveColor modifier
	PreserveColor() UISection

	// XSmall decreases a section's padding to a minimum.
	XSmall() UISection

	// Small decreases a section's padding.
	Small() UISection

	// Large increases a section's padding.
	Large() UISection

	// XLarge further increases a section's padding.
	XLarge() UISection

	// PaddingRemoveVertical removes a section's padding.
	PaddingRemoveVertical() UISection
}

type section struct {
	app.Compo

	Icontent []app.UI
	Iclass   string
}

// Section returns a UISection
func Section() UISection {
	return &section{
		Iclass: "uk-section",
	}
}

func (s *section) Class(c string) UISection {
	if s.Iclass != "" {
		s.Iclass += " "
	}

	s.Iclass += c
	return s
}

func (s *section) Content(elems ...app.UI) UISection {
	s.Icontent = app.FilterUIElems(elems...)
	return s
}

func (s *section) Default() UISection {
	return s.Class("uk-section-default")
}

func (s *section) Muted() UISection {
	return s.Class("uk-section-muted")
}

func (s *section) Primary() UISection {
	return s.Class("uk-section-primary")
}

func (s *section) Secondary() UISection {
	return s.Class("uk-section-secondary")
}

func (s *section) PreserveColor() UISection {
	return s.Class("uk-preserve-color")
}

func (s *section) XSmall() UISection {
	return s.Class("uk-section-xsmall")
}

func (s *section) Small() UISection {
	return s.Class("uk-section-small")
}

func (s *section) Large() UISection {
	return s.Class("uk-section-large")
}

func (s *section) XLarge() UISection {
	return s.Class("uk-section-xlarge")
}

func (s *section) PaddingRemoveVertical() UISection {
	return s.Class("uk-padding-remove-vertical")
}

func (s *section) Render() app.UI {
	return app.Div().
		Class(s.Iclass).
		Body(s.Icontent...)
}
