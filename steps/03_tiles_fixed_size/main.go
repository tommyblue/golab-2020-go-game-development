package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	imgSize   = 16
	numFrames = 8
)

type Game struct {
	tick  float64
	speed float64
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
	op := &ebiten.DrawImageOptions{}

	// calculate frame to show
	frameNum := int(g.tick/g.speed) % numFrames
	// as the images in the tilesheet have all the same size
	// we just move right the number of current frame * width of the image
	frameX := frameNum * imgSize
	subImg := coins.SubImage(image.Rect(frameX, 0, frameX+imgSize, imgSize)).(*ebiten.Image)

	if err := screen.DrawImage(subImg, op); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(w, h int) (screenWidth, screenHeight int) {
	return w / 2, h / 2
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Draw tiles")

	g := &Game{
		speed: 60 / 6,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
