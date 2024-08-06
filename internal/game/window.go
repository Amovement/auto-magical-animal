package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
)

func SetGameWindow(title string) {
	ebiten.SetWindowSize(consts.GameWidth, consts.GameHeight)
	ebiten.SetWindowTitle(title)
	icon, err := ebitenutil.NewImageFromURL(consts.WindowsIconImage)
	if err != nil {
		log.Panic(err)
	}
	ebiten.SetWindowIcon([]image.Image{icon.SubImage(icon.Bounds())})
}
