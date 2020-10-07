package shooter

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

type Game struct {
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func NewGame() *Game {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Shooter")
	g := &Game{}
	return g
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
