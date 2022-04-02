package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main(){
	app := app.New()
	window := app.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")

	window.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("It`s Test fyne settings linux")
		}),
	))

	window.ShowAndRun()
}
