package shooter

import (
	"log"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/objects"
	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth  = 800
	windowHeight = 600
	maxUint      = ^uint(0)
)

type Game struct {
	tick    uint
	objects []objects.Object
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.tick++
	if g.tick == maxUint {
		g.tick = 0
	}

	for _, o := range g.objects {
		o.Update(screen, g.tick)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, o := range g.objects {
		if err := o.Draw(screen); err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func NewGame() *Game {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Shooter")
	g := &Game{
		objects: []objects.Object{
			objects.NewBackground("bg_green.png"),
			objects.NewLevel1("water1.png", "duck_outline_target_white.png", 4),
			objects.NewDesk("bg_wood.png"),
			objects.NewCurtains("curtain_straight.png", "curtain.png"),
			objects.NewCrosshair("crosshair_white_large.png", "crosshair_red_large.png"),
		},
	}
	return g
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
