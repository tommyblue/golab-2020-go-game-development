package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type scene struct {
	img     *ebiten.Image
	imgPos  image.Rectangle
	onClick string
	bg      color.Color
}

type game struct {
	scenes      map[string]*scene
	activeScene string
	lastClickAt time.Time
}

const debouncer = 100 * time.Millisecond

func (g *game) Update(screen *ebiten.Image) error {
	s, ok := g.scenes[g.activeScene]
	if !ok {
		panic("unknown scene")
	}
	w, h := s.img.Size()
	sW, sH := screen.Size()
	dW := sW/2 - w/2
	dH := sH/2 - h/2
	s.imgPos = image.Rect(dW, dH, dW+w, dH+h)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && time.Now().Sub(g.lastClickAt) > debouncer {
		g.lastClickAt = time.Now()
		x, y := ebiten.CursorPosition()

		if s.imgPos.Min.X < x && s.imgPos.Min.Y < y && x < s.imgPos.Max.X && y < s.imgPos.Max.Y {
			g.activeScene = s.onClick
		}
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	s, ok := g.scenes[g.activeScene]
	if !ok {
		panic("unknown scene")
	}
	screen.Fill(s.bg)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.imgPos.Min.X), float64(s.imgPos.Min.Y))
	screen.DrawImage(s.img, op)
}

func (g *game) Layout(x, y int) (int, int) {
	return x, y
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("scenes")
	g := &game{
		scenes: make(map[string]*scene),
	}
	if err := g.addMainScene(); err != nil {
		log.Fatal(err)
	}
	if err := g.addBackScene(); err != nil {
		log.Fatal(err)
	}
	g.activeScene = "main"
	ebiten.RunGame(g)
}

func (g *game) addMainScene() error {
	return g.addScene("main", "back", startImg, color.RGBA{0, 0xff, 0, 0xff})
}
func (g *game) addBackScene() error {
	return g.addScene("back", "main", backImg, color.RGBA{0xff, 0, 0, 0xff})
}

func (g *game) addScene(key, target string, srcImg []byte, bg color.Color) error {
	rawImg, _, err := image.Decode(bytes.NewReader(srcImg))
	if err != nil {
		return err
	}

	img, err := ebiten.NewImageFromImage(rawImg, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	s := &scene{
		img:     img,
		onClick: target,
		bg:      bg,
	}
	g.scenes[key] = s

	return nil
}
