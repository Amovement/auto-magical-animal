package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
	"time"
)

type MonstersContainer struct {
	monsters []*Monster
	image    *ebiten.Image
}

type Monster struct {
	locateX     float64
	locateY     float64
	healthPoint int
}

var (
	monsterImage        *ebiten.Image
	errLoadMonsterImage error
)

func init() {
	monsterImage, errLoadMonsterImage = ebitenutil.NewImageFromURL(consts.MonsterImage)
	if errLoadMonsterImage != nil {
		log.Panic(errLoadMonsterImage)
	}
}

func NewMonstersContainer() *MonstersContainer {
	return &MonstersContainer{
		monsters: []*Monster{},
		image:    monsterImage,
	}
}

// Draw monsters
func (m *MonstersContainer) Draw(screen *ebiten.Image) {
	for _, monster := range m.monsters {
		option := &ebiten.DrawImageOptions{}
		option.GeoM.Translate(monster.locateX, monster.locateY)

		// Draw monsters
		screen.DrawImage(
			m.image,
			option,
		)
	}
}

// CreateMonster create a monster every second
//
//	monster hp will increase every round
func (m *MonstersContainer) CreateMonster() {
	// create a monster every second
	if tick == 0 {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(4) // used to check the direction of the monster
		randHealthPoint := rand.Intn(10)
		initHealthPoint := 1
		monsterHealthPoint := initHealthPoint + tickRounds + randHealthPoint
		if randNum == 0 {
			// (0, y)
			m.monsters = append(m.monsters, &Monster{
				healthPoint: monsterHealthPoint,
				locateX:     -consts.SmallUnitPx,
				locateY:     float64(rand.Intn(consts.GameHeight)),
			})
		} else if randNum == 1 {
			// (x, 0)
			m.monsters = append(m.monsters, &Monster{
				healthPoint: monsterHealthPoint,
				locateX:     float64(rand.Intn(consts.GameWidth)),
				locateY:     -consts.SmallUnitPx,
			})
		} else if randNum == 2 {
			// (x, GameHeight)
			m.monsters = append(m.monsters, &Monster{
				healthPoint: monsterHealthPoint,
				locateX:     float64(rand.Intn(consts.GameWidth)),
				locateY:     float64(consts.GameHeight + consts.SmallUnitPx),
			})
		} else if randNum == 3 {
			// (GameWidth, y)
			m.monsters = append(m.monsters, &Monster{
				healthPoint: monsterHealthPoint,
				locateX:     float64(consts.GameWidth + consts.SmallUnitPx),
				locateY:     float64(rand.Intn(consts.GameHeight)),
			})
		}
	}
}
