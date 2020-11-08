package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/rhallman96/go-tetris/game"
	"golang.org/x/image/colornames"
	"image/color"
)

const (
	gridWidth   = 264
	gridHeight  = 480
	blockWidth  = 24
	blockHeight = 24

	title = "TETRIS"
)

var colors = []color.RGBA{
	colornames.Tomato,
	colornames.Skyblue,
	colornames.Violet,
	colornames.Lime,
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

	board := game.Board{}
	board.Reset()

	imd := imdraw.New(nil)

	i := 0
	for !win.Closed() {
		i++
		if i%10 == 0 {
			if !board.DropPiece() {
				if !board.NextPiece() {
					board.Reset()
				}
				board.ClearFilledRows()
			}
		}

		if win.JustPressed(pixelgl.KeyZ) {
			board.RotatePieceLeft()
		} else if win.JustPressed(pixelgl.KeyX) {
			board.RotatePieceRight()
		}

		if win.JustPressed(pixelgl.KeyLeft) {
			board.MovePieceLeft()
		} else if win.JustPressed(pixelgl.KeyRight) {
			board.MovePieceRight()
		}

		if win.Pressed(pixelgl.KeyDown) {
			i = 9
		}

		render(&board, imd, win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func render(board *game.Board, imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Clear()
	win.Clear(colornames.Snow)
	renderGrid(board.Grid, imd)
	renderPiece(&board.Piece, imd)
	imd.Draw(win)
}

func renderGrid(grid [][]int, imd *imdraw.IMDraw) {
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

func renderBlock(x, y, value int, imd *imdraw.IMDraw) {
	if value == 0 {
		return
	}

	var xPos float64 = float64(x * blockWidth)
	var yPos float64 = float64(gridHeight - (y * blockHeight))

	imd.Color = colors[int(value-1)%len(colors)]
	imd.Push(pixel.V(xPos, yPos), pixel.V(xPos+blockWidth, yPos-blockHeight))
	imd.Rectangle(0)
}
