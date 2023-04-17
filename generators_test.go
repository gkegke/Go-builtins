package builtins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {

	type I struct {
		Start int64
		Stop  int64
		Step  int64
	}

	_inputs := []Test{
		{
			I{0, 10, 1},
			[]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			I{0, 10, 2},
			[]int64{0, 2, 4, 6, 8},
		},
		{
			I{10, 0, -2},
			[]int64{10, 8, 6, 4, 2},
		},
		{
			I{0, 10, -2},
			[]int64{},
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.(I)

		res := Range(_inp.Start, _inp.Stop, _inp.Step)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}
