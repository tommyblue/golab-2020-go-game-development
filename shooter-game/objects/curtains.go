package objects

import (
	"fmt"
	"log"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type curtains struct {
	top     string
	lateral string
}

func NewCurtains(top, lateral string) Object {
	return &curtains{
		top:     top,
		lateral: lateral,
	}
}

func (c *curtains) Update(_ *ebiten.Image, _ uint) {}

func (c *curtains) Draw(trgt *ebiten.Image) error {
	img, err := utils.GetImage(c.lateral, assets.Stall)
	if err != nil {
		return fmt.Errorf("drawing %s: %v", c.lateral, err)
	}

	op := &ebiten.DrawImageOptions{}
	// move a bit down
	op.GeoM.Translate(0, 60)
	trgt.DrawImage(img, op)

	trgtW, _ := trgt.Size()
	curtainW, _ := img.Size()
	op = &ebiten.DrawImageOptions{}
	// flip curtain and move to the right the size of the image
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(curtainW), 0)
	// then move to the right of the trgt
	op.GeoM.Translate(float64(trgtW-curtainW), 0)
	op.GeoM.Translate(0, 60)
	trgt.DrawImage(img, op)

	topImg, err := utils.GetImage(c.top, assets.Stall)
	if err != nil {
		log.Fatalf("drawing %s: %v", c.top, err)
	}
	topImgW, _ := topImg.Size()
	x := int(math.Ceil(float64(trgtW) / float64(topImgW)))
	for i := 0; i < x; i++ {
		tx := i * topImgW
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tx), 0)
		trgt.DrawImage(topImg, op)
	}

	return nil
}

func (c *curtains) OnScreen() bool {
	return true
}
