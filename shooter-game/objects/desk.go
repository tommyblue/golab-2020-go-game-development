package objects

import (
	"image/color"
	"log"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type desk struct {
	img *ebiten.Image
	w   float64
	h   float64
}

func NewDesk(imgName string) Object {
	img, err := utils.GetImage(imgName, assets.Stall)
	if err != nil {
		log.Fatalf("drawing %s: %v", imgName, err)
	}
	w, h := img.Size()

	return &desk{
		img: img,
		w:   float64(w),
		h:   float64(h),
	}
}

func (d *desk) Update(_ *ebiten.Image, _ uint) {}

func (d *desk) Draw(trgt *ebiten.Image) error {
	trgtW, trgtH := trgt.Size()
	borderH := 4
	deskH := 130

	border, _ := ebiten.NewImage(trgtW, borderH, ebiten.FilterDefault)
	border.Fill(color.RGBA{0x80, 0x57, 0x2e, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(trgtH-borderH-deskH))
	trgt.DrawImage(border, op)

	x := math.Ceil(float64(trgtW) / d.w)
	var i float64
	for i < x {
		tx := i * d.w
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(tx, float64(trgtH-deskH))
		trgt.DrawImage(d.img, op)
		i++
	}

	return nil
}

func (d *desk) OnScreen() bool {
	return true
}
