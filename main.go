package main

import (
	"time"

	preferences "yt-go/preferences"
	YtServer "yt-go/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

/* It's just good to see a dynamic widget updating so fast each second :) */
func updateTime(clock *widget.Label) {
	var formatted = time.Now().Format("Current time is : 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.NewWithID("com.yt-go.dev.preferences")
	w := a.NewWindow("Yt-Go")

	preferences.InitializePreferences(a)

	var clock = widget.NewLabel("clock_content")

	/* Initializes a Vertival Layout Container */
	var mainWindowContent *fyne.Container = container.NewVBox(
		clock,
	)

	/*
	* Updates the time displayed each second
	* Background process using Goroutines `go func() {}()`
	 */
	go func() {
		for range time.Tick(time.Second) {
			fyne.Do(func() { updateTime(clock) })
		}
	}()

	/* Applies content to the window and resizes it */
	w.SetContent(mainWindowContent)
	w.Resize(fyne.NewSize(800, 600))

	/*
	* Minimizes the app when closed
	* System Tray depending on current OS
	 */
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Yt-go",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	go func() { YtServer.StartServer() }()
	w.ShowAndRun()
}
