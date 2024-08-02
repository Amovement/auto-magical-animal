package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type User struct {
	image   *ebiten.Image
	locateX int
	locateY int
	Size    int
}

func NewUser() *User {
	userImage, err := ebitenutil.NewImageFromURL(consts.UserImage)
	if err != nil {
		log.Panic(err)
	}
	return &User{
		locateX: consts.GameWidth / 2,
		locateY: consts.GameHeight / 2,
		Size:    10,
		image:   userImage,
	}
}

// Draw 绘制用户
func (u *User) Draw(screen *ebiten.Image) {
	option := &ebiten.DrawImageOptions{}
	option.GeoM.Translate(float64(u.locateX), float64(u.locateY))

	// 绘制用户
	screen.DrawImage(
		u.image,
		option,
	)
}

// Update game state by one tick.
func (u *User) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		u.locateY--
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		u.locateY++
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		u.locateX--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		u.locateX++
	}
	// 判断边界
	if u.locateX-u.Size < 0 {
		u.locateX = u.Size
	}
	if u.locateX+u.Size > consts.GameWidth {
		u.locateX = consts.GameWidth - u.Size
	}
	if u.locateY-u.Size < 0 {
		u.locateY = u.Size
	}
	if u.locateY+u.Size > consts.GameHeight {
		u.locateY = consts.GameHeight - u.Size
	}
}
