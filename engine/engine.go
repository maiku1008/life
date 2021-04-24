package engine

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewUniverse(width, height int) *Universe {
	return &Universe{
		cells:  make([]bool, width*height),
		width:  width,
		height: height,
	}
}

type Universe struct {
	// Use a 1d slice to represent 2d data
	// https://learnsfml.com/2d-data-in-a-1d-array/
	cells  []bool
	width  int
	height int
}

// Resurrect a cell
func (u *Universe) Resurrect(x, y int) {
	if u.isInBounds(x, y) {
		u.cells[x+y*u.width] = true
	}
}

// Nuke all cells into a blank slate
func (u *Universe) Nuke() {
	u.cells = make([]bool, u.width*u.height)
}

// Print how many live cells are existing at one point
// Used for debugging
func (u *Universe) GetLiving() int {
	var count int
	for _, v := range u.cells {
		if v {
			count++
		}
	}
	return count
}

// Init initializes the universe with a random state
func (u *Universe) Init(maxLiveCells int) {
	for i := 0; i < maxLiveCells; i++ {
		x := rand.Intn(u.width)
		y := rand.Intn(u.height)
		u.Resurrect(x, y)
	}
}

// Return if a cell is alive
func (u *Universe) isAlive(x, y int) bool {
	return u.cells[x+y*u.width]
}

// Get the state of a cell's neighbours.
// Leverage Moore's neighborhood:
// https://en.wikipedia.org/wiki/Moore_neighborhood
func (u *Universe) getAliveNeighbours(x, y int) int {
	aliveNeighbours := 0
	// check the neighbour coordinates
	for nx := x - 1; nx <= x+1; nx++ {
		for ny := y - 1; ny <= y+1; ny++ {
			if !u.isInBounds(nx, ny) {
				continue
			}
			if nx == x && ny == y {
				continue
			}
			if u.isAlive(nx, ny) {
				aliveNeighbours++
			}
		}
	}
	return aliveNeighbours
}

// Implement the rules of the game of life to determine a cell's state at the next generation
func (u *Universe) GetNextGenerationState(x, y int) bool {
	alive := u.isAlive(x, y)
	aliveNeighbours := u.getAliveNeighbours(x, y)

	switch {
	// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	case alive && aliveNeighbours < 2:
		return false
	// Any live cell with two or three live neighbours lives on to the next generation.
	case alive && (aliveNeighbours == 2 || aliveNeighbours == 3):
		return true
	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	case alive && aliveNeighbours > 3:
		return false
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	case !alive && aliveNeighbours == 3:
		return true
	default:
		return false
	}
}

// Returns whether or not a set of coordinates are within the bounds of the universe
func (u *Universe) isInBounds(x, y int) bool {
	return 0 <= x && x < u.width && 0 <= y && y < u.height
}

// Update updates the universe state by one generation
func (u *Universe) Update() {
	next := make([]bool, u.width*u.height)
	for x := 0; x < u.width; x++ {
		for y := 0; y < u.height; y++ {
			next[x+y*u.width] = u.GetNextGenerationState(x, y)
		}
	}
	u.cells = next
}

// Cells returns the entire slice of cells
func (u *Universe) Cells() []bool {
	return u.cells
}
