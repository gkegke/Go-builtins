/*
builtins provides a set of useful functions and types
that enable Go to be 'easier' to use in a casual
scripting setting.
*/
package builtins

/*

Contents:

 - All
 - Any

*/

/*
Args:

  - iterable []bool

Return

  - bool

Checks whether all the values in a slice of bools
are all true.

# Gotcha

python's all([]) returns True.

Not sure this should be the case here, but since this
is meant as a python builtins replacement, I shall stick
to their plans.
*/
func All(iterable []bool) bool {

	for _, el := range iterable {

		if !el {

			return false

		}

	}

	return true
}

/*
Args:

  - iterable []bool

Return

  - bool

Checks whether at least one of the values in a
slice are true.
*/
func Any(iterable []bool) bool {

	for _, el := range iterable {

		if el {

			return true

		}

	}

	return false

}
