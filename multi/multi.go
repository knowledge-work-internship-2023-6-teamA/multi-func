package multi

import (
	"sync"
	"sync/atomic"
)

type Multi struct {
	max_count    uint32
	done_counter uint32
	m            sync.Mutex
}

func (m *Multi) Do(f func()) {
	if atomic.LoadUint32(&m.done_counter) >= m.max_count {
		return
	}
	m.m.Lock()

	if atomic.LoadUint32(&m.done_counter) < m.max_count {
		atomic.AddUint32(&m.done_counter, 1)
		m.m.Unlock()
		f()
	} else {
		m.m.Unlock()
	}
}

func NewMulti(max_count uint32) *Multi {
	return &Multi{max_count: max_count}
}
