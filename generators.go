package builtins

/*

Generators

- Range

*/

/*
Args

  - start, stop, step Number

A copy of python's range function, for the creation of
a range of numerical values.

The step argument allows the creation of more 'dynamic'
ranges of values.
*/
func Range(start, stop, step int64) []int64 {

	if step > 0 {

		res := make([]int64, (stop-start+step-1)/step)

		for i, n := 0, start; n < stop; i, n = i+1, n+step {
			res[i] = n
		}

		return res
	}

	size := (stop - start + step + 1) / step

	if size < 0 {
		size = 0
	}

	res := make([]int64, size)

	for i, n := 0, start; n > stop; i, n = i+1, n+step {
		res[i] = n
	}

	return res

}
