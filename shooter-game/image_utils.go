package shooter

import (
	"image"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/hajimehoshi/ebiten"
)

func getImage(name string, obj *assets.Object) (*ebiten.Image, error) {
	var rect image.Rectangle
	for _, img := range obj.Specs.Images {
		if img.Name == name {
			rect = image.Rect(img.X, img.Y, img.X+img.W, img.Y+img.H)
			break
		}
	}
	img := obj.Image.SubImage(rect).(*ebiten.Image)
	return img, nil
}
