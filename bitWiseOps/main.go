package main

import "fmt"

func main() {

	// <<
	x := 9
	//Agrega dos ceros a la derecha de la representacion binaria del numero
	//Cada bit agragado a la derecha es equivalente a multiplicar el numero por dos
	//Es decir que x << n = x * 2^n
	shiftedX := x << 2
	fmt.Println(multiplyBy2ToThe(5, 2))
	fmt.Println(multiplyBitWise(5, 2))
	fmt.Printf("%b << 2 = %v \n", x, shiftedX)
	
	// >>
	y := 9
	//Quida n bits de la representacion binaria del numero
	//Cada bit quitado es equivalente a dividir por 2 redondeado hacia abajo
	shiftedY := y >> 2
	fmt.Printf("%b << 2 = %b \n", y, shiftedY)

	// &
	//Toma la representacion binaria de dos numeros y aplica un AND bit por bit
	//Ejemplo:
	// 1101 & 1011 = 1 & 1, 1 & 0, 0 & 1, 1 & 1 = 1001 
	x = 5
	y = 6
	fmt.Printf("%b & %b = %b\n", x, y, x & y )

	// |
	//Lo mismo pero con OR
	//Ejemplo:
	// 1101 | 1011 = 1 | 1, 1 | 0, 0 | 1, 1 | 1 = 1111
	fmt.Printf("%b | %b = %b\n", x, y, x | y )
}

//Esto es f(n) = a * 2^n
//O(n)
func multiplyBy2ToThe(a int, n int) int {

	res := a
	i := 1

	for i <= n {
		res *= 2
		i++ 
	}

	return res
}

//Lo mismo pero aprovechando "<<""
//O(1)
func multiplyBitWise(a int, n int) int {
	return a << n//Notece como GO hace la operacion con los binarios y parcea al entero luego.
}
