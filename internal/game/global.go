package game

import "github.com/Amovement/auto-magical-animal/consts"

var (
	// tick In-game frame count statistics, 60 frames recorded as 1 second
	tick int
	// tickRounds indicates several tick cycles
	tickRounds int
	// gameStatus 0 running, game 1 end, game 2 pause
	gameStatus int
	// The bullet queue generated under the current game frame
	bulletVector []*Bullet
	// The animal queue generated under the current game frame
	animalVector []*Animal
	// Weather type look `weather.go` for details
	weatherType int
	// The score of the game. Score also used to buy animals.
	// Every 60 frames, score + 1.It means about 1s = 1 score
	score int
	// maxScore records highest score in history
	maxScore int
)

func init() {
	score = 0
	tick = 0
	tickRounds = 0
	gameStatus = 0
	bulletVector = []*Bullet{}
	animalVector = []*Animal{}
}

// SetGameStatus Set game state
//
//	gameStatus 0 running, game 1 end, game 2 pause
func SetGameStatus(status int) {
	if status == consts.GameStatusEnd {
		maxScore = score
	}
	gameStatus = status
}

// TickRunning Game frame count update
func TickRunning() {
	tick++
	if tick == consts.TimeInterval {
		tickRounds++
		tick = 0
		// Every 60 frames, score + 1.It means about 1s = 1 score
		score++
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

// RestartGlobal Restart game global variables without maxScore
func RestartGlobal() {
	ClearBulletVector()
	ClearAnimalVector()
	tick = 0
	tickRounds = 0
	score = 0
	SetGameStatus(consts.GameStatusRunning)
}
