package consts

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
	// MonsterTypePurpleVirus Deathrattle: Split into two small monsters, halving the maximum health points, less than 10 HP will die.
	//	紫色病毒 亡语: 分裂为两个小怪物, 减少最大血量一半, 不足 10 时彻底死亡
	MonsterTypePurpleVirus = 1
	// MonsterTypeZombie Deathrattle: Kill animal units around 50 px.
	//	僵尸 亡语: 击杀 50 px 附近的动物单位
	MonsterTypeZombie = 2
	MonsterTypeKappa  = 3
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
)

const (
	AnimationTypePoison = 0
)
