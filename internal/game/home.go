package game

import (
	"bytes"
	"github.com/Amovement/auto-magical-animal/assets"
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var (
	homeLocateX float64
	homeLocateY float64
)

func init() {
	homeLocateX = float64(consts.GameWidth / 2)
	homeLocateY = float64(consts.GameHeight / 2)
}

type Home struct {
	image   *ebiten.Image
	locateX float64
	locateY float64
	// Attack interval
	attackInterval int
	// The frame of the last attack
	lastAttackTime int
	// home bullet damage
	bulletDamage int
	// home bullet speed
	bulletSpeed float64
}

func NewHome() *Home {
	homeImage, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.HomeImageBytes))
	if err != nil {
		log.Panic(err)
	}
	return &Home{
		locateX:        homeLocateX,
		locateY:        homeLocateY,
		image:          homeImage,
		attackInterval: 40,
		lastAttackTime: 0,
		bulletDamage:   10,
		bulletSpeed:    5,
	}
}

// Draw Home
func (h *Home) Draw(screen *ebiten.Image) {
	option := &ebiten.DrawImageOptions{}
	option.GeoM.Translate(h.locateX, h.locateY)

	screen.DrawImage(
		h.image,
		option,
	)
}

func (h *Home) Update() {
	var attackInterval int
	if h.lastAttackTime <= tick {
		attackInterval = tick - h.lastAttackTime
	} else {
		attackInterval = tick - h.lastAttackTime + consts.TimeInterval
	}
	// A bullet can be fired if the interval is greater than the attack interval
	if attackInterval >= h.attackInterval {
		h.lastAttackTime = tick
		// create a bullet
		bullet, err := NewBullet(h.locateX, h.locateY, consts.NegativeMaxInt, consts.NegativeMaxInt, h.bulletSpeed, h.bulletDamage)
		if err != nil {
			log.Panic(err)
		}
		// add to bullet vector
		AppendBulletVector(bullet)
	}
}
