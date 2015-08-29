package pmc

import (
	"fmt"
	"hash/fnv"
	random "math/rand"
	"strconv"
)

func georand(w uint) uint {
	hasher := fnv.New64a()
	i := random.Int()
	hasher.Write([]byte(strconv.Itoa(i)))
	val := hasher.Sum64()
	// Calculate the position of the leftmost 1-bit.
	r := uint(0)
	for val&0x8000000000000000 == 0 && r < w-1 {
		r++
		val <<= 1
	}
	return uint(r)
}

func rand(m uint) uint {
	return uint(random.Int()) % m
}

func printVirtualMatrix(s *Sketch, flow []byte) {
	for i := 0.0; i < s.m; i++ {
		for j := 0.0; j < s.w; j++ {
			f := s.getHash([]byte("pmc"), uint(i), uint(j))
			pos := s.getPos(f)
			if s.B.Get(pos) == false {
				fmt.Print(0)
			} else {
				fmt.Print(1)
			}
		}
		fmt.Println("")
	}
}
