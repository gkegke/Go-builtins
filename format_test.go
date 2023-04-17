package builtins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {

	type I struct {
		S    string
		Data map[string]interface{}
	}

	_inputs := []Test{
		{
			I{
				"Hi {{ .Name }}",
				map[string]interface{}{
					"Name": "Bob",
				},
			},
			"Hi Bob",
		},
		{
			I{
				"Hi {{ .Name }}. You are {{ .Age }} years old.",
				map[string]interface{}{
					"Name": "Bob",
					"Age":  123,
				},
			},
			"Hi Bob. You are 123 years old.",
		},
	}

	for _, _test := range _inputs {

		_inp := _test.Input.(I)

		res := Format(_inp.S, _inp.Data)

		assert.Equal(
			t,
			_test.Expected,
			res,
			"",
		)
	}

}
