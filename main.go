package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var score int = 0

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	textScore := widget.NewLabel(strconv.Itoa(score))
	button := widget.NewButton("+1", func() {
		score++
		textScore.SetText(strconv.Itoa(score))
	})

	cont := container.NewGridWithRows(2, button, textScore)
	w.SetContent(cont)
	w.ShowAndRun()
}
