package game

import (
	"bytes"
	"github.com/Amovement/auto-magical-animal/assets"
	"github.com/Amovement/auto-magical-animal/consts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math/rand"
)

// Weather system
type Weather struct {
	image []*ebiten.Image
	// Weather type 0 Sunny 1 Rainy 2 snowy
	weatherType int
	// Weather name
	weatherName string
	// The weather duration indicates how many game frames have passed
	weatherDuration int
	// Background images
	backgroundImages []*ebiten.Image
}

func NewWeather() *Weather {
	var images, backgroundImages []*ebiten.Image
	weatherImage, _, _ := ebitenutil.NewImageFromReader(bytes.NewReader(assets.WeatherSunnyImageBytes))
	images = append(images, weatherImage)
	weatherImage, _, _ = ebitenutil.NewImageFromReader(bytes.NewReader(assets.BackgroundImageSunnyBytes))
	backgroundImages = append(backgroundImages, weatherImage)

	weatherImage, _, _ = ebitenutil.NewImageFromReader(bytes.NewReader(assets.WeatherRainImageBytes))
	images = append(images, weatherImage)
	weatherImage, _, _ = ebitenutil.NewImageFromReader(bytes.NewReader(assets.BackgroundImageRainBytes))
	backgroundImages = append(backgroundImages, weatherImage)

	weatherImage, _, _ = ebitenutil.NewImageFromReader(bytes.NewReader(assets.WeatherSnowImageBytes))
	images = append(images, weatherImage)
	weatherImage, _, _ = ebitenutil.NewImageFromReader(bytes.NewReader(assets.BackgroundImageSnowBytes))
	backgroundImages = append(backgroundImages, weatherImage)

	return &Weather{
		image:            images,
		backgroundImages: backgroundImages,
		weatherType:      consts.WeatherSunnyType, // Init to sunny
		weatherName:      consts.WeatherSunnyName,
		weatherDuration:  0,
	}
}

// Update Weather change
//
//	It takes at least 10 seconds for a change to occur
func (w *Weather) Update() {
	w.weatherDuration++
	// The weather can change once in 10 seconds
	if w.weatherDuration > 10*consts.TimeInterval {
		nxtWeatherType := rand.Intn(3)
		w.weatherType = nxtWeatherType
		switch w.weatherType {
		case consts.WeatherSunnyType:
			w.weatherName = consts.WeatherSunnyName
		case consts.WeatherRainType:
			w.weatherName = consts.WeatherRainName
		case consts.WeatherSnowType:
			w.weatherName = consts.WeatherSnowName
		}
		w.weatherDuration = 0
	}
	// set global weather type
	weatherType = w.weatherType
}

func (w *Weather) Draw(screen *ebiten.Image) {
	screen.DrawImage(w.backgroundImages[w.weatherType], nil)

	ebitenutil.DebugPrint(screen, "Current weather: "+w.weatherName)

	option := &ebiten.DrawImageOptions{}
	option.GeoM.Translate(0, 10)

	// draw weather
	screen.DrawImage(
		w.image[w.weatherType],
		option,
	)
}
