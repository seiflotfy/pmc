package pmc

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 10000; i++ {
		r := rand(32)
		if r >= 32 {
			fmt.Println(">>>", r)
		}
	}
}
