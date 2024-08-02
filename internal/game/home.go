package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Home struct {
	image   *ebiten.Image
	locateX int
	locateY int
}

func NewHome() *Home {
	homeImage, err := ebitenutil.NewImageFromURL(consts.HomeImage)
	if err != nil {
		log.Panic(err)
	}
	return &Home{
		locateX: consts.GameWidth / 2,
		locateY: consts.GameHeight / 2,
		image:   homeImage,
	}
}

// Draw 绘制用户
func (h *Home) Draw(screen *ebiten.Image) {
	option := &ebiten.DrawImageOptions{}
	option.GeoM.Translate(float64(h.locateX), float64(h.locateY))

	// 绘制用户
	screen.DrawImage(
		h.image,
		option,
	)
}
