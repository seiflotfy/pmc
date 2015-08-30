package pmc

import (
	"math"
	random "math/rand"
	"strconv"
	"testing"
)

func TestPMCHash(t *testing.T) {
	s, _ := New(1024, 4, 4)
	dist := make(map[uint]uint)
	for k := 0; k < 100000; k++ {
		i := float64(rand(uint(s.m)))
		j := float64(georand(uint(s.w)))
		pos := s.getPos([]byte("pmc"), i, j)
		dist[pos]++
	}
	if len(dist) > 16 {
		t.Error("Expected maximum 16 different positions, got ", len(dist))
	}
}

func TestPMCHashAdd(t *testing.T) {
	flows := make([]string, 100, 100)

	for i := 0; i < len(flows); i++ {
		flows[i] = strconv.Itoa(random.Int())
	}

	s, _ := New(8000000, 256, 32)
	for j := range flows {
		for i := 0; i < 1000000; i++ {
			if i%(j+1) == 0 {
				s.Increment([]byte(flows[j]))
			}
		}
	}

	for i, v := range flows {
		fCount := s.GetEstimate([]byte(v))
		fErr := math.Abs(100 * (1 - float64(fCount)/(1000000/float64(i+1))))
		if math.Abs(fErr) > 13 {
			t.Errorf("Expected error for flow %d '%s' <= 13%%, got %f", i, v, math.Abs(fErr))
		}
	}
}

func TestRand(t *testing.T) {
	for i := 0; i < 10000; i++ {
		r := rand(32)
		if r >= 32 {
			t.Error("Expected rand to return r < 32, got", r)
		}
	}
}
