package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	// filepath
	fileName := "./picture/sample.jpg"

	// fyne
	myApp := app.New()
	w := myApp.NewWindow("Picture")

	// Set canvas file
	image := canvas.NewImageFromFile(fileName)
	image.FillMode = canvas.ImageFillOriginal

	w.SetContent(image)
	w.ShowAndRun()
}
