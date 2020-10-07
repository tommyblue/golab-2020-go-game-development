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
