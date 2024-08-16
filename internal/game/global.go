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
	// The monster queue generated under the current game frame
	monsterVector []*Monster
	// The info queue generated under the current game frame
	infoVector []*Info
	// Weather type look `weather.go` for details
	weatherType int
	// The score of the game. Score also used to buy animals.
	// Every 60 frames, score + 1.It means about 1s = 1 score
	score int
	// maxScore records highest score in history
	maxScore int
	// bossLastCreatedTickRound records the last tick round when boss was created
	bossLastCreatedTickRound int
	// numberKeyPress Record the last number key pressed
	numberKeyPress int
)

func init() {
	score = 20
	tick = 0
	tickRounds = 0
	gameStatus = 0
	bossLastCreatedTickRound = 0
	numberKeyPress = 1
	bulletVector = []*Bullet{}
	animalVector = []*Animal{}
	monsterVector = []*Monster{}
	infoVector = []*Info{}
}

func SetNumberKeyPress(num int) {
	numberKeyPress = num
}

// SetGameStatus Set game state
//
//	gameStatus 0 running, game 1 end, game 2 pause
func SetGameStatus(status int) {
	if status == consts.GameStatusEnd {
		if score > maxScore {
			maxScore = score
		}
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

func AppendBulletVector(bullet ...*Bullet) {
	bulletVector = append(bulletVector, bullet...)
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

func AppendInfoVector(info ...*Info) {
	infoVector = append(infoVector, info...)
}

func AppendMonsterVector(monster ...*Monster) {
	monsterVector = append(monsterVector, monster...)
}

func ClearMonsterVector() {
	monsterVector = []*Monster{}
}

func ClearInfoVector() {
	infoVector = []*Info{}
}

// RestartGlobal Restart game global variables without maxScore
func RestartGlobal() {
	ClearBulletVector()
	ClearAnimalVector()
	bossLastCreatedTickRound = 0
	tick = 0
	tickRounds = 0
	score = 20
	numberKeyPress = 1
	SetGameStatus(consts.GameStatusRunning)
}
