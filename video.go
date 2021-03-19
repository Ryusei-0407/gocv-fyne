package main

import (
	"image"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"gocv.io/x/gocv"
)

func MakeMovie(ShowImg *canvas.Image) {
	movie, _ := gocv.OpenVideoCapture("./video/sakura.mp4")
	img := gocv.NewMat()
	imgDD, _ := img.ToImage()
	for {
		movie.Read(&img)
		gocv.Resize(img, &img, image.Point{250, 250}, 0, 0, gocv.InterpolationDefault)
		imgDD, _ := img.ToImage()
		ShowImg.Image = imgDD
		canvas.Refresh(ShowImg)
		time.Sleep(time.Millisecond * 33)
	}
}

func main() {
	fileName := "./picture/sample.jpg"
	img := gocv.IMRead(fileName, -1)
	gocv.Resize(img, &img, image.Point{250, 250}, 0, 0, gocv.InterpolationDefault)

	imgD, _ := img.ToImage()

	myApp := app.New()
	w := myApp.NewWindow("Image")

	ShowImg := canvas.NewImageFromFile(imgD)
	ShowImg.FillMode = canvas.ImageFillOriginal

	label := widget.NewLabel("Show Movie")
	go MakeMovie(ShowImg)

	w.SetContent(widget.NewVBox(
		ShowImg,
		label,
	))
	w.ShowAndRun()
}
