package dist

import (
	"math/rand"
	"sort"

	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/mkideal/pkg/debug"
)

type Distribution struct {
	Threshes []float64
}

func NewDistribution(weights []float64) *Distribution {
	total := 0.0
	for _, r := range weights {
		total += r
	}

	d := &Distribution{
		Threshes: make([]float64, len(weights)),
	}

	cur := 0.0
	for i, r := range weights {
		cur += r / total

		// Sometimes cur exceeds 1.0 due to floating point weirdness.
		if cur > 1.0 {
			cur = 1.0
		}

		d.Threshes[i] = cur
	}

	d.Threshes[len(weights)-1] = 1.0

	err := d.Validate()
	if err != nil {
		panic(err.Error())
	}

	return d
}

func (d *Distribution) Validate() error {
	onErr := wlerr.MakeWrapper("invalid npc distribution")

	if len(d.Threshes) == 0 {
		return onErr(nil, "invalid threshold count: have=0 want>0")
	}

	prev := 0.0
	for i, t := range d.Threshes {
		if t < prev {
			return onErr(nil,
				"entry %d (%f) is less than prev (%f)", i, t, prev)
		}

		prev = t
	}

	final := d.Threshes[len(d.Threshes)-1]
	if final != 1.0 {
		return onErr(nil, "invalid final threshold: have=%f want=1.0", final)
	}

	return nil
}

func (d *Distribution) findIdx(r float64) int {
	for i, t := range d.Threshes {
		if r < t {
			return i
		}
	}

	debug.Panicf("%f doesn't match a distribution threshold", r)
	return 0
}

func (d *Distribution) Next() int {
	r := rand.Float64()
	return d.findIdx(r)
}

func (d *Distribution) Generate(points int) []int {
	vals := make([]int, len(d.Threshes))

	for i := 0; i < points; i++ {
		vals[d.Next()]++
	}

	return vals
}

func (d *Distribution) Weights() []float64 {
	weights := make([]float64, len(d.Threshes))

	prev := 0.0
	for i, t := range d.Threshes {
		weights[i] = t - prev
		prev = t
	}

	return weights
}

// Lowest weights come first.
func (d *Distribution) SortedIndices() []int {
	type Elem struct {
		Idx    int
		Weight float64
	}

	elems := make([]Elem, len(d.Threshes))

	weights := d.Weights()
	for i := 0; i < len(elems); i++ {
		elems[i] = Elem{
			Idx:    i,
			Weight: weights[i],
		}
	}

	sort.Slice(elems, func(i int, j int) bool {
		return elems[i].Weight < elems[j].Weight
	})

	indices := make([]int, len(elems))
	for i, e := range elems {
		indices[i] = e.Idx
	}

	return indices
}
