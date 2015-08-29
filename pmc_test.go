package pmc

import (
	"fmt"
	"testing"
)

func TestPMCHash(t *testing.T) {
	s, _ := New(1024, 4, 4, 1)
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
	s, _ := New(32000000, 256, 32, 0)
	for k := 0; k < 1000000; k++ {
		s.Add(flow)
		if k%2 == 0 {
			s.Add([]byte("test"))
		}
	}
	//printVirtualMatrix(s, flow)
	fmt.Println(">>", s.GetEstimate([]byte("test")))
	fmt.Println(">>", s.GetEstimate([]byte("pmc")))
}
