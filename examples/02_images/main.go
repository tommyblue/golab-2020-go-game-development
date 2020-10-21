package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

var coin *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(coinImg))
	if err != nil {
		log.Fatal(err)
	}
	coin, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// Step 1: move
	// (0,0) is the top-left corner
	// X moves horizontally and Y vertically
	// op.GeoM.Translate(float64(100), float64(100))

	// Step 2: move to the center
	w, h := coin.Size()
	x, y := screen.Size()
	tx := x/2 - w/2
	ty := y/2 - h/2

	op.GeoM.Translate(float64(tx), float64(ty))
	if err := screen.DrawImage(coin, op); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Draw image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
