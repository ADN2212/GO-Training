package main

import (
	"fmt"
	"time"
)

var memo = make(map[int]int)

func main() {

	n := 1
	nMax := 0
	max := -100

	inicio := time.Now()

	for n < 1_000_000 {
		curr := memo_collazt(n)
		if curr > max {
			nMax = n
			max = curr
		}
		n += 1
	}

	fmt.Printf("The number that put the longest sec is %v \n", nMax)
	fmt.Printf("Tiempo de ejecución: %v\n", time.Since(inicio).Milliseconds())
}

func memo_collazt(n int) int {

	next := 0

	if n % 2 == 0 {
		next = n / 2
	} else {
		next = 3 * n + 1
	}

	if memo[next] != 0 {
		memo[n] = 1 + memo[next]
		return  memo[n]
	}

	current := n
	counter := 1
	
	for current != 1 {
		if current % 2 == 0 {
			current = current / 2
		} else {
			current = 3 * current + 1
		}

		counter += 1
	}

	memo[n] = counter
	return  counter
}
