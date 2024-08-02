package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const ()

type Game struct {
	world  *World
	user   *User
	pixels []byte
}

func (g *Game) Update() error {
	g.world.Update()
	g.user.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, consts.GameWidth*consts.GameHeight*4)
	}
	g.world.Draw(g.pixels)
	screen.WritePixels(g.pixels)

	g.user.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return consts.GameWidth, consts.GameHeight
}

func NewGame() *Game {
	world := NewWorld(consts.GameWidth, consts.GameHeight)
	return &Game{
		world: world,
		user:  NewUser(),
	}
}

func StartGame() {
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
