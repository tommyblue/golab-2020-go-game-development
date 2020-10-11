package objects

import "github.com/hajimehoshi/ebiten"

type Object interface {
	Tick(*ebiten.Image, uint) // tell the object a new tick happened
	Draw(*ebiten.Image) error // draw the object
	OnScreen() bool           // false when the object is out of the screen
}

type direction int // use custom type for direction, so we can add methods to it

const (
	right direction = 1
	left  direction = -1
	down  direction = 1
	up    direction = -1
)

func (d direction) invert() direction {
	return -d
}
