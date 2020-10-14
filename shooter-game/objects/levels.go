package objects

import (
	"log"
	"math"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type level1 struct {
	img         *ebiten.Image // level image (water waves)
	imgW        int           // width of the image
	imgH        int           // height of the image
	offsetX     float64       // horizontal offset, used to animate the image
	offsetY     float64       // vertical offset, used to animate the image
	xDirection  direction     // horizontal direction of the animation
	yDirection  direction     // vertical direction of the animation
	ducks       []*duck       // current number of ducks in the screen
	maxDucks    int           // max number of ducks in the screen
	duckImgName string        // the name of the duck img
	score       *int64
	lastClick   time.Time
}

const (
	lvl1XSpeed     = 1    // horizontal speed of the animation
	lvl1YSpeed     = 0.35 // vertical speed of the animation
	lvl1MaxOffsetX = 100  // max horizontal movement
	lvl1MaxOffsetY = 16   // max vertical movement
)

func NewLevel1(imgName, duckImgName string, maxDucks int, score *int64) Object {
	img, err := utils.GetImage(imgName, assets.Stall)
	if err != nil {
		log.Fatalf("cannot get image %s: %v", imgName, err)
	}
	w, h := img.Size()

	return &level1{
		img:         img,
		imgW:        w,
		imgH:        h,
		xDirection:  right,
		yDirection:  down,
		maxDucks:    maxDucks,
		duckImgName: duckImgName,
		score:       score,
	}
}

func (l *level1) Update(trgt *ebiten.Image, tick uint) {
	// if the current number of ducks is below the expected number, maybe generate one
	if len(l.ducks) < l.maxDucks {
		// every second there's 30% possibilities to
		// generate a missing duck
		if tick%60 == 0 && rand.Float64() < 0.3 {
			l.ducks = append(l.ducks, newDuck(l.duckImgName, l.imgH+50))
		}
	}

	// Update the tick of the ducks and check if they're still
	// on screen, removing from the list if not
	// Note: as we're playing with a slice while looping over
	// it, we use an external n counter and at the end of
	// the loop we reduce the slice to the final lenght
	// https://github.com/golang/go/wiki/SliceTricks#filter-in-place
	n := 0
	var hit bool
	clicked := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	for _, d := range l.ducks {
		d.Update(trgt, tick)
		// calculate whether the duck has been hit
		if clicked && l.isClickValid() {
			x, y := ebiten.CursorPosition()
			if d.shoot(x, y) {
				hit = true
			}
		}

		if d.onScreen {
			l.ducks[n] = d
			n++
		}
	}
	l.ducks = l.ducks[:n]

	// Increment/decrement score
	currentScore := atomic.LoadInt64(l.score)
	if clicked && hit {
		atomic.StoreInt64(l.score, currentScore+10)
	} else if clicked && len(l.ducks) > 0 && currentScore >= 5 {
		atomic.StoreInt64(l.score, currentScore-5)
	}

	// Calculate the horizontal offset of the image.
	// First the direction:
	if l.offsetX >= lvl1MaxOffsetX {
		l.xDirection = l.xDirection.invert()
	} else if l.offsetX <= 0 {
		l.xDirection = right
	}
	// Then the actual calculation
	l.offsetX = l.offsetX + float64(l.xDirection)*lvl1XSpeed

	// Same for vertical animation
	if l.offsetY >= lvl1MaxOffsetY {
		l.yDirection = up
	} else if l.offsetY <= 0 {
		l.yDirection = down
	}
	l.offsetY = l.offsetY + float64(l.yDirection)*lvl1YSpeed
}

func (l *level1) Draw(trgt *ebiten.Image) error {
	// Draw the ducks before the water because they must be below it
	for _, d := range l.ducks {
		d.Draw(trgt)
	}

	trgtW, trgtH := trgt.Size()
	// x is the number of images to draw horizontally to fill in the whole screen
	x := int(math.Ceil(float64(trgtW) / float64(l.imgW)))
	// the loop starts at -1 to add an additional element
	// out of the screen on the left, that will become visible
	// during the horizontal movement
	for i := -1; i < x; i++ {
		op := &ebiten.DrawImageOptions{}
		// horizontal offset of the image, we're using multiple images to fill in the screen
		tx := i * l.imgW
		// vertically we move the image at the bottom of the screen
		ty := trgtH - l.imgH
		op.GeoM.Translate(float64(tx), float64(ty))
		// apply offset to animate the image
		op.GeoM.Translate(l.offsetX, l.offsetY)

		trgt.DrawImage(l.img, op)
	}

	return nil
}

func (l *level1) OnScreen() bool {
	return true
}

// Don't accept duplicated clicks
func (l *level1) isClickValid() bool {
	now := time.Now()
	valid := now.Sub(l.lastClick) > 200*time.Millisecond
	l.lastClick = now
	return valid
}
