package main

import (
	"dolphinPeek/internal/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("DolphinPeek 🐬")
	myWindow.Resize(fyne.NewSize(400, 300))

	widget := widget.New()
	myWindow.SetContent(widget.GetContent())

	// Start the cycling
	go widget.Start()

	myWindow.ShowAndRun()
}
