package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	cursor                 *Cursor
	monstersContainer      *MonstersContainer
	animalsContainer       *AnimalsContainer
	home                   *Home
	weather                *Weather
	bulletPresentContainer *BulletPresentContainer
	framesLogicContainer   *FramesLogicContainer
	infoContainer          *InfoContainer
	animationContainer     *AnimationContainer
}

func (g *Game) Update() error {
	// Clear expired info
	g.infoContainer.UpdateAndClearExpiredInfo()

	if gameStatus == consts.GameStatusPause || gameStatus == consts.GameStatusEnd {
		// Monitor whether the user presses the space key when the game pauses or ends
		if gameStatus == consts.GameStatusEnd {
			g.infoContainer.GameOverTips()
		}
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			// restart game
			RestartGame(g)
		} else {
			return nil
		}
	}
	// Animation update
	g.animationContainer.UpdateAndClearExpiredAnimation()
	// Cursor update
	g.cursor.ListenMouseEvent()
	// Weather update
	g.weather.Update()
	// home and animals fire a bullet
	g.home.Update()
	// Create monster
	g.monstersContainer.CreateMonster()
	// Game frames logic update
	g.framesLogicContainer.Update(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.weather.Draw(screen)
	g.home.Draw(screen)
	g.animalsContainer.Draw(screen)
	g.monstersContainer.Draw(screen)
	g.animationContainer.Draw(screen)
	g.bulletPresentContainer.Draw(screen)
	g.cursor.Draw(screen)
	g.infoContainer.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return consts.GameWidth, consts.GameHeight
}

func NewGame() *Game {
	return &Game{
		weather:                NewWeather(),
		home:                   NewHome(),
		animalsContainer:       NewAnimalsContainer(),
		monstersContainer:      NewMonstersContainer(),
		bulletPresentContainer: NewBulletPresentContainer(),
		framesLogicContainer:   NewFramesLogicContainer(),
		cursor:                 NewCursor(),
		infoContainer:          NewInfoContainer(),
		animationContainer:     NewAnimationContainer(),
	}
}

func StartGame() {
	game := NewGame()
	game.infoContainer.InitGameTips()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func RestartGame(g *Game) {
	g.animalsContainer = NewAnimalsContainer()
	g.monstersContainer = NewMonstersContainer()
	g.bulletPresentContainer = NewBulletPresentContainer()
	g.infoContainer = NewInfoContainer()
	g.animationContainer = NewAnimationContainer()
	RestartGlobal()
	ClearBulletVector()
	ClearAnimalVector()
	ClearInfoVector()
	ClearMonsterVector()
	g.infoContainer.InitGameTips()
}
