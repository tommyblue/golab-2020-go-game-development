package assets

import (
	"bytes"
	"encoding/json"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"path"
	"runtime"

	"github.com/hajimehoshi/ebiten"
)

var (
	Hud     *Object
	Objects *Object
	Stall   *Object
)

type Object struct {
	Image *ebiten.Image
	Specs *SpriteSheet
}
type SpriteSheet struct {
	Images []ImageSpec `json:"images"`
}

type ImageSpec struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	W    int    `json:"width"`
	H    int    `json:"height"`
}

func init() {
	Hud = &Object{
		Image: newImage(hudBytes),
		Specs: buildSpecs(relativePath("./hud.json")),
	}
	Objects = &Object{
		Image: newImage(objectsBytes),
		Specs: buildSpecs(relativePath("./objects.json")),
	}
	Stall = &Object{
		Image: newImage(stallBytes),
		Specs: buildSpecs(relativePath("./stall.json")),
	}
}

func newImage(src []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(src))
	if err != nil {
		log.Fatal(err)
	}

	obj, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	return obj
}

func buildSpecs(path string) *SpriteSheet {
	j, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	s := &SpriteSheet{}
	json.Unmarshal(j, s)
	return s
}

func relativePath(filepath string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("relativePath error")
	}
	return path.Join(path.Dir(filename), filepath)
}
