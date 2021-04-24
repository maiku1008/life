package game

import (
	"life/engine"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
	title        = "micuffaro's Game of life"
)

func newGame() *game {
	u := engine.NewUniverse(screenWidth, screenHeight)
	return &game{
		universe: u,
		pixels:   make([]byte, screenWidth*screenHeight*4),
		active:   true,
	}
}

type game struct {
	universe *engine.Universe
	pixels   []byte
	active   bool
}

// Set the cell at current mouse position to alive
func (g *game) drawMouse() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.universe.Resurrect(x, y)
	}
}

// register keypresses to control the game
func (g *game) keys() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeySpace):
		g.active = true
	case ebiten.IsKeyPressed(ebiten.KeyS):
		g.active = false
	case ebiten.IsKeyPressed(ebiten.KeyK):
		g.active = false
		g.universe.Nuke()
	case ebiten.IsKeyPressed(ebiten.KeyR):
		g.active = false
		g.universe.Nuke()
		g.universe.Init((screenWidth * screenWidth) / 4)
	case ebiten.IsKeyPressed(ebiten.KeyQ):
		g.active = false
		os.Exit(0)
	}
}

func (g *game) Update() error {
	g.keys()
	if g.active {
		g.universe.Update()
	} else {
		g.drawMouse()
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	for i, v := range g.universe.Cells() {
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

// Run the game
func Run() error {
	g := newGame()
	g.universe.Init((screenWidth * screenWidth) / 4)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle(title)
	err := ebiten.RunGame(g)
	if err != nil {
		return err
	}
	return nil
}
