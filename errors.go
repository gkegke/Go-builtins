package builtins

import (
	"strings"
	"text/template"
	"time"
)

/*

Errors

//Types
Error

// Functions
Check
E

*/

// Provides common fields required for basic errors
type Error struct {
	Where string
	When  string
	What  string
	Err   error
}

const _ERR_MSG_TEMPLATE = `
Where: {{ .Where }}
When: {{ .When }}
What: {{ .What }}
Err: {{ .Err }}
`

func (e *Error) Error() string {
	t := template.Must(template.New("err").Parse(_ERR_MSG_TEMPLATE))

	sbuilder := &strings.Builder{}

	err := t.Execute(sbuilder, map[string]interface{}{
		"Where": e.Where,
		"When":  e.When,
		"What":  e.What,
		"Err":   e.Err,
	})

	Check(err)

	return sbuilder.String()
}

func Check(err error) {

	if err != nil {
		panic(err)
	}

}

// A helper function for easy creation of errors.
func E(where string, what string, err error) error {
	e := &Error{
		Where: where,
		When:  time.Now().String(),
		What:  what,
		Err:   err,
	}

	return e
}
