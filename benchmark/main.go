package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/seiflotfy/pmc"
)

func main() {
	file, _ := os.Open("zipf.csv")

	r := csv.NewReader(file)
	var expected []uint
	r.Comma = ';'

	pmc, _ := pmc.New(8000000, 64, 64)
	//cml, _ := cml.NewSketch16ForEpsilonDelta(0.00000543657, 0.99)

	dur := time.Duration(0)
	x := 0
	for {
		record, err := r.Read()
		x++
		if x == 1 {
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		id := fmt.Sprintf("flow-%s", record[0])
		counts, _ := strconv.ParseFloat(record[1], 64)
		expected = append(expected, uint(counts))
		for i := 0.0; i < counts; i++ {
			start := time.Now()
			pmc.Increment([]byte(id))
			//cml.IncreaseCount([]byte(id))
			dur += time.Since(start)
		}
	}

	for i := range expected {
		id := fmt.Sprintf("flow-%d", i)
		// flow id, expected, estimation
		est := pmc.GetEstimate([]byte(id))
		//est := cml.Frequency([]byte(id))
		fmt.Println(id, expected[i], uint(est), est/float64(expected[i]))

		if i > 10 {
			break
		}
	}
	fmt.Println("fill rate:", pmc.GetFillRate())
	//fmt.Println("fill rate:", cml.GetFillRate(), dur)

}
