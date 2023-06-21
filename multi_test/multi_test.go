package multi_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"sync"

	"github.com/knowledge-work-internship-2023-6-teamA/multi-func/multi"
)

func TestMulti(t *testing.T) {
	m := multi.NewMulti(10)
	var count int
	for i := 0; i < 100; i++ {
		m.Do(func() {
			count++
		})
	}
	if count != 10 {
		t.Errorf("count is %d, want 10", count)
	} else {
		t.Logf("count is %d\n", count)
	}
}

func TestMultiConcurrent(t *testing.T) {
	m := multi.NewMulti(10)
	var wg sync.WaitGroup
	var count uint32
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			m.Do(func() {
				atomic.AddUint32(&count, 1)
			})
			wg.Done()
		}()
	}
	wg.Wait()
	if count != 10 {
		t.Errorf("count is %d, want 10", count)
	} else {
		t.Logf("count is %d\n", count)
	}
}

// 逐次処理になっていないかテスト
func TestMultiConcurrent2(t *testing.T) {
	m := multi.NewMulti(10)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		i := i
		wg.Add(1)
		go func() {
			m.Do(func() {
				fmt.Println("start")
				time.Sleep(1 * time.Second)
				fmt.Println("end")
				fmt.Printf("%d\n", i)
			})
			wg.Done()
		}()
	}
	wg.Wait()
}
