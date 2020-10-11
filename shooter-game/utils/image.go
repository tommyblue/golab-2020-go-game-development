package utils

import (
	"fmt"
	"image"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/hajimehoshi/ebiten"
)

func GetImage(name string, obj *assets.Object) (*ebiten.Image, error) {
	var rect image.Rectangle
	var found bool
	for _, img := range obj.Specs.Images {
		if img.Name == name {
			rect = image.Rect(img.X, img.Y, img.X+img.W, img.Y+img.H)
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("not found in %s", obj.Name)
	}
	img := obj.Image.SubImage(rect).(*ebiten.Image)
	return img, nil
}
