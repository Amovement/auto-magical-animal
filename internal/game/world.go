package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// World Structure of the World background map
type World struct {
	width  int
	height int
}

// NewWorld creates a new world.
func NewWorld(width, height int) *World {
	w := &World{
		width:  width,
		height: height,
	}
	return w
}

// Draw paints current game state.
func (w *World) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 124,
		G: 235,
		B: 144,
		A: 255,
	})
}
