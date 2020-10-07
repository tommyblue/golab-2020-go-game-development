package objects

import "github.com/hajimehoshi/ebiten"

type Object interface {
	Tick(uint)                // tell the object a new tick happened
	Draw(*ebiten.Image) error // draw the object
}
