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

// Do the 64-bit functions panic? If so, don't bother testing.
var test64err = func() (err interface{}) {
	defer func() {
		err = recover()
	}()
	var x int64
	for _, f := range []func(*int64, int64) int64{
		AddInt64Relaxed, AddInt64Consume, AddInt64Acquire, AddInt64Release, AddInt64AcqRel, AddInt64SeqCst} {
		f(&x, 1)
	}
	return nil
}()

func testBinaryOpInt32(op func(int32, int32) int32, atomicFun func(*int32, int32) int32, t *testing.T) {
	var x struct {
		before int32
		i      int32
		after  int32
	}
	x.before = magic32
	x.after = magic32
	var j int32
	for delta := int32(1); delta+delta > delta; delta += delta {
		k := atomicFun(&x.i, delta)
		j = op(j, delta)
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func testBinaryOpUint32(op func(uint32, uint32) uint32, atomicFun func(*uint32, uint32) uint32, t *testing.T) {
	var x struct {
		before uint32
		i      uint32
		after  uint32
	}
	x.before = magic32
	x.after = magic32
	var j uint32
	for delta := uint32(1); delta+delta > delta; delta += delta {
		k := atomicFun(&x.i, delta)
		j = op(j, delta)
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}

}

func testBinaryOpInt64(op func(int64, int64) int64, atomicFun func(*int64, int64) int64, t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	var x struct {
		before int64
		i      int64
		after  int64
	}
	x.before = magic64
	x.after = magic64
	var j int64
	for delta := int64(1); delta+delta > delta; delta += delta {
		k := atomicFun(&x.i, delta)
		j = op(j, delta)
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func testBinaryOpUint64(op func(uint64, uint64) uint64, atomicFun func(*uint64, uint64) uint64, t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	var x struct {
		before uint64
		i      uint64
		after  uint64
	}
	x.before = magic64
	x.after = magic64
	var j uint64
	for delta := uint64(1); delta+delta > delta; delta += delta {
		k := atomicFun(&x.i, delta)
		j = op(j, delta)
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestSwapInt32(t *testing.T) {
	for _, swapFunc := range []func(*int32, int32) int32{
		SwapInt32Relaxed, SwapInt32Acquire, SwapInt32Release, SwapInt32AcqRel, SwapInt32SeqCst} {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		var j int32
		for delta := int32(1); delta+delta > delta; delta += delta {
			k := swapFunc(&x.i, delta)
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
	for _, swapFunc := range []func(*uint32, uint32) uint32{
		SwapUint32Relaxed, SwapUint32Acquire, SwapUint32Release, SwapUint32AcqRel, SwapUint32SeqCst} {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		var j uint32
		for delta := uint32(1); delta+delta > delta; delta += delta {
			k := swapFunc(&x.i, delta)
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
	for _, swapFunc := range []func(*int64, int64) int64{
		SwapInt64Relaxed, SwapInt64Acquire, SwapInt64Release, SwapInt64AcqRel, SwapInt64SeqCst} {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		var j int64
		for delta := int64(1); delta+delta > delta; delta += delta {
			k := swapFunc(&x.i, delta)
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
	for _, swapFunc := range []func(*uint64, uint64) uint64{
		SwapUint64Relaxed, SwapUint64Acquire, SwapUint64Release, SwapUint64AcqRel, SwapUint64SeqCst} {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		var j uint64
		for delta := uint64(1); delta+delta > delta; delta += delta {
			k := swapFunc(&x.i, delta)
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
	for _, swapFunc := range []func(*uintptr, uintptr) uintptr{
		SwapUintptrRelaxed, SwapUintptrAcquire, SwapUintptrRelease, SwapUintptrAcqRel, SwapUintptrSeqCst} {
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
			k := swapFunc(&x.i, delta)
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
	for _, swapFunc := range []func(*unsafe.Pointer, unsafe.Pointer) unsafe.Pointer{
		SwapPointerRelaxed, SwapPointerAcquire, SwapPointerRelease, SwapPointerAcqRel, SwapPointerSeqCst} {
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
			k := swapFunc(&x.i, unsafe.Pointer(delta))
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
	for _, addFunc := range []func(*int32, int32) int32{
		AddInt32Relaxed, AddInt32Consume, AddInt32Acquire, AddInt32Release, AddInt32AcqRel, AddInt32SeqCst} {
		testBinaryOpInt32(func(a int32, b int32) int32 { return a + b }, addFunc, t)
	}
}

func TestAddUint32(t *testing.T) {
	for _, addFunc := range []func(*uint32, uint32) uint32{
		AddUint32Relaxed, AddUint32Consume, AddUint32Acquire, AddUint32Release, AddUint32AcqRel, AddUint32SeqCst} {
		testBinaryOpUint32(func(a uint32, b uint32) uint32 { return a + b }, addFunc, t)
	}
}

func TestAddInt64(t *testing.T) {
	for _, addFunc := range []func(*int64, int64) int64{
		AddInt64Relaxed, AddInt64Consume, AddInt64Acquire, AddInt64Release, AddInt64AcqRel, AddInt64SeqCst} {
		testBinaryOpInt64(func(a int64, b int64) int64 { return a + b }, addFunc, t)
	}
}

func TestAddUint64(t *testing.T) {
	for _, addFunc := range []func(*uint64, uint64) uint64{
		AddUint64Relaxed, AddUint64Consume, AddUint64Acquire, AddUint64Release, AddUint64AcqRel, AddUint64SeqCst} {
		testBinaryOpUint64(func(a uint64, b uint64) uint64 { return a + b }, addFunc, t)
	}
}

func TestAddUintptr(t *testing.T) {
	for _, addFunc := range []func(*uintptr, uintptr) uintptr{
		AddUintptrRelaxed, AddUintptrConsume, AddUintptrAcquire, AddUintptrRelease, AddUintptrAcqRel, AddUintptrSeqCst} {
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
			k := addFunc(&x.i, delta)
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

func TestCompareAndSwapStrongInt32(t *testing.T) {
	for _, casFunc := range []func(*int32, int32, int32) bool{
		CompareAndSwapStrongInt32RelaxedRelaxed,
		CompareAndSwapStrongInt32ConsumeRelaxed, CompareAndSwapStrongInt32ConsumeConsume,
		CompareAndSwapStrongInt32AcquireRelaxed, CompareAndSwapStrongInt32AcquireConsume, CompareAndSwapStrongInt32AcquireAcquire,
		CompareAndSwapStrongInt32ReleaseRelaxed, CompareAndSwapStrongInt32ReleaseConsume, CompareAndSwapStrongInt32ReleaseAcquire,
		CompareAndSwapStrongInt32AcqRelRelaxed, CompareAndSwapStrongInt32AcqRelConsume, CompareAndSwapStrongInt32AcqRelAcquire,
		CompareAndSwapStrongInt32SeqCstRelaxed, CompareAndSwapStrongInt32SeqCstConsume, CompareAndSwapStrongInt32SeqCstAcquire, CompareAndSwapStrongInt32SeqCstSeqCst} {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		for val := int32(1); val+val > val; val += val {
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestCompareAndSwapWeakInt32(t *testing.T) {
	for _, casFunc := range []func(*int32, int32, int32) bool{
		CompareAndSwapWeakInt32RelaxedRelaxed,
		CompareAndSwapWeakInt32ConsumeRelaxed, CompareAndSwapWeakInt32ConsumeConsume,
		CompareAndSwapWeakInt32AcquireRelaxed, CompareAndSwapWeakInt32AcquireConsume, CompareAndSwapWeakInt32AcquireAcquire,
		CompareAndSwapWeakInt32ReleaseRelaxed, CompareAndSwapWeakInt32ReleaseConsume, CompareAndSwapWeakInt32ReleaseAcquire,
		CompareAndSwapWeakInt32AcqRelRelaxed, CompareAndSwapWeakInt32AcqRelConsume, CompareAndSwapWeakInt32AcqRelAcquire,
		CompareAndSwapWeakInt32SeqCstRelaxed, CompareAndSwapWeakInt32SeqCstConsume, CompareAndSwapWeakInt32SeqCstAcquire, CompareAndSwapWeakInt32SeqCstSeqCst} {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		for val := int32(1); val+val > val; val += val {
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestCompareAndSwapStrongUint32(t *testing.T) {
	for _, casFunc := range []func(*uint32, uint32, uint32) bool{
		CompareAndSwapStrongUint32RelaxedRelaxed,
		CompareAndSwapStrongUint32ConsumeRelaxed, CompareAndSwapStrongUint32ConsumeConsume,
		CompareAndSwapStrongUint32AcquireRelaxed, CompareAndSwapStrongUint32AcquireConsume, CompareAndSwapStrongUint32AcquireAcquire,
		CompareAndSwapStrongUint32ReleaseRelaxed, CompareAndSwapStrongUint32ReleaseConsume, CompareAndSwapStrongUint32ReleaseAcquire,
		CompareAndSwapStrongUint32AcqRelRelaxed, CompareAndSwapStrongUint32AcqRelConsume, CompareAndSwapStrongUint32AcqRelAcquire,
		CompareAndSwapStrongUint32SeqCstRelaxed, CompareAndSwapStrongUint32SeqCstConsume, CompareAndSwapStrongUint32SeqCstAcquire, CompareAndSwapStrongUint32SeqCstSeqCst} {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		for val := uint32(1); val+val > val; val += val {
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestCompareAndSwapWeakUint32(t *testing.T) {
	for _, casFunc := range []func(*uint32, uint32, uint32) bool{
		CompareAndSwapWeakUint32RelaxedRelaxed,
		CompareAndSwapWeakUint32ConsumeRelaxed, CompareAndSwapWeakUint32ConsumeConsume,
		CompareAndSwapWeakUint32AcquireRelaxed, CompareAndSwapWeakUint32AcquireConsume, CompareAndSwapWeakUint32AcquireAcquire,
		CompareAndSwapWeakUint32ReleaseRelaxed, CompareAndSwapWeakUint32ReleaseConsume, CompareAndSwapWeakUint32ReleaseAcquire,
		CompareAndSwapWeakUint32AcqRelRelaxed, CompareAndSwapWeakUint32AcqRelConsume, CompareAndSwapWeakUint32AcqRelAcquire,
		CompareAndSwapWeakUint32SeqCstRelaxed, CompareAndSwapWeakUint32SeqCstConsume, CompareAndSwapWeakUint32SeqCstAcquire, CompareAndSwapWeakUint32SeqCstSeqCst} {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		for val := uint32(1); val+val > val; val += val {
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magic32 || x.after != magic32 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
		}
	}
}

func TestCompareAndSwapStrongInt64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, casFunc := range []func(*int64, int64, int64) bool{
		CompareAndSwapStrongInt64RelaxedRelaxed,
		CompareAndSwapStrongInt64ConsumeRelaxed, CompareAndSwapStrongInt64ConsumeConsume,
		CompareAndSwapStrongInt64AcquireRelaxed, CompareAndSwapStrongInt64AcquireConsume, CompareAndSwapStrongInt64AcquireAcquire,
		CompareAndSwapStrongInt64ReleaseRelaxed, CompareAndSwapStrongInt64ReleaseConsume, CompareAndSwapStrongInt64ReleaseAcquire,
		CompareAndSwapStrongInt64AcqRelRelaxed, CompareAndSwapStrongInt64AcqRelConsume, CompareAndSwapStrongInt64AcqRelAcquire,
		CompareAndSwapStrongInt64SeqCstRelaxed, CompareAndSwapStrongInt64SeqCstConsume, CompareAndSwapStrongInt64SeqCstAcquire, CompareAndSwapStrongInt64SeqCstSeqCst} {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		for val := int64(1); val+val > val; val += val {
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestCompareAndSwapWeakInt64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, casFunc := range []func(*int64, int64, int64) bool{
		CompareAndSwapWeakInt64RelaxedRelaxed,
		CompareAndSwapWeakInt64ConsumeRelaxed, CompareAndSwapWeakInt64ConsumeConsume,
		CompareAndSwapWeakInt64AcquireRelaxed, CompareAndSwapWeakInt64AcquireConsume, CompareAndSwapWeakInt64AcquireAcquire,
		CompareAndSwapWeakInt64ReleaseRelaxed, CompareAndSwapWeakInt64ReleaseConsume, CompareAndSwapWeakInt64ReleaseAcquire,
		CompareAndSwapWeakInt64AcqRelRelaxed, CompareAndSwapWeakInt64AcqRelConsume, CompareAndSwapWeakInt64AcqRelAcquire,
		CompareAndSwapWeakInt64SeqCstRelaxed, CompareAndSwapWeakInt64SeqCstConsume, CompareAndSwapWeakInt64SeqCstAcquire, CompareAndSwapWeakInt64SeqCstSeqCst} {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		for val := int64(1); val+val > val; val += val {
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestCompareAndSwapStrongUint64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, casFunc := range []func(*uint64, uint64, uint64) bool{
		CompareAndSwapStrongUint64RelaxedRelaxed,
		CompareAndSwapStrongUint64ConsumeRelaxed, CompareAndSwapStrongUint64ConsumeConsume,
		CompareAndSwapStrongUint64AcquireRelaxed, CompareAndSwapStrongUint64AcquireConsume, CompareAndSwapStrongUint64AcquireAcquire,
		CompareAndSwapStrongUint64ReleaseRelaxed, CompareAndSwapStrongUint64ReleaseConsume, CompareAndSwapStrongUint64ReleaseAcquire,
		CompareAndSwapStrongUint64AcqRelRelaxed, CompareAndSwapStrongUint64AcqRelConsume, CompareAndSwapStrongUint64AcqRelAcquire,
		CompareAndSwapStrongUint64SeqCstRelaxed, CompareAndSwapStrongUint64SeqCstConsume, CompareAndSwapStrongUint64SeqCstAcquire, CompareAndSwapStrongUint64SeqCstSeqCst} {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		for val := uint64(1); val+val > val; val += val {
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestCompareAndSwapWeakUint64(t *testing.T) {
	if test64err != nil {
		t.Skipf("Skipping 64-bit tests: %v", test64err)
	}
	for _, casFunc := range []func(*uint64, uint64, uint64) bool{
		CompareAndSwapWeakUint64RelaxedRelaxed,
		CompareAndSwapWeakUint64ConsumeRelaxed, CompareAndSwapWeakUint64ConsumeConsume,
		CompareAndSwapWeakUint64AcquireRelaxed, CompareAndSwapWeakUint64AcquireConsume, CompareAndSwapWeakUint64AcquireAcquire,
		CompareAndSwapWeakUint64ReleaseRelaxed, CompareAndSwapWeakUint64ReleaseConsume, CompareAndSwapWeakUint64ReleaseAcquire,
		CompareAndSwapWeakUint64AcqRelRelaxed, CompareAndSwapWeakUint64AcqRelConsume, CompareAndSwapWeakUint64AcqRelAcquire,
		CompareAndSwapWeakUint64SeqCstRelaxed, CompareAndSwapWeakUint64SeqCstConsume, CompareAndSwapWeakUint64SeqCstAcquire, CompareAndSwapWeakUint64SeqCstSeqCst} {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		for val := uint64(1); val+val > val; val += val {
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magic64 || x.after != magic64 {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, uint64(magic64), uint64(magic64))
		}
	}
}

func TestCompareAndSwapStrongUintptr(t *testing.T) {
	for _, casFunc := range []func(*uintptr, uintptr, uintptr) bool{
		CompareAndSwapStrongUintptrRelaxedRelaxed,
		CompareAndSwapStrongUintptrConsumeRelaxed, CompareAndSwapStrongUintptrConsumeConsume,
		CompareAndSwapStrongUintptrAcquireRelaxed, CompareAndSwapStrongUintptrAcquireConsume, CompareAndSwapStrongUintptrAcquireAcquire,
		CompareAndSwapStrongUintptrReleaseRelaxed, CompareAndSwapStrongUintptrReleaseConsume, CompareAndSwapStrongUintptrReleaseAcquire,
		CompareAndSwapStrongUintptrAcqRelRelaxed, CompareAndSwapStrongUintptrAcqRelConsume, CompareAndSwapStrongUintptrAcqRelAcquire,
		CompareAndSwapStrongUintptrSeqCstRelaxed, CompareAndSwapStrongUintptrSeqCstConsume, CompareAndSwapStrongUintptrSeqCstAcquire, CompareAndSwapStrongUintptrSeqCstSeqCst} {
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
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestCompareAndSwapWeakUintptr(t *testing.T) {
	for _, casFunc := range []func(*uintptr, uintptr, uintptr) bool{
		CompareAndSwapWeakUintptrRelaxedRelaxed,
		CompareAndSwapWeakUintptrConsumeRelaxed, CompareAndSwapWeakUintptrConsumeConsume,
		CompareAndSwapWeakUintptrAcquireRelaxed, CompareAndSwapWeakUintptrAcquireConsume, CompareAndSwapWeakUintptrAcquireAcquire,
		CompareAndSwapWeakUintptrReleaseRelaxed, CompareAndSwapWeakUintptrReleaseConsume, CompareAndSwapWeakUintptrReleaseAcquire,
		CompareAndSwapWeakUintptrAcqRelRelaxed, CompareAndSwapWeakUintptrAcqRelConsume, CompareAndSwapWeakUintptrAcqRelAcquire,
		CompareAndSwapWeakUintptrSeqCstRelaxed, CompareAndSwapWeakUintptrSeqCstConsume, CompareAndSwapWeakUintptrSeqCstAcquire, CompareAndSwapWeakUintptrSeqCstSeqCst} {
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
			x.i = val
			if !casFunc(&x.i, val, val+1) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = val + 1
			if casFunc(&x.i, val, val+2) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != val+1 {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestCompareAndSwapStrongPointer(t *testing.T) {
	for _, casFunc := range []func(*unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool{
		CompareAndSwapStrongPointerRelaxedRelaxed,
		CompareAndSwapStrongPointerConsumeRelaxed, CompareAndSwapStrongPointerConsumeConsume,
		CompareAndSwapStrongPointerAcquireRelaxed, CompareAndSwapStrongPointerAcquireConsume, CompareAndSwapStrongPointerAcquireAcquire,
		CompareAndSwapStrongPointerReleaseRelaxed, CompareAndSwapStrongPointerReleaseConsume, CompareAndSwapStrongPointerReleaseAcquire,
		CompareAndSwapStrongPointerAcqRelRelaxed, CompareAndSwapStrongPointerAcqRelConsume, CompareAndSwapStrongPointerAcqRelAcquire,
		CompareAndSwapStrongPointerSeqCstRelaxed, CompareAndSwapStrongPointerSeqCstConsume, CompareAndSwapStrongPointerSeqCstAcquire, CompareAndSwapStrongPointerSeqCstSeqCst} {
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
			x.i = unsafe.Pointer(val)
			if !casFunc(&x.i, unsafe.Pointer(val), unsafe.Pointer(val+1)) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != unsafe.Pointer(val+1) {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = unsafe.Pointer(val + 1)
			if casFunc(&x.i, unsafe.Pointer(val), unsafe.Pointer(val+2)) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != unsafe.Pointer(val+1) {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}

		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestCompareAndSwapWeakPointer(t *testing.T) {
	for _, casFunc := range []func(*unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool{
		CompareAndSwapWeakPointerRelaxedRelaxed,
		CompareAndSwapWeakPointerConsumeRelaxed, CompareAndSwapWeakPointerConsumeConsume,
		CompareAndSwapWeakPointerAcquireRelaxed, CompareAndSwapWeakPointerAcquireConsume, CompareAndSwapWeakPointerAcquireAcquire,
		CompareAndSwapWeakPointerReleaseRelaxed, CompareAndSwapWeakPointerReleaseConsume, CompareAndSwapWeakPointerReleaseAcquire,
		CompareAndSwapWeakPointerAcqRelRelaxed, CompareAndSwapWeakPointerAcqRelConsume, CompareAndSwapWeakPointerAcqRelAcquire,
		CompareAndSwapWeakPointerSeqCstRelaxed, CompareAndSwapWeakPointerSeqCstConsume, CompareAndSwapWeakPointerSeqCstAcquire, CompareAndSwapWeakPointerSeqCstSeqCst} {
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
			x.i = unsafe.Pointer(val)
			if !casFunc(&x.i, unsafe.Pointer(val), unsafe.Pointer(val+1)) {
				t.Fatalf("should have swapped %#x %#x", val, val+1)
			}
			if x.i != unsafe.Pointer(val+1) {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}
			x.i = unsafe.Pointer(val + 1)
			if casFunc(&x.i, unsafe.Pointer(val), unsafe.Pointer(val+2)) {
				t.Fatalf("should not have swapped %#x %#x", val, val+2)
			}
			if x.i != unsafe.Pointer(val+1) {
				t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
			}

		}
		if x.before != magicptr || x.after != magicptr {
			t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magicptr, magicptr)
		}
	}
}

func TestLoadInt32(t *testing.T) {
	for _, loadFunc := range []func(*int32) int32{
		LoadInt32Relaxed, LoadInt32Consume, LoadInt32Acquire, LoadInt32SeqCst} {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		for delta := int32(1); delta+delta > delta; delta += delta {
			k := loadFunc(&x.i)
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
	for _, loadFunc := range []func(*uint32) uint32{
		LoadUint32Relaxed, LoadUint32Consume, LoadUint32Acquire, LoadUint32SeqCst} {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		for delta := uint32(1); delta+delta > delta; delta += delta {
			k := loadFunc(&x.i)
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
	for _, loadFunc := range []func(*int64) int64{
		LoadInt64Relaxed, LoadInt64Consume, LoadInt64Acquire, LoadInt64SeqCst} {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		for delta := int64(1); delta+delta > delta; delta += delta {
			k := loadFunc(&x.i)
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
	for _, loadFunc := range []func(*uint64) uint64{
		LoadUint64Relaxed, LoadUint64Consume, LoadUint64Acquire, LoadUint64SeqCst} {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		for delta := uint64(1); delta+delta > delta; delta += delta {
			k := loadFunc(&x.i)
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
	for _, loadFunc := range []func(*uintptr) uintptr{
		LoadUintptrRelaxed, LoadUintptrConsume, LoadUintptrAcquire, LoadUintptrSeqCst} {
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
			k := loadFunc(&x.i)
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
	for _, loadFunc := range []func(*unsafe.Pointer) unsafe.Pointer{
		LoadPointerRelaxed, LoadPointerConsume, LoadPointerAcquire, LoadPointerSeqCst} {
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
			k := loadFunc(&x.i)
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
	for _, storeFunc := range []func(*int32, int32){StoreInt32Relaxed, StoreInt32Release, StoreInt32SeqCst} {
		var x struct {
			before int32
			i      int32
			after  int32
		}
		x.before = magic32
		x.after = magic32
		v := int32(0)
		for delta := int32(1); delta+delta > delta; delta += delta {
			storeFunc(&x.i, v)
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
	for _, storeFunc := range []func(*uint32, uint32){StoreUint32Relaxed, StoreUint32Release, StoreUint32SeqCst} {
		var x struct {
			before uint32
			i      uint32
			after  uint32
		}
		x.before = magic32
		x.after = magic32
		v := uint32(0)
		for delta := uint32(1); delta+delta > delta; delta += delta {
			storeFunc(&x.i, v)
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
	for _, storeFunc := range []func(*int64, int64){StoreInt64Relaxed, StoreInt64Release, StoreInt64SeqCst} {
		var x struct {
			before int64
			i      int64
			after  int64
		}
		x.before = magic64
		x.after = magic64
		v := int64(0)
		for delta := int64(1); delta+delta > delta; delta += delta {
			storeFunc(&x.i, v)
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
	for _, storeFunc := range []func(*uint64, uint64){StoreUint64Relaxed, StoreUint64Release, StoreUint64SeqCst} {
		var x struct {
			before uint64
			i      uint64
			after  uint64
		}
		x.before = magic64
		x.after = magic64
		v := uint64(0)
		for delta := uint64(1); delta+delta > delta; delta += delta {
			storeFunc(&x.i, v)
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
	for _, storeFunc := range []func(*uintptr, uintptr){StoreUintptrRelaxed, StoreUintptrRelease, StoreUintptrSeqCst} {
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
			storeFunc(&x.i, v)
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
	for _, storeFunc := range []func(*unsafe.Pointer, unsafe.Pointer){StorePointerRelaxed, StorePointerRelease, StorePointerSeqCst} {
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
			storeFunc(&x.i, unsafe.Pointer(v))
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
	for _, testFunc := range []func(*bool) bool{
		TestAndSetRelaxed, TestAndSetConsume, TestAndSetAcquire, TestAndSetRelease, TestAndSetAcqRel, TestAndSetSeqCst} {
		var x = new(bool)
		if testFunc(x) {
			t.Fatalf("atomic flag set failed: flag=%v", *x)
		}
		if !testFunc(x) {
			t.Fatalf("atomic flag second set failed: flag=%v", *x)
		}
		if !testFunc(x) {
			t.Fatalf("atomic flag third set failed: flag=%v", *x)
		}
		ClearSeqCst(x)
		if testFunc(x) {
			t.Fatalf("atomic flag set after clear failed: flag=%v", *x)
		}
		if !testFunc(x) {
			t.Fatalf("atomic flag second set after clear failed: flag=%v", *x)
		}
	}
}

func TestAndInt32(t *testing.T) {
	for _, andFunc := range []func(*int32, int32) int32{
		AndInt32Relaxed, AndInt32Consume, AndInt32Acquire, AndInt32Release, AndInt32AcqRel, AndInt32SeqCst} {
		testBinaryOpInt32(func(a int32, b int32) int32 { return a & b }, andFunc, t)
	}
}

func TestAndUint32(t *testing.T) {
	for _, andFunc := range []func(*uint32, uint32) uint32{
		AndUint32Relaxed, AndUint32Consume, AndUint32Acquire, AndUint32Release, AndUint32AcqRel, AndUint32SeqCst} {
		testBinaryOpUint32(func(a uint32, b uint32) uint32 { return a & b }, andFunc, t)
	}
}

func TestAndInt64(t *testing.T) {
	for _, andFunc := range []func(*int64, int64) int64{
		AndInt64Relaxed, AndInt64Consume, AndInt64Acquire, AndInt64Release, AndInt64AcqRel, AndInt64SeqCst} {
		testBinaryOpInt64(func(a int64, b int64) int64 { return a & b }, andFunc, t)
	}
}

func TestAndUint64(t *testing.T) {
	for _, andFunc := range []func(*uint64, uint64) uint64{
		AndUint64Relaxed, AndUint64Consume, AndUint64Acquire, AndUint64Release, AndUint64AcqRel, AndUint64SeqCst} {
		testBinaryOpUint64(func(a uint64, b uint64) uint64 { return a & b }, andFunc, t)
	}
}

func TestOrInt32(t *testing.T) {
	for _, orFunc := range []func(*int32, int32) int32{
		OrInt32Relaxed, OrInt32Consume, OrInt32Acquire, OrInt32Release, OrInt32AcqRel, OrInt32SeqCst} {
		testBinaryOpInt32(func(a int32, b int32) int32 { return a | b }, orFunc, t)
	}
}

func TestOrUint32(t *testing.T) {
	for _, orFunc := range []func(*uint32, uint32) uint32{
		OrUint32Relaxed, OrUint32Consume, OrUint32Acquire, OrUint32Release, OrUint32AcqRel, OrUint32SeqCst} {
		testBinaryOpUint32(func(a uint32, b uint32) uint32 { return a | b }, orFunc, t)
	}
}

func TestOrInt64(t *testing.T) {
	for _, orFunc := range []func(*int64, int64) int64{
		OrInt64Relaxed, OrInt64Consume, OrInt64Acquire, OrInt64Release, OrInt64AcqRel, OrInt64SeqCst} {
		testBinaryOpInt64(func(a int64, b int64) int64 { return a | b }, orFunc, t)
	}
}

func TestOrUint64(t *testing.T) {
	for _, orFunc := range []func(*uint64, uint64) uint64{
		OrUint64Relaxed, OrUint64Consume, OrUint64Acquire, OrUint64Release, OrUint64AcqRel, OrUint64SeqCst} {
		testBinaryOpUint64(func(a uint64, b uint64) uint64 { return a | b }, orFunc, t)
	}
}

func TestXorInt32(t *testing.T) {
	for _, xorFunc := range []func(*int32, int32) int32{
		XorInt32Relaxed, XorInt32Consume, XorInt32Acquire, XorInt32Release, XorInt32AcqRel, XorInt32SeqCst} {
		testBinaryOpInt32(func(a int32, b int32) int32 { return a ^ b }, xorFunc, t)
	}
}

func TestXorUint32(t *testing.T) {
	for _, xorFunc := range []func(*uint32, uint32) uint32{
		XorUint32Relaxed, XorUint32Consume, XorUint32Acquire, XorUint32Release, XorUint32AcqRel, XorUint32SeqCst} {
		testBinaryOpUint32(func(a uint32, b uint32) uint32 { return a ^ b }, xorFunc, t)
	}
}

func TestXorInt64(t *testing.T) {
	for _, xorFunc := range []func(*int64, int64) int64{
		XorInt64Relaxed, XorInt64Consume, XorInt64Acquire, XorInt64Release, XorInt64AcqRel, XorInt64SeqCst} {
		testBinaryOpInt64(func(a int64, b int64) int64 { return a ^ b }, xorFunc, t)
	}
}

func TestXorUint64(t *testing.T) {
	for _, xorFunc := range []func(*uint64, uint64) uint64{
		XorUint64Relaxed, XorUint64Consume, XorUint64Acquire, XorUint64Release, XorUint64AcqRel, XorUint64SeqCst} {
		testBinaryOpUint64(func(a uint64, b uint64) uint64 { return a ^ b }, xorFunc, t)
	}
}

func TestNandInt32(t *testing.T) {
	for _, nandFunc := range []func(*int32, int32) int32{
		NandInt32Relaxed, NandInt32Consume, NandInt32Acquire, NandInt32Release, NandInt32AcqRel, NandInt32SeqCst} {
		testBinaryOpInt32(func(a int32, b int32) int32 { return ^(a & b) }, nandFunc, t)
	}
}

func TestNandUint32(t *testing.T) {
	for _, nandFunc := range []func(*uint32, uint32) uint32{
		NandUint32Relaxed, NandUint32Consume, NandUint32Acquire, NandUint32Release, NandUint32AcqRel, NandUint32SeqCst} {
		testBinaryOpUint32(func(a uint32, b uint32) uint32 { return ^(a & b) }, nandFunc, t)
	}
}

func TestNandInt64(t *testing.T) {
	for _, nandFunc := range []func(*int64, int64) int64{
		NandInt64Relaxed, NandInt64Consume, NandInt64Acquire, NandInt64Release, NandInt64AcqRel, NandInt64SeqCst} {
		testBinaryOpInt64(func(a int64, b int64) int64 { return ^(a & b) }, nandFunc, t)
	}
}

func TestNandUint64(t *testing.T) {
	for _, nandFunc := range []func(*uint64, uint64) uint64{
		NandUint64Relaxed, NandUint64Consume, NandUint64Acquire, NandUint64Release, NandUint64AcqRel, NandUint64SeqCst} {
		testBinaryOpUint64(func(a uint64, b uint64) uint64 { return ^(a & b) }, nandFunc, t)
	}
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
		old := uint32(SwapInt32SeqCst(addr, int32(new)))
		if old>>16 != old<<16>>16 {
			panic(fmt.Sprintf("SwapInt32 is not atomic: %v", old))
		}
	}
}

func hammerSwapUint32(addr *uint32, count int) {
	seed := int(uintptr(unsafe.Pointer(&count)))
	for i := 0; i < count; i++ {
		new := uint32(seed+i)<<16 | uint32(seed+i)<<16>>16
		old := SwapUint32SeqCst(addr, new)
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
		old := SwapUintptrSeqCst(addr, new)
		if old>>16 != old<<16>>16 {
			panic(fmt.Sprintf("SwapUintptr is not atomic: %#08x", old))
		}
	}
}

func hammerAddInt32(uaddr *uint32, count int) {
	addr := (*int32)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		AddInt32SeqCst(addr, 1)
	}
}

func hammerAddUint32(addr *uint32, count int) {
	for i := 0; i < count; i++ {
		AddUint32SeqCst(addr, 1)
	}
}

func hammerAddUintptr32(uaddr *uint32, count int) {
	// only safe when uintptr is 32-bit.
	// not called on 64-bit systems.
	addr := (*uintptr)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		AddUintptrSeqCst(addr, 1)
	}
}

func hammerCompareAndSwapInt32(uaddr *uint32, count int) {
	addr := (*int32)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		for {
			v := LoadInt32SeqCst(addr)
			if CompareAndSwapStrongInt32SeqCstSeqCst(addr, v, v+1) {
				break
			}
		}
	}
}

func hammerCompareAndSwapUint32(addr *uint32, count int) {
	for i := 0; i < count; i++ {
		for {
			v := LoadUint32SeqCst(addr)
			if CompareAndSwapStrongUint32SeqCstSeqCst(addr, v, v+1) {
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
			v := LoadUintptrSeqCst(addr)
			if CompareAndSwapStrongUintptrSeqCstSeqCst(addr, v, v+1) {
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
		old := uint64(SwapInt64SeqCst(addr, int64(new)))
		if old>>32 != old<<32>>32 {
			panic(fmt.Sprintf("SwapInt64 is not atomic: %v", old))
		}
	}
}

func hammerSwapUint64(addr *uint64, count int) {
	seed := int(uintptr(unsafe.Pointer(&count)))
	for i := 0; i < count; i++ {
		new := uint64(seed+i)<<32 | uint64(seed+i)<<32>>32
		old := SwapUint64SeqCst(addr, new)
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
			old := SwapUintptrSeqCst(addr, new)
			if old>>32 != old<<32>>32 {
				panic(fmt.Sprintf("SwapUintptr is not atomic: %v", old))
			}
		}
	}
}

func hammerAddInt64(uaddr *uint64, count int) {
	addr := (*int64)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		AddInt64SeqCst(addr, 1)
	}
}

func hammerAddUint64(addr *uint64, count int) {
	for i := 0; i < count; i++ {
		AddUint64SeqCst(addr, 1)
	}
}

func hammerAddUintptr64(uaddr *uint64, count int) {
	// only safe when uintptr is 64-bit.
	// not called on 32-bit systems.
	addr := (*uintptr)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		AddUintptrSeqCst(addr, 1)
	}
}

func hammerCompareAndSwapInt64(uaddr *uint64, count int) {
	addr := (*int64)(unsafe.Pointer(uaddr))
	for i := 0; i < count; i++ {
		for {
			v := LoadInt64SeqCst(addr)
			if CompareAndSwapStrongInt64SeqCstSeqCst(addr, v, v+1) {
				break
			}
		}
	}
}

func hammerCompareAndSwapUint64(addr *uint64, count int) {
	for i := 0; i < count; i++ {
		for {
			v := LoadUint64SeqCst(addr)
			if CompareAndSwapStrongUint64SeqCstSeqCst(addr, v, v+1) {
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
			v := LoadUintptrSeqCst(addr)
			if CompareAndSwapStrongUintptrSeqCstSeqCst(addr, v, v+1) {
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
	v := LoadInt32SeqCst(addr)
	vlo := v & ((1 << 16) - 1)
	vhi := v >> 16
	if vlo != vhi {
		t.Fatalf("Int32: %#x != %#x", vlo, vhi)
	}
	new := v + 1 + 1<<16
	if vlo == 1e4 {
		new = 0
	}
	StoreInt32SeqCst(addr, new)
}

func hammerStoreLoadUint32(t *testing.T, paddr unsafe.Pointer) {
	addr := (*uint32)(paddr)
	v := LoadUint32SeqCst(addr)
	vlo := v & ((1 << 16) - 1)
	vhi := v >> 16
	if vlo != vhi {
		t.Fatalf("Uint32: %#x != %#x", vlo, vhi)
	}
	new := v + 1 + 1<<16
	if vlo == 1e4 {
		new = 0
	}
	StoreUint32SeqCst(addr, new)
}

func hammerStoreLoadInt64(t *testing.T, paddr unsafe.Pointer) {
	addr := (*int64)(paddr)
	v := LoadInt64SeqCst(addr)
	vlo := v & ((1 << 32) - 1)
	vhi := v >> 32
	if vlo != vhi {
		t.Fatalf("Int64: %#x != %#x", vlo, vhi)
	}
	new := v + 1 + 1<<32
	StoreInt64SeqCst(addr, new)
}

func hammerStoreLoadUint64(t *testing.T, paddr unsafe.Pointer) {
	addr := (*uint64)(paddr)
	v := LoadUint64SeqCst(addr)
	vlo := v & ((1 << 32) - 1)
	vhi := v >> 32
	if vlo != vhi {
		t.Fatalf("Uint64: %#x != %#x", vlo, vhi)
	}
	new := v + 1 + 1<<32
	StoreUint64SeqCst(addr, new)
}

func hammerStoreLoadUintptr(t *testing.T, paddr unsafe.Pointer) {
	addr := (*uintptr)(paddr)
	v := LoadUintptrSeqCst(addr)
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
	StoreUintptrSeqCst(addr, new)
}

func hammerStoreLoadPointer(t *testing.T, paddr unsafe.Pointer) {
	addr := (*unsafe.Pointer)(paddr)
	v := uintptr(LoadPointerSeqCst(addr))
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
	StorePointerSeqCst(addr, unsafe.Pointer(new))
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
				StoreInt32SeqCst(&X[me], i)
				my := LoadInt32SeqCst(&X[he])
				StoreInt32SeqCst(&ack[me][i%3], my)
				for w := 1; LoadInt32SeqCst(&ack[he][i%3]) == -1; w++ {
					if w%1000 == 0 {
						runtime.Gosched()
					}
				}
				his := LoadInt32SeqCst(&ack[he][i%3])
				if (my != i && my != i-1) || (his != i && his != i-1) {
					t.Errorf("invalid values: %d/%d (%d)", my, his, i)
					break
				}
				if my != i && his != i {
					t.Errorf("store/load are not sequentially consistent: %d/%d (%d)", my, his, i)
					break
				}
				StoreInt32SeqCst(&ack[me][(i-1)%3], -1)
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
				StoreInt64SeqCst(&X[me], i)
				my := LoadInt64SeqCst(&X[he])
				StoreInt64SeqCst(&ack[me][i%3], my)
				for w := 1; LoadInt64SeqCst(&ack[he][i%3]) == -1; w++ {
					if w%1000 == 0 {
						runtime.Gosched()
					}
				}
				his := LoadInt64SeqCst(&ack[he][i%3])
				if (my != i && my != i-1) || (his != i && his != i-1) {
					t.Errorf("invalid values: %d/%d (%d)", my, his, i)
					break
				}
				if my != i && his != i {
					t.Errorf("store/load are not sequentially consistent: %d/%d (%d)", my, his, i)
					break
				}
				StoreInt64SeqCst(&ack[me][(i-1)%3], -1)
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
					StoreInt32Release(&X.signal, i)
				} else {
					for w := 1; LoadInt32Acquire(&X.signal) != i; w++ {
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
					StoreInt64Release(&X.signal, i)
				} else {
					for w := 1; LoadInt64Acquire(&X.signal) != i; w++ {
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

	shouldPanic(t, "LoadUint64", func() { LoadUint64SeqCst(p) })
	shouldPanic(t, "StoreUint64", func() { StoreUint64SeqCst(p, 1) })
	shouldPanic(t, "CompareAndSwapUint64", func() { CompareAndSwapStrongUint64SeqCstSeqCst(p, 1, 2) })
	shouldPanic(t, "AddUint64", func() { AddUint64SeqCst(p, 3) })
}

func TestSimpleSpinlock(t *testing.T) {
	const rts = 10
	a := 0
	sl := new(bool)
	var w sync.WaitGroup
	for i := 0; i < rts; i++ {
		w.Add(1)
		go func() {
			for TestAndSetAcquire(sl) {
				runtime.Gosched()
			}
			defer w.Done()
			defer ClearRelease(sl)
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
		func() { CompareAndSwapInt32SeqCstSeqCst(nil, 0, 0, false) },
		func() { CompareAndSwapInt64SeqCstSeqCst(nil, 0, 0, false) },
		func() { CompareAndSwapUint32SeqCstSeqCst(nil, 0, 0, false) },
		func() { CompareAndSwapUint64SeqCstSeqCst(nil, 0, 0, false) },
		func() { CompareAndSwapUintptrSeqCstSeqCst(nil, 0, 0, false) },
		func() { CompareAndSwapPointerSeqCstSeqCst(nil, nil, nil, false) },
		func() { SwapInt32SeqCst(nil, 0) },
		func() { SwapUint32SeqCst(nil, 0) },
		func() { SwapInt64SeqCst(nil, 0) },
		func() { SwapUint64SeqCst(nil, 0) },
		func() { SwapUintptrSeqCst(nil, 0) },
		func() { SwapPointerSeqCst(nil, nil) },
		func() { AddInt32SeqCst(nil, 0) },
		func() { AddUint32SeqCst(nil, 0) },
		func() { AddInt64SeqCst(nil, 0) },
		func() { AddUint64SeqCst(nil, 0) },
		func() { AddUintptrSeqCst(nil, 0) },
		func() { LoadInt32SeqCst(nil) },
		func() { LoadInt64SeqCst(nil) },
		func() { LoadUint32SeqCst(nil) },
		func() { LoadUint64SeqCst(nil) },
		func() { LoadUintptrSeqCst(nil) },
		func() { LoadPointerSeqCst(nil) },
		func() { StoreInt32SeqCst(nil, 0) },
		func() { StoreInt64SeqCst(nil, 0) },
		func() { StoreUint32SeqCst(nil, 0) },
		func() { StoreUint64SeqCst(nil, 0) },
		func() { StoreUintptrSeqCst(nil, 0) },
		func() { StorePointerSeqCst(nil, nil) },
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
