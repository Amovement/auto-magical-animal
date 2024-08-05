package game

import "github.com/Amovement/auto-magical-animal/consts"

var (
	// tick In-game frame count statistics, 60 frames recorded as 1 second
	tick int
	// gameStatus 0 running, game 1 end, game 2 pause
	gameStatus int
	// The bullet queue generated under the current game frame
	bulletVector []*Bullet
	// The animal queue generated under the current game frame
	animalVector []*Animal
)

func init() {
	tick = 0
	gameStatus = 0
	bulletVector = []*Bullet{}
	animalVector = []*Animal{}
}

// SetGameStatus Set game state
//
//	gameStatus 0 running, game 1 end, game 2 pause
func SetGameStatus(status int) {
	gameStatus = status
}

// TickRunning Game frame count update
func TickRunning() {
	tick++
	if tick == consts.TimeInterval {
		tick = 0
	}
}

func AppendBulletVector(bullet *Bullet) {
	bulletVector = append(bulletVector, bullet)
}

func ClearBulletVector() {
	bulletVector = []*Bullet{}
}

func AppendAnimalVector(animal *Animal) {
	animalVector = append(animalVector, animal)
}

func ClearAnimalVector() {
	animalVector = []*Animal{}
}
