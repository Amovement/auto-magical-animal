package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Info is a module used to display global information within the game
//
//	like score and maxScore and game status and some other information
//	Only painting is included, but regardless of updates to these attributes, updates are the responsibility of other struct
type Info struct {
}

func NewInfo() *Info {
	return &Info{}
}

func (i *Info) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Your Score: "+fmt.Sprint(score), 0, 40)
	ebitenutil.DebugPrintAt(screen, "History Max Score: "+fmt.Sprint(maxScore), 0, 60)
}
