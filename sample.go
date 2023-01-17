package stream

import (
	"math/rand"
	"time"
)

// Sample picks n pseudo-randomly chosen input items.  Different executions
// of a Sample filter will chose different items.
func Sample[T streamable](n int) Filter[T] {
	return SampleWithSeed[T](n, time.Now().UnixNano())
}

// SampleWithSeed picks n pseudo-randomly chosen input items. It uses
// seed as the argument for its random number generation and therefore
// different executions of SampleWithSeed with the same arguments will
// chose the same items.
func SampleWithSeed[T streamable](n int, seed int64) Filter[T] {
	return FilterFunc[T](func(arg Arg[T]) error {
		// Could speed this up by using Algorithm Z from Vitter.
		r := rand.New(rand.NewSource(seed))
		reservoir := make([]T, 0, n)
		i := 0
		for s := range arg.In {
			if i < n {
				reservoir = append(reservoir, s)
			} else {
				j := r.Intn(i + 1)
				if j < n {
					reservoir[j] = s
				}
			}
			i++
		}
		for _, s := range reservoir {
			arg.Out <- s
		}
		return nil
	})
}
