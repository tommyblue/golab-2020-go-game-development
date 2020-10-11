package objects

import (
	"fmt"
	"image/color"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type desk struct {
	name string
}

func NewDesk(img string) Object {
	return &desk{
		name: img,
	}
}

func (d *desk) Tick(_ *ebiten.Image, _ uint) {}

func (d *desk) Draw(trgt *ebiten.Image) error {
	trgtW, trgtH := trgt.Size()
	borderH := 4
	deskH := 130

	border, _ := ebiten.NewImage(trgtW, borderH, ebiten.FilterDefault)
	border.Fill(color.RGBA{0x80, 0x57, 0x2e, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(trgtH-borderH-deskH))
	trgt.DrawImage(border, op)

	deskBg, err := utils.GetImage(d.name, assets.Stall)
	if err != nil {
		return fmt.Errorf("drawing %s: %v", d.name, err)
	}
	deskW, _ := deskBg.Size()

	x := int(math.Ceil(float64(trgtW) / float64(deskW)))
	for i := 0; i < x; i++ {
		tx := i * deskW
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tx), float64(trgtH-deskH))
		trgt.DrawImage(deskBg, op)
	}

	return nil
}

func (d *desk) OnScreen() bool {
	return true
}
