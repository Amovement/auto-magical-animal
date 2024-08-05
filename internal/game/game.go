package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	world                  *World
	cursor                 *Cursor
	monstersContainer      *MonstersContainer
	animalsContainer       *AnimalsContainer
	home                   *Home
	weather                *Weather
	bulletPresentContainer *BulletPresentContainer
	framesLogicContainer   *FramesLogicContainer
}

func (g *Game) Update() error {
	if gameStatus == consts.GameStatusPause || gameStatus == consts.GameStatusEnd {
		return nil
	}
	// Cursor update
	g.cursor.ListenMouseEvent()
	// Weather update
	g.weather.Update()
	// home and animals fire a bullet
	g.home.Update()
	g.animalsContainer.Update()
	// Create monster
	g.monstersContainer.CreateMonster()
	// Game frames logic update
	g.framesLogicContainer.Update(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.weather.Draw(screen)
	g.home.Draw(screen)
	g.animalsContainer.Draw(screen)
	g.monstersContainer.Draw(screen)
	g.bulletPresentContainer.Draw(screen)
	g.cursor.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return consts.GameWidth, consts.GameHeight
}

func NewGame() *Game {
	return &Game{
		world:                  NewWorld(consts.GameWidth, consts.GameHeight),
		weather:                NewWeather(),
		home:                   NewHome(),
		animalsContainer:       NewAnimalsContainer(),
		monstersContainer:      NewMonstersContainer(),
		bulletPresentContainer: NewBulletPresentContainer(),
		framesLogicContainer:   NewFramesLogicContainer(),
		cursor:                 NewCursor(),
	}
}

func StartGame() {
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
