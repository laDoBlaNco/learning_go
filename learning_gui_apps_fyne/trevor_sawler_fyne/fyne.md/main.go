package main

import (
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// the config we are going to use to share  information between different parts of the app
type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (app *config) createMenuItems(win fyne.Window) {
	// here we see a typical process of developing menus in fyne. You start by creating each layer
	// of menu (Mainmenu/subsequent menus 'file, etc'/subsequent menu items 'save, open, etc')
	// Then you place them all in each other where they go. items in the menu which goes in its
	// parent menu which goes in the main menu, etc.
	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(win)) // we could do inline functions or...

	saveMenuItem := fyne.NewMenuItem("Save", app.saveFunc(win)) // ...create a func that returns a func

	app.SaveMenuItem = saveMenuItem // this is so we can enable and disable from another item
	app.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

// Funcs that return the funcs needed for our menuitems
func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if write == nil {
				// user cancelled
				return
			}
			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with a .md ext!", win)
				return
			}
			//save the file
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false

		}, win)
		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}

// creating a filter var at package level since we'll use in two funcs
var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (app *config) saveFunc(win fyne.Window) func() {
	return func() {
		if app.CurrentFile != nil {
			write, err := storage.Writer(app.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
			}

			write.Write([]byte(app.EditWidget.Text))
			defer write.Close()
		}
	}
}

func (app *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if read == nil {
				return
			}
			defer read.Close()

			data, err := ioutil.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			app.EditWidget.SetText(string(data))
			app.CurrentFile = read.URI()
			win.SetTitle(win.Title() + " - " + read.URI().Name())
			app.SaveMenuItem.Disabled = false

		}, win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

// here we have a package level var
var cfg config

// Below in our main func we first start with the comments of the steps we will take. And for the
// vast majority of fyne apps we will create, it'll be these same 5 steps.
func main() {

	// create a fyne app
	a := app.New()
	a.Settings().SetTheme(&myTheme{})

	// create a window for the app
	win := a.NewWindow("Markdown")

	// get the user interface - rather than creating a function that gets the widgets and wraps
	// them in a container and returns it. its much more useful to simply return the actual
	// widgets
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	// show window and run app
	win.Resize(fyne.Size{Width: 500, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}
