package pmc

import (
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
	r := uint(1)
	for val&0x8000000000000000 == 0 && r <= w {
		r++
		val <<= 1
	}
	return uint(r - 1)
}

func rand(m uint) uint {
	return uint(random.Int()) % m
}
