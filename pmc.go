package pmc

import (
	"errors"
	"math"
	"strconv"

	"code.google.com/p/gofarmhash"

	"github.com/lukut/bitmaps"
)

/*
Sketch ...
*/
type Sketch struct {
	l      float64
	m      float64
	w      float64
	bitmap bitmaps.Bitmap
}

/*
New returns a PMC Sketch with the properties:
l = total number of bits for sketch
m = total number of rows for each flow
w = total number of columns for each flow
*/
func New(l uint, m uint, w uint) (*Sketch, error) {
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

/*
NewForMaxFlows returns a PMC Sketch adapted to the size of the max number of
flows expected.
*/
func NewForMaxFlows(maxFlows uint) (*Sketch, error) {
	l := maxFlows * 32
	return New(l, 256, 32)
}

/*
It is straightforward to use any uniformly distributed hash function with
sufficiently random output in the role of H: the input parameters can
simply be concatenated to a single bit string.
*/
func (sketch *Sketch) getPos(f []byte, i, j uint) uint {
	s := strconv.Itoa(int(i)) + string(f) + strconv.Itoa(int(j))
	hash := farmhash.Hash64([]byte(s))
	return uint(hash) % uint(sketch.l)
}

/*
Increment the count of the flow by 1
*/
func (sketch *Sketch) Increment(flow []byte) {
	i := rand(uint(sketch.m))
	j := georand(uint(sketch.w))
	pos := sketch.getPos(flow, i, j)
	sketch.bitmap.Set(pos, true)
}

func (sketch *Sketch) getZSum(flow []byte) uint {
	z := 0.0
	for i := 0.0; i < sketch.m; i++ {
		j := 0.0
		for j < sketch.w {
			pos := sketch.getPos(flow, uint(i), uint(j))
			if sketch.bitmap.Get(pos) == false {
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
		pos := sketch.getPos(flow, uint(i), 0)
		if sketch.bitmap.Get(pos) == false {
			k++
		}
	}
	return k
}

func (sketch *Sketch) getP() float64 {
	ones := float64(0)
	for i := uint(0); i < uint(sketch.bitmap.Size()); i++ {
		if sketch.bitmap.Get(i) == true {
			ones++
		}
	}
	return ones / float64(sketch.l)
}

/*
GetEstimate returns the estimated count of a given flow
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
	return uint(m * math.Pow(2, z/m) / rho(p))
}
