package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	world             *World
	monstersContainer *MonstersContainer
	Home              *Home
	pixels            []byte
}

func (g *Game) Update() error {
	g.world.Update()
	g.monstersContainer.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, consts.GameWidth*consts.GameHeight*4)
	}
	g.world.Draw(screen)
	g.Home.Draw(screen)
	g.monstersContainer.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return consts.GameWidth, consts.GameHeight
}

func NewGame() *Game {
	world := NewWorld(consts.GameWidth, consts.GameHeight)
	home := NewHome()
	return &Game{
		world:             world,
		Home:              home,
		monstersContainer: NewMonstersContainer(),
	}
}

func StartGame() {
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
