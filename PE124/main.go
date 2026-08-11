package main

import (
	"cmp"
	"fmt"
	"slices"
	"sync"
	"time"
)

type pair struct {
	n      uint64
	radVal uint64
}

func main() {
	start := time.Now()

	// rads := []pair{}
	// limit := uint64(100_000)

	// n := uint64(1)
	// for n <= limit {
	// 	rads = append(rads, pair{
	// 		n:      n,
	// 		radVal: rad(n),
	// 	})
	// 	n += 1
	// }

	rads1 := []pair{}
	rads2 := []pair{}
	rads3 := []pair{}
	rads4 := []pair{}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		n := uint64(1)
		for n <= 25_000 {
			rads1 = append(rads1, pair{
				n:      n,
				radVal: rad(n),
			})
			n += 1
		}
		wg.Done()
		//fmt.Println(n)
		fmt.Println("Firts go rutine completed")
	}()

	wg.Add(1)
	go func() {
		n := uint64(25_001)
		for n <= 50_000 {
			rads2 = append(rads2, pair{
				n:      n,
				radVal: rad(n),
			})
			n += 1
		}
		wg.Done()
		//fmt.Println(n)
		fmt.Println("Second go rutine completed")
	}()

	wg.Add(1)
	go func() {
		n := uint64(50_001)
		for n <= 75_000 {
			rads3 = append(rads3, pair{
				n:      n,
				radVal: rad(n),
			})
			n += 1
		}
		wg.Done()
		//fmt.Println(n)
		fmt.Println("Tercera go rutine completed")
	}()

	wg.Add(1)
	go func() {
		n := uint64(75_001)
		for n <= 100_000 {
			rads4 = append(rads4, pair{
				n:      n,
				radVal: rad(n),
			})
			n += 1
		}
		wg.Done()
		//fmt.Println(n)
		fmt.Println("Cuarta go rutine completed")
	}()

	wg.Wait()
	rads := append(rads1, rads2...)
	rads = append(rads, rads3...)
	rads = append(rads, rads4...)
	//fmt.Printf("len = %v \n", len(rads)) //100_000
	//fmt.Println("Despues de que terminaron las go rutines")
	slices.SortFunc(rads, func(p1, p2 pair) int {
		if p1.radVal == p2.radVal {
			return cmp.Compare(p1.n, p2.n)
		}
		return cmp.Compare(p1.radVal, p2.radVal)
	})

	k := 10_000
	fmt.Println(rads[k-1].n) //21417
	fmt.Printf("Completado en %v segundos\n", time.Since(start).Seconds())
}

// O(n)
func rad(n uint64) uint64 {
	factors := []uint64{}

	for n%2 == 0 {
		if !slices.Contains(factors, 2) {
			factors = append(factors, 2)
		}
		n /= 2
	}

	d := uint64(3)
	for d*d <= n {
		for n%d == 0 {
			if !slices.Contains(factors, d) {
				factors = append(factors, d)
			}
			n /= d
		}
		d += 2
	}

	if n > 1 {
		if !slices.Contains(factors, d) {
			factors = append(factors, n)
		}
	}

	product := uint64(1)
	for _, f := range factors {
		product *= f
	}

	return product
}
