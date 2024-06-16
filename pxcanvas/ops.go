package pxcanvas

import (
	"fyne.io/fyne/v2"
)

func (pxCanvas *PxCanvas) scale(direction int) {
	if direction > 0 {
		pxCanvas.PxSize++
	} else if direction < 0 && pxCanvas.PxSize > 1 {
		pxCanvas.PxSize--
	} else {
		pxCanvas.PxSize = 10
	}
}

func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff

	pxCanvas.Refresh()
}
