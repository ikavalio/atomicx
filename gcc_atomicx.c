#include <stdint.h>
#include <stdbool.h>

#define _STRINGIFY2_(x) #x
#define _STRINGIFY_(x) _STRINGIFY2_(x)
#define GOSYM_PREFIX _STRINGIFY_(__USER_LABEL_PREFIX__)

const int modes[] = {
    __ATOMIC_RELAXED,
	__ATOMIC_CONSUME,
	__ATOMIC_ACQUIRE,
	__ATOMIC_RELEASE,
	__ATOMIC_ACQ_REL,
	__ATOMIC_SEQ_CST
};

#define TOKENPASTE(x, y) x ## y
#define TOKENPASTE2(x, y) TOKENPASTE(x, y)

#define TOKENPASTE2_STR(x, y) _STRINGIFY_(x ## y)

#define DECL_ATOMIC_LOAD(type, name) \
extern type TOKENPASTE2(Load, name) (type*, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(Load, name)); \
type TOKENPASTE2(Load, name) (type *addr, int mode) { \
    return __atomic_load_n(addr, modes[mode]); \
} \

#define DECL_ATOMIC_STORE(type, name) \
extern void TOKENPASTE2(Store, name) (type*, type, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(Store, name)); \
void TOKENPASTE2(Store, name) (type *addr, type new, int mode) { \
    __atomic_store_n(addr, new, modes[mode]); \
} \

#define DECL_ATOMIC_SWAP(type, name) \
extern type TOKENPASTE2(Swap, name) (type*, type, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(Swap, name)); \
type TOKENPASTE2(Swap, name) (type *addr, type new, int mode) { \
    return __atomic_exchange_n(addr, new, modes[mode]); \
} \

#define DECL_ATOMIC_CAS_STRONG(type, name) \
extern _Bool TOKENPASTE2(CompareAndSwapStrong, name) (type*, type, type, int mode) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(CompareAndSwapStrong, name)); \
extern _Bool TOKENPASTE2(CompareAndSwapStrong2, name) (type*, type, type, int, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(CompareAndSwapStrong2, name)); \
_Bool TOKENPASTE2(CompareAndSwapStrong, name) (type *addr, type expected, type new, int mode) { \
    return __atomic_compare_exchange_n(addr, &expected, new, false, modes[mode], modes[mode]); \
} \
_Bool TOKENPASTE2(CompareAndSwapStrong2, name) (type *addr, type expected, type new, int modeSuccess, int modeFailure) { \
    return __atomic_compare_exchange_n(addr, &expected, new, false, modes[modeSuccess], modes[modeFailure]); \
} \

#define DECL_ATOMIC_CAS_WEAK(type, name) \
extern _Bool TOKENPASTE2(CompareAndSwapWeak, name) (type*, type, type, int mode) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(CompareAndSwapWeak, name)); \
extern _Bool TOKENPASTE2(CompareAndSwapWeak2, name) (type*, type, type, int, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(CompareAndSwapWeak2, name)); \
_Bool TOKENPASTE2(CompareAndSwapWeak, name) (type *addr, type expected, type new, int mode) { \
    return __atomic_compare_exchange_n(addr, &expected, new, true, modes[mode], modes[mode]); \
} \
_Bool TOKENPASTE2(CompareAndSwapWeak2, name) (type *addr, type expected, type new, int modeSuccess, int modeFailure) { \
    return __atomic_compare_exchange_n(addr, &expected, new, true, modes[modeSuccess], modes[modeFailure]); \
} \

#define DECL_ADD_FETCH(type, name) \
extern type TOKENPASTE2(Add, name) (type*, type, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(Add, name)); \
type TOKENPASTE2(Add, name) (type *addr, type value, int mode) { \
    return __atomic_add_fetch(addr, value, modes[mode]); \
} \

#define DECL_AND_FETCH(type, name) \
extern type TOKENPASTE2(And, name) (type*, type, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(And, name)); \
type TOKENPASTE2(And, name) (type *addr, type value, int mode) { \
    return __atomic_and_fetch(addr, value, modes[mode]); \
} \

#define DECL_OR_FETCH(type, name) \
extern type TOKENPASTE2(Or, name) (type*, type, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(Or, name)); \
type TOKENPASTE2(Or, name) (type *addr, type value, int mode) { \
    return __atomic_or_fetch(addr, value, modes[mode]); \
} \

#define DECL_XOR_FETCH(type, name) \
extern type TOKENPASTE2(Xor, name) (type*, type, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(Xor, name)); \
type TOKENPASTE2(Xor, name) (type *addr, type value, int mode) { \
    return __atomic_xor_fetch(addr, value, modes[mode]); \
} \

#define DECL_NAND_FETCH(type, name) \
extern type TOKENPASTE2(Nand, name) (type*, type, int) \
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx." TOKENPASTE2_STR(Nand, name)); \
type TOKENPASTE2(Nand, name) (type *addr, type value, int mode) { \
    return __atomic_nand_fetch(addr, value, modes[mode]); \
} \

extern _Bool TestAndSet(_Bool*, int)
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx.TestAndSet");
_Bool TestAndSet(_Bool *ptr, int mode) {
	return __atomic_test_and_set(ptr, modes[mode]);
}

extern void Clear(_Bool*, int)
    __asm__(GOSYM_PREFIX "github_com_ikavalio_atomicx.Clear");
void Clear(_Bool *ptr, int mode) {
	__atomic_clear(ptr, modes[mode]);
}

DECL_ATOMIC_LOAD(int32_t, Int32)
DECL_ATOMIC_LOAD(int64_t, Int64)
DECL_ATOMIC_LOAD(uint32_t, Uint32)
DECL_ATOMIC_LOAD(uint64_t, Uint64)
DECL_ATOMIC_LOAD(uintptr_t, Uintptr)
DECL_ATOMIC_LOAD(void*, Pointer)

DECL_ATOMIC_STORE(int32_t, Int32)
DECL_ATOMIC_STORE(int64_t, Int64)
DECL_ATOMIC_STORE(uint32_t, Uint32)
DECL_ATOMIC_STORE(uint64_t, Uint64)
DECL_ATOMIC_STORE(uintptr_t, Uintptr)
DECL_ATOMIC_STORE(void*, Pointer)

DECL_ATOMIC_SWAP(int32_t, Int32)
DECL_ATOMIC_SWAP(int64_t, Int64)
DECL_ATOMIC_SWAP(uint32_t, Uint32)
DECL_ATOMIC_SWAP(uint64_t, Uint64)
DECL_ATOMIC_SWAP(uintptr_t, Uintptr)
DECL_ATOMIC_SWAP(void*, Pointer)

DECL_ATOMIC_CAS_STRONG(int32_t, Int32)
DECL_ATOMIC_CAS_STRONG(int64_t, Int64)
DECL_ATOMIC_CAS_STRONG(uint32_t, Uint32)
DECL_ATOMIC_CAS_STRONG(uint64_t, Uint64)
DECL_ATOMIC_CAS_STRONG(uintptr_t, Uintptr)
DECL_ATOMIC_CAS_STRONG(void*, Pointer)

DECL_ATOMIC_CAS_WEAK(int32_t, Int32)
DECL_ATOMIC_CAS_WEAK(int64_t, Int64)
DECL_ATOMIC_CAS_WEAK(uint32_t, Uint32)
DECL_ATOMIC_CAS_WEAK(uint64_t, Uint64)
DECL_ATOMIC_CAS_WEAK(uintptr_t, Uintptr)
DECL_ATOMIC_CAS_WEAK(void*, Pointer)

DECL_ADD_FETCH(int32_t, Int32)
DECL_ADD_FETCH(int64_t, Int64)
DECL_ADD_FETCH(uint32_t, Uint32)
DECL_ADD_FETCH(uint64_t, Uint64)
DECL_ADD_FETCH(uintptr_t, Uintptr)

DECL_AND_FETCH(int32_t, Int32)
DECL_AND_FETCH(int64_t, Int64)
DECL_AND_FETCH(uint32_t, Uint32)
DECL_AND_FETCH(uint64_t, Uint64)

DECL_OR_FETCH(int32_t, Int32)
DECL_OR_FETCH(int64_t, Int64)
DECL_OR_FETCH(uint32_t, Uint32)
DECL_OR_FETCH(uint64_t, Uint64)

DECL_XOR_FETCH(int32_t, Int32)
DECL_XOR_FETCH(int64_t, Int64)
DECL_XOR_FETCH(uint32_t, Uint32)
DECL_XOR_FETCH(uint64_t, Uint64)

DECL_NAND_FETCH(int32_t, Int32)
DECL_NAND_FETCH(int64_t, Int64)
DECL_NAND_FETCH(uint32_t, Uint32)
DECL_NAND_FETCH(uint64_t, Uint64)
