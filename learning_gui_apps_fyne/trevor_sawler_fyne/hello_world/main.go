package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {

	// 'a' is my fyne application. You will always have 1 application var
	a := app.New()
	// Every application must have at least 1 window, mine is 'w'
	w := a.NewWindow("Hello, World!") // Hello World is the title

	output, entry, btn := myApp.makeUI()

	// The above would work alone, but we want to start putting stuff in our window
	// Then we put content in the window (container or widgets)
	// Typically if we have more than one thing (widget) in our app, we would need
	// a wrapper of all those widgets and that's called a 'container'
	w.SetContent(container.NewVBox(output, entry, btn))
	// w.ShowAndRun() // this is the shortcut, but we can also run it like...
	w.SetTitle(output.Text)
	w.Resize(fyne.Size{Width: 500, Height: 500})
	w.ShowAndRun() // everything now stops in the program and simply runs event loop.

}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello, World!")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance
	app.output = output
	return output, entry, btn
}
