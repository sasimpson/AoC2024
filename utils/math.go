package utils

import "golang.org/x/exp/constraints"

type summable interface {
	constraints.Float | constraints.Integer
}

func Sum[T summable](vals []T) T {
	var sum T
	for _, v := range vals {
		sum = sum + v
	}
	return sum
}
