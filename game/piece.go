package game

type Shape []Point

type Point struct {
	X, Y int
}

type Piece struct {
	Value uint8

	coords        Point
	rotationIndex int
	rotations     []Shape
}

// Rotations are pre-computed for each unique kind of piece.
var pieceRotations = [][]Shape{

	// Square
	Shape{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}.identityRotations(),

	// Line
	Shape{
		{1, 0},
		{1, 1},
		{1, 2},
		{1, 3},
	}.flipRotations(),

	// T-Shape
	Shape{
		{0, 1},
		{1, 1},
		{2, 1},
		{1, 2},
	}.squareRotations(),

	// L-Shape
	Shape{
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 2},
	}.squareRotations(),

	// Flipped L-Shape
	Shape{
		{1, 0},
		{1, 1},
		{1, 2},
		{0, 2},
	}.squareRotations(),
}

func (piece *Piece) GetShape() Shape {
	if len(piece.rotations) == 0 {
		return Shape{}
	}

	rotation := piece.rotations[piece.rotationIndex%len(piece.rotations)]
	shape := make(Shape, len(rotation))
	copy(shape, rotation)

	for i, _ := range shape {
		shape[i].X += piece.coords.X
		shape[i].Y += piece.coords.Y
	}
	return shape
}

func (piece *Piece) overlapsGrid(grid [][]uint8) bool {
	shape := piece.GetShape()
	for _, point := range shape {
		if point.X < 0 || point.X >= len(grid[0]) ||
			point.Y < 0 || point.Y >= len(grid) ||
			grid[point.Y][point.X] != 0 {
			return true
		}
	}
	return false
}

func (shape Shape) squareRotations() []Shape {
	transforms := []Shape{shape}
	for i := 0; i < 3; i++ {
		next := make(Shape, len(shape))
		for j, point := range transforms[i] {
			next[j] = Point{2 - point.Y, point.X}
		}
		transforms = append(transforms, next)
	}

	return transforms
}

func (shape Shape) flipRotations() []Shape {
	transforms := []Shape{shape}
	flip := make(Shape, len(shape))
	copy(flip, shape)
	for i, point := range shape {
		flip[i].X = point.Y
		flip[i].Y = point.X
	}
	transforms = append(transforms, flip)
	return transforms
}

func (shape Shape) identityRotations() []Shape {
	return []Shape{shape}
}
