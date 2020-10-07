package shooter

import (
	"log"
	"math"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/hajimehoshi/ebiten"
)

func (g *Game) drawBackground(screen *ebiten.Image) {
	name := "bg_wood.png"
	img, err := getImage(name, assets.Stall)
	if err != nil {
		log.Fatalf("drawing %s: %v", name, err)
	}

	// as the background is smaller than the screen, it must be
	// drawn multiple times. Let's calculate the numbers
	screenW, screenH := screen.Size()
	bgW, bgH := img.Size()
	x := int(math.Ceil(float64(screenW) / float64(bgW)))
	y := int(math.Ceil(float64(screenH) / float64(bgH)))

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			op := &ebiten.DrawImageOptions{}
			tx := i * bgW
			ty := j * bgH
			op.GeoM.Translate(float64(tx), float64(ty))
			screen.DrawImage(img, op)
		}
	}
}

func (g *Game) drawCurtains(screen *ebiten.Image) {
	name := "curtain.png"
	img, err := getImage(name, assets.Stall)
	if err != nil {
		log.Fatalf("drawing %s: %v", name, err)
	}

	op := &ebiten.DrawImageOptions{}
	// move a bit down
	op.GeoM.Translate(0, 60)
	screen.DrawImage(img, op)

	screenW, _ := screen.Size()
	curtainW, _ := img.Size()
	op = &ebiten.DrawImageOptions{}
	// flip curtain and move to the right the size of the image
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(curtainW), 0)
	// then move to the right of the screen
	op.GeoM.Translate(float64(screenW-curtainW), 0)
	op.GeoM.Translate(0, 60)
	screen.DrawImage(img, op)

	name = "curtain_straight.png"
	topImg, err := getImage(name, assets.Stall)
	if err != nil {
		log.Fatalf("drawing %s: %v", name, err)
	}
	topImgW, _ := topImg.Size()
	x := int(math.Ceil(float64(screenW) / float64(topImgW)))
	for i := 0; i < x; i++ {
		tx := i * topImgW
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tx), 0)
		screen.DrawImage(topImg, op)
	}

}
