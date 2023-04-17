package builtins

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {

	_inputs := []Test{
		{
			0.0,
			0.0,
		},
		{
			1.0,
			1.0,
		},
		{
			-1.0,
			1.0,
		},
		{
			-123.123341,
			123.123341,
		},
		{
			123.123341,
			123.123341,
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.(float64)

		res := Abs(_inp)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestRound(t *testing.T) {

	type I struct {
		N      float64
		Digits float64
	}

	_inputs := []Test{
		{
			I{3.14, 0},
			3.14,
		},
		{
			I{3.14, 1},
			3.1,
		},
		{
			I{31.14, -1},
			30.0,
		},
		{
			I{431.14, -2},
			400.0,
		},
		{
			I{9123431.145235513, 6},
			9123431.145236,
		},
		{
			I{9123431.145235513, -4},
			9120000.0,
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.(I)

		res := Round(_inp.N, _inp.Digits)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestDivMod(t *testing.T) {

	type I struct {
		x float64
		y float64
	}

	type O struct {
		Modulus   float64
		Remainder float64
	}

	_nan := math.NaN()

	_inputs := []Test{
		{
			I{0.0, 0.0},
			O{_nan, _nan},
		},
		{
			I{10.0, 1.0},
			O{10.0, 0.0},
		},
		{
			I{10.0, 2.0},
			O{5.0, 0.0},
		},
		{
			I{10.0, 2.0},
			O{5.0, 0.0},
		},
		{
			I{10.5, 2.0},
			O{5.0, 0.5},
		},
	}

	for i, _test := range _inputs {

		_inp := _test.Input.(I)

		m, r := DivMod(_inp.x, _inp.y)

		if i == 0 {

			assert.True(t, math.IsNaN(m))
			assert.True(t, math.IsNaN(r))
			continue
		}

		assert.Equal(
			t,
			_test.Expected.(O),
			O{m, r},
			"",
		)
	}

}

func TestBin(t *testing.T) {

	_inputs := []Test{
		{
			0,
			"0",
		},
		{
			1,
			"1",
		},
		{
			100,
			"1100100",
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.(int)

		res := Bin(int64(_inp))

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestHex(t *testing.T) {

	_inputs := []Test{
		{
			0,
			"0",
		},
		{
			1,
			"1",
		},
		{
			100,
			"64",
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.(int)

		res := Hex(int64(_inp))

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}

func TestOct(t *testing.T) {

	_inputs := []Test{
		{
			0,
			"0",
		},
		{
			1,
			"1",
		},
		{
			100,
			"144",
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.(int)

		res := Oct(int64(_inp))

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}
