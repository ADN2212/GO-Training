package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"sync"
	"time"
)

type stat struct {
	min   float64
	sum   float64
	max   float64
	count int
}

type cityRow struct {
	name string
	temp float64
}

type interval struct {
	start int64
	end   int64
}

func main() {
	start := time.Now()

	fmt.Println("Abriendo al chivo ..")
	f, err := os.Open("measurements.txt")
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		fmt.Println("Cerrando al chivo")
		f.Close()
	}()

	var wg sync.WaitGroup
	var mu sync.Mutex
	
	//La idea aqui es dividir el slice en varias partes para poder enviar una gorutine por cada intervalo.
	intervals := []interval{
		{
			start: 0,
			end:   2_000_000_004,
		},
		{
			start: 2_000_000_005,
			end:   4_000_000_002,
		},
		{
			start: 4_000_000_003,
			end:   6_000_000_005,
		},
		{
			start: 6_000_000_007,
			end:   8_000_000_000,
		},
		{
			start: 8_000_000_002,
			end:   10_000_000_000,
		},
		{
			start: 10_000_000_001,
			end:   12_850_000_000,
		},
	}

	statsMaps := []map[string]stat{}
	for _, interval := range intervals {
		wg.Add(1)
		//Por cada itervalo se lanza una gorutine que ejecuta el mismo algoritmo que se haria en un single treat.
		go func() {
			section := csv.NewReader(io.NewSectionReader(f, interval.start, interval.end))
			section.Comma = ';'
			statsMap := map[string]stat{}

			for {
				row, err := section.Read()
				if errors.Is(err, io.EOF) {
					break
				}

				if len(row) != 2 {
					fmt.Println("Broken line skipped ...")
					fmt.Println(row)
					fmt.Println("-----------------------------------")
					continue
				}

				name := row[0]
				tempStr := row[1]

				cityStats, ok := statsMap[name]
				val, err := strconv.ParseFloat(tempStr, 64)
				if err != nil {
					panic(err.Error())
				}

				if !ok {
					statsMap[name] = stat{
						min:   val,
						sum:   val,
						max:   val,
						count: 1,
					}
				} else {
					cityStats.count += 1
					cityStats.sum += val
					if val > cityStats.max {
						cityStats.max = val
					}
					if val < cityStats.min {
						cityStats.min = val
					}
					statsMap[name] = cityStats
				}

			}
			//Si no se lockea la operacion append  el algoritmo deja de ser correcto, varios append al mismo tiempo generan una race condition en le array.
			mu.Lock()
			statsMaps = append(statsMaps, statsMap)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	//El hecho de que se lanzen varios go rutines hace que sea necesario hacer merge a los resultados, cosa que tambien toma tiempo.
	result := map[string]stat{}
	for _, m := range statsMaps {
		for key, st := range m {
			totalStat, ok := result[key]
			if !ok {
				result[key] = st
			} else {
				//Ojo: aqui se esta repitiendo trabajo que ya se hice en cada una de las gorutine.
				newMin := totalStat.min
				if st.min < newMin {
					newMin = st.min
				}
				newMax := totalStat.max
				if st.max > newMin {
					newMax = st.max
				}
				result[key] = stat{
					sum:   totalStat.sum + st.sum,
					count: totalStat.count + st.count,
					min:   newMin,
					max:   newMax,
				}
			}
		}
	}

	names := []string{}
	for name := range result {
		names = append(names, name)
	}

	//Lo preferible serai que este sort se haga "in place"
	slices.Sort(names) //Esto no deberia costar mucho porque las keys anidadas son relativamente pocas ...
	currMean := float64(0)

	for _, name := range names {
		stat, _ := result[name]
		currMean = stat.sum / float64(stat.count)
		fmt.Printf("%v;%v;%v;%v \n", name, stat.min, currMean, stat.max)
	}

	fmt.Printf("Un billon(1,000,000,000) de lineas procesadas en %f segundos \n", time.Since(start).Seconds())
}
