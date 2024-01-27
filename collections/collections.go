package collections

func OfSlice[K any](items ...K) []K {
	return items
}

func FromKeys[K comparable](keys ...K) map[K]rune {
	result := make(map[K]rune, len(keys))
	for _, k := range keys {
	 result[k] = 0
	}
	return result
}
