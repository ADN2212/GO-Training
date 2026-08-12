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

type interval struct {
	start uint64
	end   uint64
}

func main() {
	start := time.Now()

	intervals := []interval{
		{
			start: 1,
			end:   25_000,
		},
		{
			start: 25_001,
			end:   50_000,
		},
		{
			start: 50_001,
			end:   75_000,
		},
		{
			start: 75_001,
			end:   100_000,
		},
	}

	rads := []pair{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, inter := range intervals {
		wg.Add(1)
		go func() {
			n := uint64(inter.start)
			for n <= inter.end {
				mu.Lock()
				rads = append(rads, pair{
					n:      n,
					radVal: rad(n),
				})
				mu.Unlock()
				n += 1
			}
			wg.Done()
			fmt.Printf("Go routine no %v completed \n", i+1)
		}()
	}
	wg.Wait()
	
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
