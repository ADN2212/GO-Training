package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"time"
)

type stast struct {
	min   float64
	sum   float64
	max   float64
	count int
}

func main() {
	start := time.Now()

	f, err := os.Open("measurements.txt")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Abriendo al chivo ..")
	defer func() {
		fmt.Println("Cerrando al chivo")
		f.Close()
	}()


	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'


	csvReader.InputOffset()

	statsMap := map[string]stast{}
	//i := 1
	fmt.Println("Starting loop ...")
	for {

		row, err := csvReader.Read()
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
			//fmt.Printf("Adding %v to stats map \n", name)
			statsMap[name] = stast{
				min:   val,
				sum:   val,
				max:   val,
				count: 1,
			}
		} else {
			//fmt.Printf("Updating %v value \n", name)
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
		//i += 1
	}
	fmt.Println("Loop ended ...")

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
