package main

import (
	"fmt"

	"yt-go/internal/preferences"
	"yt-go/internal/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("com.yt-go.dev.preferences")
	w := a.NewWindow("Yt-Go")

	var undoButton = widget.NewButton("Undo them", func() { a.Preferences().SetBool("CONFIGURED", false) })

	/* Initializes a Vertival Layout Container */
	var mainWindowContent *fyne.Container = container.NewPadded(
		container.NewVBox(
			undoButton,
		),
	)

	/* Applies content to the window and resizes it */
	w.SetContent(mainWindowContent)
	w.Resize(fyne.NewSize(1280, 576))
	w.SetMaster()
	w.Show()

	var preferencesExist bool = a.Preferences().BoolWithFallback("CONFIGURED", false)
	if !preferencesExist {
		fmt.Println("Preferences not found\nInitializing...")

		preferences.InitializePreferences(a)

		var preferencesDialog *dialog.CustomDialog
		var label = widget.NewLabel("Preferences not found. Yt-Go won't work without.")
		var confirmButton = widget.NewButton("Configure them", func() {
			a.Preferences().SetBool("CONFIGURED", true)
			preferencesDialog.Dismiss()
		})
		var mainContent = container.NewVBox(
			label,
			confirmButton,
		)

		preferencesDialog = dialog.NewCustomWithoutButtons("Preferences", mainContent, w)
		preferencesDialog.Show()
	}

	/*
	 *
	 * Minimizes the app when closed
	 * System Tray depending on current OS
	 *
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

	go func() { server.StartServer() }()
	a.Run()

	if server.StopServer() == nil {
		fmt.Println("Server shutdown gracefully. App closed")
	}
}
