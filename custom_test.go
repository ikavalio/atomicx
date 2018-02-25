package atomicx

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync/atomic"
	"testing"
)

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
		StoreInt32SeqCst(pa, 0)
		StoreInt32SeqCst(pb, 0)
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
			a = (int)(LoadInt32Relaxed(pa))
			b = (int)(*pb) // not atomic
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			*pb = 1 // not atomic
			StoreInt32Relaxed(pa, 1)
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
			a = (int)(LoadInt32SeqCst(pa))
			b = (int)(LoadInt32SeqCst(pb))
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			StoreInt32SeqCst(pb, 1)
			StoreInt32SeqCst(pa, 1)
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
			a = (int)(LoadInt32Acquire(pa))
			b = (int)(*pb) // not atomic
			return
		}),
		testWriter(func(pa *int32, pb *int32) {
			*pb = 1 // not atomic
			StoreInt32Release(pa, 1)
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
