package game

import (
	"fmt"
	"github.com/Amovement/auto-magical-animal/consts"
	"math"
)

// FramesLogicContainer Logical frame processor
//
//	Handle collisions, global parameters, etc
type FramesLogicContainer struct {
}

func NewFramesLogicContainer() *FramesLogicContainer {
	return &FramesLogicContainer{}
}

func (f *FramesLogicContainer) Update(game *Game) {
	//f.printGameLog(game)
	// Update current frame
	TickRunning()

	f.addBulletInVector2PresentContainer(game)
	f.addAnimalInVector2PresentContainer(game)
	f.addInfoInVector2InfoContainer(game)

	f.bulletsMove(game)
	f.monstersBulletsCollision(game)
	f.monsterMove(game)
}

func (f *FramesLogicContainer) monsterMove(game *Game) {
	// The monster moves Home
	for i := 0; i < len(game.monstersContainer.monsters); i++ {
		monster := game.monstersContainer.monsters[i]

		// Calculate the distance traveled
		xDelta := math.Abs(monster.locateX - homeLocateX)
		yDelta := math.Abs(monster.locateY - homeLocateY)
		localDistance := math.Sqrt(xDelta*xDelta + yDelta*yDelta)
		if localDistance <= 1 {
			// If the gap is less than 1, you have reached Home
			// game end
			SetGameStatus(consts.GameStatusEnd)
			return
		}
		moveRate := 1.0
		moveX := moveRate / localDistance * xDelta
		moveY := moveRate / localDistance * yDelta

		// Updated monster move to Home
		if monster.locateX < homeLocateX {
			monster.locateX += moveX
		}
		if monster.locateX > homeLocateX {
			monster.locateX -= moveX
		}
		if monster.locateY < homeLocateY {
			monster.locateY += moveY
		}
		if monster.locateY > homeLocateY {
			monster.locateY -= moveY
		}
	}
}

func (f *FramesLogicContainer) bulletsMove(game *Game) {
	for i := 0; i < len(game.bulletPresentContainer.bullets); i++ {
		bullet := game.bulletPresentContainer.bullets[i]

		// Calculate the distance traveled
		xDelta := math.Abs(bullet.locateX - bullet.TargetX)
		yDelta := math.Abs(bullet.locateY - bullet.TargetY)
		localDistance := math.Sqrt(xDelta*xDelta + yDelta*yDelta)
		if localDistance <= 1 {
			// If the gap is less than 1, bullet have reached target
		}
		moveRate := bullet.speed
		moveX := moveRate / localDistance * xDelta
		moveY := moveRate / localDistance * yDelta
		if bullet.locateX < bullet.TargetX {
			bullet.locateX += moveX
		}
		if bullet.locateX > bullet.TargetX {
			bullet.locateX -= moveX
		}
		if bullet.locateY < bullet.TargetY {
			bullet.locateY += moveY
		}
		if bullet.locateY > bullet.TargetY {
			bullet.locateY -= moveY
		}
	}
}

// monstersBulletsCollision Check if the bullet hit the monster
func (f *FramesLogicContainer) monstersBulletsCollision(game *Game) {
	// todo
	// A blocking algorithm can be added to optimize acceleration

	for i := 0; i < len(game.bulletPresentContainer.bullets); i++ {
		bullet := game.bulletPresentContainer.bullets[i]
		for j := 0; j < len(game.monstersContainer.monsters); j++ {
			monster := game.monstersContainer.monsters[j]
			xDelta := math.Abs(bullet.locateX - monster.locateX)
			yDelta := math.Abs(bullet.locateY - monster.locateY)
			if xDelta <= consts.SmallUnitPx && yDelta <= consts.SmallUnitPx {
				// hit
				monster.healthPoint -= bullet.damage
				// set this bullet speed to 0
				// then remove the bullet at a later stage
				bullet.speed = 0
			}
		}
		// check if bullet hit target
		if bullet.speed != 0 {
			xDelta := math.Abs(bullet.locateX - bullet.TargetX)
			yDelta := math.Abs(bullet.locateY - bullet.TargetY)
			if xDelta <= consts.SmallUnitPx && yDelta <= consts.SmallUnitPx {
				bullet.speed = 0
			}
		}
	}

	// remove dead monsters
	var newMonsters []*Monster
	for i := 0; i < len(game.monstersContainer.monsters); i++ {
		monster := game.monstersContainer.monsters[i]
		if monster.healthPoint > 0 {
			newMonsters = append(newMonsters, monster)
		} else if monster.healthPoint <= 0 {
			// Kill a monster will add score
			score++
		}
	}
	game.monstersContainer.monsters = newMonsters

	// remove bullets that speed is 0
	var newBullets []*Bullet
	for i := 0; i < len(game.bulletPresentContainer.bullets); i++ {
		bullet := game.bulletPresentContainer.bullets[i]
		if bullet.speed > 0 {
			newBullets = append(newBullets, bullet)
		}
	}
	game.bulletPresentContainer.bullets = newBullets
}

// addBulletInVector2PresentContainer Add bullets to the bullet container
func (f *FramesLogicContainer) addBulletInVector2PresentContainer(game *Game) {
	if game.bulletPresentContainer.bullets == nil {
		game.bulletPresentContainer.bullets = []*Bullet{}
	}
	for index, v := range bulletVector {
		if v.TargetX == consts.NegativeMaxInt && v.TargetY == consts.NegativeMaxInt {
			// bullet is no target, find the nearest monster
			targetX, targetY := f.findNearestMonster(v.locateX, v.locateY, game.monstersContainer.monsters)
			bulletVector[index].TargetX = targetX
			bulletVector[index].TargetY = targetY
		}
		game.bulletPresentContainer.bullets = append(game.bulletPresentContainer.bullets, bulletVector[index])
	}
	ClearBulletVector()
}

// findNearestMonster help the `bullet without a target` find the nearest monster
func (f *FramesLogicContainer) findNearestMonster(localX, localY float64, monsters []*Monster) (float64, float64) {
	minDistance := float64(consts.MaxInt)
	targetX, targetY := float64(consts.NegativeMaxInt), float64(consts.NegativeMaxInt)
	for _, monster := range monsters {
		xDelta := math.Abs(monster.locateX - localX)
		yDelta := math.Abs(monster.locateY - localY)
		localDistance := math.Sqrt(xDelta*xDelta + yDelta*yDelta)
		if localDistance < minDistance {
			minDistance = localDistance
			targetX = monster.locateX
			targetY = monster.locateY
		}
	}
	return targetX, targetY
}

// addAnimalInVector2PresentContainer Add animals to the animal present container
func (f *FramesLogicContainer) addAnimalInVector2PresentContainer(game *Game) {
	if game.animalsContainer.animals == nil {
		game.animalsContainer.animals = []*Animal{}
	}
	for index, _ := range animalVector {
		// Check for animals that get too close
		tooCloseTag := false
		for _, animalExisted := range game.animalsContainer.animals {
			xDelta := math.Abs(animalVector[index].locateX - animalExisted.locateX)
			yDelta := math.Abs(animalVector[index].locateY - animalExisted.locateY)
			if xDelta <= consts.SmallUnitPx && yDelta <= consts.SmallUnitPx {
				tooCloseTag = true
				break
			}
		}
		// If the animal is too close, skip
		if tooCloseTag {
			continue
		} else {
			if score >= 20 { // Check if the score is enough to buy an animal
				game.animalsContainer.animals = append(game.animalsContainer.animals, animalVector[index])
				score -= 20
			}
		}
	}
	ClearAnimalVector()
}

func (f *FramesLogicContainer) printGameLog(game *Game) {
	//for i := 0; i < len(game.bulletPresentContainer.bullets); i++ {
	//	fmt.Printf("bullets: %+v\n", game.bulletPresentContainer.bullets[i])
	//}
	for i := 0; i < len(bulletVector); i++ {
		fmt.Printf("bulletVector: %+v\n", bulletVector[i])
	}
}

func (f *FramesLogicContainer) addInfoInVector2InfoContainer(game *Game) {
	if game.infoContainer.infoBoxes == nil {
		game.infoContainer.infoBoxes = []*Info{}
	}
	for _, v := range infoVector {
		game.infoContainer.infoBoxes = append(game.infoContainer.infoBoxes, v)
	}
	ClearInfoVector()
}
