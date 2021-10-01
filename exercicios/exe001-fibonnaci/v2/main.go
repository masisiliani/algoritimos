package main

var (
	Counter map[int]int = make(map[int]int)
)

func Fibonacci(number int) int {
	//initialize cache variable
	baseCases := map[int]int{
		0: 0,
		1: 1,
	}

	return computeCache(number, baseCases)
}

func computeCache(number int, cache map[int]int) int {
	if value, found := cache[number]; found {
		return value
	}
	cache[number] = computeCache(number-1, cache) + computeCache(number-2, cache)
	Counter[number]++
	return cache[number]
}
