package parallels

import (
	"sync"
)

func ParallelMap[I any, O any](list []I, mapper func(I) O) []O {
	result := make([]O, len(list))
	var wg sync.WaitGroup
	mapperFunc := func(element I, index int) {
		defer wg.Done()
		// This should be thread safe since only one coroutine is writing to an particular index
		// This implementation saves significant overhead when compared to using channels
		// Please be extra careful about race conditions when modifying this code
		result[index] = mapper(element)
	}

	wg.Add(len(list))
	for index, element := range list {
		go mapperFunc(element, index)
	}
	wg.Wait()

	return result
}

func ParallelFilter[I any](list []I, filter func(I) bool) []I {
	result := make([]I, 0, len(list))
	filterList := ParallelMap(list, filter)
	for index, filter := range filterList {
		if filter {
			result = append(result, list[index])
		}
	}
	return result
}
