package multi_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/knowledge-work-internship-2023-6-teamA/multi-func/multi"
)

func TestMultiFunc(t *testing.T) {
	count := 0
	f := multi.MultiFunc(func() {
		fmt.Println("hello world")
		count++
	}, 10)

	for i := 0; i < 100; i++ {
		f()
	}
	if count != 10 {
		t.Errorf("count is %d, want 10", count)
	} else {
		t.Logf("count is %d\n", count)
	}
}

func TestMultiFuncConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	var count uint32
	f := multi.MultiFunc(func() {
		fmt.Println("hello world")
		atomic.AddUint32(&count, 1)
	}, 10)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			f()
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
func TestMultiFuncConcurrent2(t *testing.T) {
	var wg sync.WaitGroup
	var count uint32
	f := multi.MultiFunc(func() {
		fmt.Println("start")
		time.Sleep(1 * time.Second)
		panic("ぱにっく")
		// runtime.Goexit()
		atomic.AddUint32(&count, 1)
		fmt.Println("end")
	}, 10)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			f()
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
