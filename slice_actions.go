// @Title
// @Description
// @Author
// @Update
package builtins

import (
	"sort"

	. "golang.org/x/exp/constraints"
)

/*

Slice Actions

Functions

 - Reversed
 - ReversedInPlace
 - Sorted
 - SortedInPlace
 - Pop
 - PopLeft
 - Zip
 - SlidingWindow

*/

func Reversed[T any](s []T) []T {

	res := make([]T, len(s))

	for i, j := 0, len(s)-1; j >= 0; i, j = i+1, j-1 {
		res[i] = s[j]
	}

	return res
}

func ReversedInPlace[T any](s []T) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}

func Sorted[T Ordered](s []T, ascending bool) []T {

	a := make([]T, len(s))

	copy(a, s)

	if ascending {

		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})

	} else {

		sort.Slice(a, func(i, j int) bool {
			return a[i] > a[j]
		})

	}

	return a

}

func SortedInPlace[T Ordered](s []T, ascending bool) {

	if ascending {

		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})

	} else {

		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})

	}

}

func Pop[T any](s []T) (T, []T) {
	return s[len(s)-1], s[:len(s)-1]
}

func PopLeft[T any](s []T) (T, []T) {
	return s[0], s[1:]
}

func Zip[T, O any](one []T, two []O) [][2]interface{} {

	_min := Min([]int{len(one), len(two)})
	res := make([][2]interface{}, _min)

	for i := 0; i < _min; i++ {

		res[i] = [2]interface{}{one[i], two[i]}

	}

	return res

}

// Havn't tested this much, due to realizing it probably won't be utilised much
func SlidingWindow[T any](size int, s []T) [][]T {

	if len(s) <= size {
		return [][]T{s}
	}

	r := make([][]T, 0, len(s)-size+1)

	for i, j := 0, size; j <= len(s); i, j = i+1, j+1 {
		r = append(r, s[i:j])
	}

	return r

}
