package objects

import "github.com/hajimehoshi/ebiten"

type Object interface {
	Tick(uint)                // tell the object a new tick happened
	Draw(*ebiten.Image) error // draw the object
}

const (
	right float64 = 1
	left  float64 = -1
	down  float64 = 1
	up    float64 = -1
)
