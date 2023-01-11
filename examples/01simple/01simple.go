package main

import (
	"fmt"

	"github.com/maseology/goHydro/grid"
	"github.com/maseology/golia"
	"github.com/maseology/mmio"
)

// simple tilted plane starting with a uniform depth
// model run to steady state where water pools to lowest point
func main() {
	tt := mmio.NewTimer()
	defer tt.Print("solution complete")

	nr, nc := 5, 4
	gd, cid := grid.NewDefinition("test", nr, nc, 100.), 0
	mz, mh, mn := make(map[int]float64, nr*nc), make(map[int]float64, nr*nc), make(map[int]float64, nr*nc)
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			mz[cid] = 12. - float64(i*j)
			mh[cid] = 0.5 + mz[cid]
			mn[cid] = .05
			cid++
		}
	}
	printGrid(mz, nr, nc)
	printGrid(mh, nr, nc)
	// printGrid(mn, nr, nc)

	lia := golia.Domain{Alpha: .15, Theta: .7}
	lia.Build(gd, mz, mh, mn)
	out := lia.SolveSteadyState()

	printGrid(out, nr, nc)
}

func printGrid(m map[int]float64, nr, nc int) {
	cid := 0
	println()
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if v, ok := m[cid]; ok {
				fmt.Printf("%10.4f ", v)
			} else {
				panic("printGrid error 1")
			}
			cid++
		}
		println()
	}
	println()
}
