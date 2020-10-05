package main

import (
	"bytes"
	"encoding/json"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type framesSpec struct {
	Frames []frameSpec `json:"frames"`
}

type frameSpec struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type Game struct {
	tick      float64
	speed     float64
	frames    []frameSpec
	numFrames int
}

var coins *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(coinImg))
	if err != nil {
		log.Fatal(err)
	}

	coins, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// calculate frame to show
	frameNum := int(g.tick*g.speed/10) % g.numFrames
	frame := g.frames[frameNum]

	// select the frame in the tileset image
	subImg := coins.SubImage(image.Rect(frame.X, frame.Y, frame.X+frame.W, frame.H)).(*ebiten.Image)

	// move to the center of the screen
	x, y := screen.Size()
	tx := x/2 - frame.W/2
	ty := y/2 - frame.H/2
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tx), float64(ty))

	if err := screen.DrawImage(subImg, op); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w / 2, h / 2
}

// buildFrames read the specs of the frames from
// a json file and stores them into the game object
func (g *Game) buildFrames(path string) error {
	j, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fSpec := &framesSpec{}
	json.Unmarshal(j, fSpec)

	g.frames = fSpec.Frames
	g.numFrames = len(g.frames)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expecting json spec path as command line argument")
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Draw tiles with different sizes")

	g := &Game{
		speed: 2,
	}

	if err := g.buildFrames(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
