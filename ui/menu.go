package ui

import (
	"errors"
	"image"
	"image/png"
	"os"
	"pixl/utils"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Creates a menu item for creating a new image with specified width and
// height in a GUI application.
func BuildMenuItem(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)

			if err != nil {
				return errors.New("must be a positive integer")
			}

			if width <= 0 {
				return errors.New("must be greater than 0")
			}

			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}

		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0

				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid width"), app.PixlWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)
				}

				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid height"), app.PixlWindow)
				} else {
					pixelHeight, _ = strconv.Atoi(heightEntry.Text)
				}

				app.PixlCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PixlWindow)
	})
}

// Saves the pixel data from a canvas to a PNG file using a file save dialog.
func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, e error) {
		if uri == nil {
			return
		} else {
			err := png.Encode(uri, app.PixlCanvas.PixelData)

			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
			app.State.SetFilePath(uri.URI().Path())
		}
	}, app.PixlWindow)
}

// Creates a "Save As..." menu item that opens a save file dialog.
func BuildSaveAsMenuItem(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As...", func() {
		saveFileDialog(app)
	})
}

// Creates a menu item for saving an image to a file.
func BuildSaveMenuItem(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
		} else {
			tryClose := func(fh *os.File) {
				err := fh.Close()
				if err != nil {
					dialog.ShowError(err, app.PixlWindow)
				}
			}

			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)

			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}

			err = png.Encode(fh, app.PixlCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}
		}
	})
}

// Creates a menu item for opening a file, loading the image, setting
// file path, and updating swatches based on image colors.
func BuildOpenMenuItem(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Open", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, e error) {
			if uri == nil {
				return
			}

			img, _, err := image.Decode(uri)

			if err != nil {
				dialog.ShowError(err, app.PixlWindow)
				return
			}

			app.PixlCanvas.LoadImage(img)
			app.State.SetFilePath(uri.URI().Path())
			imgColors := utils.GetImageColors(img)

			i := 0
			for c := range imgColors {
				if i == len(app.Swatches) {
					break
				}
				app.Swatches[i].SetColor(c)
				i++
			}
		}, app.PixlWindow)
	})
}

func BuildMenu(app *AppInit) *fyne.Menu {
	return fyne.NewMenu("File",
		BuildMenuItem(app),
		BuildOpenMenuItem(app),
		BuildSaveMenuItem(app),
		BuildSaveAsMenuItem(app),
	)
}

func SetupMenu(app *AppInit) {
	menu := BuildMenu(app)
	mainMenu := fyne.NewMainMenu(menu)
	app.PixlWindow.SetMainMenu(mainMenu)
}
