package sudoku

import (
	"fmt"
	"sync"
	"time"
)

//Game struct
type Game struct {
	Board       *board          `json:"board"`
	Level       ComplexityLevel `json:"level"`
	StartTime   time.Time       `json:"start_time"`
	LockedCoord []*Coordinate   `json:"locked_coordinates"`
	complexity  *complexity
	mutex       sync.Mutex
}

//NewGame creates new game with the specific complexity
func NewGame(level ComplexityLevel) *Game {
	g := &Game{
		Board: initGame(level),
		Level: level,
	}

	if level != EmptyLevel {
		b := *g.Board
		for !g.CanBeSolved() {
			g.Board = initGame(level)
			b = *g.Board
		}
		g.Board = &b
	}

	g.lockInitialCoordinates()
	return g
}

//IsSolved returns true if the board doesn't have empty fields
func (g *Game) IsSolved() bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if g.IsEmpty(x, y) {
				return false
			}
		}
	}
	return true
}

func (g *Game) lockInitialCoordinates() {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if g.IsEmpty(x, y) {
				continue
			}

			g.LockCoordinateXY(x, y)
		}
	}
}

//Solve solves the current board
func (g *Game) Solve() {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if !g.IsEmpty(x, y) {
				continue
			}

			for n := 1; n < 10; n++ {
				if g.IsValid(x, y, n) {
					g.Set(x, y, n)
					g.Solve()
					if !g.IsSolved() {
						g.Set(x, y, 0)
					}
				}
			}
			return

		}
	}

}

//Validate validates if the solutions is correct
func (g *Game) Validate() *ValidationErrors {
	errs := new(ValidationErrors)
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			switch {
			case g.IsEmpty(x, y):
				errs.appendError(emptyError, NewErrorCoordinate(x, y))
			case g.IsXValid(x, y, g.Get(x, y)):
				errs.appendError(invalidRow, NewErrorCoordinate(x, y))
			case g.IsYValid(x, y, g.Get(x, y)):
				errs.appendError(invalidColumn, NewErrorCoordinate(x, y))
			case g.IsSquareValid(x, y, g.Get(x, y)):
				errs.appendError(invalidSquare, NewErrorCoordinate(x, y))
			}
		}
	}

	return errs
}

//Set sets the value in the coordinate
func (g *Game) Set(x, y, n int) {
	g.Board.set(x, y, n)
}

//Get return the value in the coordinate
func (g *Game) Get(x, y int) int {
	return g.Board.get(x, y)
}

//IsXValid validate if n is valid in the row
func (g *Game) IsXValid(x, y, n int) bool {
	return g.Board.isXValid(x, y, n)
}

//IsYValid validate if n is valid in the column
func (g *Game) IsYValid(x, y, n int) bool {
	return g.Board.isYValid(x, y, n)
}

//IsSquareValid validate if n is valid in the square
func (g *Game) IsSquareValid(x, y, n int) bool {
	return g.Board.isSquareValid(x, y, n)
}

//IsValid validate if n is valid in the row, column and square
func (g *Game) IsValid(x, y, n int) bool {
	return g.Board.isValid(x, y, n)
}

//IsEmpty validate if the coordinate has value
func (g *Game) IsEmpty(x, y int) bool {
	return g.Board.isEmpty(x, y)
}

//IsCoordinateLockedXY returns if the coordinate x,y is locked
func (g *Game) IsCoordinateLockedXY(x, y int) bool {
	return g.IsCoordinateLocked(NewCoordinate(x, y))
}

//IsCoordinateLocked returns if the coordinate x,y is locked
func (g *Game) IsCoordinateLocked(c *Coordinate) bool {
	for _, c := range g.LockedCoord {
		if c.Equals(c) {
			return true
		}
	}

	return false
}

//LockCoordinateXY add coordinate to the lock list
func (g *Game) LockCoordinateXY(x, y int) {
	g.LockCoordinate(NewCoordinate(x, y))
}

//LockCoordinate add coordinate to the lock list
func (g *Game) LockCoordinate(c *Coordinate) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.IsCoordinateLocked(c) {
		return
	}

	g.LockedCoord = append(g.LockedCoord, c)
}

//CanBeSolved returns true if the board can be solved
func (g *Game) CanBeSolved() bool {
	g.Solve()

	return g.IsSolved()
}

//SetDataRow sets the data to the specified row when the game is Empty
func (g *Game) SetDataRow(row int, data [9]int) error {
	if g.Level != EmptyLevel {
		return fmt.Errorf("This game cannot be modified")
	}

	g.Board[row] = data

	return nil
}
