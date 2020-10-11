package objects

import (
	"log"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

const (
	duckName        = "duck_outline_target_white.png"
	ducksXSpeed     = 1.5 // horizontal speed
	ducksYSpeed     = 0.6 // vertical speed
	ducksMaxOffsetY = 16  // max vertical movement for animation
)

type duck struct {
	img            *ebiten.Image
	offsetX        float64   // horizontal position
	offsetY        float64   // vertical position
	initialOffsetY float64   // initial vertical position. set by the caller
	yDirection     direction // whether the vertical animation is going up or down
	onScreen       bool      // true when the image is visible in the screen
}

// newDuck generates a new duck with an initial vertical position
func newDuck(initialOffsetY int) *duck {
	img, err := utils.GetImage(duckName, assets.Objects)
	if err != nil {
		log.Fatalf("drawing %s: %v", duckName, err)
	}

	w, _ := img.Size()

	return &duck{
		img:            img,
		initialOffsetY: float64(initialOffsetY),
		offsetX:        float64(-w),
		offsetY:        0,
		yDirection:     down,
		onScreen:       true,
	}
}

func (d *duck) Tick(screen *ebiten.Image, _ uint) {
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
