package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
)

func SetGameWindow(title string) {
	ebiten.SetWindowSize(consts.GameWidth, consts.GameHeight)
	ebiten.SetWindowTitle(title)
}
