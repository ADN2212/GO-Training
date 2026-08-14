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

type stast struct {
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

	intervals := []interval{
		{
			start: 0,
			end:   5_000_000_007,
		},
		{
			start: 5_000_000_008,
			end:   10_000_000_000,
		},
		{
			start: 10_000_000_001,
			end:   12_850_000_000,
		},
	}

	i := 1

	statsMap := map[string]stast{}

	for _, interval := range intervals {
		wg.Add(1)
		go func() {
			section := csv.NewReader(io.NewSectionReader(f, interval.start, interval.end))
			section.Comma = ';'

			for {
				row, err := section.Read()
				if errors.Is(err, io.EOF) {
					break
				}

				name := row[0]
				tempStr := row[1]

				mu.Lock()
				cityStats, ok := statsMap[name]
				val, err := strconv.ParseFloat(tempStr, 64)
				if err != nil {
					panic(err.Error())
				}

				if !ok {
					statsMap[name] = stast{
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
				mu.Unlock()
			}

			wg.Done()
			fmt.Printf("Go routine No. %d completada. \n", i)
			mu.Lock()
			i = i + 1
			mu.Unlock()
		}()
	}
	wg.Wait()

	names := []string{}
	for name := range statsMap {
		names = append(names, name)
	}

	slices.Sort(names) //Esto no deberia costar mucho porque las keys anidadas son relativamente pocas ...
	currMean := float64(0)

	for _, name := range names {
		stat, _ := statsMap[name]
		currMean = stat.sum / float64(stat.count)
		fmt.Printf("%v;%v;%v;%v \n", name, stat.min, currMean, stat.max)
	}

	fmt.Printf("Completado en %f segundos \n", time.Since(start).Seconds())
}
