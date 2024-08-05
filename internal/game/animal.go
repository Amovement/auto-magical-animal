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
	animalType int // 0: cat
	// Attack interval
	attackInterval int
	// The frame of the last attack
	lastAttackTime int
	// bullet damage
	bulletDamage int
	// bullet speed
	bulletSpeed float64
}

type AnimalsContainer struct {
	animals []*Animal
}

var (
	errAnimalImage error
	animalImageCat *ebiten.Image
)

func init() {
	animalImageCat, errAnimalImage = ebitenutil.NewImageFromURL(consts.AnimalImageCat)
	if errAnimalImage != nil {
		panic(errAnimalImage)
	}
}

func NewAnimal(animalType int, locateX, locateY float64) *Animal {
	animal := &Animal{
		locateX:        locateX,
		locateY:        locateY,
		animalType:     animalType,
		lastAttackTime: tick,
	}
	switch animalType {
	case consts.AnimalTypeCat:
		animal.image = animalImageCat
		animal.attackInterval = 40
		animal.bulletDamage = 20
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

func (ac *AnimalsContainer) Update() {
	for index, animal := range ac.animals {
		var attackInterval int
		if animal.lastAttackTime <= tick {
			attackInterval = tick - animal.lastAttackTime
		} else {
			attackInterval = tick - animal.lastAttackTime + consts.TimeInterval
		}
		// A bullet can be fired if the interval is greater than the attack interval
		if attackInterval >= animal.attackInterval {
			ac.animals[index].lastAttackTime = tick
			// create a bullet
			bullet, err := NewBullet(animal.locateX, animal.locateY, consts.NegativeMaxInt, consts.NegativeMaxInt, animal.bulletSpeed, animal.bulletDamage)
			if err != nil {
				log.Panic(err)
			}
			// add to bullet vector
			AppendBulletVector(bullet)
		}
	}
}
