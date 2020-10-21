package objects

import (
	"log"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type crosshair struct {
	img        *ebiten.Image
	clickedImg *ebiten.Image
	w          float64
	h          float64
	clickedW   float64
	clickedH   float64
	x          int
	y          int
	clicked    bool
}

func NewCrosshair(imgName, clickedImgName string) Object {
	img, err := utils.GetImage(imgName, assets.Hud)
	if err != nil {
		log.Fatalf("drawing %s: %v", imgName, err)
	}
	w, h := img.Size()

	clickedImg, err := utils.GetImage(clickedImgName, assets.Hud)
	if err != nil {
		log.Fatalf("drawing %s: %v", clickedImgName, err)
	}
	clickedW, clickedH := clickedImg.Size()

	return &crosshair{
		img:        img,
		clickedImg: clickedImg,
		w:          float64(w),
		h:          float64(h),
		clickedW:   float64(clickedW),
		clickedH:   float64(clickedH),
	}
}

func (c *crosshair) Update(_ *ebiten.Image, _ uint) {
	c.x, c.y = ebiten.CursorPosition()
	c.clicked = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func (c *crosshair) Draw(trgt *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))
	if c.clicked {
		op.GeoM.Translate(-c.clickedW/2, -c.clickedH/2)
		trgt.DrawImage(c.clickedImg, op)
	} else {
		op.GeoM.Translate(-c.w/2, -c.h/2)
		trgt.DrawImage(c.img, op)
	}
	return nil
}

func (c *crosshair) OnScreen() bool {
	return true
}
