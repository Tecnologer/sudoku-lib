package sudoku

import (
	"math/rand"
	"time"
)

func getRandRange(min, max int) int {
	max = max * 10
	min = min * 10
	rand.Seed(time.Now().UnixNano() + int64(min-max))
	return (rand.Intn(max-min) + min) / 10
}
func getCoordinate() (int, int) {
	return rand.Intn(9), rand.Intn(9)
}

func getSquareOffset(i int) int {
	return int(i/3) * 3
}
