package builtins

import (
	"golang.org/x/exp/constraints"
)

// Generics type for numerical data
type Number interface {
	constraints.Float | constraints.Integer
}
