package objects

import (
	"fmt"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type background struct {
	name string
}

func NewBackground(img string) Object {
	return &background{
		name: img,
	}
}

func (b *background) Update(_ *ebiten.Image, tick uint) {}

func (b *background) Draw(trgt *ebiten.Image) error {
	img, err := utils.GetImage(b.name, assets.Stall)
	if err != nil {
		return fmt.Errorf("drawing %s: %v", b.name, err)
	}

	// as the background is smaller than the trgt, it must be
	// drawn multiple times. Let's calculate the numbers
	trgtW, trgtH := trgt.Size()
	bgW, bgH := img.Size()
	x := int(math.Ceil(float64(trgtW) / float64(bgW)))
	y := int(math.Ceil(float64(trgtH) / float64(bgH)))

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			op := &ebiten.DrawImageOptions{}
			tx := i * bgW
			ty := j * bgH
			op.GeoM.Translate(float64(tx), float64(ty))
			trgt.DrawImage(img, op)
		}
	}

	return nil
}

func (b *background) OnScreen() bool {
	return true
}
