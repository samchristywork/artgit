package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("GUI Application")

	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() {}),
		fyne.NewMenuItem("Open", func() {}),
		fyne.NewMenuItem("Save", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Quit", func() { app.Quit() }),
	)
	mainMenu := fyne.NewMainMenu(fileMenu)
	window.SetMainMenu(mainMenu)

	window.ShowAndRun()
}
