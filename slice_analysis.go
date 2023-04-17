// @Title
// @Description
// @Author
// @Update
package builtins

import (
	. "golang.org/x/exp/constraints"
)

/*

SLICE ANALYSIS

Max
Min
Sum
Mean
Counter
MostCommon

*/

func Max[T Ordered](s []T) T {

	if len(s) == 0 {

		var zero T

		return zero

	}

	_max := s[0]

	for _, v := range s {

		if _max < v {
			_max = v
		}

	}

	return _max

}

func Min[T Ordered](s []T) T {

	if len(s) == 0 {

		var zero T

		return zero

	}

	_min := s[0]

	for _, v := range s {

		if _min > v {
			_min = v
		}

	}

	return _min

}

func Sum[T Number](s []T) T {

	if len(s) == 0 {
		var zero T
		return zero
	}

	var _sum T

	for _, n := range s {

		_sum += n

	}

	return _sum

}

/*
Arg

	s []T

Returns

	map[T]int

Creates a map that contains the unique values in a slice, and their
number of occurances.

[1,1,2,3] => map[int]{1: 2, 2: 1, 3: 1}
*/
func Counter[T comparable](s []T) map[T]int {

	c := make(map[T]int)

	for _, v := range s {

		c[v] = c[v] + 1

	}

	return c

}

// helper function, better to create your own for more advanced queries
func MostCommon[T comparable](s []T) []T {

	c := Counter(s)

	_max := 0
	_res := []T{}

	for k, v := range c {

		if v == _max {
			_res = append(_res, k)
		} else if v > _max {
			_res = []T{k}
			_max = v
		}

	}

	return _res
}
