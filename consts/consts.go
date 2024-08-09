package consts

const (
	GameWidth  = 640
	GameHeight = 640

	// Game resource address

	HomeImage        = "https://s3.bmp.ovh/imgs/2024/08/02/3ed11e6d363cfbc4.png"
	BulletImage      = "https://s3.bmp.ovh/imgs/2024/08/07/b329884e9e562b10.png"
	CursorImage      = "https://s3.bmp.ovh/imgs/2024/08/05/2404d029c1889ec3.png"
	WindowsIconImage = "https://s3.bmp.ovh/imgs/2024/08/02/ced70c73b509f81f.png"
	InfoBoxImage     = "https://s3.bmp.ovh/imgs/2024/08/07/39040a86ce9ffa55.png"

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

var (
	MonsterImage = []string{
		"https://s3.bmp.ovh/imgs/2024/08/02/21d9e687d9cff87d.png", // ghost
		"https://s3.bmp.ovh/imgs/2024/08/08/416ff8227a235c8d.png", // Purple virus
		"https://s3.bmp.ovh/imgs/2024/08/08/33e24322451b3cec.png", // Zombie
		"https://s3.bmp.ovh/imgs/2024/08/08/c90f2f81cd55ce08.png", // Kappa 河童
		"https://s3.bmp.ovh/imgs/2024/08/08/ddefef04522f2060.png", // Boss UFO
	}
)

// weather
const (
	WeatherSunnyImage = "https://s3.bmp.ovh/imgs/2024/08/05/5fe75e67b3570194.png"
	WeatherRainImage  = "https://s3.bmp.ovh/imgs/2024/08/05/c534acd3e42f9fee.png"
	WeatherSnowImage  = "https://s3.bmp.ovh/imgs/2024/08/05/ea03ffc115168ef7.png"
	WeatherSunnyName  = "sunny"
	WeatherRainName   = "rain"
	WeatherSnowName   = "snow"
	WeatherSunnyType  = 0
	WeatherRainType   = 1
	WeatherSnowType   = 2

	BackgroundImageSunny = "https://s3.bmp.ovh/imgs/2024/08/06/e41999dc8211aead.png"
	BackgroundImageRain  = "https://s3.bmp.ovh/imgs/2024/08/06/75fc34bf5847e53d.png"
	BackgroundImageSnow  = "https://s3.bmp.ovh/imgs/2024/08/06/5d59c29c96f5dd4b.png"
)

// Animal resource
const (
	AnimalTypeByWeather = -1 // 随天气选择

	AnimalTypeCat  = 0
	AnimalImageCat = "https://s3.bmp.ovh/imgs/2024/08/05/2e5b945128128d57.png"

	AnimalTypeFish  = 1
	AnimalImageFish = "https://s3.bmp.ovh/imgs/2024/08/06/bfc20fd8665638c9.png"

	AnimalTypePenguin  = 2
	AnimalImagePenguin = "https://s3.bmp.ovh/imgs/2024/08/06/273e243c2c2cd1cf.png"
)

var (
	AnimationImage = []string{
		"https://s3.bmp.ovh/imgs/2024/08/08/be8f4d8c87cb5b34.png", // Poison animation
	}
)

const (
	AnimationTypePoison = 0
)
