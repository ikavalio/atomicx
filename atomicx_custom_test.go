package atomicx

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

func basicAtomicityTestSuite(order MemOrder, t *testing.T) {
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
	orders := [...]MemOrder{OrderRelaxed, OrderConsume, OrderAcquire, OrderRelease, OrderAcqRel, OrderSeqCst}
	for _, order := range orders {
		basicAtomicityTestSuite(order, t)
	}
}

type testReader func(*int32, *int32) (int, int)
type testWriter func(*int32, *int32)

func basicMemoryOrderTest(order MemOrder, reader testReader, writer testWriter) [4]int {
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
		StoreInt32(pa, 0, OrderSeqCst)
		StoreInt32(pb, 0, OrderSeqCst)
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
		OrderRelaxed,
		testReader(func(pa *int32, pb *int32) (a int, b int) {
			a = (int)(LoadInt32(pa, OrderRelaxed))
			b = (int)(LoadInt32(pb, OrderRelaxed))
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			StoreInt32(pa, 1, OrderRelaxed)
			StoreInt32(pb, 1, OrderRelaxed)
		}),
	)

	var buff bytes.Buffer
	buff.WriteString("Not triggered cases: ")
	failedCases := 0
	for i, times := range res {
		a, b := i>>1, i&1
		str := fmt.Sprintf("(%d,%d) ", a, b)
		if times == 0 {
			failedCases++
			buff.WriteString(str)
		}
		t.Log(str + " = " + strconv.Itoa(times))
	}
	if failedCases > 0 {
		t.Error(buff.String())
	}
}

func TestMemoryOrderSeqCst(t *testing.T) {
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	res := basicMemoryOrderTest(
		OrderRelaxed,
		testReader(func(pa *int32, pb *int32) (a int, b int) {
			a = (int)(LoadInt32(pa, OrderSeqCst))
			b = (int)(LoadInt32(pb, OrderSeqCst))
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			StoreInt32(pb, 1, OrderSeqCst)
			StoreInt32(pa, 1, OrderSeqCst)
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
		OrderRelaxed,
		testReader(func(pa *int32, pb *int32) (a int, b int) {
			a = (int)(LoadInt32(pa, OrderAcquire))
			b = (int)(*pb) // not atomic
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			*pb = 1 // not atomic
			StoreInt32(pa, 1, OrderRelease)
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
