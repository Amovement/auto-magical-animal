package consts

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

const (
	GameWidth  = 640
	GameHeight = 640

	// TimeInterval 60 frames
	TimeInterval = 60

	// SmallUnitPx Indicates a small unit pixel
	SmallUnitPx = 25

	// GameStatusRunning Game status running
	GameStatusRunning = 0
	// GameStatusEnd Game status end
	GameStatusEnd = 1
	// GameStatusPause Game status pause
	GameStatusPause = 2

	NegativeMaxInt = -1 << 63
	MaxInt         = 1<<63 - 1
)

const (
	// MonsterTypeNormalGhost normal monster
	MonsterTypeNormalGhost = 0
	// MonsterTypePurpleVirus Deathrattle: Split into two small monsters
	//	紫色病毒 亡语: 分裂为两个小怪物
	MonsterTypePurpleVirus = 1
	// MonsterTypeZombie Deathrattle: Kill animal units around 50 px.
	//	僵尸 亡语: 击杀 50 px 附近的动物单位
	MonsterTypeZombie = 2
	// MonsterTypeKappa Kappa will use a hook to pull the nearest small animal. Kappa will restore 15 health when hit, but only the health of a normal monster.
	//	河童 使用钩子拉取最近的小动物 河童将会在受伤时恢复 15 点生命，但是生命值仅为普通怪物的血量
	MonsterTypeKappa = 3
	// MonsterTypeBossUFO Boss UFO every 3 tickRound generates a zombie elite monster MonsterTypeZombie
	//	Boss UFO 每 3 tickRound 生成一个精英僵尸怪物 MonsterTypeZombie
	MonsterTypeBossUFO = 4
)

// weather
const (
	WeatherSunnyName = "sunny"
	WeatherRainName  = "rain"
	WeatherSnowName  = "snow"
	WeatherSunnyType = 0
	WeatherRainType  = 1
	WeatherSnowType  = 2
)

// Animal resource
const (
	AnimalTypeByWeather = -1 // 随天气选择
	AnimalTypeCat       = 0
	AnimalTypeFish      = 1
	AnimalTypePenguin   = 2
	AnimalTypeCactus    = 3
	AnimalTypeMonkey    = 4
	AnimalTypeHorse     = 5
)

const (
	AnimationTypePoison     = 0
	AnimationTypeGreenHeart = 1
	AnimationTypeFire       = 2
)

var (
	WhiteImage    *ebiten.Image
	WhiteSubImage *ebiten.Image
	ColorRed      = color.RGBA{R: 255, G: 0, B: 0, A: 255}
)

func init() {
	WhiteImage = ebiten.NewImage(3, 3)
	WhiteImage.Fill(color.White)
	WhiteSubImage = WhiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
}
