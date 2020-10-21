package assets

import (
	"bytes"
	"encoding/json"
	"image"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"path"
	"runtime"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

var (
	Hud          *Object
	Objects      *Object
	Stall        *Object
	audioContext *audio.Context
)

type Object struct {
	Name  string
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
		Name:  "HUD",
		Image: newImage(hudBytes),
		Specs: buildSpecs(relativePath("./hud.json")),
	}
	Objects = &Object{
		Name:  "Objects",
		Image: newImage(objectsBytes),
		Specs: buildSpecs(relativePath("./objects.json")),
	}
	Stall = &Object{
		Name:  "Stall",
		Image: newImage(stallBytes),
		Specs: buildSpecs(relativePath("./stall.json")),
	}
	var err error
	audioContext, err = audio.NewContext(44100)
	if err != nil {
		log.Fatalf("cannot initialize sound context: %v", err)
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

func LoadWavPlayer(src []byte) *audio.Player {
	s, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(src))
	if err != nil {
		log.Fatal(err)
	}
	return loadPlayer(s)
}

func LoadOggPlayer(src []byte) *audio.Player {
	s, err := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(src))
	if err != nil {
		log.Fatal(err)
	}
	return loadPlayer(s)
}

func loadPlayer(sound io.ReadCloser) *audio.Player {
	p, err := audio.NewPlayer(audioContext, sound)
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func BackgroundMusicPlayer() (*audio.Player, error) {
	oggS, err := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(RagtimeSound))
	if err != nil {
		return nil, err
	}
	s := audio.NewInfiniteLoop(oggS, oggS.Length())

	player, err := audio.NewPlayer(audioContext, s)
	if err != nil {
		return nil, err
	}

	return player, nil
}
