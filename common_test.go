package builtins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {

	_inputs := []Test{
		{
			[]bool{},
			true,
		},
		{
			[]bool{true},
			true,
		},
		{
			[]bool{true, false},
			false,
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]bool)

		res := All(_inp)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestAny(t *testing.T) {

	_inputs := []Test{
		{
			[]bool{},
			false,
		},
		{
			[]bool{true},
			true,
		},
		{
			[]bool{true, false},
			true,
		},
		{
			[]bool{false, false},
			false,
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.([]bool)

		res := Any(_inp)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}
