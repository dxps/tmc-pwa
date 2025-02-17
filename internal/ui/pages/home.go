package pages

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Homepage struct {
	app.Compo
}

func (h *Homepage) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Homepage"),
		app.A().Href("/about").Text("About"),
	)
}
