package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the screen with red
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})

	// Create a smaller green image (screen/2) and draw it on top of the above
	w, h := screen.Size()
	i1, _ := ebiten.NewImage(w/2, h/2, ebiten.FilterDefault)
	i1.Fill(color.RGBA{0, 0xff, 0, 0xff})
	screen.DrawImage(i1, nil)
	i1w, i1h := i1.Size()

	// Create an even smaller blue rectangle, a bit transparent. Then move and rotate it before
	// drawing it over the screen image
	i2, _ := ebiten.NewImage(w/3, h/3, ebiten.FilterDefault)
	i2.Fill(color.RGBA{0, 0, 0xff, 0x88})
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(i1w), float64(i1h))
	opts.GeoM.Rotate(0.5)
	opts.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(i2, opts)

	// Extra exercise: draw i2 over i1, not over screen
	// Tip: changing screen to i1 at line 34 isn't enough
	// because i1 has already been disposed
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Colors and ImageOptions")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
