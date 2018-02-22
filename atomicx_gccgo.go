package atomicx

import "unsafe"

// https://gcc.gnu.org/onlinedocs/gcc/_005f_005fatomic-Builtins.html

// MemoryOrder represents C++11 memory order
type MemoryOrder int

const (
	// MemoryOrderRelaxed implies no inter-thread ordering constraints.
	MemoryOrderRelaxed MemoryOrder = iota
	// MemoryOrderConsume is deprecated and equals to OrderAcquire
	MemoryOrderConsume
	// MemoryOrderAcquire creates an inter-thread happens-before constraint from the release (or stronger) semantic store to this acquire load.
	MemoryOrderAcquire
	// MemoryOrderRelease creates an inter-thread happens-before constraint to acquire (or stronger) semantic loads that read from this release store.
	MemoryOrderRelease
	// MemoryOrderAcqRel combines the effects of both __ATOMIC_ACQUIRE and __ATOMIC_RELEASE.
	MemoryOrderAcqRel
	// MemoryOrderSeqCst enforces total ordering with all other sequentially consistent operations.
	MemoryOrderSeqCst
)

// LoadInt32 atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
//extern github_com_ikavalio_atomicx.LoadInt32
func LoadInt32(addr *int32, order MemoryOrder) int32

// LoadInt64 atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
//extern github_com_ikavalio_atomicx.LoadInt64
func LoadInt64(addr *int64, order MemoryOrder) int64

// LoadPointer atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
//extern github_com_ikavalio_atomicx.LoadPointer
func LoadPointer(addr *unsafe.Pointer, order MemoryOrder) unsafe.Pointer

// LoadUint32 atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
//extern github_com_ikavalio_atomicx.LoadUint32
func LoadUint32(addr *uint32, order MemoryOrder) uint32

// LoadUint64 atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
//extern github_com_ikavalio_atomicx.LoadUint64
func LoadUint64(addr *uint64, order MemoryOrder) uint64

// LoadUintptr atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
//extern github_com_ikavalio_atomicx.LoadUintptr
func LoadUintptr(addr *uintptr, order MemoryOrder) uintptr

// StoreInt32 atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
//extern github_com_ikavalio_atomicx.StoreInt32
func StoreInt32(addr *int32, value int32, order MemoryOrder)

// StoreInt64 atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
//extern github_com_ikavalio_atomicx.StoreInt64
func StoreInt64(addr *int64, value int64, order MemoryOrder)

// StorePointer atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
//extern github_com_ikavalio_atomicx.StorePointer
func StorePointer(addr *unsafe.Pointer, value unsafe.Pointer, order MemoryOrder)

// StoreUint32 atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
//extern github_com_ikavalio_atomicx.StoreUint32
func StoreUint32(addr *uint32, value uint32, order MemoryOrder)

// StoreUint64 atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
//extern github_com_ikavalio_atomicx.StoreUint64
func StoreUint64(addr *uint64, value uint64, order MemoryOrder)

// StoreUintptr atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
//extern github_com_ikavalio_atomicx.StoreUintptr
func StoreUintptr(addr *uintptr, value uintptr, order MemoryOrder)

// SwapInt32 atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
//extern github_com_ikavalio_atomicx.SwapInt32
func SwapInt32(addr *int32, value int32, order MemoryOrder) int32

// SwapInt64 atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
//extern github_com_ikavalio_atomicx.SwapInt64
func SwapInt64(addr *int64, value int64, order MemoryOrder) int64

// SwapPointer atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
//extern github_com_ikavalio_atomicx.SwapPointer
func SwapPointer(addr *unsafe.Pointer, value unsafe.Pointer, order MemoryOrder) unsafe.Pointer

// SwapUint32 atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
//extern github_com_ikavalio_atomicx.SwapUint32
func SwapUint32(addr *uint32, value uint32, order MemoryOrder) uint32

// SwapUint64 atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
//extern github_com_ikavalio_atomicx.SwapUint64
func SwapUint64(addr *uint64, value uint64, order MemoryOrder) uint64

// SwapUintptr atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
//extern github_com_ikavalio_atomicx.SwapUintptr
func SwapUintptr(addr *uintptr, value uintptr, order MemoryOrder) uintptr

// CompareAndSwapStrongInt32 executes the compare-and-swap operation for an int32 value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32
func CompareAndSwapStrongInt32(addr *int32, old, new int32, order MemoryOrder) bool

// CompareAndSwapStrongInt64 executes the compare-and-swap operation for an int64 value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64
func CompareAndSwapStrongInt64(addr *int64, old, new int64, order MemoryOrder) bool

// CompareAndSwapStrongPointer executes the compare-and-swap operation for an unsafe.Pointer value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointer
func CompareAndSwapStrongPointer(addr *unsafe.Pointer, old, new unsafe.Pointer, order MemoryOrder) bool

// CompareAndSwapStrongUint32 executes the compare-and-swap operation for an uint32 value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32
func CompareAndSwapStrongUint32(addr *uint32, old, new uint32, order MemoryOrder) bool

// CompareAndSwapStrongUint64 executes the compare-and-swap operation for an uint64 value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64
func CompareAndSwapStrongUint64(addr *uint64, old, new uint64, order MemoryOrder) bool

// CompareAndSwapStrongUintptr executes the compare-and-swap operation for an uintptr value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptr
func CompareAndSwapStrongUintptr(addr *uintptr, old, new uintptr, order MemoryOrder) bool

// CompareAndSwapStrong2Int32 executes the compare-and-swap operation for an int32 value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrong2Int32
func CompareAndSwapStrong2Int32(addr *int32, old, new int32, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapStrong2Int64 executes the compare-and-swap operation for an int64 value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrong2Int64
func CompareAndSwapStrong2Int64(addr *int64, old, new int64, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapStrong2Pointer executes the compare-and-swap operation for an unsafe.Pointer value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrong2Pointer
func CompareAndSwapStrong2Pointer(addr *unsafe.Pointer, old, new unsafe.Pointer, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapStrong2Uint32 executes the compare-and-swap operation for an uint32 value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrong2Uint32
func CompareAndSwapStrong2Uint32(addr *uint32, old, new uint32, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapStrong2Uint64 executes the compare-and-swap operation for an uint64 value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrong2Uint64
func CompareAndSwapStrong2Uint64(addr *uint64, old, new uint64, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapStrong2Uintptr executes the compare-and-swap operation for an uintptr value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrong2Uintptr
func CompareAndSwapStrong2Uintptr(addr *uintptr, old, new uintptr, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapWeakInt32 executes the compare-and-swap operation for an int32 value. CAS may fail spuriously.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32
func CompareAndSwapWeakInt32(addr *int32, old, new int32, order MemoryOrder) bool

// CompareAndSwapWeakInt64 executes the compare-and-swap operation for an int64 value. CAS may fail spuriously.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64
func CompareAndSwapWeakInt64(addr *int64, old, new int64, order MemoryOrder) bool

// CompareAndSwapWeakPointer executes the compare-and-swap operation for an unsafe.Pointer value. CAS may fail spuriously.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointer
func CompareAndSwapWeakPointer(addr *unsafe.Pointer, old, new unsafe.Pointer, order MemoryOrder) bool

// CompareAndSwapWeakUint32 executes the compare-and-swap operation for an uint32 value. CAS may fail spuriously.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32
func CompareAndSwapWeakUint32(addr *uint32, old, new uint32, order MemoryOrder) bool

// CompareAndSwapWeakUint64 executes the compare-and-swap operation for an uint64 value. CAS may fail spuriously.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64
func CompareAndSwapWeakUint64(addr *uint64, old, new uint64, order MemoryOrder) bool

// CompareAndSwapWeakUintptr executes the compare-and-swap operation for an uintptr value. CAS may fail spuriously.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptr
func CompareAndSwapWeakUintptr(addr *uintptr, old, new uintptr, order MemoryOrder) bool

// CompareAndSwapWeak2Int32 executes the compare-and-swap operation for an int32 value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeak2Int32
func CompareAndSwapWeak2Int32(addr *int32, old, new int32, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapWeak2Int64 executes the compare-and-swap operation for an int64 value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeak2Int64
func CompareAndSwapWeak2Int64(addr *int64, old, new int64, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapWeak2Pointer executes the compare-and-swap operation for an unsafe.Pointer value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeak2Pointer
func CompareAndSwapWeak2Pointer(addr *unsafe.Pointer, old, new unsafe.Pointer, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapWeak2Uint32 executes the compare-and-swap operation for an uint32 value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeak2Uint32
func CompareAndSwapWeak2Uint32(addr *uint32, old, new uint32, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapWeak2Uint64 executes the compare-and-swap operation for an uint64 value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeak2Uint64
func CompareAndSwapWeak2Uint64(addr *uint64, old, new uint64, orderSuccess, orderFailure MemoryOrder) bool

// CompareAndSwapWeak2Uintptr executes the compare-and-swap operation for an uintptr value.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeak2Uintptr
func CompareAndSwapWeak2Uintptr(addr *uintptr, old, new uintptr, orderSuccess, orderFailure MemoryOrder) bool

// AddInt32 atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.AddInt32
func AddInt32(addr *int32, value int32, order MemoryOrder) int32

// AddInt64 atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.AddInt64
func AddInt64(addr *int64, value int64, order MemoryOrder) int64

// AddUint32 atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.AddUint32
func AddUint32(addr *uint32, value uint32, order MemoryOrder) uint32

// AddUint64 atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.AddUint64
func AddUint64(addr *uint64, value uint64, order MemoryOrder) uint64

// AddUintptr atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.AddUintptr
func AddUintptr(addr *uintptr, value uintptr, order MemoryOrder) uintptr

// AndInt32 does atomic bitwise-and between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.AndInt32
func AndInt32(addr *int32, value int32, order MemoryOrder) int32

// AndInt64 does atomic bitwise-and between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.AndInt64
func AndInt64(addr *int64, value int64, order MemoryOrder) int64

// AndUint32 does atomic bitwise-and between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.AndUint32
func AndUint32(addr *uint32, value uint32, order MemoryOrder) uint32

// AndUint64 does atomic bitwise-and between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.AndUint64
func AndUint64(addr *uint64, value uint64, order MemoryOrder) uint64

// OrInt32 does atomic bitwise-or between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.OrInt32
func OrInt32(addr *int32, value int32, order MemoryOrder) int32

// OrInt64 does atomic bitwise-or between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.OrInt64
func OrInt64(addr *int64, value int64, order MemoryOrder) int64

// OrUint32 does atomic bitwise-or between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.OrUint32
func OrUint32(addr *uint32, value uint32, order MemoryOrder) uint32

// OrUint64 does atomic bitwise-or between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.OrUint64
func OrUint64(addr *uint64, value uint64, order MemoryOrder) uint64

// XorInt32 does atomic bitwise-xor between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.XorInt32
func XorInt32(addr *int32, value int32, order MemoryOrder) int32

// XorInt64 does atomic bitwise-xor between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.XorInt64
func XorInt64(addr *int64, value int64, order MemoryOrder) int64

// XorUint32 does atomic bitwise-xor between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.XorUint32
func XorUint32(addr *uint32, value uint32, order MemoryOrder) uint32

// XorUint64 does atomic bitwise-xor between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.XorUint64
func XorUint64(addr *uint64, value uint64, order MemoryOrder) uint64

// NandInt32 does atomic bitwise-nand between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.NandInt32
func NandInt32(addr *int32, value int32, order MemoryOrder) int32

// NandInt64 does atomic bitwise-nand between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.NandInt64
func NandInt64(addr *int64, value int64, order MemoryOrder) int64

// NandUint32 does atomic bitwise-nand between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.NandUint32
func NandUint32(addr *uint32, value uint32, order MemoryOrder) uint32

// NandUint64 does atomic bitwise-nand between delta and *addr and returns the new value.
// Valid memory orders: all.
//extern github_com_ikavalio_atomicx.NandUint64
func NandUint64(addr *uint64, value uint64, order MemoryOrder) uint64

// TestAndSet does atomic test-and-set operation on the bool *addr (atomically *addr = *addr ? *addr : true).
// Returns the old value of *addr.
// Valid memory orders:  all.
//extern github_com_ikavalio_atomicx.TestAndSet
func TestAndSet(addr *bool, order MemoryOrder) bool

// Clear does atomic clear operation on the bool *addr. Should be used in conjunction with TestAndSet.
// Valid memory orders:  OrderRelaxed, OrderSeqCst and OrderRelease.
//extern github_com_ikavalio_atomicx.Clear
func Clear(addr *bool, order MemoryOrder)
