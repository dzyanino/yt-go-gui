package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	fmt.Println("Here")

	w.SetContent(widget.NewLabel("Hello Fyne!"))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
