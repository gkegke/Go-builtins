// @Title
// @Description
// @Author
// @Update

package builtins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {

	_inputs := []Test{
		{
			[]int{0, 10, 1},
			10,
		},
		{
			[]int{-10, 10, 1},
			10,
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]int)

		res := Max(_inp)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestMin(t *testing.T) {

	_inputs := []Test{
		{
			[]int{0, 10, 1},
			0,
		},
		{
			[]int{-10, 10, 1},
			-10,
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]int)

		res := Min(_inp)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestSum(t *testing.T) {

	_inputs := []Test{
		{
			[]int{0, 10, 1},
			11,
		},
		{
			[]int{-10, 10, 1},
			1,
		},
	}

	_inputs_floats := []Test{
		{
			[]float64{-10.5, 10.5, 1.0},
			1.0,
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]int)

		res := Sum(_inp)

		assert.EqualValues(
			t,
			_test.Expected,
			res,
			"",
		)

	}

	for _, _test := range _inputs_floats {

		_inp := _test.Input.([]float64)

		res := Sum(_inp)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestCounter(t *testing.T) {

	_inputs := []Test{
		{
			[]int{0, 10, 2},
			map[int]int{
				0:  1,
				10: 1,
				2:  1,
			},
		},
		{
			[]int{0, 10, 10, 2},
			map[int]int{
				0:  1,
				10: 2,
				2:  1,
			},
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]int)

		res := Counter(_inp)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestMostCommon(t *testing.T) {

	_inputs := []Test{
		{
			[]int{0, 10, 2},
			[]int{0, 10, 2},
		},
		{
			[]int{0, 10, 10, 2},
			[]int{10},
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]int)

		res := MostCommon(_inp)

		assert.ElementsMatch(
			t,
			_test.Expected,
			res,
			"",
		)

	}

}
