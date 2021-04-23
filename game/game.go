package game

import (
	"life/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
	title        = "Game of life"
)

func newGame(initialCells int) *game {
	u := engine.NewUniverse(screenWidth, screenHeight)
	u.Init(initialCells)
	return &game{
		universe: u,
	}
}

type game struct {
	universe *engine.Universe
	pixels   []byte
}

func (g *game) Update() error {
	g.universe.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, screenWidth*screenHeight*4)
	}
	for i, v := range g.universe.Print() {
		if v {
			g.pixels[4*i] = 0
			g.pixels[4*i+1] = 0
			g.pixels[4*i+2] = 0
			g.pixels[4*i+3] = 0
		} else {
			g.pixels[4*i] = 0xff
			g.pixels[4*i+1] = 0xff
			g.pixels[4*i+2] = 0xff
			g.pixels[4*i+3] = 0xff
		}
	}
	screen.ReplacePixels(g.pixels)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func Start(initialCells int) error {
	g := newGame(initialCells)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle(title)
	err := ebiten.RunGame(g)
	if err != nil {
		return err
	}
	return nil
}
