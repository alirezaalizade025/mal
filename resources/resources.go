package resources

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"time"
)

var Window fyne.Window

func App() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Fading Fyne App")

	logo := widget.NewLabel("Mal")
	logo.TextStyle = fyne.TextStyle{Bold: true}

	content := container.New(layout.NewVBoxLayout(),
		logo,
	)

	fullScreenContent := container.New(layout.NewCenterLayout(),
		content,
	)

	myWindow.SetContent(fullScreenContent)
	myWindow.Resize(fyne.NewSize(800, 600)) // Set the desired width and height

	myWindow.Show()

	Window = myWindow

	// Start a timer to trigger the fade effect after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)

		Login()
	}()

	Window.ShowAndRun()
}
