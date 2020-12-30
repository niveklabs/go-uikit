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
			uikit.Section().
				XSmall().
				Content(
					uikit.Container().
						Content(
							app.H1().Text("Sections"),
						),
				),
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
				Secondary().
				Content(
					uikit.Container().
						Content(
							app.P().Text("Section Secondary!"),
						),
				),

			uikit.Section().
				Default().
				Content(
					uikit.Container().
						Content(
							app.H1().Text("Alerts"),
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
							app.H1().Text("Accordions"),
							uikit.Grid().
								Class("uk-child-width-expand").
								Divider().
								Content(
									app.Div().
										Body(
											app.P().Text("With multiple select"),
											uikit.Accordion().
												Option("multiple", true).
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
									app.Div().
										Body(
											app.P().Text("With single select"),
											uikit.Accordion().
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
						),
				),
			uikit.Section().
				Content(
					uikit.Container().
						Content(
							app.H1().
								Text("Cards"),
							uikit.Grid().
								Class("uk-child-width-expand uk-text-center").
								Content(
									app.Div().Body(
										uikit.Card().
											Default().
											Content(
												app.H3().
													Class("uk-card-title").
													Text("Default"),
												app.P().
													Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
											).
											Hover(),
									),
									app.Div().Body(
										uikit.Card().
											Primary().
											Small().
											Content(
												app.H3().
													Class("uk-card-title").
													Text("Small"),
												app.P().
													Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
											).
											Hover(),
									),
									app.Div().Body(
										uikit.Card().
											Secondary().
											Large().
											Content(
												app.H3().
													Class("uk-card-title").
													Text("Large"),
												app.P().
													Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
											).
											Hover(),
									),
								),
						),
				),
		)
}
