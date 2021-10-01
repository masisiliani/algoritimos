package main

var (
	Counter map[int]int = make(map[int]int)
)

func Fibonacci(number int) int {
	// variável Counter é usada para sabermos quantas vezes a função Fibonacci(number) é chamada para determinado numero n
	Counter[number]++

	//Essa condição valida se o número passado por parâmtero é uma das constantes
	//definidas pela fórmula de Fibonacci. Se o valor de "number" for 0 ou 1,
	//o mesmo valor é retornado
	if number <= 1 {
		return number
	}

	return Fibonacci(number-1) + Fibonacci(number-2)
}
