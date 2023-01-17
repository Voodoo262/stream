package stream

// First yields the first n items that it receives.
func First[T streamable](n int) Filter[T] {
	return FilterFunc[T](func(arg Arg[T]) error {
		seen := 0
		for s := range arg.In {
			if seen >= n {
				break
			}
			arg.Out <- s
			seen++
		}
		return nil
	})
}

// DropFirst yields all items except for the first n items that it receives.
func DropFirst[T string](n int) Filter[T] {
	return FilterFunc[T](func(arg Arg[T]) error {
		seen := 0
		for s := range arg.In {
			if seen >= n {
				arg.Out <- s
			}
			seen++
		}
		return nil
	})
}

// ring is a circular buffer.
type ring[T streamable] struct {
	buf     []T
	next, n int
}

func newRing[T streamable](n int) *ring[T] { return &ring[T]{buf: make([]T, n)} }
func (r *ring[T]) empty() bool             { return r.n == 0 }
func (r *ring[T]) full() bool              { return r.n == len(r.buf) }
func (r *ring[T]) pushBack(s T) {
	r.buf[r.next] = s
	r.next = (r.next + 1) % len(r.buf)
	if r.n < len(r.buf) {
		r.n++
	}
}
func (r *ring[T]) popFront() T {
	// The addition of len(r.buf) is so that the dividend is
	// non-negative and therefore remainder is non-negative.
	first := (r.next - r.n + len(r.buf)) % len(r.buf)
	s := r.buf[first]
	r.n--
	return s
}

// Last yields the last n items that it receives.
func Last[T streamable](n int) Filter[T] {
	return FilterFunc[T](func(arg Arg[T]) error {
		r := newRing[T](n)
		for s := range arg.In {
			r.pushBack(s)
		}
		for !r.empty() {
			arg.Out <- r.popFront()
		}
		return nil
	})
}

// DropLast yields all items except for the last n items that it receives.
func DropLast[T streamable](n int) Filter[T] {
	return FilterFunc[T](func(arg Arg[T]) error {
		r := newRing[T](n)
		for s := range arg.In {
			if r.full() {
				arg.Out <- r.popFront()
			}
			r.pushBack(s)
		}
		return nil
	})
}
