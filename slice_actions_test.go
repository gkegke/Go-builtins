// @Title
// @Description
// @Author
// @Update

package builtins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversed(t *testing.T) {

	_inputs := []Test{
		{
			[]int{0, 10, 1},
			[]int{1, 10, 0},
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]int)

		res := Reversed(_inp)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestReversedInPlace(t *testing.T) {

	_inputs := []Test{
		{
			[]int{0, 10, 1},
			[]int{1, 10, 0},
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]int)

		ReversedInPlace(_inp)

		assert.Equal(
			t,
			_test.Expected,
			_inp,
			"",
		)
	}

}

func TestSorted(t *testing.T) {

	_inputs_asc := []Test{
		{
			[]int{0, 10, 1},
			[]int{0, 1, 10},
		},
		{
			[]int{0, -10, 1},
			[]int{-10, 0, 1},
		},
	}

	_inputs_desc := []Test{
		{
			[]int{0, 10, 1},
			[]int{10, 1, 0},
		},
		{
			[]int{0, -10, 1},
			[]int{1, 0, -10},
		},
	}

	for _, _test := range _inputs_asc {

		_inp := _test.Input.([]int)

		res := Sorted(_inp, true)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

	for _, _test := range _inputs_desc {

		_inp := _test.Input.([]int)

		res := Sorted(_inp, false)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestSortedInPlace(t *testing.T) {

	_inputs_asc := []Test{
		{
			[]int{0, 10, 1},
			[]int{0, 1, 10},
		},
		{
			[]int{0, -10, 1},
			[]int{-10, 0, 1},
		},
	}

	_inputs_desc := []Test{
		{
			[]int{0, 10, 1},
			[]int{10, 1, 0},
		},
		{
			[]int{0, -10, 1},
			[]int{1, 0, -10},
		},
	}

	for _, _test := range _inputs_asc {

		_inp := _test.Input.([]int)

		SortedInPlace(_inp, true)

		assert.Equal(
			t,
			_test.Expected,
			_inp,
			"",
		)
	}

	for _, _test := range _inputs_desc {

		_inp := _test.Input.([]int)

		SortedInPlace(_inp, false)

		assert.Equal(
			t,
			_test.Expected,
			_inp,
			"",
		)
	}

}

func TestPop(t *testing.T) {

	_inputs := []Test{
		{
			[]int{0, 10, 1},
			[]int{0, 10},
		},
		{
			[]int{0, -10, 1},
			[]int{0, -10},
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]int)

		temp := _inp[len(_inp)-1]

		n, _inp := Pop(_inp)

		assert.Equal(
			t,
			temp,
			n,
			"",
		)

		assert.Equal(
			t,
			_test.Expected,
			_inp,
			"",
		)
	}

}

func TestPopLeft(t *testing.T) {

	_inputs := []Test{
		{
			[]int{0, 10, 1},
			[]int{10, 1},
		},
		{
			[]int{0, -10, 1},
			[]int{-10, 1},
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]int)

		temp := _inp[0]

		n, _inp := PopLeft(_inp)

		assert.Equal(
			t,
			temp,
			n,
			"",
		)

		assert.Equal(
			t,
			_test.Expected,
			_inp,
			"",
		)
	}

}

func TestZip(t *testing.T) {

	type I struct {
		A []int
		B []string
	}

	_inputs := []Test{
		{
			I{
				[]int{1, 2, 3},
				[]string{"hi", "there"},
			},
			[][2]interface{}{
				{1, "hi"},
				{2, "there"},
			},
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.(I)

		res := Zip(_inp.A, _inp.B)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}
