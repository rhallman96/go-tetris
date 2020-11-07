package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gridWidth   = 240
	gridHeight  = 480
	blockWidth  = 24
	blockHeight = 24

	title = "go-tetris"
)

var colors = []uint32{
	0xff8080, // red
	0x80ff80, // green
	0x8080ff, // blue
}

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		gridWidth, gridHeight, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	board := &Board{}
	board.Reset()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
		renderBoard(board, surface)
		window.UpdateSurface()
	}
}

func renderBoard(board *Board, surface *sdl.Surface) {
	rect := sdl.Rect{0, 0, gridWidth, gridHeight}
	surface.FillRect(&rect, 0xffffffff)

	renderGrid(board.Grid, surface)
	renderPiece(&board.Piece, surface)
}

func renderGrid(grid [][]BlockValue, surface *sdl.Surface) {
	for i, row := range grid {
		for j, value := range row {
			renderBlock(j, i, value, surface)
		}
	}
}

func renderPiece(piece *Piece, surface *sdl.Surface) {
	shape := piece.GetShape()
	for _, point := range shape {
		renderBlock(point.X, point.Y, piece.Value, surface)
	}
}

func renderBlock(x, y int, value BlockValue, surface *sdl.Surface) {
	if value == 0 {
		return
	}

	var xPos int32 = (int32)(x * blockWidth)
	var yPos int32 = (int32)(y * blockHeight)
	color := colors[int(value-1)%len(colors)]
	rect := sdl.Rect{xPos, yPos, blockWidth, blockHeight}
	surface.FillRect(&rect, color)
}
