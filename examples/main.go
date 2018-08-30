// Package main provides various examples of Fyne API capabilities
package main

import "flag"
import "fmt"

import "github.com/fyne-io/fyne/examples/apps"

import "github.com/fyne-io/fyne"
import "github.com/fyne-io/fyne/layout"
import "github.com/fyne-io/fyne/theme"
import W "github.com/fyne-io/fyne/widget"

func blogApp(app fyne.App) {
	apps.Blog(app)
}

func calcApp(app fyne.App) {
	apps.Calculator(app)
}

func canvasApp(app fyne.App) {
	apps.Canvas(app)
}

func clockApp(app fyne.App) {
	apps.Clock(app)
}

func fractalApp(app fyne.App) {
	apps.Fractal(app)
}

func layoutApp(app fyne.App) {
	apps.Layout(app)
}

func lifeApp(app fyne.App) {
	apps.Life(app)
}

func appButton(app fyne.App, label string, onClick func(fyne.App)) *W.Button {
	return &W.Button{Text: label, OnTapped: func() {
		onClick(app)
	}}
}

func welcome(app fyne.App) {
	w := app.NewWindow("Examples")
	w.SetContent(&W.List{Children: []fyne.CanvasObject{
		&W.Label{Text: "Fyne Examples!"},

		appButton(app, "Blog", blogApp),
		appButton(app, "Calculator", calcApp),
		appButton(app, "Clock", clockApp),
		appButton(app, "Fractal", fractalApp),
		appButton(app, "Canvas", canvasApp),
		appButton(app, "Layout", layoutApp),
		appButton(app, "Life", lifeApp),

		layout.NewSpacer(),
		&W.Entry{},
		&W.Check{Text: "Check", OnChanged: func(on bool) { fmt.Println("checked", on) }},
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			&W.Button{Text: "Dark", OnTapped: func() {
				fyne.GetSettings().SetTheme("dark")
			}},
			&W.Button{Text: "Light", OnTapped: func() {
				fyne.GetSettings().SetTheme("light")
			}},
		),
		&W.Button{Text: "Quit", Icon: theme.CancelIcon(), OnTapped: func() {
			app.Quit()
		}},
	}})
	w.Show()
}

func main() {
	app := apps.NewApp()

	var ex string
	flag.StringVar(&ex, "example", "", "Launch an app directly (blog,calculator,canvas,clock)")
	flag.Parse()

	switch ex {
	case "blog":
		blogApp(app)
	case "calculator":
		calcApp(app)
	case "canvas":
		canvasApp(app)
	case "clock":
		clockApp(app)
	case "fractal":
		fractalApp(app)
	case "layout":
		layoutApp(app)
	case "life":
		lifeApp(app)
	default:
		welcome(app)
	}
}
