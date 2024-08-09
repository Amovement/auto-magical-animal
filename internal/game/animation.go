package game

import (
	"bytes"
	"github.com/Amovement/auto-magical-animal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type AnimationContainer struct {
	animationBoxes []*Animation
}

type Animation struct {
	image        *ebiten.Image
	durationTick int
	locationX    float64
	locationY    float64
}

var (
	animationImages []*ebiten.Image
)

func init() {
	for i := 0; i < 1; i++ {
		newImage, _, errLoad := ebitenutil.NewImageFromReader(bytes.NewReader(assets.AnimationImagesBytes[i]))
		if errLoad != nil {
			log.Panic(errLoad)
		}
		animationImages = append(animationImages, newImage)
	}
}

func NewAnimationContainer() *AnimationContainer {
	return &AnimationContainer{
		animationBoxes: []*Animation{},
	}
}

func (an *AnimationContainer) UpdateAndClearExpiredAnimation() {
	var newAnimationBoxes []*Animation
	for index := range an.animationBoxes {
		an.animationBoxes[index].durationTick--
		if an.animationBoxes[index].durationTick > 0 {
			newAnimationBoxes = append(newAnimationBoxes, an.animationBoxes[index])
		}
	}
	an.animationBoxes = newAnimationBoxes
}

func (an *AnimationContainer) Draw(screen *ebiten.Image) {
	// show animations
	for _, animation := range an.animationBoxes {
		option := &ebiten.DrawImageOptions{}
		option.GeoM.Translate(animation.locationX, animation.locationY)
		screen.DrawImage(
			animation.image,
			option,
		)
	}
}

func NewAnimation(image *ebiten.Image, durationTick int, locationX, locationY float64) *Animation {
	return &Animation{
		image:        image,
		durationTick: durationTick,
		locationX:    locationX,
		locationY:    locationY,
	}
}

func (an *AnimationContainer) AddAnimation(animationType int, durationTick int, locationX, locationY float64) {
	an.animationBoxes = append(an.animationBoxes, NewAnimation(animationImages[animationType], durationTick, locationX, locationY))
}
