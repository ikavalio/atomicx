package atomicx

// https://gcc.gnu.org/onlinedocs/gcc/_005f_005fatomic-Builtins.html

// TODO: __atomic_thread_fence

/*
#include <stdint.h>
#include <stdbool.h>

typedef void* goptr_t;

#define TOKENPASTE(x, y) x ## y
#define TOKENPASTE2(x, y) TOKENPASTE(x, y)

#define ATOMIC_LOAD(TYPE) \
TYPE TOKENPASTE2(atomic_load_, TYPE) (TYPE *p, enum MemOrder order) { \
	return __atomic_load_n(p, order); \
} \

#define ATOMIC_STORE(TYPE) \
void TOKENPASTE2(atomic_store_, TYPE) (TYPE *p, TYPE v, enum MemOrder order) { \
	__atomic_store_n(p, v, order); \
} \

#define ATOMIC_ADD_FETCH(TYPE) \
TYPE TOKENPASTE2(atomic_add_fetch_, TYPE) (TYPE *p, TYPE v, enum MemOrder order) { \
	return __atomic_add_fetch(p, v, order); \
} \

#define ATOMIC_COMPARE_EXCHANGE(TYPE) \
bool TOKENPASTE2(atomic_compare_exchange_, TYPE) (TYPE *p, TYPE *exp, TYPE what, bool weak, enum MemOrder orderSuccess, enum MemOrder orderFailure) { \
	return __atomic_compare_exchange_n(p, exp, what, weak, orderSuccess, orderFailure); \
} \

#define ATOMIC_EXCHANGE(TYPE) \
TYPE TOKENPASTE2(atomic_exchange_, TYPE) (TYPE *p, TYPE what, enum MemOrder order) { \
	return __atomic_exchange_n(p, what, order); \
} \

#define ATOMIC_AND_FETCH(TYPE) \
TYPE TOKENPASTE2(atomic_and_fetch_, TYPE) (TYPE *p, TYPE v, enum MemOrder order) { \
	return __atomic_and_fetch(p, v, order); \
} \

#define ATOMIC_OR_FETCH(TYPE) \
TYPE TOKENPASTE2(atomic_or_fetch_, TYPE) (TYPE *p, TYPE v, enum MemOrder order) { \
	return __atomic_or_fetch(p, v, order); \
} \

#define ATOMIC_XOR_FETCH(TYPE) \
TYPE TOKENPASTE2(atomic_xor_fetch_, TYPE) (TYPE *p, TYPE v, enum MemOrder order) { \
	return __atomic_xor_fetch(p, v, order); \
} \

#define ATOMIC_NAND_FETCH(TYPE) \
TYPE TOKENPASTE2(atomic_nand_fetch_, TYPE) (TYPE *p, TYPE v, enum MemOrder order) { \
	return __atomic_nand_fetch(p, v, order); \
} \

enum MemOrder {
	OrderRelaxed = __ATOMIC_RELAXED,
	OrderConsume = __ATOMIC_CONSUME,
	OrderAcquire = __ATOMIC_ACQUIRE,
	OrderRelease = __ATOMIC_RELEASE,
	OrderAcqRel  = __ATOMIC_ACQ_REL,
	OrderSeqCst  = __ATOMIC_SEQ_CST
};

ATOMIC_LOAD(int32_t)
ATOMIC_LOAD(int64_t)
ATOMIC_LOAD(uint32_t)
ATOMIC_LOAD(uint64_t)
ATOMIC_LOAD(uintptr_t)
ATOMIC_LOAD(goptr_t)

ATOMIC_STORE(int32_t)
ATOMIC_STORE(int64_t)
ATOMIC_STORE(uint32_t)
ATOMIC_STORE(uint64_t)
ATOMIC_STORE(uintptr_t)
ATOMIC_STORE(goptr_t)

ATOMIC_ADD_FETCH(int32_t)
ATOMIC_ADD_FETCH(int64_t)
ATOMIC_ADD_FETCH(uint32_t)
ATOMIC_ADD_FETCH(uint64_t)
ATOMIC_ADD_FETCH(uintptr_t)

ATOMIC_COMPARE_EXCHANGE(int32_t)
ATOMIC_COMPARE_EXCHANGE(int64_t)
ATOMIC_COMPARE_EXCHANGE(uint32_t)
ATOMIC_COMPARE_EXCHANGE(uint64_t)
ATOMIC_COMPARE_EXCHANGE(uintptr_t)
ATOMIC_COMPARE_EXCHANGE(goptr_t)

ATOMIC_EXCHANGE(int32_t)
ATOMIC_EXCHANGE(int64_t)
ATOMIC_EXCHANGE(uint32_t)
ATOMIC_EXCHANGE(uint64_t)
ATOMIC_EXCHANGE(uintptr_t)
ATOMIC_EXCHANGE(goptr_t)

bool atomic_test_and_set(bool *ptr, enum MemOrder order) {
	return __atomic_test_and_set(ptr, order);
}

void atomic_clear(bool *ptr, enum MemOrder order) {
	__atomic_clear(ptr, order);
}

ATOMIC_AND_FETCH(int32_t)
ATOMIC_AND_FETCH(int64_t)
ATOMIC_AND_FETCH(uint32_t)
ATOMIC_AND_FETCH(uint64_t)

ATOMIC_OR_FETCH(int32_t)
ATOMIC_OR_FETCH(int64_t)
ATOMIC_OR_FETCH(uint32_t)
ATOMIC_OR_FETCH(uint64_t)

ATOMIC_XOR_FETCH(int32_t)
ATOMIC_XOR_FETCH(int64_t)
ATOMIC_XOR_FETCH(uint32_t)
ATOMIC_XOR_FETCH(uint64_t)

ATOMIC_NAND_FETCH(int32_t)
ATOMIC_NAND_FETCH(int64_t)
ATOMIC_NAND_FETCH(uint32_t)
ATOMIC_NAND_FETCH(uint64_t)
*/
import "C"

import "unsafe"

// MemOrder represents C++11 memory order
type MemOrder C.enum_MemOrder

const (
	// OrderRelaxed implies no inter-thread ordering constraints.
	OrderRelaxed MemOrder = C.OrderRelaxed
	// OrderConsume is deprecated and equals to OrderAcquire
	OrderConsume MemOrder = C.OrderConsume
	// OrderAcquire creates an inter-thread happens-before constraint from the release (or stronger) semantic store to this acquire load.
	OrderAcquire MemOrder = C.OrderAcquire
	// OrderRelease creates an inter-thread happens-before constraint to acquire (or stronger) semantic loads that read from this release store.
	OrderRelease MemOrder = C.OrderRelease
	// OrderAcqRel combines the effects of both __ATOMIC_ACQUIRE and __ATOMIC_RELEASE.
	OrderAcqRel MemOrder = C.OrderAcqRel
	// OrderSeqCst enforces total ordering with all other sequentially consistent operations.
	OrderSeqCst MemOrder = C.OrderSeqCst
)

func (order MemOrder) String() string {
	switch order {
	case OrderRelaxed:
		return "MemOrder Relaxed"
	case OrderConsume:
		return "MemOrder Consume"
	case OrderAcquire:
		return "MemOrder Acquire"
	case OrderRelease:
		return "MemOrder Release"
	case OrderAcqRel:
		return "MemOrder Acquire-Release"
	case OrderSeqCst:
		return "MemOrder Sequentially Consistent"
	default:
		return "MemOrder Unknown"
	}
}

func (order MemOrder) asC() C.enum_MemOrder {
	return (C.enum_MemOrder)(order)
}

func cptrI32(addr *int32) *C.int32_t {
	return (*C.int32_t)(unsafe.Pointer(addr))
}

func cptrI64(addr *int64) *C.int64_t {
	return (*C.int64_t)(unsafe.Pointer(addr))
}

func cptrU32(addr *uint32) *C.uint32_t {
	return (*C.uint32_t)(unsafe.Pointer(addr))
}

func cptrU64(addr *uint64) *C.uint64_t {
	return (*C.uint64_t)(unsafe.Pointer(addr))
}

func cptrUptr(addr *uintptr) *C.uintptr_t {
	return (*C.uintptr_t)(unsafe.Pointer(addr))
}

func cptrPtr(addr *unsafe.Pointer) *C.goptr_t {
	return (*C.goptr_t)(addr)
}

func cptrBool(addr *bool) *C._Bool {
	return (*C._Bool)(unsafe.Pointer(addr))
}

// AddInt32 atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
func AddInt32(addr *int32, delta int32, order MemOrder) int32 {
	return (int32)(C.atomic_add_fetch_int32_t(cptrI32(addr), (C.int32_t)(delta), order.asC()))
}

// AddInt64 atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
func AddInt64(addr *int64, delta int64, order MemOrder) int64 {
	return (int64)(C.atomic_add_fetch_int64_t(cptrI64(addr), (C.int64_t)(delta), order.asC()))
}

// AddUint32 atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
func AddUint32(addr *uint32, delta uint32, order MemOrder) uint32 {
	return (uint32)(C.atomic_add_fetch_uint32_t(cptrU32(addr), (C.uint32_t)(delta), order.asC()))
}

// AddUint64 atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
func AddUint64(addr *uint64, delta uint64, order MemOrder) uint64 {
	return (uint64)(C.atomic_add_fetch_uint64_t(cptrU64(addr), (C.uint64_t)(delta), order.asC()))
}

// AddUintptr atomically adds delta to *addr and returns the new value.
// Valid memory orders: all.
func AddUintptr(addr *uintptr, delta uintptr, order MemOrder) uintptr {
	return (uintptr)(C.atomic_add_fetch_uintptr_t(cptrUptr(addr), (C.uintptr_t)(delta), order.asC()))
}

// CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
// If weak flag is set to true CAS may fail spuriously.
// Valid memory orders: all.
func CompareAndSwapInt32(addr *int32, old, new int32, weak bool, order MemOrder) bool {
	return CompareAndSwap2Int32(addr, old, new, weak, order, order)
}

// CompareAndSwap2Int32 executes the compare-and-swap operation for an int32 value.
// If weak is true CAS may fail spuriously.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
func CompareAndSwap2Int32(addr *int32, old, new int32, weak bool, orderSuccess MemOrder, orderFailure MemOrder) bool {
	return (bool)(C.atomic_compare_exchange_int32_t(cptrI32(addr), cptrI32(&old), (C.int32_t)(new), (C._Bool)(weak), orderSuccess.asC(), orderFailure.asC()))
}

// CompareAndSwapInt64 executes the compare-and-swap operation for an int64 value.
// If weak flag is set to true CAS may fail spuriously.
// Valid memory orders: all.
func CompareAndSwapInt64(addr *int64, old, new int64, weak bool, order MemOrder) bool {
	return CompareAndSwap2Int64(addr, old, new, weak, order, order)
}

// CompareAndSwap2Int64 executes the compare-and-swap operation for an int64 value.
// If weak is true CAS may fail spuriously.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
func CompareAndSwap2Int64(addr *int64, old, new int64, weak bool, orderSuccess MemOrder, orderFailure MemOrder) bool {
	return (bool)(C.atomic_compare_exchange_int64_t(cptrI64(addr), cptrI64(&old), (C.int64_t)(new), (C._Bool)(weak), orderSuccess.asC(), orderFailure.asC()))
}

// CompareAndSwapPointer executes the compare-and-swap operation for an unsafe.Pointer value.
// If weak flag is set to true CAS may fail spuriously.
// Valid memory orders: all.
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer, weak bool, order MemOrder) bool {
	return CompareAndSwap2Pointer(addr, old, new, weak, order, order)
}

// CompareAndSwap2Pointer executes the compare-and-swap operation for an unsafe.Pointer value.
// If weak is true CAS may fail spuriously.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
func CompareAndSwap2Pointer(addr *unsafe.Pointer, old, new unsafe.Pointer, weak bool, orderSuccess MemOrder, orderFailure MemOrder) bool {
	return (bool)(C.atomic_compare_exchange_goptr_t(cptrPtr(addr), cptrPtr(&old), (C.goptr_t)(new), (C._Bool)(weak), orderSuccess.asC(), orderFailure.asC()))
}

// CompareAndSwapUint32 executes the compare-and-swap operation for an uint32 value.
// If weak flag is set to true CAS may fail spuriously.
// Valid memory orders: all.
func CompareAndSwapUint32(addr *uint32, old, new uint32, weak bool, order MemOrder) bool {
	return CompareAndSwap2Uint32(addr, old, new, weak, order, order)
}

// CompareAndSwap2Uint32 executes the compare-and-swap operation for an uint32 value.
// If weak is true CAS may fail spuriously.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
func CompareAndSwap2Uint32(addr *uint32, old, new uint32, weak bool, orderSuccess MemOrder, orderFailure MemOrder) bool {
	return (bool)(C.atomic_compare_exchange_uint32_t(cptrU32(addr), cptrU32(&old), (C.uint32_t)(new), (C._Bool)(weak), orderSuccess.asC(), orderFailure.asC()))
}

// CompareAndSwapUint64 executes the compare-and-swap operation for an uint64 value.
// If weak flag is set to true CAS may fail spuriously.
// Valid memory orders: all.
func CompareAndSwapUint64(addr *uint64, old, new uint64, weak bool, order MemOrder) bool {
	return CompareAndSwap2Uint64(addr, old, new, weak, order, order)
}

// CompareAndSwap2Uint64 executes the compare-and-swap operation for an uint64 value.
// If weak is true CAS may fail spuriously.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
func CompareAndSwap2Uint64(addr *uint64, old, new uint64, weak bool, orderSuccess MemOrder, orderFailure MemOrder) bool {
	return (bool)(C.atomic_compare_exchange_uint64_t(cptrU64(addr), cptrU64(&old), (C.uint64_t)(new), (C._Bool)(weak), orderSuccess.asC(), orderFailure.asC()))
}

// CompareAndSwapUintptr executes the compare-and-swap operation for an uintptr value.
// If weak flag is set to true CAS may fail spuriously.
// Valid memory orders: all.
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr, weak bool, order MemOrder) bool {
	return CompareAndSwap2Uintptr(addr, old, new, weak, order, order)
}

// CompareAndSwap2Uintptr executes the compare-and-swap operation for an uintptr value.
// If weak is true CAS may fail spuriously.
// Valid memory orders for orderSuccess: all.
// Valid memory orders for orderFailure: cannot be OrderRelease nor OrderAcqRel. It also cannot be a stronger order than that specified by orderSuccess.
func CompareAndSwap2Uintptr(addr *uintptr, old, new uintptr, weak bool, orderSuccess MemOrder, orderFailure MemOrder) bool {
	return (bool)(C.atomic_compare_exchange_uintptr_t(cptrUptr(addr), cptrUptr(&old), (C.uintptr_t)(new), (C._Bool)(weak), orderSuccess.asC(), orderFailure.asC()))
}

// LoadInt32 atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
func LoadInt32(addr *int32, order MemOrder) int32 {
	return (int32)(C.atomic_load_int32_t(cptrI32(addr), order.asC()))
}

// LoadInt64 atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
func LoadInt64(addr *int64, order MemOrder) int64 {
	return (int64)(C.atomic_load_int64_t(cptrI64(addr), order.asC()))
}

// LoadPointer atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
func LoadPointer(addr *unsafe.Pointer, order MemOrder) unsafe.Pointer {
	return (unsafe.Pointer)(C.atomic_load_goptr_t(cptrPtr(addr), order.asC()))
}

// LoadUint32 atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
func LoadUint32(addr *uint32, order MemOrder) uint32 {
	return (uint32)(C.atomic_load_uint32_t(cptrU32(addr), order.asC()))
}

// LoadUint64 atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
func LoadUint64(addr *uint64, order MemOrder) uint64 {
	return (uint64)(C.atomic_load_uint64_t(cptrU64(addr), order.asC()))
}

// LoadUintptr atomically loads *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst, OrderAcquire and OrderConsume.
func LoadUintptr(addr *uintptr, order MemOrder) uintptr {
	return (uintptr)(C.atomic_load_uintptr_t(cptrUptr(addr), order.asC()))
}

// StoreInt32 atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
func StoreInt32(addr *int32, val int32, order MemOrder) {
	C.atomic_store_int32_t(cptrI32(addr), (C.int32_t)(val), order.asC())
}

// StoreInt64 atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
func StoreInt64(addr *int64, val int64, order MemOrder) {
	C.atomic_store_int64_t(cptrI64(addr), (C.int64_t)(val), order.asC())
}

// StorePointer atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer, order MemOrder) {
	C.atomic_store_goptr_t(cptrPtr(addr), (C.goptr_t)(val), order.asC())
}

// StoreUint32 atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
func StoreUint32(addr *uint32, val uint32, order MemOrder) {
	C.atomic_store_uint32_t(cptrU32(addr), (C.uint32_t)(val), order.asC())
}

// StoreUint64 atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
func StoreUint64(addr *uint64, val uint64, order MemOrder) {
	C.atomic_store_uint64_t(cptrU64(addr), (C.uint64_t)(val), order.asC())
}

// StoreUintptr atomically stores val into *addr.
// Valid memory orders: OrderRelaxed, OrderSeqCst and OrderRelease.
func StoreUintptr(addr *uintptr, val uintptr, order MemOrder) {
	C.atomic_store_uintptr_t(cptrUptr(addr), (C.uintptr_t)(val), order.asC())
}

// SwapInt32 atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
func SwapInt32(addr *int32, new int32, order MemOrder) int32 {
	return (int32)(C.atomic_exchange_int32_t(cptrI32(addr), (C.int32_t)(new), order.asC()))
}

// SwapInt64 atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
func SwapInt64(addr *int64, new int64, order MemOrder) int64 {
	return (int64)(C.atomic_exchange_int64_t(cptrI64(addr), (C.int64_t)(new), order.asC()))
}

// SwapPointer atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer, order MemOrder) unsafe.Pointer {
	return (unsafe.Pointer)(C.atomic_exchange_goptr_t(cptrPtr(addr), (C.goptr_t)(new), order.asC()))
}

// SwapUint32 atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
func SwapUint32(addr *uint32, new uint32, order MemOrder) uint32 {
	return (uint32)(C.atomic_exchange_uint32_t(cptrU32(addr), (C.uint32_t)(new), order.asC()))
}

// SwapUint64 atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
func SwapUint64(addr *uint64, new uint64, order MemOrder) uint64 {
	return (uint64)(C.atomic_exchange_uint64_t(cptrU64(addr), (C.uint64_t)(new), order.asC()))
}

// SwapUintptr atomically stores new into *addr and returns the previous *addr value.
// Valid memory orders:  OrderRelaxed, OrderSeqCst, OrderAcquire, OrderRelease, and OrderAcqRel.
func SwapUintptr(addr *uintptr, new uintptr, order MemOrder) uintptr {
	return (uintptr)(C.atomic_exchange_uintptr_t(cptrUptr(addr), (C.uintptr_t)(new), order.asC()))
}

// TestAndSet does atomic test-and-set operation on the bool *addr (atomically *addr = *addr ? *addr : true).
// Returns the old value of *addr.
// Valid memory orders:  all.
func TestAndSet(addr *bool, order MemOrder) bool {
	return (bool)(C.atomic_test_and_set(cptrBool(addr), order.asC()))
}

// Clear does atomic clear operation on the bool *addr. Should be used in conjunction with TestAndSet.
// Valid memory orders:  OrderRelaxed, OrderSeqCst and OrderRelease.
func Clear(addr *bool, order MemOrder) {
	C.atomic_clear(cptrBool(addr), order.asC())
}

// AndInt32 does atomic bitwise-and between delta and *addr and returns the new value.
// Valid memory orders: all.
func AndInt32(addr *int32, delta int32, order MemOrder) int32 {
	return (int32)(C.atomic_and_fetch_int32_t(cptrI32(addr), (C.int32_t)(delta), order.asC()))
}

// AndInt64 does atomic bitwise-and between delta and *addr and returns the new value.
// Valid memory orders: all.
func AndInt64(addr *int64, delta int64, order MemOrder) int64 {
	return (int64)(C.atomic_and_fetch_int64_t(cptrI64(addr), (C.int64_t)(delta), order.asC()))
}

// AndUint32 does atomic bitwise-and between delta and *addr and returns the new value.
// Valid memory orders: all.
func AndUint32(addr *uint32, delta uint32, order MemOrder) uint32 {
	return (uint32)(C.atomic_and_fetch_uint32_t(cptrU32(addr), (C.uint32_t)(delta), order.asC()))
}

// AndUint64 does atomic bitwise-and between delta and *addr and returns the new value.
// Valid memory orders: all.
func AndUint64(addr *uint64, delta uint64, order MemOrder) uint64 {
	return (uint64)(C.atomic_and_fetch_uint64_t(cptrU64(addr), (C.uint64_t)(delta), order.asC()))
}

// OrInt32 does atomic bitwise-or between delta and *addr and returns the new value.
// Valid memory orders: all.
func OrInt32(addr *int32, delta int32, order MemOrder) int32 {
	return (int32)(C.atomic_or_fetch_int32_t(cptrI32(addr), (C.int32_t)(delta), order.asC()))
}

// OrInt64 does atomic bitwise-or between delta and *addr and returns the new value.
// Valid memory orders: all.
func OrInt64(addr *int64, delta int64, order MemOrder) int64 {
	return (int64)(C.atomic_or_fetch_int64_t(cptrI64(addr), (C.int64_t)(delta), order.asC()))
}

// OrUint32 does atomic bitwise-or between delta and *addr and returns the new value.
// Valid memory orders: all.
func OrUint32(addr *uint32, delta uint32, order MemOrder) uint32 {
	return (uint32)(C.atomic_or_fetch_uint32_t(cptrU32(addr), (C.uint32_t)(delta), order.asC()))
}

// OrUint64 does atomic bitwise-or between delta and *addr and returns the new value.
// Valid memory orders: all.
func OrUint64(addr *uint64, delta uint64, order MemOrder) uint64 {
	return (uint64)(C.atomic_or_fetch_uint64_t(cptrU64(addr), (C.uint64_t)(delta), order.asC()))
}

// XorInt32 does atomic bitwise-xor between delta and *addr and returns the new value.
// Valid memory orders: all.
func XorInt32(addr *int32, delta int32, order MemOrder) int32 {
	return (int32)(C.atomic_xor_fetch_int32_t(cptrI32(addr), (C.int32_t)(delta), order.asC()))
}

// XorInt64 does atomic bitwise-xor between delta and *addr and returns the new value.
// Valid memory orders: all.
func XorInt64(addr *int64, delta int64, order MemOrder) int64 {
	return (int64)(C.atomic_xor_fetch_int64_t(cptrI64(addr), (C.int64_t)(delta), order.asC()))
}

// XorUint32 does atomic bitwise-xor between delta and *addr and returns the new value.
// Valid memory orders: all.
func XorUint32(addr *uint32, delta uint32, order MemOrder) uint32 {
	return (uint32)(C.atomic_xor_fetch_uint32_t(cptrU32(addr), (C.uint32_t)(delta), order.asC()))
}

// XorUint64 does atomic bitwise-xor between delta and *addr and returns the new value.
// Valid memory orders: all.
func XorUint64(addr *uint64, delta uint64, order MemOrder) uint64 {
	return (uint64)(C.atomic_xor_fetch_uint64_t(cptrU64(addr), (C.uint64_t)(delta), order.asC()))
}

// NandInt32 does atomic bitwise-nand between delta and *addr and returns the new value.
// Valid memory orders: all.
func NandInt32(addr *int32, delta int32, order MemOrder) int32 {
	return (int32)(C.atomic_nand_fetch_int32_t(cptrI32(addr), (C.int32_t)(delta), order.asC()))
}

// NandInt64 does atomic bitwise-nand between delta and *addr and returns the new value.
// Valid memory orders: all.
func NandInt64(addr *int64, delta int64, order MemOrder) int64 {
	return (int64)(C.atomic_nand_fetch_int64_t(cptrI64(addr), (C.int64_t)(delta), order.asC()))
}

// NandUint32 does atomic bitwise-nand between delta and *addr and returns the new value.
// Valid memory orders: all.
func NandUint32(addr *uint32, delta uint32, order MemOrder) uint32 {
	return (uint32)(C.atomic_nand_fetch_uint32_t(cptrU32(addr), (C.uint32_t)(delta), order.asC()))
}

// NandUint64 does atomic bitwise-nand between delta and *addr and returns the new value.
// Valid memory orders: all.
func NandUint64(addr *uint64, delta uint64, order MemOrder) uint64 {
	return (uint64)(C.atomic_nand_fetch_uint64_t(cptrU64(addr), (C.uint64_t)(delta), order.asC()))
}
