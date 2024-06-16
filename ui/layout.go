package ui

import "fyne.io/fyne/v2/container"

// The Setup function initializes and sets up the user interface components for Pixl.
func Setup(app *AppInit) {
	SetupMenu(app)
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(
		nil,
		swatchesContainer,
		nil,
		colorPicker,
		app.PixlCanvas,
	)

	app.PixlWindow.SetContent(appLayout)
}
