package atomicx

// All tests are taken from gccgo's sync/atomic implementation
// https://github.com/golang/gofrontend/blob/master/libgo/go/sync/atomic/atomic_test.go

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

// Tests of correct behavior, without contention.
// (Does the function work as advertised?)
//
// Test that the Add functions add correctly.
// Test that the CompareAndSwap functions actually
// do the comparison and the swap correctly.
//
// The loop over power-of-two values is meant to
// ensure that the operations apply to the full word size.
// The struct fields x.before and x.after check that the
// operations do not extend past the full word size.

const (
	magic32 = 0xdedbeef
	magic64 = 0xdeddeadbeefbeef
)

var memOrders = [...]MemOrder{OrderRelaxed, OrderConsume, OrderAcquire, OrderRelease, OrderAcqRel, OrderSeqCst}

// Do the 64-bit functions panic? If so, don't bother testing.
var test64err = func() (err interface{}) {
	defer func() {
		err = recover()
	}()
	var x int64
	for _, order := range memOrders {
		AddInt64(&x, 1, order)
	}
	return nil
}()

func testBinaryOpInt32(op func(int32, int32) int32, atomicFun func(*int32, int32, MemOrder) int32, t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		var j int32
		for delta := int32(1); delta+delta > delta; delta += delta {
			k := atomicFun(&x.i, delta, order)
			j = op(j, delta)
			if x.i != j || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func testBinaryOpUint32(op func(uint32, uint32) uint32, atomicFun func(*uint32, uint32, MemOrder) uint32, t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		var j uint32
		for delta := uint32(1); delta+delta > delta; delta += delta {
			k := atomicFun(&x.i, delta, order)
			j = op(j, delta)
			if x.i != j || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func testBinaryOpInt64(op func(int64, int64) int64, atomicFun func(*int64, int64, MemOrder) int64, t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		var j int64
		for delta := int64(1); delta+delta > delta; delta += delta {
			k := atomicFun(&x.i, delta, order)
			j = op(j, delta)
			if x.i != j || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
		}
	}
}

func testBinaryOpUint64(op func(uint64, uint64) uint64, atomicFun func(*uint64, uint64, MemOrder) uint64, t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		var j uint64
		for delta := uint64(1); delta+delta > delta; delta += delta {
			k := atomicFun(&x.i, delta, order)
			j = op(j, delta)
			if x.i != j || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
		}
	}
}

func TestSwapInt32(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		var j int32
		for delta := int32(1); delta+delta > delta; delta += delta {
			k := SwapInt32(&x.i, delta, order)
			if x.i != delta || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
			j = delta
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestSwapUint32(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		var j uint32
		for delta := uint32(1); delta+delta > delta; delta += delta {
			k := SwapUint32(&x.i, delta, order)
			if x.i != delta || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
			j = delta
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestSwapInt64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		var j int64
		for delta := int64(1); delta+delta > delta; delta += delta {
			k := SwapInt64(&x.i, delta, order)
			if x.i != delta || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
			j = delta
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestSwapUint64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		var j uint64
		for delta := uint64(1); delta+delta > delta; delta += delta {
			k := SwapUint64(&x.i, delta, order)
			if x.i != delta || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
			j = delta
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestSwapUintptr(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uintptr
			i      uintptr
			after  uintptr
		}
		var m uint64 = magic64
		magicptr := uintptr(m)
		x.before = magicptr
		x.after = magicptr
		var j uintptr
		for delta := uintptr(1); delta+delta > delta; delta += delta {
			k := SwapUintptr(&x.i, delta, order)
			if x.i != delta || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
			j = delta
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestSwapPointer(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uintptr
			i      unsafe.Pointer
			after  uintptr
		}
		var m uint64 = magic64
		magicptr := uintptr(m)
		x.before = magicptr
		x.after = magicptr
		var j uintptr
		for delta := uintptr(1 << 16); delta+delta > delta; delta += delta {
			k := SwapPointer(&x.i, unsafe.Pointer(delta), order)
			if uintptr(x.i) != delta || uintptr(k) != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
			j = delta
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestAddInt32(t *testing.T) {
	testBinaryOpInt32(func(a int32, b int32) int32 { return a + b }, AddInt32, t)
}

func TestAddUint32(t *testing.T) {
	testBinaryOpUint32(func(a uint32, b uint32) uint32 { return a + b }, AddUint32, t)
}

func TestAddInt64(t *testing.T) {
	testBinaryOpInt64(func(a int64, b int64) int64 { return a + b }, AddInt64, t)
}

func TestAddUint64(t *testing.T) {
	testBinaryOpUint64(func(a uint64, b uint64) uint64 { return a + b }, AddUint64, t)
}

func TestAddUintptr(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uintptr
			i      uintptr
			after  uintptr
		}
		var m uint64 = magic64
		magicptr := uintptr(m)
		x.before = magicptr
		x.after = magicptr
		var j uintptr
		for delta := uintptr(1); delta+delta > delta; delta += delta {
			k := AddUintptr(&x.i, delta, order)
			j += delta
			if x.i != j || k != j {
				t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
			}
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestCompareAndSwapInt32(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		for val := int32(1); val+val > val; val += val {
			for _, isWeak := range []bool{false, true} {
				x.i = val
				if !CompareAndSwapInt32(&x.i, val, val+1, isWeak, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if CompareAndSwapInt32(&x.i, val, val+2, isWeak, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestCompareAndSwap2Int32(t *testing.T) {
	for _, order1 := range memOrders {
		for _, order2 := range memOrders {
			var x struct {
				before int32
				i      int32
				after  int32
			}
			x.before = magic32
			x.after = magic32
			for val := int32(1); val+val > val; val += val {
				for _, isWeak := range []bool{false, true} {
					x.i = val
					if !CompareAndSwap2Int32(&x.i, val, val+1, isWeak, order1, order2) {
						t.Fatalf("should have swapped %#x %#x", val, val+1)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
					x.i = val + 1
					if CompareAndSwap2Int32(&x.i, val, val+2, isWeak, order1, order2) {
						t.Fatalf("should not have swapped %#x %#x", val, val+2)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
				}
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
	}
}

func TestCompareAndSwapUint32(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		for val := uint32(1); val+val > val; val += val {
			for _, isWeak := range []bool{false, true} {
				x.i = val
				if !CompareAndSwapUint32(&x.i, val, val+1, isWeak, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if CompareAndSwapUint32(&x.i, val, val+2, isWeak, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestCompareAndSwap2Uint32(t *testing.T) {
	for _, order1 := range memOrders {
		for _, order2 := range memOrders {
			var x struct {
				before uint32
				i      uint32
				after  uint32
			}
			x.before = magic32
			x.after = magic32
			for val := uint32(1); val+val > val; val += val {
				for _, isWeak := range []bool{false, true} {
					x.i = val
					if !CompareAndSwap2Uint32(&x.i, val, val+1, isWeak, order1, order2) {
						t.Fatalf("should have swapped %#x %#x", val, val+1)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
					x.i = val + 1
					if CompareAndSwap2Uint32(&x.i, val, val+2, isWeak, order1, order2) {
						t.Fatalf("should not have swapped %#x %#x", val, val+2)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
				}
			}
			if x.before != magic32 || x.after != magic32 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
			}
		}
	}
}

func TestCompareAndSwapInt64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		for val := int64(1); val+val > val; val += val {
			for _, isWeak := range []bool{false, true} {
				x.i = val
				if !CompareAndSwapInt64(&x.i, val, val+1, isWeak, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if CompareAndSwapInt64(&x.i, val, val+2, isWeak, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestCompareAndSwap2Int64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order1 := range memOrders {
		for _, order2 := range memOrders {
			var x struct {
				before int64
				i      int64
				after  int64
			}
			x.before = magic64
			x.after = magic64
			for val := int64(1); val+val > val; val += val {
				for _, isWeak := range []bool{false, true} {
					x.i = val
					if !CompareAndSwap2Int64(&x.i, val, val+1, isWeak, order1, order2) {
						t.Fatalf("should have swapped %#x %#x", val, val+1)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
					x.i = val + 1
					if CompareAndSwap2Int64(&x.i, val, val+2, isWeak, order1, order2) {
						t.Fatalf("should not have swapped %#x %#x", val, val+2)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
				}
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
			}
		}
	}
}

func testCompareAndSwapUint64(t *testing.T, cas func(*uint64, uint64, uint64, bool, MemOrder) bool) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		for val := uint64(1); val+val > val; val += val {
			for _, isWeak := range []bool{false, true} {
				x.i = val
				if !cas(&x.i, val, val+1, isWeak, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if cas(&x.i, val, val+2, isWeak, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestCompareAndSwapUint64(t *testing.T) {
	testCompareAndSwapUint64(t, CompareAndSwapUint64)
}

func TestCompareAndSwap2Uint64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order1 := range memOrders {
		for _, order2 := range memOrders {
			var x struct {
				before uint64
				i      uint64
				after  uint64
			}
			x.before = magic64
			x.after = magic64
			for val := uint64(1); val+val > val; val += val {
				for _, isWeak := range []bool{false, true} {
					x.i = val
					if !CompareAndSwap2Uint64(&x.i, val, val+1, isWeak, order1, order2) {
						t.Fatalf("should have swapped %#x %#x", val, val+1)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
					x.i = val + 1
					if CompareAndSwap2Uint64(&x.i, val, val+2, isWeak, order1, order2) {
						t.Fatalf("should not have swapped %#x %#x", val, val+2)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
				}
			}
			if x.before != magic64 || x.after != magic64 {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
			}
		}
	}
}

func TestCompareAndSwapUintptr(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uintptr
			i      uintptr
			after  uintptr
		}
		var m uint64 = magic64
		magicptr := uintptr(m)
		x.before = magicptr
		x.after = magicptr
		for val := uintptr(1); val+val > val; val += val {
			for _, isWeak := range []bool{false, true} {
				x.i = val
				if !CompareAndSwapUintptr(&x.i, val, val+1, isWeak, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = val + 1
				if CompareAndSwapUintptr(&x.i, val, val+2, isWeak, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != val+1 {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestCompareAndSwap2Uintptr(t *testing.T) {
	for _, order1 := range memOrders {
		for _, order2 := range memOrders {
			var x struct {
				before uintptr
				i      uintptr
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			for val := uintptr(1); val+val > val; val += val {
				for _, isWeak := range []bool{false, true} {
					x.i = val
					if !CompareAndSwap2Uintptr(&x.i, val, val+1, isWeak, order1, order2) {
						t.Fatalf("should have swapped %#x %#x", val, val+1)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
					x.i = val + 1
					if CompareAndSwap2Uintptr(&x.i, val, val+2, isWeak, order1, order2) {
						t.Fatalf("should not have swapped %#x %#x", val, val+2)
					}
					if x.i != val+1 {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
				}
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
	}
}

func TestCompareAndSwapPointer(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uintptr
			i      unsafe.Pointer
			after  uintptr
		}
		var m uint64 = magic64
		magicptr := uintptr(m)
		x.before = magicptr
		x.after = magicptr
		for val := uintptr(1 << 16); val+val > val; val += val {
			for _, isWeak := range []bool{false, true} {
				x.i = unsafe.Pointer(val)
				if !CompareAndSwapPointer(&x.i, unsafe.Pointer(val), unsafe.Pointer(val+1), isWeak, order) {
					t.Fatalf("should have swapped %#x %#x", val, val+1)
				}
				if x.i != unsafe.Pointer(val+1) {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
				x.i = unsafe.Pointer(val + 1)
				if CompareAndSwapPointer(&x.i, unsafe.Pointer(val), unsafe.Pointer(val+2), isWeak, order) {
					t.Fatalf("should not have swapped %#x %#x", val, val+2)
				}
				if x.i != unsafe.Pointer(val+1) {
					t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
				}
			}
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestCompareAndSwap2Pointer(t *testing.T) {
	for _, order1 := range memOrders {
		for _, order2 := range memOrders {
			var x struct {
				before uintptr
				i      unsafe.Pointer
				after  uintptr
			}
			var m uint64 = magic64
			magicptr := uintptr(m)
			x.before = magicptr
			x.after = magicptr
			for val := uintptr(1 << 16); val+val > val; val += val {
				for _, isWeak := range []bool{false, true} {
					x.i = unsafe.Pointer(val)
					if !CompareAndSwap2Pointer(&x.i, unsafe.Pointer(val), unsafe.Pointer(val+1), isWeak, order1, order2) {
						t.Fatalf("should have swapped %#x %#x", val, val+1)
					}
					if x.i != unsafe.Pointer(val+1) {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
					x.i = unsafe.Pointer(val + 1)
					if CompareAndSwap2Pointer(&x.i, unsafe.Pointer(val), unsafe.Pointer(val+2), isWeak, order1, order2) {
						t.Fatalf("should not have swapped %#x %#x", val, val+2)
					}
					if x.i != unsafe.Pointer(val+1) {
						t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
					}
				}
			}
			if x.before != magicptr || x.after != magicptr {
				t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
			}
		}
	}
}

func TestLoadInt32(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		for delta := int32(1); delta+delta > delta; delta += delta {
			k := LoadInt32(&x.i, order)
			if k != x.i {
				t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
			}
			x.i += delta
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestLoadUint32(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		for delta := uint32(1); delta+delta > delta; delta += delta {
			k := LoadUint32(&x.i, order)
			if k != x.i {
				t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
			}
			x.i += delta
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestLoadInt64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		for delta := int64(1); delta+delta > delta; delta += delta {
			k := LoadInt64(&x.i, order)
			if k != x.i {
				t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
			}
			x.i += delta
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestLoadUint64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		for delta := uint64(1); delta+delta > delta; delta += delta {
			k := LoadUint64(&x.i, order)
			if k != x.i {
				t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
			}
			x.i += delta
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestLoadUintptr(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uintptr
			i      uintptr
			after  uintptr
		}
		var m uint64 = magic64
		magicptr := uintptr(m)
		x.before = magicptr
		x.after = magicptr
		for delta := uintptr(1); delta+delta > delta; delta += delta {
			k := LoadUintptr(&x.i, order)
			if k != x.i {
				t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
			}
			x.i += delta
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestLoadPointer(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uintptr
			i      unsafe.Pointer
			after  uintptr
		}
		var m uint64 = magic64
		magicptr := uintptr(m)
		x.before = magicptr
		x.after = magicptr
		for delta := uintptr(1 << 16); delta+delta > delta; delta += delta {
			k := LoadPointer(&x.i, order)
			if k != x.i {
				t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
			}
			x.i = unsafe.Pointer(uintptr(x.i) + delta)
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestStoreInt32(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		v := int32(0)
		for delta := int32(1); delta+delta > delta; delta += delta {
			StoreInt32(&x.i, v, order)
			if x.i != v {
				t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
			}
			v += delta
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestStoreUint32(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		v := uint32(0)
		for delta := uint32(1); delta+delta > delta; delta += delta {
			StoreUint32(&x.i, v, order)
			if x.i != v {
				t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
			}
			v += delta
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestStoreInt64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		v := int64(0)
		for delta := int64(1); delta+delta > delta; delta += delta {
			StoreInt64(&x.i, v, order)
			if x.i != v {
				t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
			}
			v += delta
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestStoreUint64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, order := range memOrders {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		v := uint64(0)
		for delta := uint64(1); delta+delta > delta; delta += delta {
			StoreUint64(&x.i, v, order)
			if x.i != v {
				t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
			}
			v += delta
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestStoreUintptr(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uintptr
			i      uintptr
			after  uintptr
		}
		var m uint64 = magic64
		magicptr := uintptr(m)
		x.before = magicptr
		x.after = magicptr
		v := uintptr(0)
		for delta := uintptr(1); delta+delta > delta; delta += delta {
			StoreUintptr(&x.i, v, order)
			if x.i != v {
				t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
			}
			v += delta
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestStorePointer(t *testing.T) {
	for _, order := range memOrders {
		var x struct {
			before uintptr
			i      unsafe.Pointer
			after  uintptr
		}
		var m uint64 = magic64
		magicptr := uintptr(m)
		x.before = magicptr
		x.after = magicptr
		v := unsafe.Pointer(uintptr(0))
		for delta := uintptr(1 << 16); delta+delta > delta; delta += delta {
			StorePointer(&x.i, unsafe.Pointer(v), order)
			if x.i != v {
				t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
			}
			v = unsafe.Pointer(uintptr(v) + delta)
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestTestAndSetClear(t *testing.T) {
	for _, order := range memOrders {
		var x = new(bool)
		if TestAndSet(x, order) {
			t.Fatalf("atomic flag set failed: flag=%v", *x)
		}
		if !TestAndSet(x, order) {
			t.Fatalf("atomic flag second set failed: flag=%v", *x)
		}
		if !TestAndSet(x, order) {
			t.Fatalf("atomic flag third set failed: flag=%v", *x)
		}
		Clear(x, order)
		if TestAndSet(x, order) {
			t.Fatalf("atomic flag set after clear failed: flag=%v", *x)
		}
		if !TestAndSet(x, order) {
			t.Fatalf("atomic flag second set after clear failed: flag=%v", *x)
		}
	}
}

func TestAndInt32(t *testing.T) {
	testBinaryOpInt32(func(a int32, b int32) int32 { return a & b }, AndInt32, t)
}

func TestAndUint32(t *testing.T) {
	testBinaryOpUint32(func(a uint32, b uint32) uint32 { return a & b }, AndUint32, t)
}

func TestAndInt64(t *testing.T) {
	testBinaryOpInt64(func(a int64, b int64) int64 { return a & b }, AndInt64, t)
}

func TestAndUint64(t *testing.T) {
	testBinaryOpUint64(func(a uint64, b uint64) uint64 { return a & b }, AndUint64, t)
}

func TestOrInt32(t *testing.T) {
	testBinaryOpInt32(func(a int32, b int32) int32 { return a | b }, OrInt32, t)
}

func TestOrUint32(t *testing.T) {
	testBinaryOpUint32(func(a uint32, b uint32) uint32 { return a | b }, OrUint32, t)
}

func TestOrInt64(t *testing.T) {
	testBinaryOpInt64(func(a int64, b int64) int64 { return a | b }, OrInt64, t)
}

func TestOrUint64(t *testing.T) {
	testBinaryOpUint64(func(a uint64, b uint64) uint64 { return a | b }, OrUint64, t)
}

func TestXorInt32(t *testing.T) {
	testBinaryOpInt32(func(a int32, b int32) int32 { return a ^ b }, XorInt32, t)
}

func TestXorUint32(t *testing.T) {
	testBinaryOpUint32(func(a uint32, b uint32) uint32 { return a ^ b }, XorUint32, t)
}

func TestXorInt64(t *testing.T) {
	testBinaryOpInt64(func(a int64, b int64) int64 { return a ^ b }, XorInt64, t)
}

func TestXorUint64(t *testing.T) {
	testBinaryOpUint64(func(a uint64, b uint64) uint64 { return a ^ b }, XorUint64, t)
}

func TestNandInt32(t *testing.T) {
	testBinaryOpInt32(func(a int32, b int32) int32 { return ^(a & b) }, NandInt32, t)
}

func TestNandUint32(t *testing.T) {
	testBinaryOpUint32(func(a uint32, b uint32) uint32 { return ^(a & b) }, NandUint32, t)
}

func TestNandInt64(t *testing.T) {
	testBinaryOpInt64(func(a int64, b int64) int64 { return ^(a & b) }, NandInt64, t)
}

func TestNandUint64(t *testing.T) {
	testBinaryOpUint64(func(a uint64, b uint64) uint64 { return ^(a & b) }, NandUint64, t)
}

// Tests of correct behavior, with contention.
// (Is the function atomic?)
//
// For compatibility reasons all operations have SeqCst memory order
// and CASes are strong.
//
// For each function, we write a "hammer" function that repeatedly
// uses the atomic operation to add 1 to a value. After running
// multiple hammers in parallel, check that we end with the correct
// total.
// Swap can't add 1, so it uses a different scheme.
// The functions repeatedly generate a pseudo-random number such that
// low bits are equal to high bits, swap, check that the old value
// has low and high bits equal.

var hammer32 = map[string]func(*uint32, int){
	"SwapInt32":             hammerSwapInt32,
	"SwapUint32":            hammerSwapUint32,
	"SwapUintptr":           hammerSwapUintptr32,
	"AddInt32":              hammerAddInt32,
	"AddUint32":             hammerAddUint32,
	"AddUintptr":            hammerAddUintptr32,
	"CompareAndSwapInt32":   hammerCompareAndSwapInt32,
	"CompareAndSwapUint32":  hammerCompareAndSwapUint32,
	"CompareAndSwapUintptr": hammerCompareAndSwapUintptr32,
}

func init() {
	var v uint64 = 1 << 50
	if uintptr(v) != 0 {
		// 64-bit system; clear uintptr tests
		delete(hammer32, "SwapUintptr")
		delete(hammer32, "AddUintptr")
		delete(hammer32, "CompareAndSwapUintptr")
	}
}

func hammerSwapInt32(uaddr *uint32, count int) {
	addr := (*int32)(unsafe.Pointer(uaddr))
	seed := int(uintptr(unsafe.Pointer(&count)))
	for i := 0; i < count; i++ {
		new := uint32(seed+i)<<16 | uint32(seed+i)<<16>>16
		old := uint32(SwapInt32(addr, int32(new), OrderSeqCst))
		if old>>16 != old<<16>>16 {
			panic(fmt.Sprintf("SwapInt32 is not atomic: %v", old))
		}
	}
}

func hammerSwapUint32(addr *uint32, count int) {
	seed := int(uintptr(unsafe.Pointer(&count)))
	for i := 0; i < count; i++ {
		new := uint32(seed+i)<<16 | uint32(seed+i)<<16>>16
		old := SwapUint32(addr, new, OrderSeqCst)
		if old>>16 != old<<16>>16 {
			panic(fmt.Sprintf("SwapUint32 is not atomic: %v", old))
		}
	}
}

func hammerSwapUintptr32(uaddr *uint32, count int) {
	// only safe when uintptr is 32-bit.
	// not called on 64-bit systems.
	addr := (*uintptr)(unsafe.Pointer(uaddr))
	seed := int(uintptr(unsafe.Pointer(&count)))
	for i := 0; i < count; i++ {
		new := uintptr(seed+i)<<16 | uintptr(seed+i)<<16>>16
		old := SwapUintptr(addr, new, OrderSeqCst)
		if old>>16 != old<<16>>16 {
			panic(fmt.Sprintf("SwapUintptr is not atomic: %#08x", old))
		}
	}
}

func hammerAddInt32(uaddr *uint32, count int) {
	addr := (*int32)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		AddInt32(addr, 1, OrderSeqCst)
	}
}

func hammerAddUint32(addr *uint32, count int) {
	for i := 0; i < count; i++ {
		AddUint32(addr, 1, OrderSeqCst)
	}
}

func hammerAddUintptr32(uaddr *uint32, count int) {
	// only safe when uintptr is 32-bit.
	// not called on 64-bit systems.
	addr := (*uintptr)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		AddUintptr(addr, 1, OrderSeqCst)
	}
}

func hammerCompareAndSwapInt32(uaddr *uint32, count int) {
	addr := (*int32)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		for {
			v := LoadInt32(addr, OrderSeqCst)
			if CompareAndSwapInt32(addr, v, v+1, false, OrderSeqCst) {
				break
			}
		}
	}
}

func hammerCompareAndSwapUint32(addr *uint32, count int) {
	for i := 0; i < count; i++ {
		for {
			v := LoadUint32(addr, OrderSeqCst)
			if CompareAndSwapUint32(addr, v, v+1, false, OrderSeqCst) {
				break
			}
		}
	}
}

func hammerCompareAndSwapUintptr32(uaddr *uint32, count int) {
	// only safe when uintptr is 32-bit.
	// not called on 64-bit systems.
	addr := (*uintptr)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		for {
			v := LoadUintptr(addr, OrderSeqCst)
			if CompareAndSwapUintptr(addr, v, v+1, false, OrderSeqCst) {
				break
			}
		}
	}
}

func TestHammer32(t *testing.T) {
	const p = 4
	n := 100000
	if testing.Short() {
		n = 1000
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(p))

	for name, testf := range hammer32 {
		c := make(chan int)
		var val uint32
		for i := 0; i < p; i++ {
			go func() {
				defer func() {
					if err := recover(); err != nil {
						t.Error(err.(string))
					}
					c <- 1
				}()
				testf(&val, n)
			}()
		}
		for i := 0; i < p; i++ {
			<-c
		}
		if !strings.HasPrefix(name, "Swap") && val != uint32(n)*p {
			t.Fatalf("%s: val=%d want %d", name, val, n*p)
		}
	}
}

var hammer64 = map[string]func(*uint64, int){
	"SwapInt64":             hammerSwapInt64,
	"SwapUint64":            hammerSwapUint64,
	"SwapUintptr":           hammerSwapUintptr64,
	"AddInt64":              hammerAddInt64,
	"AddUint64":             hammerAddUint64,
	"AddUintptr":            hammerAddUintptr64,
	"CompareAndSwapInt64":   hammerCompareAndSwapInt64,
	"CompareAndSwapUint64":  hammerCompareAndSwapUint64,
	"CompareAndSwapUintptr": hammerCompareAndSwapUintptr64,
}

func init() {
	var v uint64 = 1 << 50
	if uintptr(v) == 0 {
		// 32-bit system; clear uintptr tests
		delete(hammer64, "SwapUintptr")
		delete(hammer64, "AddUintptr")
		delete(hammer64, "CompareAndSwapUintptr")
	}
}

func hammerSwapInt64(uaddr *uint64, count int) {
	addr := (*int64)(unsafe.Pointer(uaddr))
	seed := int(uintptr(unsafe.Pointer(&count)))
	for i := 0; i < count; i++ {
		new := uint64(seed+i)<<32 | uint64(seed+i)<<32>>32
		old := uint64(SwapInt64(addr, int64(new), OrderSeqCst))
		if old>>32 != old<<32>>32 {
			panic(fmt.Sprintf("SwapInt64 is not atomic: %v", old))
		}
	}
}

func hammerSwapUint64(addr *uint64, count int) {
	seed := int(uintptr(unsafe.Pointer(&count)))
	for i := 0; i < count; i++ {
		new := uint64(seed+i)<<32 | uint64(seed+i)<<32>>32
		old := SwapUint64(addr, new, OrderSeqCst)
		if old>>32 != old<<32>>32 {
			panic(fmt.Sprintf("SwapUint64 is not atomic: %v", old))
		}
	}
}

const arch32 = unsafe.Sizeof(uintptr(0)) == 4

func hammerSwapUintptr64(uaddr *uint64, count int) {
	// only safe when uintptr is 64-bit.
	// not called on 32-bit systems.
	if !arch32 {
		addr := (*uintptr)(unsafe.Pointer(uaddr))
		seed := int(uintptr(unsafe.Pointer(&count)))
		for i := 0; i < count; i++ {
			new := uintptr(seed+i)<<32 | uintptr(seed+i)<<32>>32
			old := SwapUintptr(addr, new, OrderSeqCst)
			if old>>32 != old<<32>>32 {
				panic(fmt.Sprintf("SwapUintptr is not atomic: %v", old))
			}
		}
	}
}

func hammerAddInt64(uaddr *uint64, count int) {
	addr := (*int64)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		AddInt64(addr, 1, OrderSeqCst)
	}
}

func hammerAddUint64(addr *uint64, count int) {
	for i := 0; i < count; i++ {
		AddUint64(addr, 1, OrderSeqCst)
	}
}

func hammerAddUintptr64(uaddr *uint64, count int) {
	// only safe when uintptr is 64-bit.
	// not called on 32-bit systems.
	addr := (*uintptr)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		AddUintptr(addr, 1, OrderSeqCst)
	}
}

func hammerCompareAndSwapInt64(uaddr *uint64, count int) {
	addr := (*int64)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		for {
			v := LoadInt64(addr, OrderSeqCst)
			if CompareAndSwapInt64(addr, v, v+1, false, OrderSeqCst) {
				break
			}
		}
	}
}

func hammerCompareAndSwapUint64(addr *uint64, count int) {
	for i := 0; i < count; i++ {
		for {
			v := LoadUint64(addr, OrderSeqCst)
			if CompareAndSwapUint64(addr, v, v+1, false, OrderSeqCst) {
				break
			}
		}
	}
}

func hammerCompareAndSwapUintptr64(uaddr *uint64, count int) {
	// only safe when uintptr is 64-bit.
	// not called on 32-bit systems.
	addr := (*uintptr)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		for {
			v := LoadUintptr(addr, OrderSeqCst)
			if CompareAndSwapUintptr(addr, v, v+1, false, OrderSeqCst) {
				break
			}
		}
	}
}

func TestHammer64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	const p = 4
	n := 100000
	if testing.Short() {
		n = 1000
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(p))

	for name, testf := range hammer64 {
		c := make(chan int)
		var val uint64
		for i := 0; i < p; i++ {
			go func() {
				defer func() {
					if err := recover(); err != nil {
						t.Error(err.(string))
					}
					c <- 1
				}()
				testf(&val, n)
			}()
		}
		for i := 0; i < p; i++ {
			<-c
		}
		if !strings.HasPrefix(name, "Swap") && val != uint64(n)*p {
			t.Fatalf("%s: val=%d want %d", name, val, n*p)
		}
	}
}

func hammerStoreLoadInt32(t *testing.T, paddr unsafe.Pointer) {
	addr := (*int32)(paddr)
	v := LoadInt32(addr, OrderSeqCst)
	vlo := v & ((1 << 16) - 1)
	vhi := v >> 16
	if vlo != vhi {
		t.Fatalf("Int32: %#x != %#x", vlo, vhi)
	}
	new := v + 1 + 1<<16
	if vlo == 1e4 {
		new = 0
	}
	StoreInt32(addr, new, OrderSeqCst)
}

func hammerStoreLoadUint32(t *testing.T, paddr unsafe.Pointer) {
	addr := (*uint32)(paddr)
	v := LoadUint32(addr, OrderSeqCst)
	vlo := v & ((1 << 16) - 1)
	vhi := v >> 16
	if vlo != vhi {
		t.Fatalf("Uint32: %#x != %#x", vlo, vhi)
	}
	new := v + 1 + 1<<16
	if vlo == 1e4 {
		new = 0
	}
	StoreUint32(addr, new, OrderSeqCst)
}

func hammerStoreLoadInt64(t *testing.T, paddr unsafe.Pointer) {
	addr := (*int64)(paddr)
	v := LoadInt64(addr, OrderSeqCst)
	vlo := v & ((1 << 32) - 1)
	vhi := v >> 32
	if vlo != vhi {
		t.Fatalf("Int64: %#x != %#x", vlo, vhi)
	}
	new := v + 1 + 1<<32
	StoreInt64(addr, new, OrderSeqCst)
}

func hammerStoreLoadUint64(t *testing.T, paddr unsafe.Pointer) {
	addr := (*uint64)(paddr)
	v := LoadUint64(addr, OrderSeqCst)
	vlo := v & ((1 << 32) - 1)
	vhi := v >> 32
	if vlo != vhi {
		t.Fatalf("Uint64: %#x != %#x", vlo, vhi)
	}
	new := v + 1 + 1<<32
	StoreUint64(addr, new, OrderSeqCst)
}

func hammerStoreLoadUintptr(t *testing.T, paddr unsafe.Pointer) {
	addr := (*uintptr)(paddr)
	v := LoadUintptr(addr, OrderSeqCst)
	new := v
	if arch32 {
		vlo := v & ((1 << 16) - 1)
		vhi := v >> 16
		if vlo != vhi {
			t.Fatalf("Uintptr: %#x != %#x", vlo, vhi)
		}
		new = v + 1 + 1<<16
		if vlo == 1e4 {
			new = 0
		}
	} else {
		vlo := v & ((1 << 32) - 1)
		vhi := v >> 32
		if vlo != vhi {
			t.Fatalf("Uintptr: %#x != %#x", vlo, vhi)
		}
		inc := uint64(1 + 1<<32)
		new = v + uintptr(inc)
	}
	StoreUintptr(addr, new, OrderSeqCst)
}

func hammerStoreLoadPointer(t *testing.T, paddr unsafe.Pointer) {
	addr := (*unsafe.Pointer)(paddr)
	v := uintptr(LoadPointer(addr, OrderSeqCst))
	new := v
	if arch32 {
		vlo := v & ((1 << 16) - 1)
		vhi := v >> 16
		if vlo != vhi {
			t.Fatalf("Pointer: %#x != %#x", vlo, vhi)
		}
		new = v + 1 + 1<<16
		if vlo == 1e4 {
			new = 0
		}
	} else {
		vlo := v & ((1 << 32) - 1)
		vhi := v >> 32
		if vlo != vhi {
			t.Fatalf("Pointer: %#x != %#x", vlo, vhi)
		}
		inc := uint64(1 + 1<<32)
		new = v + uintptr(inc)
	}
	StorePointer(addr, unsafe.Pointer(new), OrderSeqCst)
}

func TestHammerStoreLoad(t *testing.T) {
	var tests []func(*testing.T, unsafe.Pointer)
	tests = append(tests, hammerStoreLoadInt32, hammerStoreLoadUint32,
		hammerStoreLoadUintptr, hammerStoreLoadPointer)
	if test64err == nil {
		tests = append(tests, hammerStoreLoadInt64, hammerStoreLoadUint64)
	}
	n := int(1e6)
	if testing.Short() {
		n = int(1e4)
	}
	const procs = 8
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(procs))
	for _, tt := range tests {
		c := make(chan int)
		var val uint64
		for p := 0; p < procs; p++ {
			go func() {
				for i := 0; i < n; i++ {
					tt(t, unsafe.Pointer(&val))
				}
				c <- 1
			}()
		}
		for p := 0; p < procs; p++ {
			<-c
		}
	}
}

func TestStoreLoadSeqCst32(t *testing.T) {
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	N := int32(1e3)
	if testing.Short() {
		N = int32(1e2)
	}
	c := make(chan bool, 2)
	X := [2]int32{}
	ack := [2][3]int32{{-1, -1, -1}, {-1, -1, -1}}
	for p := 0; p < 2; p++ {
		go func(me int) {
			he := 1 - me
			for i := int32(1); i < N; i++ {
				StoreInt32(&X[me], i, OrderSeqCst)
				my := LoadInt32(&X[he], OrderSeqCst)
				StoreInt32(&ack[me][i%3], my, OrderSeqCst)
				for w := 1; LoadInt32(&ack[he][i%3], OrderSeqCst) == -1; w++ {
					if w%1000 == 0 {
						runtime.Gosched()
					}
				}
				his := LoadInt32(&ack[he][i%3], OrderSeqCst)
				if (my != i && my != i-1) || (his != i && his != i-1) {
					t.Errorf("invalid values: %d/%d (%d)", my, his, i)
					break
				}
				if my != i && his != i {
					t.Errorf("store/load are not sequentially consistent: %d/%d (%d)", my, his, i)
					break
				}
				StoreInt32(&ack[me][(i-1)%3], -1, OrderSeqCst)
			}
			c <- true
		}(p)
	}
	<-c
	<-c
}

func TestStoreLoadSeqCst64(t *testing.T) {
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	N := int64(1e3)
	if testing.Short() {
		N = int64(1e2)
	}
	c := make(chan bool, 2)
	X := [2]int64{}
	ack := [2][3]int64{{-1, -1, -1}, {-1, -1, -1}}
	for p := 0; p < 2; p++ {
		go func(me int) {
			he := 1 - me
			for i := int64(1); i < N; i++ {
				StoreInt64(&X[me], i, OrderSeqCst)
				my := LoadInt64(&X[he], OrderSeqCst)
				StoreInt64(&ack[me][i%3], my, OrderSeqCst)
				for w := 1; LoadInt64(&ack[he][i%3], OrderSeqCst) == -1; w++ {
					if w%1000 == 0 {
						runtime.Gosched()
					}
				}
				his := LoadInt64(&ack[he][i%3], OrderSeqCst)
				if (my != i && my != i-1) || (his != i && his != i-1) {
					t.Errorf("invalid values: %d/%d (%d)", my, his, i)
					break
				}
				if my != i && his != i {
					t.Errorf("store/load are not sequentially consistent: %d/%d (%d)", my, his, i)
					break
				}
				StoreInt64(&ack[me][(i-1)%3], -1, OrderSeqCst)
			}
			c <- true
		}(p)
	}
	<-c
	<-c
}

func TestStoreLoadRelAcq32(t *testing.T) {
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	N := int32(1e3)
	if testing.Short() {
		N = int32(1e2)
	}
	c := make(chan bool, 2)
	type Data struct {
		signal int32
		pad1   [128]int8
		data1  int32
		pad2   [128]int8
		data2  float32
	}
	var X Data
	for p := int32(0); p < 2; p++ {
		go func(p int32) {
			for i := int32(1); i < N; i++ {
				if (i+p)%2 == 0 {
					X.data1 = i
					X.data2 = float32(i)
					StoreInt32(&X.signal, i, OrderRelease)
				} else {
					for w := 1; LoadInt32(&X.signal, OrderAcquire) != i; w++ {
						if w%1000 == 0 {
							runtime.Gosched()
						}
					}
					d1 := X.data1
					d2 := X.data2
					if d1 != i || d2 != float32(i) {
						t.Errorf("incorrect data: %d/%g (%d)", d1, d2, i)
						break
					}
				}
			}
			c <- true
		}(p)
	}
	<-c
	<-c
}

func TestStoreLoadRelAcq64(t *testing.T) {
	if runtime.NumCPU() == 1 {
		t.Skipf("Skipping test on %v processor machine", runtime.NumCPU())
	}
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	N := int64(1e3)
	if testing.Short() {
		N = int64(1e2)
	}
	c := make(chan bool, 2)
	type Data struct {
		signal int64
		pad1   [128]int8
		data1  int64
		pad2   [128]int8
		data2  float64
	}
	var X Data
	for p := int64(0); p < 2; p++ {
		go func(p int64) {
			for i := int64(1); i < N; i++ {
				if (i+p)%2 == 0 {
					X.data1 = i
					X.data2 = float64(i)
					StoreInt64(&X.signal, i, OrderRelease)
				} else {
					for w := 1; LoadInt64(&X.signal, OrderAcquire) != i; w++ {
						if w%1000 == 0 {
							runtime.Gosched()
						}
					}
					d1 := X.data1
					d2 := X.data2
					if d1 != i || d2 != float64(i) {
						t.Errorf("incorrect data: %d/%g (%d)", d1, d2, i)
						break
					}
				}
			}
			c <- true
		}(p)
	}
	<-c
	<-c
}

func shouldPanic(t *testing.T, name string, f func()) {
	defer func() {
		if recover() == nil {
			t.Errorf("%s did not panic", name)
		}
	}()
	f()
}

func TestUnaligned64(t *testing.T) {
	// Unaligned 64-bit atomics on 32-bit systems are
	// a continual source of pain. Test that on 32-bit systems they crash
	// instead of failing silently.

	switch runtime.GOARCH {
	default:
		if !arch32 {
			t.Skip("test only runs on 32-bit systems")
		}
	case "amd64p32":
		// amd64p32 can handle unaligned atomics.
		t.Skipf("test not needed on %v", runtime.GOARCH)
	}

	x := make([]uint32, 4)
	p := (*uint64)(unsafe.Pointer(&x[1])) // misaligned

	shouldPanic(t, "LoadUint64", func() { LoadUint64(p, OrderSeqCst) })
	shouldPanic(t, "StoreUint64", func() { StoreUint64(p, 1, OrderSeqCst) })
	shouldPanic(t, "CompareAndSwapUint64", func() { CompareAndSwapUint64(p, 1, 2, false, OrderSeqCst) })
	shouldPanic(t, "AddUint64", func() { AddUint64(p, 3, OrderSeqCst) })
}

func TestSimpleSpinlock(t *testing.T) {
	const rts = 10
	a := 0
	sl := new(bool)
	var w sync.WaitGroup
	for i := 0; i < rts; i++ {
		w.Add(1)
		go func() {
			for TestAndSet(sl, OrderAcquire) {
				runtime.Gosched()
			}
			defer w.Done()
			defer Clear(sl, OrderRelease)
			a++
		}()
	}
	w.Wait()
	if a != rts {
		t.Errorf("a expected = %d, actual = %d", rts, a)
	}
}

/* Current behavior: Don't check for nil just fail with SIGSEGV
func TestNilDeref(t *testing.T) {
	funcs := [...]func(){
		func() { CompareAndSwapInt32(nil, 0, 0, false, OrderSeqCst) },
		func() { CompareAndSwapInt64(nil, 0, 0, false, OrderSeqCst) },
		func() { CompareAndSwapUint32(nil, 0, 0, false, OrderSeqCst) },
		func() { CompareAndSwapUint64(nil, 0, 0, false, OrderSeqCst) },
		func() { CompareAndSwapUintptr(nil, 0, 0, false, OrderSeqCst) },
		func() { CompareAndSwapPointer(nil, nil, nil, false, OrderSeqCst) },
		func() { CompareAndSwap2Int32(nil, 0, 0, false, OrderSeqCst, OrderSeqCst) },
		func() { CompareAndSwap2Int64(nil, 0, 0, false, OrderSeqCst, OrderSeqCst) },
		func() { CompareAndSwap2Uint32(nil, 0, 0, false, OrderSeqCst, OrderSeqCst) },
		func() { CompareAndSwap2Uint64(nil, 0, 0, false, OrderSeqCst, OrderSeqCst) },
		func() { CompareAndSwap2Uintptr(nil, 0, 0, false, OrderSeqCst, OrderSeqCst) },
		func() { CompareAndSwap2Pointer(nil, nil, nil, false, OrderSeqCst, OrderSeqCst) },
		func() { SwapInt32(nil, 0, OrderSeqCst) },
		func() { SwapUint32(nil, 0, OrderSeqCst) },
		func() { SwapInt64(nil, 0, OrderSeqCst) },
		func() { SwapUint64(nil, 0, OrderSeqCst) },
		func() { SwapUintptr(nil, 0, OrderSeqCst) },
		func() { SwapPointer(nil, nil, OrderSeqCst) },
		func() { AddInt32(nil, 0, OrderSeqCst) },
		func() { AddUint32(nil, 0, OrderSeqCst) },
		func() { AddInt64(nil, 0, OrderSeqCst) },
		func() { AddUint64(nil, 0, OrderSeqCst) },
		func() { AddUintptr(nil, 0, OrderSeqCst) },
		func() { LoadInt32(nil, OrderSeqCst) },
		func() { LoadInt64(nil, OrderSeqCst) },
		func() { LoadUint32(nil, OrderSeqCst) },
		func() { LoadUint64(nil, OrderSeqCst) },
		func() { LoadUintptr(nil, OrderSeqCst) },
		func() { LoadPointer(nil, OrderSeqCst) },
		func() { StoreInt32(nil, 0, OrderSeqCst) },
		func() { StoreInt64(nil, 0, OrderSeqCst) },
		func() { StoreUint32(nil, 0, OrderSeqCst) },
		func() { StoreUint64(nil, 0, OrderSeqCst) },
		func() { StoreUintptr(nil, 0, OrderSeqCst) },
		func() { StorePointer(nil, nil, OrderSeqCst) },
	}
	for _, f := range funcs {
		func() {
			defer func() {
				runtime.GC()
				recover()
			}()
			f()
		}()
	}
}
*/
