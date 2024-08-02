package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// World 地图的结构
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

// Update game state by one tick.
//
//	更新内置时间
func (w *World) Update() {
	tick++
	if tick == consts.TimeInterval {
		tick = 0
	}
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
