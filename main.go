package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/rhallman96/go-tetris/game"
	"golang.org/x/image/colornames"
	"image/color"
	"time"
)

const (
	blockWidth  = 24
	blockHeight = 24
	gridWidth   = blockWidth * game.BoardWidth
	gridHeight  = blockHeight * game.BoardHeight
	title       = "TETRIS"

	// board updates per second
	logicalFramerate   uint64 = 25
	logicalFramerateMS uint64 = (1000 / logicalFramerate)

	// piece falling rates
	startDropTicks = 12
	finalDropTicks = 2
)

var colors = []color.RGBA{
	colornames.Tomato,
	colornames.Skyblue,
	colornames.Violet,
	colornames.Lime,
}

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, gridWidth, gridHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	board := &game.Board{}
	board.Reset()
	imd := imdraw.New(nil)

	prevTime := time.Now()
	var timeCounter uint64 = 0
	var frameCounter uint64 = 0
	for !win.Closed() {
		timeCounter += uint64(time.Since(prevTime).Milliseconds())
		prevTime = time.Now()
		logicalTicks := timeCounter / logicalFramerateMS
		timeCounter %= logicalFramerateMS
		frameCounter += logicalTicks

		tick(board, win, logicalTicks, frameCounter)

		render(board, imd, win)
		win.Update()
	}
}

func tick(board *game.Board, win *pixelgl.Window, logicalTicks, frameCounter uint64) {
	for ; logicalTicks > 0; logicalTicks-- {
		shouldUpdate := (frameCounter%dropTicks(board) == 0) || win.Pressed(pixelgl.KeyDown)
		if shouldUpdate {
			if !board.DropPiece() && !board.NextPiece() {
				board.Reset()
			}
		}
	}

	if win.JustPressed(pixelgl.KeyLeft) {
		board.MovePieceLeft()
	}

	if win.JustPressed(pixelgl.KeyRight) {
		board.MovePieceRight()
	}

	if win.JustPressed(pixelgl.KeyZ) {
		board.RotatePieceLeft()
	}

	if win.JustPressed(pixelgl.KeyX) {
		board.RotatePieceRight()
	}

	if win.JustPressed(pixelgl.KeyUp) {
		board.QuickDropPiece()
		board.NextPiece()
	}
}

func dropTicks(board *game.Board) uint64 {
	level := board.Level()
	if level >= startDropTicks-finalDropTicks {
		return finalDropTicks
	}
	return uint64(startDropTicks - level)
}

func render(board *game.Board, imd *imdraw.IMDraw, win *pixelgl.Window) {
	win.Clear(colornames.Snow)
	imd.Clear()
	renderGrid(board.Grid, imd)
	renderPiece(&board.Piece, imd)
	imd.Draw(win)
}

func renderGrid(grid [][]uint8, imd *imdraw.IMDraw) {
	for i, row := range grid {
		for j, value := range row {
			renderBlock(j, i, value, imd)
		}
	}
}

func renderPiece(piece *game.Piece, imd *imdraw.IMDraw) {
	shape := piece.GetShape()
	for _, point := range shape {
		renderBlock(point.X, point.Y, piece.Value, imd)
	}
}

func renderBlock(x, y int, value uint8, imd *imdraw.IMDraw) {
	if value == 0 {
		return
	}

	var xPos float64 = float64(x * blockWidth)
	var yPos float64 = float64(gridHeight - (y * blockHeight))

	imd.Color = colors[int(value-1)%len(colors)]
	imd.Push(pixel.V(xPos, yPos), pixel.V(xPos+blockWidth, yPos-blockHeight))
	imd.Rectangle(0)
}
