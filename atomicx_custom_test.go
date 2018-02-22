package atomicx

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func basicAtomicityTestSuite(order MemoryOrder, t *testing.T) {
	const (
		repeat     = 100
		iterations = 100
		goroutines = 100
		totalExp   = iterations * goroutines
	)

	for k := 0; k < repeat; k++ {
		var wg sync.WaitGroup
		pa := new(int64)
		pb := new(int64)

		for i := 0; i < goroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < iterations; j++ {
					AddInt64(pa, 1, order)
					AddInt64(pb, 1, order)
				}
			}()
		}

		wg.Wait()
		a := (int)(LoadInt64(pa, order))
		b := (int)(LoadInt64(pb, order))
		if totalExp != a || totalExp != b {
			t.Errorf("Atomicity is violated for order %s. Got a = %d and b = %d, expected %d", order, a, b, totalExp)
			return
		}
	}
}

func TestAtomicity(t *testing.T) {
	// some of the orders may not be applicable for atomic operations in atomicityTestSuite, but they still must be atomic
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	orders := [...]MemoryOrder{
		MemoryOrderRelaxed,
		MemoryOrderConsume,
		MemoryOrderAcquire,
		MemoryOrderRelease,
		MemoryOrderAcqRel,
		MemoryOrderSeqCst,
	}
	for _, order := range orders {
		basicAtomicityTestSuite(order, t)
	}
}

type testReader func(*int32, *int32) (int, int)
type testWriter func(*int32, *int32)

func basicMemoryOrderTest(reader testReader, writer testWriter) [4]int {
	pa, pb := new(int32), new(int32)
	beginWr, beginRd, endWr, endRd := make(chan bool), make(chan bool), make(chan bool), make(chan bool)
	results := make(chan [4]int)
	iterations := 1000000

	// writer
	go func() {
		for i := 0; i < iterations; i++ {
			<-beginWr
			writer(pa, pb)
			endWr <- true
		}
	}()

	// reader
	go func() {
		m := [4]int{}
		for i := 0; i < iterations; i++ {
			<-beginRd
			a, b := reader(pa, pb)
			m[2*a+b]++
			endRd <- true
		}
		results <- m
	}()

	for i := 0; i < iterations; i++ {
		StoreInt32(pa, 0, MemoryOrderSeqCst)
		StoreInt32(pb, 0, MemoryOrderSeqCst)
		beginWr <- true
		beginRd <- true
		<-endWr
		<-endRd
	}

	return <-results
}

func TestMemoryOrderRelaxed(t *testing.T) {
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	res := basicMemoryOrderTest(
		testReader(func(pa *int32, pb *int32) (a int, b int) {
			a = (int)(LoadInt32(pa, MemoryOrderRelaxed))
			b = (int)(*pb) // not atomic
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			*pb = 1 // not atomic
			StoreInt32(pa, 1, MemoryOrderRelaxed)
		}),
	)

	var buff bytes.Buffer
	buff.WriteString("Not triggered cases: ")
	for i, times := range res {
		a, b := i>>1, i&1
		str := fmt.Sprintf("(%d,%d) ", a, b)
		if times == 0 {
			buff.WriteString(str)
		}
		t.Log(str + " = " + strconv.Itoa(times))
	}
}

func TestMemoryOrderSeqCst(t *testing.T) {
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	res := basicMemoryOrderTest(
		testReader(func(pa *int32, pb *int32) (a int, b int) {
			a = (int)(LoadInt32(pa, MemoryOrderSeqCst))
			b = (int)(LoadInt32(pb, MemoryOrderSeqCst))
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			StoreInt32(pb, 1, MemoryOrderSeqCst)
			StoreInt32(pa, 1, MemoryOrderSeqCst)
		}),
	)

	for i, times := range res {
		a, b := i>>1, i&1
		t.Log(fmt.Sprintf("(%d,%d) = %d", a, b, times))
	}
	if res[2] != 0 { // (1, 0) = 1 * 2 + 0 = 2 is not allowed
		t.Errorf("State (1, 0) is forbidden, but was seen %d times", res[2])
	}
}

func TestMemoryOrderAcqRel(t *testing.T) {
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	res := basicMemoryOrderTest(
		testReader(func(pa *int32, pb *int32) (a int, b int) {
			a = (int)(LoadInt32(pa, MemoryOrderAcquire))
			b = (int)(*pb) // not atomic
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			*pb = 1 // not atomic
			StoreInt32(pa, 1, MemoryOrderRelease)
		}),
	)

	for i, times := range res {
		a, b := i>>1, i&1
		t.Log(fmt.Sprintf("(%d,%d) = %d", a, b, times))
	}
	if res[2] != 0 { // (1, 0) = 1 * 2 + 0 = 2 is not allowed
		t.Errorf("State (1, 0) is forbidden, but was seen %d times", res[2])
	}
}

func TestMemoryOrderBaseline(t *testing.T) {
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	res := basicMemoryOrderTest(
		testReader(func(pa *int32, pb *int32) (a int, b int) {
			a = (int)(atomic.LoadInt32(pa))
			b = (int)(atomic.LoadInt32(pb))
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			atomic.StoreInt32(pb, 1)
			atomic.StoreInt32(pa, 1)
		}),
	)

	for i, times := range res {
		a, b := i>>1, i&1
		t.Log(fmt.Sprintf("(%d,%d) = %d", a, b, times))
	}
	if res[2] != 0 { // (1, 0) = 1 * 2 + 0 = 2 is not allowed
		t.Errorf("State (1, 0) is forbidden, but was seen %d times", res[2])
	}
}
