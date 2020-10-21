package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	audioContext *audio.Context
	click        *audio.Player
)

const debouncer = 100 * time.Millisecond

type game struct {
	lastClickAt time.Time
}

func (g *game) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyA) && time.Now().Sub(g.lastClickAt) > debouncer {
		log.Printf("A pressed")
		click.Rewind()
		click.Play()
		g.lastClickAt = time.Now()
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Click A to play a sound")
}

func (g *game) Layout(w, h int) (int, int) {
	return w, h
}

func main() {
	audioContext, _ = audio.NewContext(44100)
	oggS, _ := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(RagtimeSound))

	s := audio.NewInfiniteLoop(oggS, oggS.Length())

	background, _ := audio.NewPlayer(audioContext, s)

	sound, _ := wav.Decode(audioContext, audio.BytesReadSeekCloser(ClickSound))
	click, _ = audio.NewPlayer(audioContext, sound)

	background.Play()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("sounds")
	ebiten.RunGame(&game{})
}
