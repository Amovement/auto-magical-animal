package consts

const (
	GameWidth  = 640
	GameHeight = 640

	// Game resource address

	HomeImage    = "https://s3.bmp.ovh/imgs/2024/08/02/3ed11e6d363cfbc4.png"
	MonsterImage = "https://s3.bmp.ovh/imgs/2024/08/02/21d9e687d9cff87d.png"
	BulletImage  = "https://s3.bmp.ovh/imgs/2024/08/05/28af2d2a0b3f732f.png"
	CursorImage  = "https://s3.bmp.ovh/imgs/2024/08/05/2404d029c1889ec3.png"

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
)

// Animal resource
const (
	AnimalTypeCat  = 0
	AnimalImageCat = "https://s3.bmp.ovh/imgs/2024/08/05/2e5b945128128d57.png"
)
