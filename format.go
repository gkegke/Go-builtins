package builtins

import (
	"strings"
	"text/template"
)

/*

Complex String

Format

*/

/*
Format takes a string, parses it for identifiers, and
replaces them with related data.

Intended to be be very similar to Python's string format function.

IMPORTANT: Uses text/template, so all it's indentifiers work.

Given map[string]{ "Name": "Bob" }

"{{ .Name }}" => "Bob"

and

"Hi my name is {{ .Name }}" => "Hi my name is Bob"
*/
func Format(fstring string, data map[string]interface{}) string {
	t := template.Must(template.New("fstring").Parse(fstring))

	sbuilder := &strings.Builder{}

	if err := t.Execute(sbuilder, data); err != nil {
		panic(err)
	}

	return sbuilder.String()
}
