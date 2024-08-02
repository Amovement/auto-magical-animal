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
	locateX float64
	locateY float64
}

func NewMonstersContainer() *MonstersContainer {
	monsterImage, err := ebitenutil.NewImageFromURL(consts.MonsterImage)
	if err != nil {
		log.Panic(err)
	}
	return &MonstersContainer{
		monsters: []*Monster{},
		image:    monsterImage,
	}
}

// Draw 绘制怪物
func (m *MonstersContainer) Draw(screen *ebiten.Image) {
	for _, monster := range m.monsters {
		option := &ebiten.DrawImageOptions{}
		option.GeoM.Translate(monster.locateX, monster.locateY)

		// 绘制怪物
		screen.DrawImage(
			m.image,
			option,
		)
	}
}

// Update game state by one tick.
func (m *MonstersContainer) Update() {
	if tick == 0 {
		// 随机生产一只怪物
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(2)
		if randNum == 0 {
			// (0, y)
			m.monsters = append(m.monsters, &Monster{
				locateX: 0,
				locateY: float64(rand.Intn(consts.GameHeight)),
			})
		} else {
			// (x, 0)
			m.monsters = append(m.monsters, &Monster{
				locateX: float64(rand.Intn(consts.GameWidth)),
				locateY: 0,
			})
		}
	}
}
