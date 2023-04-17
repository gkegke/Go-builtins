package builtins

import (
	"fmt"
	"math"
	"strconv"
)

/*

Contents:

- Abs
- Round
- DivMod

- Bin
- Hex
- Oct

*/

func Abs[T Number](n T) T {

	if n >= 0 {
		return n
	}

	return -n

}

func Round(n float64, digits float64) float64 {

	if digits == 0 {

		return n

	}

	tenp := math.Pow(10, digits)

	return math.RoundToEven(n*tenp) / tenp

}

func DivMod(x, y float64) (float64, float64) {

	return math.Floor(x / y), math.Mod(x, y)

}

func Bin(n int64) string {
	return strconv.FormatInt(n, 2)
}

func Hex(n int64) string {

	return fmt.Sprintf("%x", n)

}

func Oct(n int64) string {

	return fmt.Sprintf("%o", n)

}
