package game

import (
	"bytes"
	"github.com/Amovement/auto-magical-animal/assets"
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Cursor struct {
	image   *ebiten.Image
	locateX float64
	locateY float64
}

func NewCursor() *Cursor {
	cursorImage, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.CursorImageBytes))
	if err != nil {
		log.Panic(err)
	}
	return &Cursor{
		image:   cursorImage,
		locateX: consts.GameHeight / 3,
		locateY: consts.GameHeight / 3,
	}
}

// ListenMouseEvent Listen for mouse events
func (c *Cursor) ListenMouseEvent() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		c.locateX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		c.locateX += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		c.locateY -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		c.locateY += 5
	}
	c.RecordNumberPress()
	// Mouse click
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		clickX, clickY := ebiten.CursorPosition()
		c.locateX = float64(clickX) - 13
		c.locateY = float64(clickY) - 13
	}
	// Press space to create an animal
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		animalNew := NewAnimal(consts.AnimalTypeByWeather, c.locateX, c.locateY)
		AppendAnimalVector(animalNew)
	}
}

func (c *Cursor) Draw(screen *ebiten.Image) {
	option := &ebiten.DrawImageOptions{}
	option.GeoM.Translate(c.locateX, c.locateY)

	// Draw cursor
	screen.DrawImage(
		c.image,
		option,
	)
}

func (c *Cursor) RecordNumberPress() {
	if ebiten.IsKeyPressed(ebiten.Key1) {
		SetNumberKeyPress(1)
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		SetNumberKeyPress(2)
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		SetNumberKeyPress(3)
	}
	if ebiten.IsKeyPressed(ebiten.Key4) {
		SetNumberKeyPress(4)
	}
	if ebiten.IsKeyPressed(ebiten.Key5) {
		SetNumberKeyPress(5)
	}
	if ebiten.IsKeyPressed(ebiten.Key6) {
		SetNumberKeyPress(6)
	}
	if ebiten.IsKeyPressed(ebiten.Key7) {
		SetNumberKeyPress(7)
	}
	if ebiten.IsKeyPressed(ebiten.Key8) {
		SetNumberKeyPress(8)
	}
}
