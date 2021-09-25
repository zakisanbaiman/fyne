package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main()  {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(
		widget.NewLabel("Hello Fyne!2"),
	)

	w.ShowAndRun()
}
