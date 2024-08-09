package game

import (
	"bytes"
	"github.com/Amovement/auto-magical-animal/assets"
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
)

func SetGameWindow(title string) {
	ebiten.SetWindowSize(consts.GameWidth, consts.GameHeight)
	ebiten.SetWindowTitle(title)
	icon, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.WindowsIconImageBytes))
	if err != nil {
		log.Panic(err)
	}
	ebiten.SetWindowIcon([]image.Image{icon.SubImage(icon.Bounds())})
}
