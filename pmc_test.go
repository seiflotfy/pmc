package pmc

import (
	"math"
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
	flows := make([]string, 5, 5)
	s, _ := NewForMaxFlows(100000)
	for j := range flows {
		flows[j] = "flow" + strconv.Itoa(j)
		for i := 0; i < 1000000; i++ {
			if i%(j+1) == 0 {
				s.Increment([]byte(flows[j]))
			}
		}
	}
	for i, v := range flows {
		fCount := s.GetEstimate([]byte(v))
		fErr := 100 - (100 * (1 - float64(fCount)/(10000000/float64(i+1))))
		if math.Abs(fErr) > 11 {
			t.Errorf("Expected error for flow '%s' <= 10%%, got %f", v, math.Abs(fErr))
		}
	}
}
