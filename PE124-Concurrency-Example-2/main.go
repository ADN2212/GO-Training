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

	//Dependiendo del tamanio de estos intervalos las go routines que los ejecuten terminaran antes o despues
	//vease como en este caso el primer intervalo es el mas largo y por tanto su gr el la ultima en terminar.
	intervals := []interval{
		{
			start: 1,
			end:   50_000,
		},
		{
			start: 50_001,
			end:   60_000,
		},
		{
			start: 60_001,
			end:   70_000,
		},
		{
			start: 70_001,
			end:   100_000,
		},
	}

	//rads := []pair{}
	rads := make([]pair, 100_000)//Declarar el slice de esta manera permite quitar el mutex
	var wg sync.WaitGroup
	//var mu sync.Mutex

	//Ojo: al parecer esta version es mas lenta que la de las go routines separadas
	for i, inter := range intervals {
		wg.Add(1)
		go func() {
			n := uint64(inter.start)
			for n <= inter.end {

				// mu.Lock() //Todo bloque de codigo entre esta linea y UnLock esta protegida againt race conditios
				// rads = append(rads, pair{
				// 	n:      n,
				// 	radVal: rad(n),
				// })
				// mu.Unlock()

				//Usar estas lineas cuando se declare el slice de ante mano.
				//como el slice tiene un tamanio fijo de ante mano no hay que usar append
				//lo cual elimina la posibilidad de que hayan rece conditions
				//por por tanto no se hace necesario el uso de mutex 
				rads[n-1].n = n
				rads[n-1].radVal = rad(n)
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
