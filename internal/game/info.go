package game

import (
	"bytes"
	"fmt"
	"github.com/Amovement/auto-magical-animal/assets"
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// InfoContainer is a module used to display global information within the game
//
//	like score and maxScore and game status and some other information
//	Only painting is included, but regardless of updates to these attributes, updates are the responsibility of other struct
type InfoContainer struct {
	boxImage  *ebiten.Image
	infoBoxes []*Info
	boxHeight float64
	boxWidth  float64
}

type Info struct {
	message      string
	durationTick int
}

var (
	infoContainerImage        *ebiten.Image
	errLoadInfoContainerImage error
)

func init() {
	infoContainerImage, _, errLoadInfoContainerImage = ebitenutil.NewImageFromReader(bytes.NewReader(assets.InfoBoxImageBytes))
	if errLoadInfoContainerImage != nil {
		log.Panic(errLoadInfoContainerImage)
	}
}

func NewInfoContainer() *InfoContainer {
	return &InfoContainer{
		boxImage:  infoContainerImage,
		infoBoxes: []*Info{},
		boxWidth:  336,
		boxHeight: 145,
	}
}

func (i *InfoContainer) UpdateAndClearExpiredInfo() {
	var newInfoBoxes []*Info
	for index := range i.infoBoxes {
		i.infoBoxes[index].durationTick--
		if i.infoBoxes[index].durationTick > 0 {
			newInfoBoxes = append(newInfoBoxes, i.infoBoxes[index])
		}
	}
	i.infoBoxes = newInfoBoxes
}

func (i *InfoContainer) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Your Score: "+fmt.Sprint(score), 0, 40)
	ebitenutil.DebugPrintAt(screen, "History Max Score: "+fmt.Sprint(maxScore), 0, 60)
	ebitenutil.DebugPrintAt(screen, "Boss coming after: "+fmt.Sprint(bossLastCreatedTickRound+90-tickRounds), 0, 80)

	// show info boxes
	if len(i.infoBoxes) > 0 {
		option := &ebiten.DrawImageOptions{}
		option.GeoM.Translate((consts.GameWidth-i.boxWidth)/2, consts.GameHeight-i.boxHeight-100)
		option.ColorScale.SetA(0.5)
		screen.DrawImage(
			i.boxImage,
			option,
		)
	}
	for ind, info := range i.infoBoxes {
		ebitenutil.DebugPrintAt(screen, info.message, int(consts.GameWidth-i.boxWidth)/2+50, int(consts.GameHeight-i.boxHeight-45)+ind*15)
	}
}

// InitGameTips When the game starts, the game tips are displayed
func (i *InfoContainer) InitGameTips() {
	i.infoBoxes = append(i.infoBoxes, []*Info{
		{
			message:      "Use `W` `A` `S` `D` to control your focus.",
			durationTick: 60 * 15,
		},
		{
			message:      "Left mouse button is also ok.",
			durationTick: 60 * 15,
		},
		{
			message:      "Then press `Space key` to place an animal.",
			durationTick: 60 * 15,
		},
		{
			message:      "This is going to cost 20 points.",
			durationTick: 60 * 15,
		},
	}...)
}

// GameOverTips When the game ends, the game tips are displayed
func (i *InfoContainer) GameOverTips() {
	i.infoBoxes = append(i.infoBoxes, []*Info{
		{
			message:      "Game Over! Press `r` to restart.",
			durationTick: 1,
		},
	}...)
}
