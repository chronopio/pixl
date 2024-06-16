package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject
}

func (renderer *PxCanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	renderer.canvasCursor = objects
}

func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingAreaSize
}

// Is responsible for returning a list of `fyne.CanvasObject` that should be rendered on the canvas.
func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)

	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}

	objects = append(objects, renderer.canvasImage)
	objects = append(objects, renderer.canvasCursor...)

	return objects
}

func (renderer *PxCanvasRenderer) Destroy() {}

func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size)
	renderer.LayoutBorder(size)
}

// Is responsible for updating the canvas image based on the pixel data in the `PxCanvas` struct.
func (renderer *PxCanvasRenderer) Refresh() {
	if renderer.pxCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.pxCanvas.PixelData)

		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain

		renderer.pxCanvas.reloadImage = false
	}

	renderer.Layout(renderer.pxCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}

// This function `LayoutCanvas` is responsible for positioning and resizing the canvas image based on
// the pixel data in the `PxCanvas` struct.
func (renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := renderer.pxCanvas.PxCols
	imgPxHeight := renderer.pxCanvas.PxRows

	pxSize := renderer.pxCanvas.PxSize

	renderer.canvasImage.Move(fyne.NewPos(renderer.pxCanvas.CanvasOffset.X, renderer.pxCanvas.CanvasOffset.Y))

	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}

// Is responsible for positioning and sizing the border lines around the canvas image.
// It calculates the positions of the four border lines (left, top, right, bottom) based
// on the canvas offset, image height, and image width. Each border line is defined by two points
// (Position1 and Position2) that determine its start and end positions on the canvas. The function
// sets these positions for each border line to create a complete border around the canvas image.
func (renderer *PxCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.pxCanvas.CanvasOffset
	imgHeight := renderer.pxCanvas.Size().Height
	imgWidth := renderer.pxCanvas.Size().Width

	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)

	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)

	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)
}
