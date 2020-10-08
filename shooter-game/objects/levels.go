package objects

import (
	"fmt"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type level1 struct {
	name       string
	tick       uint
	offsetX    float64
	offsetY    float64
	xSpeed     float64
	ySpeed     float64
	xDirection float64
	yDirection float64
	maxOffsetX float64
	maxOffsetY float64
}

func NewLevel1(img string) Object {
	return &level1{
		name:       img,
		xDirection: right,
		yDirection: down,
		xSpeed:     1,
		ySpeed:     0.35,
		maxOffsetX: 100,
		maxOffsetY: 16,
	}
}

func (l *level1) Tick(tick uint) {
	l.tick = tick
	if l.offsetX >= l.maxOffsetX {
		l.xDirection = left
	} else if l.offsetX <= 0 {
		l.xDirection = right
	}
	l.offsetX = l.offsetX + l.xDirection*l.xSpeed

	if l.offsetY >= l.maxOffsetY {
		l.yDirection = up
	} else if l.offsetY <= 0 {
		l.yDirection = down
	}
	l.offsetY = l.offsetY + l.yDirection*l.ySpeed
}

func (l *level1) Draw(trgt *ebiten.Image) error {
	trgtW, trgtH := trgt.Size()

	imgW1, err := utils.GetImage(l.name, assets.Stall)
	if err != nil {
		return fmt.Errorf("drawing %s: %v", l.name, err)
	}

	imgW1W, imgW1H := imgW1.Size()
	x := int(math.Ceil(float64(trgtW) / float64(imgW1W)))
	// the loop starts at -1 to add an additional element
	// out of the screen on the left, that will become visible
	// during the movement
	for i := -1; i < x; i++ {
		tx := i * imgW1W
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tx), float64(trgtH-imgW1H))
		op.GeoM.Translate(l.offsetX, l.offsetY)
		trgt.DrawImage(imgW1, op)
	}

	return nil
}
