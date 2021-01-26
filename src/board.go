package sudoku

import "sync"

type board [9][9]int

func initGame(level ComplexityLevel) *board {
	c := buildComplexity(level)
	b := newBoard()

	if level == EmptyLevel {
		return &b
	}

	var wg sync.WaitGroup
	wg.Add(9)

	ch := make(chan *complexData)
	for i := 0; i < 9; i++ {
		go func() {
			for data := range ch {

				xReseted, yReseted := false, false

				for i := 0; i < data.randCount; i++ {
					x := getRandRange(data.xMax-2, data.xMax)
					y := getRandRange(data.yMax-2, data.yMax)

					for ; x <= data.xMax; x++ {
						if b.isEmpty(x, y) {
							break
						}

						for ; y <= data.yMax; y++ {
							if b.isEmpty(x, y) {
								break
							}

							if y >= data.yMax && !yReseted {
								y = data.yMax - 2
								yReseted = true
							}
						}

						if b.isEmpty(x, y) && y <= data.yMax {
							break
						}

						if x >= data.xMax && !xReseted {
							x = data.xMax - 2
							xReseted = true
						}

						if xReseted && yReseted {
							wg.Done()
							return
						}
					}

					isReseted := false
					for n := getRandRange(1, 9); n < 10; n++ {
						if !b.isValid(x, y, n) {
							if n == 9 {
								if isReseted {
									break
								}

								n = 0
								isReseted = true
							}
							continue
						}

						b.set(x, y, n)
						break
					}
				}
			}
			wg.Done()
		}()
	}

	for _, data := range c {
		ch <- data
	}
	close(ch)
	wg.Wait()
	return &b
}

func newBoard() board {
	empty := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b := [9][9]int{
		// {7, 0, 0, 2, 3, 1, 0, 8, 9},
		// {0, 0, 0, 0, 6, 0, 4, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {6, 5, 7, 0, 9, 8, 3, 2, 1},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {4, 3, 9, 1, 8, 2, 0, 6, 0},
		// {0, 0, 0, 0, 0, 0, 5, 1, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for x := 0; x < 9; x++ {
		b[x] = empty
	}

	return board(b)
}

//methods

//isEmpty returns if the coordinate has value
func (b *board) isEmpty(x, y int) bool {
	if x > 8 || y > 8 {
		return false
	}
	return b[x][y] == 0
}

func (b *board) isValid(x, y, v int) bool {
	return b.isYValid(x, y, v) && b.isXValid(x, y, v) && b.isSquareValid(x, y, v)
}

func (b *board) isYValid(x, y, v int) bool {
	for i := 0; i < 9; i++ {
		if i == x {
			continue
		}
		if b[i][y] == v {
			return false
		}
	}

	return true
}

func (b *board) isXValid(x, y, v int) bool {
	for i := 0; i < 9; i++ {
		if i == y {
			continue
		}

		if b[x][i] == v {
			return false
		}
	}

	return true
}

func (b *board) isSquareValid(x, y, v int) bool {
	xOffset := getSquareOffset(x)
	yOffset := getSquareOffset(y)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == x && j == y {
				continue
			}

			if b[i+xOffset][j+yOffset] == v {
				return false
			}
		}
	}
	return true
}

//get returns the value in the coordinate
func (b *board) get(x, y int) int {
	return b[x][y]
}

//set sets the value in the coordinate
func (b *board) set(x, y, v int) {
	b[x][y] = v
}
