package intutils

import (
	"math"
	"strconv"

	"github.com/sunilbpandey/go-toolkit/errorutils"
)

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	errorutils.Check(err)
	return i
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
