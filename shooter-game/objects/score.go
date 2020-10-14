package objects

import (
	"fmt"
	"log"
	"math"
	"strings"
	"sync/atomic"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game/assets"
	"github.com/develersrl/golab2020-go-game-dev/shooter-game/utils"
	"github.com/hajimehoshi/ebiten"
)

type score struct {
	scoreImg   *ebiten.Image
	dotsImg    *ebiten.Image
	numImgs    []*ebiten.Image
	scoreValue *int64
}

func NewScore(scoreImgName, dotsImgName, numImgName string, scoreValue *int64) Object {
	scoreImg, err := utils.GetImage(scoreImgName, assets.Hud)
	if err != nil {
		log.Fatalf("drawing %s: %v", scoreImgName, err)
	}
	dotsImg, err := utils.GetImage(dotsImgName, assets.Hud)
	if err != nil {
		log.Fatalf("drawing %s: %v", dotsImgName, err)
	}

	s := &score{
		scoreImg:   scoreImg,
		dotsImg:    dotsImg,
		scoreValue: scoreValue,
	}

	for i := 0; i < 10; i++ {
		imgName := strings.Replace(numImgName, "$", fmt.Sprintf("%d", i), 1)
		img, err := utils.GetImage(imgName, assets.Hud)
		if err != nil {
			log.Fatalf("drawing %s: %v", imgName, err)
		}
		s.numImgs = append(s.numImgs, img)
	}

	return s
}

func (s *score) Update(_ *ebiten.Image, tick uint) {}

func (s *score) Draw(trgt *ebiten.Image) error {
	_, h := trgt.Size()
	baseH := h - 20

	var lastW float64 = 20

	scoreW, scoreH := s.scoreImg.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(lastW, float64(baseH-scoreH))
	trgt.DrawImage(s.scoreImg, op)

	lastW = lastW + float64(scoreW)

	dotsW, dotsH := s.dotsImg.Size()
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(lastW, float64(baseH-dotsH-scoreH/2+dotsH/2))
	trgt.DrawImage(s.dotsImg, op)

	lastW = lastW + float64(dotsW) + 10

	num := atomic.LoadInt64(s.scoreValue)
	var i int64 = 1
	digits := []int64{}
	for {
		d := s.digit(num, i)
		digits = append([]int64{d}, digits...)
		num = num / 10
		if num == 0 {
			break
		}
	}

	for _, d := range digits {
		nW, nH := s.numImgs[d].Size()
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(lastW, float64(baseH-nH-scoreH/2+nH/2))
		trgt.DrawImage(s.numImgs[d], op)

		lastW = lastW + float64(nW)
	}
	return nil
}

func (s *score) OnScreen() bool {
	return true
}

func (s *score) digit(num, place int64) int64 {
	if num == 0 {
		return 0
	}
	r := num % int64(math.Pow(10, float64(place)))
	return r / int64(math.Pow(10, float64(place-1)))
}
