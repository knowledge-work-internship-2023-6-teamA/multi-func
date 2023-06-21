package multi

// OnceFunc returns a function that invokes f only once. The returned function
// may be called concurrently.
//
// If f panics, the returned function will panic with the same value on every call.
func MultiFunc(f func(), max_count uint32) func() {
	var (
		valid bool
		p     any
	)
	multi := NewMulti(max_count)
	// Construct the inner closure just once to reduce costs on the fast path.
	g := func() {
		defer func() {
			// panicしなくてもcall ok?
			p = recover()
			if !valid {
				// Re-panic immediately so on the first call the user gets a
				// complete stack trace into f.
				panic(p)
			}
		}()
		f()
		valid = true // Set only if f does not panic
	}
	return func() {
		multi.Do(g)
		if !valid {
			panic(p)
		}
	}
}
