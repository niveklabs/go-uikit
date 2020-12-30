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

func (p *page) highlightCode() {
	app.Dispatch(func() {
		app.Window().Get("Prism").Call("highlightAll")
	})
}

func (p *page) OnMount(ctx app.Context) {
	p.Update()
	go p.load(ctx)
}

func (p *page) load(ctx app.Context) {
	defer app.Dispatch(func() {
		p.Update()
		p.highlightCode()
	})
}

func (p *page) Render() app.UI {
	return app.Div().
		Body(
			uikit.Section().
				XSmall().
				Content(
					uikit.Container().
						Content(
							app.H1().Text("Section"),
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
			uikit.Container().
				Class("uk-padding uk-padding-remove-bottom").
				Content(
					app.Pre().
						Body(
							app.Code().
								Class("language-go").
								Text(`
func (p *page) Render() app.UI {
	return app.Div().
		Body(
			uikit.Section().
			Primary().
			Content(
				uikit.Container().
					Content(
						app.P().Text("Section Primary!"),
					),
			),
		)
}
									`),
						),
				),
			uikit.Section().
				Default().
				Content(
					uikit.Container().
						Content(
							app.H1().Text("Alert"),
							uikit.Grid().
								Class("uk-child-width-expand").
								Content(
									app.Div().Body(
										uikit.Alert().
											Primary().
											Close(true).
											Content(
												app.P().Text("I am a primary alert!"),
											),
									),
									app.Div().Body(
										uikit.Alert().
											Success().
											Close(true).
											Content(
												app.P().Text("I am a success alert!"),
											),
									),
									app.Div().Body(
										uikit.Alert().
											Warning().
											Close(true).
											Content(
												app.P().Text("I am a warning alert!"),
											),
									),
									app.Div().Body(
										uikit.Alert().
											Danger().
											Close(true).
											Content(
												app.P().Text("I am a danger alert!"),
											),
									),
								),
						),
				),
			uikit.Container().
				Content(
					app.Pre().
						Body(
							app.Code().
								Class("language-go").
								Text(`
func (p *page) Render() app.UI {
	return app.Div().
		Body(
			uikit.Alert().
				Danger().
				Close(true).
				Content(
					app.P().Text("I am a danger alert!"),
				),
			)
}
							`),
						),
				),
			uikit.Section().
				Default().
				Content(
					uikit.Container().
						Content(
							app.H1().Text("Accordion"),
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
			uikit.Container().
				Content(
					app.Pre().
						Body(
							app.Code().
								Class("language-go").
								Text(`
func (p *page) Render() app.UI {
	return app.Div().
		Body(
			uikit.Accordion().
			Content(
				uikit.AccordionItem().
					Title("Item 1").
					Content(
						app.P().Text("Lorem ipsum dolor sit amet."),
					),
				uikit.AccordionItem().
					Title("Item 2").
					Content(
						app.P().Text("Lorem ipsum dolor sit amet."),
					),
			),
}
							`),
						),
				),
			uikit.Section().
				Content(
					uikit.Container().
						Content(
							app.H1().
								Text("Card"),
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
			uikit.Container().
				Content(
					app.Pre().
						Body(
							app.Code().
								Class("language-go").
								Text(`
func (p *page) Render() app.UI {
	return app.Div().
		Body(
			uikit.Card().
				Default().
				Content(
					app.H3().
						Class("uk-card-title").
						Text("Default"),
					app.P().
						Text("Lorem ipsum dolor sit amet."),
				).
				Hover(),
			)
}
							`),
						),
				),
		)
}
