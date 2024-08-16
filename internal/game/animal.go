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
)

type Animal struct {
	image      *ebiten.Image
	locateX    float64
	locateY    float64
	animalType int // 0: cat 1: fish 2: penguin
	// Attack interval
	attackInterval int
	// The frame of the last attack
	lastAttackTime int
	// The tick round of the last attack
	lastAttackTickRound int
	// bullet damage
	bulletDamage int
	// bullet speed
	bulletSpeed float64
	healthPoint int
	moveSpeed   float64
}

type AnimalsContainer struct {
	animals []*Animal
}

var (
	animalImagesArr     []*ebiten.Image
	weatherAnimalsArray [][]int
)

func init() {
	for index := 0; index < 6; index++ {
		animalImage, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.AnimalImagesBytes[index]))
		if err != nil {
			log.Panic(err)
		}
		animalImagesArr = append(animalImagesArr, animalImage)
	}

	for index := 0; index < 3; index++ {
		weatherAnimalsArray = append(weatherAnimalsArray, []int{})
		if index == consts.WeatherSunnyType {
			weatherAnimalsArray[index] = append(weatherAnimalsArray[index], consts.AnimalTypeMonkey)
			weatherAnimalsArray[index] = append(weatherAnimalsArray[index], consts.AnimalTypeCat)
		}
		if index == consts.WeatherRainType {
			weatherAnimalsArray[index] = append(weatherAnimalsArray[index], consts.AnimalTypeFish)
			weatherAnimalsArray[index] = append(weatherAnimalsArray[index], consts.AnimalTypeCactus)
		}
		if index == consts.WeatherSnowType {
			weatherAnimalsArray[index] = append(weatherAnimalsArray[index], consts.AnimalTypePenguin)
			weatherAnimalsArray[index] = append(weatherAnimalsArray[index], consts.AnimalTypeHorse)
		}
	}
}

func NewAnimal(animalType int, locateX, locateY float64) *Animal {
	animal := &Animal{
		locateX:             locateX,
		locateY:             locateY,
		animalType:          animalType,
		lastAttackTime:      tick,
		lastAttackTickRound: tickRounds,
		healthPoint:         100,
	}
	if animal.animalType == consts.AnimalTypeByWeather {
		if len(weatherAnimalsArray[weatherType]) >= numberKeyPress && numberKeyPress > 0 {
			animal.animalType = weatherAnimalsArray[weatherType][numberKeyPress-1]
		} else {
			animal.animalType = weatherAnimalsArray[weatherType][0]
		}
	}
	switch animal.animalType {
	case consts.AnimalTypeCat:
		animal.image = animalImagesArr[consts.AnimalTypeCat]
		animal.attackInterval = 30
		animal.bulletDamage = 10
		animal.bulletSpeed = 5
	case consts.AnimalTypeFish:
		animal.image = animalImagesArr[consts.AnimalTypeFish]
		animal.attackInterval = 90
		animal.bulletDamage = 60
		animal.bulletSpeed = 5
	case consts.AnimalTypePenguin:
		animal.image = animalImagesArr[consts.AnimalTypePenguin]
		animal.attackInterval = 90
		animal.bulletDamage = 120
		animal.bulletSpeed = 5
	case consts.AnimalTypeCactus:
		animal.image = animalImagesArr[consts.AnimalTypeCactus]
		animal.attackInterval = 45
		animal.bulletDamage = 60
		animal.bulletSpeed = 5
	case consts.AnimalTypeMonkey:
		animal.image = animalImagesArr[consts.AnimalTypeMonkey]
		animal.attackInterval = 90
		animal.bulletDamage = 120
		randSpeed := rand.Float64() * 5
		animal.moveSpeed = randSpeed
		if animal.moveSpeed < 2 {
			animal.moveSpeed += 2
		}
		animal.bulletSpeed = 5
	case consts.AnimalTypeHorse:
		animal.image = animalImagesArr[consts.AnimalTypeHorse]
		animal.attackInterval = 10
		animal.bulletDamage = 10000
		animal.bulletSpeed = 5
		animal.moveSpeed = 2
	}
	return animal
}

func NewAnimalsContainer() *AnimalsContainer {
	return &AnimalsContainer{
		animals: []*Animal{},
	}
}

func (ac *AnimalsContainer) Draw(screen *ebiten.Image) {
	if ac.animals == nil {
		return
	}
	for _, animal := range ac.animals {
		option := &ebiten.DrawImageOptions{}
		option.GeoM.Translate(animal.locateX, animal.locateY)

		// Draw bullet
		screen.DrawImage(
			animal.image,
			option,
		)
	}
}

func (ani *Animal) SurvivalMove(game *Game) {
	if ani.animalType == consts.AnimalTypeMonkey || ani.animalType == consts.AnimalTypeHorse {
		monsterNearest := ani.findNearestMonster(game.monstersContainer.monsters)
		if monsterNearest == nil {
			return
		}
		// Calculate the distance traveled
		xDelta := monsterNearest.locateX - ani.locateX
		yDelta := monsterNearest.locateY - ani.locateY
		localDistance := math.Sqrt(xDelta*xDelta + yDelta*yDelta)
		moveRate := ani.moveSpeed
		moveX := (moveRate / localDistance) * xDelta
		moveY := (moveRate / localDistance) * yDelta
		ani.AnimalMove(moveX, moveY)
	}
}

func (ani *Animal) SurvivalSkill(game *Game) {
	var bullets []*Bullet
	if ani.animalType == consts.AnimalTypeCat {
		// Cat fire bullets that automatically capture the enemy
		// create a bullet
		bullet, err := NewBullet(ani.locateX, ani.locateY, consts.NegativeMaxInt, consts.NegativeMaxInt, ani.bulletSpeed, ani.bulletDamage)
		if err != nil {
			log.Panic(err)
		}
		bullets = append(bullets, bullet)
	} else if ani.animalType == consts.AnimalTypeFish {
		// Fish fire bullets that four directions
		bullet, _ := NewBullet(ani.locateX, ani.locateY, ani.locateX, consts.GameHeight, ani.bulletSpeed, ani.bulletDamage)
		bullets = append(bullets, bullet)
		bullet, _ = NewBullet(ani.locateX, ani.locateY, ani.locateX, 0, ani.bulletSpeed, ani.bulletDamage)
		bullets = append(bullets, bullet)
		bullet, _ = NewBullet(ani.locateX, ani.locateY, 0, ani.locateY, ani.bulletSpeed, ani.bulletDamage)
		bullets = append(bullets, bullet)
		bullet, _ = NewBullet(ani.locateX, ani.locateY, consts.GameWidth, ani.locateY, ani.bulletSpeed, ani.bulletDamage)
		bullets = append(bullets, bullet)
	} else if ani.animalType == consts.AnimalTypePenguin {
		// penguin fire bullets around itself
		directionX := []float64{3 * consts.SmallUnitPx, -3 * consts.SmallUnitPx, 0}
		directionY := []float64{3 * consts.SmallUnitPx, -3 * consts.SmallUnitPx, 0}
		for _, dx := range directionX {
			for _, dy := range directionY {
				if dx == 0 && dy == 0 {
					continue
				}
				bullet, _ := NewBullet(ani.locateX, ani.locateY, ani.locateX+dx, ani.locateY+dy, ani.bulletSpeed, ani.bulletDamage)
				bullets = append(bullets, bullet)
			}
		}
	} else if ani.animalType == consts.AnimalTypeCactus {
		// Cactus fire bullets around itself
		directionX := []float64{3 * consts.SmallUnitPx, -3 * consts.SmallUnitPx, 0}
		directionY := []float64{3 * consts.SmallUnitPx, -3 * consts.SmallUnitPx, 0}
		for _, dx := range directionX {
			for _, dy := range directionY {
				if dx == 0 && dy == 0 {
					continue
				}
				bullet, _ := NewBullet(ani.locateX, ani.locateY, ani.locateX+dx, ani.locateY+dy, ani.bulletSpeed, ani.bulletDamage)
				bullets = append(bullets, bullet)
			}
		}
	} else if ani.animalType == consts.AnimalTypeMonkey {
		// Cactus fire bullets around itself
		directionX := []float64{1.5 * consts.SmallUnitPx, -1.5 * consts.SmallUnitPx, 0}
		directionY := []float64{1.5 * consts.SmallUnitPx, -1.5 * consts.SmallUnitPx, 0}
		for _, dx := range directionX {
			for _, dy := range directionY {
				if dx == 0 && dy == 0 {
					continue
				}
				bullet, _ := NewBullet(ani.locateX, ani.locateY, ani.locateX+dx, ani.locateY+dy, ani.bulletSpeed, ani.bulletDamage)
				bullets = append(bullets, bullet)
			}
		}
	} else if ani.animalType == consts.AnimalTypeHorse {
		for j := 0; j < len(game.monstersContainer.monsters); j++ {
			monster := game.monstersContainer.monsters[j]
			if ani.checkCollision(monster, ani) {
				// hit
				monster.healthPoint -= ani.bulletDamage
				monster.SkillsWhenInjured(game)
				ani.healthPoint = 0
				game.animationContainer.AddAnimation(consts.AnimationTypeFire, consts.TimeInterval, monster.locateX, monster.locateY)
			}
		}
	}

	if len(bullets) > 0 {
		// add to bullet vector
		AppendBulletVector(bullets...)
	}
}

func (ani *Animal) AnimalMove(dx, dy float64) {
	ani.locateX += dx
	ani.locateY += dy
}

// findNearestMonster help someone find the nearest Monster
func (ani *Animal) findNearestMonster(monster []*Monster) *Monster {
	if len(monster) == 0 {
		return nil
	}
	minDistance := float64(consts.MaxInt)
	targetIndex := 0
	for index, mon := range monster {
		xDelta := math.Abs(mon.locateX - ani.locateX)
		yDelta := math.Abs(mon.locateY - ani.locateX)
		localDistance := math.Sqrt(xDelta*xDelta + yDelta*yDelta)
		if localDistance < minDistance {
			minDistance = localDistance
			targetIndex = index
		}
	}
	return monster[targetIndex]
}

// checkCollision Check if the animal hit the monster
func (ani *Animal) checkCollision(m *Monster, b *Animal) bool {
	monsterX1 := m.locateX - 10
	monsterY1 := m.locateY - 10
	monsterX2 := m.locateX + m.collisionSizeX + 10
	monsterY2 := m.locateY + m.collisionSizeY + 10
	bulletX1 := b.locateX
	bulletY1 := b.locateY
	if bulletX1 >= monsterX1 && bulletX1 <= monsterX2 && bulletY1 >= monsterY1 && bulletY1 <= monsterY2 {
		return true
	}
	return false
}

func (ani *Animal) Deathrattle(game *Game) {
	if ani.healthPoint > 0 {
		return
	}
	if ani.animalType == consts.AnimalTypeHorse {
		// Kill monster units around 150 px.
		for index := 0; index < len(game.monstersContainer.monsters); index++ {
			mons := game.monstersContainer.monsters[index]
			if math.Abs(mons.locateX-ani.locateX) <= 150 && math.Abs(mons.locateY-ani.locateY) <= 150 {
				mons.healthPoint = mons.healthPoint - ani.bulletDamage
			}
		}
	}

}
