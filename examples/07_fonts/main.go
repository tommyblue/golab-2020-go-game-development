package main

import (
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

const (
	sampleText = `Hello, Gophers!`
	dpi        = 72
	fontSize   = 36
)

type game struct{}

var mplusNormalFont font.Face

func (g *game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// calculate the rectangle containing the text
	bounds := text.BoundString(mplusNormalFont, sampleText)
	// write moving the text down by its height
	text.Draw(screen, sampleText, mplusNormalFont, 10, bounds.Dy(), color.White)
}

func (g *game) Layout(x, y int) (int, int) {
	return x, y
}

func init() {
	tt, err := truetype.Parse(Font)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("inputs")
	ebiten.RunGame(&game{})
}
