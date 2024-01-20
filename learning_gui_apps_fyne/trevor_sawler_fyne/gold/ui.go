package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// get the current price of gold - we do this with a separate function that returns
	// 3 objects. We could do it all here, but it'll be harder to test later.
	openPrice, currentPrice, priceChange := app.getPriceText()
	// put price information into a container
	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)

	app.PriceContainer = priceContent

	// get toolbar
	toolBar := app.getToolBar() // NOTE: here I'm starting to see the patterns. Even
	// with the separate files they are all in package 'main' and all connected with methods
	// off of *Config. This connects everything as all data from all packages are connecting
	// back to the storage struct and then I have access throughout the app.
	app.ToolBar = toolBar

	priceTabContent := app.pricesTab()

	// get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), canvas.NewText("Holdings content goes here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(priceContent, toolBar, tabs)

	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 5) { // this is the first time I've ever seen this
			// construct of 'for range' its an infinite range loop ;-o
			app.refreshPriceContent()
		}
	}()

}

func (app *Config) refreshPriceContent() {
	app.InfoLog.Print("refreshing prices")

	open, current, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()
}
