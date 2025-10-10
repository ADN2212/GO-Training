package main

import "fmt"

func main() {
	top := 100
	var currT int 
	counter := 0

	for n := 2; n <= top; n++ {
		currT = t(n)
		if isPrime(currT) {
			fmt.Printf("t(%v) = %v\n", n, currT)
			counter += 1
		} else {
			fmt.Printf("t(%v) = %v NO\n", n, currT)
		}
	} 

	fmt.Println(counter)
}


func t(n int) int {
	return 2 * (n * n) - 1
}

func isPrime(n int) bool {
	
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n % i == 0 {
			return false
		}
	}

	return true
}
