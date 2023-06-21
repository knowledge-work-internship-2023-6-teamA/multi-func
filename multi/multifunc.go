package multi

import (
	"sync"
	"sync/atomic"
)

func MultiFunc(f func(), max_count uint32) func() {
	var done_counter uint32
	var m sync.Mutex

	return func() {
		if atomic.LoadUint32(&done_counter) >= max_count {
			return
		}
		m.Lock()

		if atomic.LoadUint32(&done_counter) < max_count {
			atomic.AddUint32(&done_counter, 1)
			m.Unlock()
			f()
		} else {
			m.Unlock()
		}
	}
}
