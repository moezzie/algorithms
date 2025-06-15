package display

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func Display(myImage image.Image) {
	a := app.New()
	w := a.NewWindow("Images")

	img := canvas.NewImageFromImage(myImage)
	w.SetContent(img)
	w.Resize(getSize(myImage))

	w.ShowAndRun()
}

func getSize(img image.Image) fyne.Size {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	return fyne.NewSize(max(256, float32(width)), max(256, float32(height)))
}
