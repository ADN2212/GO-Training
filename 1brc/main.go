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

	statsMaps := []map[string]stast{}
	for _, interval := range intervals {
		wg.Add(1)
		go func() {
			section := csv.NewReader(io.NewSectionReader(f, interval.start, interval.end))
			section.Comma = ';'
			statsMap := map[string]stast{}

			for {
				row, err := section.Read()
				if errors.Is(err, io.EOF) {
					break
				}

				name := row[0]
				tempStr := row[1]

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

			}
			mu.Lock()
			statsMaps = append(statsMaps, statsMap)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	result := map[string]stast{}
	for _, m := range statsMaps {
		for key, stat := range m {
			totalStat, ok := result[key]
			if !ok {
				result[key] = stat
			} else {
				newMin := totalStat.min
				if stat.min < newMin {
					newMin = stat.min
				}
				newMax := totalStat.max
				if stat.max > newMin {
					newMax = stat.max
				}
				result[key] = stast{
					sum:   totalStat.sum + stat.sum,
					count: totalStat.count + stat.count,
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

	slices.Sort(names) //Esto no deberia costar mucho porque las keys anidadas son relativamente pocas ...
	currMean := float64(0)

	for _, name := range names {
		stat, _ := result[name]
		currMean = stat.sum / float64(stat.count)
		fmt.Printf("%v;%v;%v;%v \n", name, stat.min, currMean, stat.max)
	}

	fmt.Printf("Un billon(1,000,000,000) de lineas procesadas en %f segundos \n", time.Since(start).Seconds())
}
