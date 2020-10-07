package objects

import (
	"fmt"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type level1 struct {
	name string
	tick uint
}

func NewLevel1(img string) Object {
	return &level1{
		name: img,
	}
}

func (l *level1) Tick(tick uint) {
	l.tick = tick
}

func (l *level1) Draw(trgt *ebiten.Image) error {
	trgtW, trgtH := trgt.Size()
	var offsetX float64 = 0
	var offsetY float64 = 0

	imgW1, err := utils.GetImage(l.name, assets.Stall)
	if err != nil {
		return fmt.Errorf("drawing %s: %v", l.name, err)
	}

	imgW1W, imgW1H := imgW1.Size()
	x := int(math.Ceil(float64(trgtW) / float64(imgW1W)))
	for i := 0; i < x; i++ {
		tx := i * imgW1W
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tx), float64(trgtH-imgW1H))
		op.GeoM.Translate(offsetX, offsetY)
		trgt.DrawImage(imgW1, op)
	}

	return nil
}
