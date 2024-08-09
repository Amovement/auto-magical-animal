package game

import (
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
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
}

type AnimalsContainer struct {
	animals []*Animal
}

var (
	errAnimalImage     error
	animalImageCat     *ebiten.Image
	animalImageFish    *ebiten.Image
	animalImagePenguin *ebiten.Image
)

func init() {
	animalImageCat, errAnimalImage = ebitenutil.NewImageFromURL(consts.AnimalImageCat)
	if errAnimalImage != nil {
		panic(errAnimalImage)
	}
	animalImageFish, errAnimalImage = ebitenutil.NewImageFromURL(consts.AnimalImageFish)
	if errAnimalImage != nil {
		panic(errAnimalImage)
	}
	animalImagePenguin, errAnimalImage = ebitenutil.NewImageFromURL(consts.AnimalImagePenguin)
	if errAnimalImage != nil {
		panic(errAnimalImage)
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
		animal.animalType = weatherType
	}
	switch animal.animalType {
	case consts.AnimalTypeCat:
		animal.image = animalImageCat
		animal.attackInterval = 30
		animal.bulletDamage = 10
		animal.bulletSpeed = 5
	case consts.AnimalTypeFish:
		animal.image = animalImageFish
		animal.attackInterval = 90
		animal.bulletDamage = 60
		animal.bulletSpeed = 5
	case consts.AnimalTypePenguin:
		animal.image = animalImagePenguin
		animal.attackInterval = 90
		animal.bulletDamage = 120
		animal.bulletSpeed = 5
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

// Update animals fire
func (ac *AnimalsContainer) Update() {
	for index, animal := range ac.animals {
		var attackInterval int
		if animal.lastAttackTickRound < tickRounds {
			attackInterval = tick + (consts.TimeInterval - animal.lastAttackTime) + consts.TimeInterval*(tickRounds-animal.lastAttackTickRound)
		} else {
			attackInterval = tick - animal.lastAttackTime
		}
		// A bullet can be fired if the interval is greater than the attack interval
		if attackInterval >= animal.attackInterval {
			ac.animals[index].lastAttackTime = tick
			ac.animals[index].lastAttackTickRound = tickRounds
			var bullets []*Bullet

			if animal.animalType == consts.AnimalTypeCat {
				// Cat fire bullets that automatically capture the enemy
				// create a bullet
				bullet, err := NewBullet(animal.locateX, animal.locateY, consts.NegativeMaxInt, consts.NegativeMaxInt, animal.bulletSpeed, animal.bulletDamage)
				if err != nil {
					log.Panic(err)
				}
				bullets = append(bullets, bullet)
			} else if animal.animalType == consts.AnimalTypeFish {
				// Fish fire bullets that four directions
				bullet, _ := NewBullet(animal.locateX, animal.locateY, animal.locateX, consts.GameHeight, animal.bulletSpeed, animal.bulletDamage)
				bullets = append(bullets, bullet)
				bullet, _ = NewBullet(animal.locateX, animal.locateY, animal.locateX, 0, animal.bulletSpeed, animal.bulletDamage)
				bullets = append(bullets, bullet)
				bullet, _ = NewBullet(animal.locateX, animal.locateY, 0, animal.locateY, animal.bulletSpeed, animal.bulletDamage)
				bullets = append(bullets, bullet)
				bullet, _ = NewBullet(animal.locateX, animal.locateY, consts.GameWidth, animal.locateY, animal.bulletSpeed, animal.bulletDamage)
				bullets = append(bullets, bullet)
			} else if animal.animalType == consts.AnimalTypePenguin {
				// penguin fire bullets around itself
				directionX := []float64{50, -50, 0}
				directionY := []float64{50, -50, 0}
				for _, dx := range directionX {
					for _, dy := range directionY {
						if dx == 0 && dy == 0 {
							continue
						}
						bullet, _ := NewBullet(animal.locateX, animal.locateY, animal.locateX+dx, animal.locateY+dy, animal.bulletSpeed, animal.bulletDamage)
						bullets = append(bullets, bullet)
					}
				}
			}

			// add to bullet vector
			AppendBulletVector(bullets...)
		}
	}
}
