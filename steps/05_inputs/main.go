package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type game struct {
	lastClickAt  time.Time // 0-value of time is 0001-01-01 00:00:00 +0000 UTC
	x, y         int
	currentColor int
}

var colors = []color.Color{
	color.RGBA{0, 0, 0, 0xff},
	color.RGBA{0xff, 0, 0, 0xff},
	color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0, 0, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0, 0xff},
	color.RGBA{0xff, 0, 0xff, 0xff},
	color.RGBA{0, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

const debouncer = 100 * time.Millisecond

func (g *game) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyA) && time.Now().Sub(g.lastClickAt) > debouncer {
		log.Printf("A pressed")
		g.lastClickAt = time.Now()
		g.currentColor = (g.currentColor + 1) % len(colors)
	}
	g.x, g.y = ebiten.CursorPosition()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("press A to switch background color. Mouse position: %s", g.mousePosition())
	screen.Fill(colors[g.currentColor])
	ebitenutil.DebugPrint(screen, msg)
}

func (g *game) Layout(x, y int) (int, int) {
	return x, y
}

func (g *game) mousePosition() string {
	return fmt.Sprintf("(%d, %d)", g.x, g.y)
}

func main() {
	log.Println(time.Time{})
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("inputs")
	ebiten.RunGame(&game{})
}
