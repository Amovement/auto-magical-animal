package game

import (
	"bytes"
	"github.com/Amovement/auto-magical-animal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type AnimationContainer struct {
	vertices       []ebiten.Vertex
	indices        []uint16
	animationBoxes []*Animation
	Lines          []*Line
}

type Animation struct {
	image        *ebiten.Image
	durationTick int
	locationX    float64
	locationY    float64
}

type Line struct {
	c0x, c0y, c1x, c1y float32
	r                  float32
	color              color.Color
	durationTick       int
}

var (
	animationImages []*ebiten.Image
)

func init() {
	for i := 0; i < 3; i++ {
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

	var newLines []*Line
	for index := range an.Lines {
		an.Lines[index].durationTick--
		if an.Lines[index].durationTick > 0 {
			newLines = append(newLines, an.Lines[index])
		}
	}
	an.Lines = newLines
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

	for _, line := range an.Lines {
		vector.StrokeLine(screen, line.c0x, line.c0y, line.c1x, line.c1y, line.r, line.color, true)
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

func (an *AnimationContainer) drawLine(c0x, c0y, c1x, c1y, r float64, color color.Color, tick int) {
	an.Lines = append(an.Lines, &Line{
		c0x:          float32(c0x),
		c0y:          float32(c0y),
		c1x:          float32(c1x),
		c1y:          float32(c1y),
		r:            float32(r),
		color:        color,
		durationTick: tick,
	})
}
