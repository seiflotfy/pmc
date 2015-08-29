package pmc

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"code.google.com/p/gofarmhash"

	"github.com/lukut/bitmaps"
)

/*
Sketch ...
*/
type Sketch struct {
	l float64
	m float64
	w float64
	B bitmaps.Bitmap
}

/*
New ...
*/
func New(l uint, m uint, w uint, maxFlows uint) (*Sketch, error) {
	if l == 0 {
		return nil, errors.New("Expected l > 0, got 0")
	}
	if m == 0 {
		return nil, errors.New("Expected m > 0, got 0")
	}
	if w == 0 {
		return nil, errors.New("Expected w > 0, got 0")
	}
	return &Sketch{float64(l), float64(m), float64(w), make(bitmaps.Bitmap, l/8)}, nil
}

func (sketch *Sketch) getHash(flow []byte, i, j uint) []byte {
	bi := strconv.Itoa(int(i))
	bj := strconv.Itoa(int(j))
	f := append(flow, bi...)
	f = append(f, bj...)
	return f
}

func (sketch *Sketch) getPos(f []byte) uint {
	hf := uint(farmhash.Hash64(f))
	return hf % uint(sketch.l)
}

/*
Add ...
*/
func (sketch *Sketch) Add(flow []byte) {
	i := rand(uint(sketch.m))
	j := georand(uint(sketch.w))
	f := sketch.getHash(flow, i, j)
	pos := sketch.getPos(f)
	sketch.B.Set(pos, true)
}

func (sketch *Sketch) getZSum(flow []byte) uint {
	z := 0.0
	for i := 0.0; i < sketch.m; i++ {
		j := 0.0
		for j < sketch.w {
			f := sketch.getHash(flow, uint(i), uint(j))
			pos := sketch.getPos(f)
			if sketch.B.Get(pos) == false {
				break
			}
			j++
		}
		z += j
	}
	return uint(z)
}

func (sketch *Sketch) getEmptyRows(flow []byte) uint {
	k := uint(0)
	for i := 0.0; i < sketch.m; i++ {
		f := sketch.getHash(flow, uint(i), 0)
		pos := sketch.getPos(f)
		if sketch.B.Get(pos) == false {
			k++
		}
	}
	return k
}

func (sketch *Sketch) getP() float64 {
	ones := float64(0)
	for i := uint(0); i < uint(sketch.B.Size()); i++ {
		if sketch.B.Get(i) == true {
			ones++
		}
	}
	return ones / float64(sketch.l)
}

/*
GetEstimate ...
*/
func (sketch *Sketch) GetEstimate(flow []byte) uint {
	m := sketch.m
	p := sketch.getP()
	k := float64(sketch.getEmptyRows(flow))
	// Use const due to quick conversion against 0.78 (n = 1000000.0)
	//n := -2 * m * math.Log((k)/(m*(1-p)))
	n := 100000.0

	// Dealing with small multiplicities
	if k/(1-p) > 0.3*m {
		return uint(-2 * m * math.Log(k/(m*(1-p))))
	}

	qk := func(k, n, p float64) float64 {
		result := 1.0
		for i := 1.0; i <= k; i++ {
			result *= (1.0 - math.Pow(1.0-math.Pow(2, -i), n)*(1.0-p))
		}
		return result
	}

	E := func(n, p float64) float64 {
		result := float64(0)
		for k := 1.0; k <= sketch.w; k++ {
			result += (k * (qk(k, n, p) - qk(k+1, n, p)))
		}
		return result
	}
	rho := func(p float64) float64 {
		return math.Pow(2, E(n, p)) / n
	}

	z := float64(sketch.getZSum(flow))

	fmt.Println("p:", p)
	fmt.Println("k:", k)
	fmt.Println("n:", n)
	fmt.Println("z:", z)
	return uint(m * math.Pow(2, z/m) / rho(p))
}
