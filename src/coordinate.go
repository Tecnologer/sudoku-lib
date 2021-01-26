package sudoku

//Coordinate struct
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//NewCoordinate creates new coordinate instance
func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{x, y}
}

//EqualsXY returns true x,y is the coordinate
func (c *Coordinate) EqualsXY(x, y int) bool {
	return c.X == x && c.Y == y
}

//Equals compare two instance of coordinate, returns true if has the same value
func (c *Coordinate) Equals(c2 *Coordinate) bool {
	return c.X == c2.X && c.Y == c2.Y
}
