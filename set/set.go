package set

import (
	"slices"

	"github.com/mohitudupa/snowshoes/collections"
)

func Unique[K comparable](a []K) []K {
	set := collections.FromKeys(a...)

	result := make([]K, 0, len(set))
	for k := range set {
		result = append(result, k)
	}

	return result
}

func Intersection[K comparable](a, b []K) []K {
	right := collections.FromKeys(b...)
	result := make([]K, 0, min(len(a), len(b)))

	for _, k := range a {
		if _, ok := right[k]; ok {
			result = append(result, k)
		}
	}

	return result
}

func Difference[K comparable](a, b []K) []K {
	right := collections.FromKeys(b...)
	result := make([]K, 0, len(a))

	for _, k := range a {
		if _, ok := right[k]; !ok {
			result = append(result, k)
		}
	}

	return result
}

func Union[K comparable](a, b []K) []K {
	left := collections.FromKeys(a...)
	slices.Grow(a, len(b))

	for _, k := range b {
		if _, ok := left[k]; !ok {
			a = append(a, k)
		}
	}

	return a
}
