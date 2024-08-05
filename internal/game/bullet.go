package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	locateX float64
	locateY float64
	TargetX float64
	TargetY float64
	image   *ebiten.Image
	speed   float64
	damage  int
}

// BulletPresentContainer indicates the number of bullets in the game
type BulletPresentContainer struct {
	bullets []*Bullet
}

var (
	errBulletImage error
	bulletImage    *ebiten.Image
)

func init() {
	bulletImage, errBulletImage = ebitenutil.NewImageFromURL(consts.BulletImage)
	if errBulletImage != nil {
		panic(errBulletImage)
	}
}

func NewBullet(locateX, locateY float64, targetX, targetY float64, speed float64, damage int) (*Bullet, error) {
	return &Bullet{
		image:   bulletImage,
		locateX: locateX,
		locateY: locateY,
		speed:   speed,
		damage:  damage,
		TargetX: targetX,
		TargetY: targetY,
	}, nil
}

func NewBulletPresentContainer() *BulletPresentContainer {
	return &BulletPresentContainer{
		bullets: []*Bullet{},
	}
}

func (bpc *BulletPresentContainer) Draw(screen *ebiten.Image) {
	if bpc.bullets == nil {
		return
	}
	for _, bullet := range bpc.bullets {
		option := &ebiten.DrawImageOptions{}
		option.GeoM.Translate(bullet.locateX, bullet.locateY)

		// Draw bullet
		screen.DrawImage(
			bullet.image,
			option,
		)
	}
}
