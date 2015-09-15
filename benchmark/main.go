package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/seiflotfy/pmc"
)

func main() {
	file, _ := os.Open("zipf.csv")

	r := csv.NewReader(file)
	var expected []uint
	r.Comma = ';'

	pmc, _ := pmc.New(8000000, 256, 320)

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
			pmc.Increment([]byte(id))
		}
	}

	for i := range expected {
		id := fmt.Sprintf("flow-%d", i)
		// flow id, expected, estimation
		est := pmc.GetEstimate([]byte(id))
		fmt.Println(id, expected[i], uint(est), est/float64(expected[i]))
		if i == 10 {
			break
		}
	}
	fmt.Println("fill rate:", pmc.GetFillRate())
}
