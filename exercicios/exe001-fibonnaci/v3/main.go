package main

func Fibonacci(number int) int {
	if number <= 1 {
		return number
	}

	cache := make(map[int]int)
	// preencher a cache com os valores iniciais
	cache[0] = 0
	cache[1] = 1

	//interação para preencher até o index passado por parametro
	for currentIndice := 2; currentIndice <= number; currentIndice++ {
		lastIndice := currentIndice - 1
		penultIndice := currentIndice - 2

		cache[currentIndice] = cache[lastIndice] + cache[penultIndice]
	}

	return cache[number]

}
