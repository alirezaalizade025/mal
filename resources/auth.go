package resources

import (
	"fmt"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Login() {
	myWindow := Window

	// Username entry field
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	// Password entry field
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	// Login button
	loginButton := widget.NewButton("Login", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text
		fmt.Printf("Username: %s\nPassword: %s\n", username, password)
	})

	// Create a form container with padding
	formContainer := container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("Login"),
			layout.NewSpacer(),
		),

		usernameEntry,
		passwordEntry,
		fyne.CanvasObject(&layout.Spacer{}),
		loginButton,
		fyne.CanvasObject(&layout.Spacer{}),
	)

	// Set the form container as the window's content
	myWindow.SetContent(formContainer)

}
