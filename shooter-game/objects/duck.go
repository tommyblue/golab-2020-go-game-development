package objects

import (
	"log"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

const (
	ducksXSpeed     = 1.5 // horizontal speed
	ducksYSpeed     = 0.6 // vertical speed
	ducksMaxOffsetY = 16  // max vertical movement for animation
)

type duck struct {
	img            *ebiten.Image
	h              int
	w              int
	offsetX        float64   // horizontal position
	offsetY        float64   // vertical position
	initialOffsetY float64   // initial vertical position. set by the caller
	yDirection     direction // whether the vertical animation is going up or down
	onScreen       bool      // true when the image is visible in the screen
}

// newDuck generates a new duck with an initial vertical position
func newDuck(duckImgName string, initialOffsetY int) *duck {
	img, err := utils.GetImage(duckImgName, assets.Objects)
	if err != nil {
		log.Fatalf("drawing %s: %v", duckImgName, err)
	}

	w, h := img.Size()

	return &duck{
		img:            img,
		w:              w,
		h:              h,
		initialOffsetY: float64(initialOffsetY),
		offsetX:        float64(-w),
		offsetY:        0,
		yDirection:     down,
		onScreen:       true,
	}
}

func (d *duck) Update(screen *ebiten.Image, _ uint) {
	// horizontal movement
	d.offsetX = d.offsetX + ducksXSpeed

	// when the duck is over the screen size, it's no more visible
	screenW, _ := screen.Size()
	if d.offsetX > float64(screenW) {
		d.onScreen = false
	}

	// calculate the vertical direction and offset (for animation)
	if ducksMaxOffsetY-math.Abs(d.offsetY) < 0 {
		d.yDirection = d.yDirection.invert()
	}
	d.offsetY = d.offsetY + float64(d.yDirection)*ducksYSpeed

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		clickX, clickY := ebiten.CursorPosition()

		x := int(d.offsetX)
		y := int(d.offsetY + d.initialOffsetY)

		// Approximate the duck to its rectangle, though there're transparent
		// pixels. For better results we can either approximate the duck to other
		// shapes (like a rectangle+circle) or use image.At() to understand
		// if a transparent pixel was hit
		if clickX >= x && clickX <= x+d.w && clickY >= y && clickY <= y+d.h {
			d.onScreen = false
		}

	}
}

func (d *duck) Draw(trgt *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(d.offsetX, d.offsetY+d.initialOffsetY)
	trgt.DrawImage(d.img, op)
	return nil
}

func (d *duck) OnScreen() bool {
	return d.onScreen
}
