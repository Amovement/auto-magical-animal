package game

import (
	"bytes"
	"github.com/Amovement/auto-magical-animal/assets"
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
	"math/rand"
	"time"
)

type MonstersContainer struct {
	monsters []*Monster
}

type Monster struct {
	maxHealthPoint int
	healthPoint    int
	image          *ebiten.Image
	locateX        float64
	locateY        float64
	// comeFromX is the positive or negative x-coordinate of the monster spawn point with respect to home
	comeFromX int
	// comeFromY is the positive or negative y-coordinate of the monster spawn point with respect to home
	comeFromY      int
	monsterType    int
	speed          float64
	collisionSizeX float64
	collisionSizeY float64
}

var (
	monsterImage []*ebiten.Image
)

func init() {
	for i := 0; i < 5; i++ {
		newImage, _, errLoadMonsterImage := ebitenutil.NewImageFromReader(bytes.NewReader(assets.MonsterImagesBytes[i]))
		if errLoadMonsterImage != nil {
			log.Panic(errLoadMonsterImage)
		}
		monsterImage = append(monsterImage, newImage)
	}
}

func NewMonstersContainer() *MonstersContainer {
	return &MonstersContainer{
		monsters: []*Monster{},
	}
}

// Draw monsters
func (m *MonstersContainer) Draw(screen *ebiten.Image) {
	for _, monster := range m.monsters {
		option := &ebiten.DrawImageOptions{}
		option.GeoM.Translate(monster.locateX, monster.locateY)

		// Draw monsters
		screen.DrawImage(
			monster.image,
			option,
		)
	}
}

// CreateMonster create a monster every second
//
//	Monster hp will increase every tickRound
//	Now the monster hp is 1 + tickRounds + randHealthPoint
//	Monster type is random, 0: normal, 1 and 2: elite, 3: boss.
//	And it will 90% create a normal monster, 10% elite monster.
//	Every 60 tickRound create a boss.
func (m *MonstersContainer) CreateMonster() {
	// create a monster every second
	rand.Seed(time.Now().UnixNano())
	newMonster := &Monster{
		speed:          1.0,
		collisionSizeX: consts.SmallUnitPx,
		collisionSizeY: consts.SmallUnitPx,
	}

	// hp
	randHealthPoint := rand.Intn(10)
	initHealthPoint := 1
	monsterHealthPoint := initHealthPoint + tickRounds + randHealthPoint
	newMonster.maxHealthPoint = monsterHealthPoint
	newMonster.healthPoint = monsterHealthPoint
	// type
	monsterTypeRand := rand.Intn(100)
	var monsterTypeNow int
	if monsterTypeRand < 90 {
		// normal monster
		monsterTypeNow = consts.MonsterTypeNormalGhost
	} else if monsterTypeRand >= 90 && monsterTypeRand < 95 {
		// elite monster
		monsterTypeNow = consts.MonsterTypePurpleVirus
		// double hp
		newMonster.maxHealthPoint = monsterHealthPoint * 2
		newMonster.healthPoint = monsterHealthPoint * 2
	} else if monsterTypeRand >= 95 && monsterTypeRand < 100 {
		// elite monster
		monsterTypeNow = consts.MonsterTypeZombie
		// double hp
		newMonster.maxHealthPoint = monsterHealthPoint * 2
		newMonster.healthPoint = monsterHealthPoint * 2
	}
	// It's time to create a boss
	if bossLastCreatedTickRound+90-tickRounds == 0 {
		bossLastCreatedTickRound = tickRounds
		monsterTypeNow = consts.MonsterTypeBossUFO
		// hp * 10
		newMonster.maxHealthPoint = monsterHealthPoint * 10
		newMonster.healthPoint = monsterHealthPoint * 10
		newMonster.speed = 0.3
		// boss Size * 3
		newMonster.collisionSizeX = consts.SmallUnitPx * 3
		newMonster.collisionSizeY = consts.SmallUnitPx * 3
	}
	newMonster.monsterType = monsterTypeNow
	newMonster.image = monsterImage[newMonster.monsterType]

	// confirm direction
	randNum := rand.Intn(4) // used to check the direction of the monster
	if randNum == 0 {
		// (0, y)
		newMonster.locateX = -consts.SmallUnitPx
		newMonster.locateY = float64(rand.Intn(consts.GameHeight))
		newMonster.comeFromX = -1
		if newMonster.locateY > consts.GameHeight/2 {
			newMonster.comeFromY = 1
		} else {
			newMonster.comeFromY = -1
		}
	} else if randNum == 1 {
		// (x, 0)
		newMonster.locateX = float64(rand.Intn(consts.GameWidth))
		newMonster.locateY = -consts.SmallUnitPx
		newMonster.comeFromY = -1
		if newMonster.locateX > consts.GameWidth/2 {
			newMonster.comeFromX = 1
		} else {
			newMonster.comeFromX = -1
		}
	} else if randNum == 2 {
		// (x, GameHeight)
		newMonster.locateX = float64(rand.Intn(consts.GameWidth))
		newMonster.locateY = float64(consts.GameHeight + consts.SmallUnitPx)
		newMonster.comeFromY = 1
		if newMonster.locateX > consts.GameWidth/2 {
			newMonster.comeFromX = 1
		} else {
			newMonster.comeFromX = -1
		}
	} else if randNum == 3 {
		// (GameWidth, y)
		newMonster.locateX = float64(consts.GameWidth + consts.SmallUnitPx)
		newMonster.locateY = float64(rand.Intn(consts.GameHeight))
		newMonster.comeFromX = 1
		if newMonster.locateY > consts.GameHeight/2 {
			newMonster.comeFromY = 1
		} else {
			newMonster.comeFromY = -1
		}
	}

	// It's time to create a monster
	if tick == 0 {
		m.monsters = append(m.monsters, newMonster)
	}
}

// NewMonster return a monster struct
func NewMonster(monsterType int, maxHealthPoint, healthPoint int, speed, locateX, locateY float64, comeFromX, comeFromY int) *Monster {
	return &Monster{
		maxHealthPoint: maxHealthPoint,
		healthPoint:    healthPoint,
		image:          monsterImage[monsterType],
		locateX:        locateX,
		locateY:        locateY,
		comeFromX:      comeFromX,
		comeFromY:      comeFromY,
		monsterType:    monsterType,
		speed:          speed,
	}
}

// SurvivalSkill triggers when the monster is alive
func (m *Monster) SurvivalSkill(game *Game) {
	if m.healthPoint <= 0 {
		return
	}
	switch m.monsterType {
	case consts.MonsterTypeBossUFO:
		if tickRounds%3 == 0 && tick == 0 {
			// Boss UFO every 3 tickRound generates a zombie elite monster
			createdMonster := NewMonster(consts.MonsterTypeZombie, m.maxHealthPoint/10, tickRounds, 1,
				m.locateX-float64(consts.SmallUnitPx*m.comeFromX), m.locateY-float64(consts.SmallUnitPx*m.comeFromY), m.comeFromX, m.comeFromY)
			AppendMonsterVector(createdMonster)
		}
	}
}

// Deathrattle triggers when the monster dies
func (m *Monster) Deathrattle(game *Game) {
	if m.healthPoint > 0 {
		return
	}
	// Deathrattle
	if m.monsterType == consts.MonsterTypePurpleVirus {
		// MonsterTypePurpleVirus skill need to be implemented
		// Split into two small monsters, halving the maximum health points, less than 10 HP will die.
		newPurpleVirusHealthPoint := m.maxHealthPoint / 2
		revivedMonster := NewMonster(consts.MonsterTypeNormalGhost, m.maxHealthPoint, newPurpleVirusHealthPoint, m.speed*1.1,
			m.locateX+float64(m.comeFromX*consts.SmallUnitPx), m.locateY, m.comeFromX, m.comeFromY)
		AppendMonsterVector(revivedMonster)
		revivedMonster = NewMonster(consts.MonsterTypeNormalGhost, m.maxHealthPoint, newPurpleVirusHealthPoint, m.speed*1.1,
			m.locateX, m.locateY+float64(m.comeFromY*consts.SmallUnitPx), m.comeFromX, m.comeFromY)
		AppendMonsterVector(revivedMonster)
	} else if m.monsterType == consts.MonsterTypeZombie {
		// Kill animal units around 50 px.
		for indexAnimal := 0; indexAnimal < len(game.animalsContainer.animals); indexAnimal++ {
			animal := game.animalsContainer.animals[indexAnimal]
			if math.Abs(animal.locateX-m.locateX) <= consts.SmallUnitPx*2 && math.Abs(animal.locateY-m.locateY) <= consts.SmallUnitPx*2 {
				animal.healthPoint = 0
				game.animationContainer.AddAnimation(consts.AnimationTypePoison, 60*1.5, animal.locateX, animal.locateY)
			}
		}
	}
}

// MonsterAnimation shows monster animation every tickRounds
func (m *Monster) MonsterAnimation(game *Game) {
	if m.monsterType == consts.MonsterTypeZombie {
		if tickRounds%3 == 0 && tick == 0 {
			dxArr := []int{1, -1, 0}
			dyArr := []int{1, -1, 0}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					dx := dxArr[i] * consts.SmallUnitPx
					dy := dyArr[j] * consts.SmallUnitPx
					game.animationContainer.AddAnimation(consts.AnimationTypePoison, 60*1.5, m.locateX+float64(dx), m.locateY+float64(dy))
				}
			}
		}
	}

}
