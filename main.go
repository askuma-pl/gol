package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	scrWidth  = 256
	scrHeight = 256
	winWidht  = 1024
	winHeight = 1024
)

func RandBool() bool {
	return rand.Intn(2) == 1
}

type Cell struct {
	X          int
	Y          int
	Live       bool
	Color      color.RGBA
	LifeCycles int
}

func NewCell(x, y int) Cell {
	return Cell{
		X:          x,
		Y:          y,
		Live:       RandBool(),
		Color:      color.RGBA{0, 255, 0, 255},
		LifeCycles: 100,
	}
}

type Game struct {
	StartTime time.Time
	Cells     [][]Cell
}

func NewGame() *Game {
	cells := [][]Cell{}
	for x := 0; x < scrWidth; x++ {
		cellsRow := []Cell{}
		for y := 0; y < scrHeight; y++ {
			nc := NewCell(x, y)
			cellsRow = append(cellsRow, nc)
		}
		cells = append(cells, cellsRow)
	}

	return &Game{
		StartTime: time.Now(),
		Cells:     cells,
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, cr := range g.Cells {
		for _, c := range cr {
			if c.Live {
				drawRectangle(screen, c)
			}
		}
	}
	time.Sleep(100 * time.Millisecond)
}

func drawRectangle(screen *ebiten.Image, c Cell) {
	screen.Set(c.X, c.Y, c.Color)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scrWidth, scrHeight
}

func (g *Game) Update() error {
	// Your game's update logic
	cycles := 0
	newCells := [][]Cell{}
	for i, cr := range g.Cells {
		newCells = append(newCells, []Cell{})
		for _, c := range cr {
			newCell := c
			newCell.Live = g.checkIfCellLive(c)
			newCell.LifeCycles--
			newCells[i] = append(newCells[i], newCell)
		}
		cycles++
	}

	log.Println("Time elapsed:", time.Since(g.StartTime), "Cycles: ", cycles)
	g.Cells = newCells

	return nil
}

func (g *Game) checkIfCellLive(c Cell) bool {
	counter := 0
	if g.checkNW(c) {
		counter++
	}
	if g.checkN(c) {
		counter++
	}
	if g.checkNE(c) {
		counter++
	}
	if g.checkE(c) {
		counter++
	}
	if g.checkSE(c) {
		counter++
	}
	if g.checkS(c) {
		counter++
	}
	if g.checkSW(c) {
		counter++
	}
	if g.checkW(c) {
		counter++
	}

	if c.LifeCycles <= 0 {
		return false
	}

	if !c.Live && counter == 3 {
		return true
	} else if c.Live && (counter == 2 || counter == 3) {
		return true
	} else {
		return false
	}

}

func (g *Game) checkNW(c Cell) bool {
	x := c.X - 1
	y := c.Y - 1

	if g.checkXOutOfBouds(x) {
		x = len(g.Cells[0]) - 1
	}
	if g.checkYOutOfBouds(y) {
		y = len(g.Cells) - 1
	}

	return g.Cells[x][y].Live
}
func (g *Game) checkN(c Cell) bool {
	x := c.X
	y := c.Y - 1

	if g.checkYOutOfBouds(y) {
		y = len(g.Cells) - 1
	}

	return g.Cells[x][y].Live
}
func (g *Game) checkNE(c Cell) bool {
	x := c.X + 1
	y := c.Y - 1

	if g.checkXOutOfBouds(x) {
		x = 0
	}

	if g.checkYOutOfBouds(y) {
		y = len(g.Cells) - 1
	}

	return g.Cells[x][y].Live
}
func (g *Game) checkE(c Cell) bool {
	x := c.X + 1
	y := c.Y

	if g.checkXOutOfBouds(x) {
		x = 0
	}

	return g.Cells[x][y].Live
}
func (g *Game) checkSE(c Cell) bool {
	x := c.X + 1
	y := c.Y + 1

	if g.checkXOutOfBouds(x) {
		x = 0
	}
	if g.checkYOutOfBouds(y) {
		y = 0
	}

	return g.Cells[x][y].Live
}

func (g *Game) checkS(c Cell) bool {
	x := c.X
	y := c.Y + 1

	if g.checkYOutOfBouds(y) {
		y = 0
	}

	return g.Cells[x][y].Live
}

func (g *Game) checkSW(c Cell) bool {
	x := c.X - 1
	y := c.Y + 1

	if g.checkXOutOfBouds(x) {
		x = len(g.Cells[0]) - 1
	}
	if g.checkYOutOfBouds(y) {
		y = 0
	}

	return g.Cells[x][y].Live
}
func (g *Game) checkW(c Cell) bool {
	x := c.X - 1
	y := c.Y

	if g.checkXOutOfBouds(x) {
		x = len(g.Cells[0]) - 1
	}

	return g.Cells[x][y].Live
}

func (g *Game) checkXOutOfBouds(x int) bool {
	if x < 0 || x >= len(g.Cells[0]) {
		return true
	}
	return false
}

func (g *Game) checkYOutOfBouds(y int) bool {
	if y < 0 || y >= len(g.Cells) {
		return true
	}
	return false
}

func main() {
	ebiten.SetWindowSize(winWidht, winHeight)
	ebiten.SetWindowTitle("Game of Julek")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
