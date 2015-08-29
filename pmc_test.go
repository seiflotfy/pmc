package pmc

import (
	"math"
	"testing"
)

func TestPMCHash(t *testing.T) {
	s, _ := New(1024, 4, 4)
	dist := make(map[uint]uint)
	for k := 0; k < 100000; k++ {
		i := rand(uint(s.m))
		j := georand(uint(s.w))
		pos := s.getPos([]byte("pmc"), i, j)
		dist[pos]++
	}
	if len(dist) > 16 {
		t.Error("Expected maximum 16 different positions, got ", len(dist))
	}
}

func TestPMCHashAdd(t *testing.T) {
	flow := []byte("pmc")
	halfFlow := []byte("halfpmc")
	s, _ := NewForMaxFlows(10000000)
	//start := time.Now()
	for k := 0; k < 1000000; k++ {
		s.Increment(flow)
		if k%2 == 0 {
			s.Increment(halfFlow)
		}
	}
	//elapsed := time.Since(start)
	//fmt.Println(1500000, "x Incrememt took", elapsed)

	//start = time.Now()
	hfCount := s.GetEstimate(halfFlow)
	fCount := s.GetEstimate(flow)

	fErr := 100 * (1 - float64(fCount)/1000000)
	hfErr := 100 * (1 - float64(hfCount)/500000)
	if math.Abs(fErr) > 10 {
		t.Errorf("Expected error for flow 'flow' <= 10%%, got %f", math.Abs(fErr))
	}
	if math.Abs(hfErr) > 10 {
		t.Errorf("Expected error for flow 'flow' <= 10%%, got %f", math.Abs(hfErr))
	}

	//elapsed = time.Since(start)
	//fmt.Println("GetEstimate took", elapsed)
}
