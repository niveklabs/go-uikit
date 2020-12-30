package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	uikit "github.com/niveklabs/go-uikit"
)

func pages() map[string]func() app.UI {
	return map[string]func() app.UI{
		"": newPage,
	}
}

type page struct {
	app.Compo
}

func newPage() app.UI {
	return &page{}
}

func (p *page) Render() app.UI {
	return app.Div().
		Body(
			app.H4().Text("Section"),
			uikit.Section().
				Muted().
				XSmall().
				Content(
					uikit.Container().
						Content(
							app.P().Text("XSmall section Muted!"),
						),
				),
			uikit.Section().
				Primary().
				Content(
					uikit.Container().
						Content(
							app.P().Text("Section Primary!"),
						),
				),

			uikit.Section().
				Default().
				Content(
					uikit.Container().
						Content(
							app.H4().Text("Alerts"),
							uikit.Alert().
								Primary().
								Close(true).
								Content(
									app.P().Text("I am a primary alert!"),
								),
							uikit.Alert().
								Success().
								Close(true).
								Content(
									app.P().Text("I am a success alert!"),
								),
							uikit.Alert().
								Warning().
								Close(true).
								Content(
									app.P().Text("I am a warning alert!"),
								),
							uikit.Alert().
								Danger().
								Close(true).
								Content(
									app.P().Text("I am a danger alert!"),
								),
						),
				),
			uikit.Section().
				Default().
				Content(
					uikit.Container().
						Content(
							app.H4().Text("Accordion"),
							uikit.Accordion().
								Option("multiple", true).
								Class("uk-width-1-2").
								Content(
									uikit.AccordionItem().
										Title("Item 1").
										Content(
											app.P().Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
										),
									uikit.AccordionItem().
										Title("Item 2").
										Content(
											app.P().Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
										),
								),
						),
				),
		)
}
