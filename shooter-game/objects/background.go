package objects

import (
	"log"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type background struct {
	img *ebiten.Image
	w   int
	h   int
}

func NewBackground(imgName string) Object {
	img, err := utils.GetImage(imgName, assets.Stall)
	if err != nil {
		log.Fatalf("drawing %s: %v", imgName, err)
	}
	w, h := img.Size()

	return &background{
		img: img,
		w:   w,
		h:   h,
	}
}

func (b *background) Update(_ *ebiten.Image, tick uint) {}

func (b *background) Draw(trgt *ebiten.Image) error {
	// as the background is smaller than the trgt, it must be
	// drawn multiple times. Let's calculate the numbers
	trgtW, trgtH := trgt.Size()
	x := int(math.Ceil(float64(trgtW) / float64(b.w)))
	y := int(math.Ceil(float64(trgtH) / float64(b.h)))

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			op := &ebiten.DrawImageOptions{}
			tx := i * b.w
			ty := j * b.h
			op.GeoM.Translate(float64(tx), float64(ty))
			trgt.DrawImage(b.img, op)
		}
	}

	return nil
}

func (b *background) OnScreen() bool {
	return true
}
