#include <stdint.h>
#include <stdbool.h>

// https://gcc.gnu.org/onlinedocs/gcc/_005f_005fatomic-Builtins.html

#define GOPKG_PREFIX "github_com_ikavalio_atomicx."

#define STR2(x) #x
#define STR(x) STR2(x)

#define __CONCAT3__(x, y, z) x ## y ## z
#define CONCAT3(x, y, z) __CONCAT3__(x, z, y)

#define __CONCAT4__(x, y, z, w) x ## y ## z ## w
#define CONCAT4(x, y, z, w) __CONCAT4__(x, w, y, z)

#define DECL_STORE(type, typeName, atomic, atomicName) \
extern void CONCAT3(Store, atomicName, typeName) (type*, type) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(Store, atomicName, typeName))); \
\
inline void CONCAT3(Store, atomicName, typeName) (type *num, type x) { \
    __atomic_store_n(num, x, atomic); \
} \

#define DECL_LOAD(type, typeName, atomic, atomicName) \
extern type CONCAT3(Load, atomicName, typeName) (type*) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(Load, atomicName, typeName))); \
\
inline type CONCAT3(Load, atomicName, typeName) (type *num) { \
    return  __atomic_load_n(num, atomic); \
} \

#define DECL_SWAP(type, typeName, atomic, atomicName) \
extern type CONCAT3(Swap, atomicName, typeName) (type*, type) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(Swap, atomicName, typeName))); \
\
inline type CONCAT3(Swap, atomicName, typeName) (type *num, type value) { \
    return __atomic_exchange_n(num, value, atomic); \
} \

#define DECL_CAS(type, typeName, atomic1, atomicName1, atomic2, atomicName2) \
extern bool CONCAT4(CompareAndSwapStrong, atomicName1, atomicName2, typeName) (type*, type, type) \
    __asm__(GOPKG_PREFIX STR(CONCAT4(CompareAndSwapStrong, atomicName1, atomicName2, typeName))); \
extern bool CONCAT4(CompareAndSwapWeak, atomicName1, atomicName2, typeName) (type*, type, type) \
    __asm__(GOPKG_PREFIX STR(CONCAT4(CompareAndSwapWeak, atomicName1, atomicName2, typeName))); \
\
inline bool CONCAT4(CompareAndSwapStrong, atomicName1, atomicName2, typeName) (type *ptr, type exp, type upd) { \
    return __atomic_compare_exchange_n(ptr, &exp, upd, false, atomic1, atomic2); \
} \
inline bool CONCAT4(CompareAndSwapWeak, atomicName1, atomicName2, typeName) (type *ptr, type exp, type upd) { \
    return __atomic_compare_exchange_n(ptr, &exp, upd, true, atomic1, atomic2); \
} \

#define DECL_ADD(type, typeName, atomic, atomicName) \
extern type CONCAT3(Add, atomicName, typeName) (type*, type) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(Add, atomicName, typeName))); \
\
inline type CONCAT3(Add, atomicName, typeName) (type *num, type value) { \
    return __atomic_add_fetch(num, value, atomic); \
} \

#define DECL_AND(type, typeName, atomic, atomicName) \
extern type CONCAT3(And, atomicName, typeName) (type*, type) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(And, atomicName, typeName))); \
\
inline type CONCAT3(And, atomicName, typeName) (type *num, type value) { \
    return __atomic_and_fetch(num, value, atomic); \
} \

#define DECL_OR(type, typeName, atomic, atomicName) \
extern type CONCAT3(Or, atomicName, typeName) (type*, type) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(Or, atomicName, typeName))); \
\
inline type CONCAT3(Or, atomicName, typeName) (type *num, type value) { \
    return __atomic_or_fetch(num, value, atomic); \
} \

#define DECL_XOR(type, typeName, atomic, atomicName) \
extern type CONCAT3(Xor, atomicName, typeName) (type*, type) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(Xor, atomicName, typeName))); \
\
inline type CONCAT3(Xor, atomicName, typeName) (type *num, type value) { \
    return __atomic_xor_fetch(num, value, atomic); \
} \

#define DECL_NAND(type, typeName, atomic, atomicName) \
extern type CONCAT3(Nand, atomicName, typeName) (type*, type) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(Nand, atomicName, typeName))); \
\
inline type CONCAT3(Nand, atomicName, typeName) (type *num, type value) { \
    return __atomic_nand_fetch(num, value, atomic); \
} \

#define DECL_TEST_AND_SET(atomic, atomicName) \
extern bool CONCAT3(TestAndSet, atomicName,) (bool*) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(TestAndSet, atomicName,))); \
\
inline bool CONCAT3(TestAndSet, atomicName,) (bool *num) { \
    return __atomic_test_and_set(num, atomic); \
} \

#define DECL_CLEAR(atomic, atomicName) \
extern void CONCAT3(Clear, atomicName,) (bool*) \
    __asm__(GOPKG_PREFIX STR(CONCAT3(Clear, atomicName,))); \
\
inline void CONCAT3(Clear, atomicName,) (bool *num) { \
    return __atomic_clear(num, atomic); \
} \

#define DECL_STORE_FUNCS(args...) \
DECL_STORE(args, __ATOMIC_RELAXED, Relaxed) \
DECL_STORE(args, __ATOMIC_RELEASE, Release) \
DECL_STORE(args, __ATOMIC_SEQ_CST, SeqCst) \

#define DECL_LOAD_FUNCS(args...) \
DECL_LOAD(args, __ATOMIC_RELAXED, Relaxed) \
DECL_LOAD(args, __ATOMIC_CONSUME, Consume) \
DECL_LOAD(args, __ATOMIC_ACQUIRE, Acquire) \
DECL_LOAD(args, __ATOMIC_SEQ_CST, SeqCst) \

#define DECL_SWAP_FUNCS(args...) \
DECL_SWAP(args, __ATOMIC_RELAXED, Relaxed) \
DECL_SWAP(args, __ATOMIC_ACQUIRE, Acquire) \
DECL_SWAP(args, __ATOMIC_RELEASE, Release) \
DECL_SWAP(args, __ATOMIC_ACQ_REL, AcqRel) \
DECL_SWAP(args, __ATOMIC_SEQ_CST, SeqCst) \

#define DECL_CAS_FUNCS(args...) \
DECL_CAS(args, __ATOMIC_RELAXED, Relaxed, __ATOMIC_RELAXED, Relaxed) \
DECL_CAS(args, __ATOMIC_CONSUME, Consume, __ATOMIC_RELAXED, Relaxed) \
DECL_CAS(args, __ATOMIC_CONSUME, Consume, __ATOMIC_CONSUME, Consume) \
DECL_CAS(args, __ATOMIC_ACQUIRE, Acquire, __ATOMIC_RELAXED, Relaxed) \
DECL_CAS(args, __ATOMIC_ACQUIRE, Acquire, __ATOMIC_CONSUME, Consume) \
DECL_CAS(args, __ATOMIC_ACQUIRE, Acquire, __ATOMIC_ACQUIRE, Acquire) \
DECL_CAS(args, __ATOMIC_RELEASE, Release, __ATOMIC_RELAXED, Relaxed) \
DECL_CAS(args, __ATOMIC_RELEASE, Release, __ATOMIC_CONSUME, Consume) \
DECL_CAS(args, __ATOMIC_RELEASE, Release, __ATOMIC_ACQUIRE, Acquire) \
DECL_CAS(args, __ATOMIC_ACQ_REL, AcqRel, __ATOMIC_RELAXED, Relaxed) \
DECL_CAS(args, __ATOMIC_ACQ_REL, AcqRel, __ATOMIC_CONSUME, Consume) \
DECL_CAS(args, __ATOMIC_ACQ_REL, AcqRel, __ATOMIC_ACQUIRE, Acquire) \
DECL_CAS(args, __ATOMIC_SEQ_CST, SeqCst, __ATOMIC_RELAXED, Relaxed) \
DECL_CAS(args, __ATOMIC_SEQ_CST, SeqCst, __ATOMIC_CONSUME, Consume) \
DECL_CAS(args, __ATOMIC_SEQ_CST, SeqCst, __ATOMIC_ACQUIRE, Acquire) \
DECL_CAS(args, __ATOMIC_SEQ_CST, SeqCst, __ATOMIC_SEQ_CST, SeqCst) \

#define DECL_ADD_FUNCS(args...) \
DECL_ADD(args, __ATOMIC_RELAXED, Relaxed) \
DECL_ADD(args, __ATOMIC_CONSUME, Consume) \
DECL_ADD(args, __ATOMIC_ACQUIRE, Acquire) \
DECL_ADD(args, __ATOMIC_RELEASE, Release) \
DECL_ADD(args, __ATOMIC_ACQ_REL, AcqRel) \
DECL_ADD(args, __ATOMIC_SEQ_CST, SeqCst) \

#define DECL_AND_FUNCS(args...) \
DECL_AND(args, __ATOMIC_RELAXED, Relaxed) \
DECL_AND(args, __ATOMIC_CONSUME, Consume) \
DECL_AND(args, __ATOMIC_ACQUIRE, Acquire) \
DECL_AND(args, __ATOMIC_RELEASE, Release) \
DECL_AND(args, __ATOMIC_ACQ_REL, AcqRel) \
DECL_AND(args, __ATOMIC_SEQ_CST, SeqCst) \

#define DECL_OR_FUNCS(args...) \
DECL_OR(args, __ATOMIC_RELAXED, Relaxed) \
DECL_OR(args, __ATOMIC_CONSUME, Consume) \
DECL_OR(args, __ATOMIC_ACQUIRE, Acquire) \
DECL_OR(args, __ATOMIC_RELEASE, Release) \
DECL_OR(args, __ATOMIC_ACQ_REL, AcqRel) \
DECL_OR(args, __ATOMIC_SEQ_CST, SeqCst) \

#define DECL_XOR_FUNCS(args...) \
DECL_XOR(args, __ATOMIC_RELAXED, Relaxed) \
DECL_XOR(args, __ATOMIC_CONSUME, Consume) \
DECL_XOR(args, __ATOMIC_ACQUIRE, Acquire) \
DECL_XOR(args, __ATOMIC_RELEASE, Release) \
DECL_XOR(args, __ATOMIC_ACQ_REL, AcqRel) \
DECL_XOR(args, __ATOMIC_SEQ_CST, SeqCst) \

#define DECL_NAND_FUNCS(args...) \
DECL_NAND(args, __ATOMIC_RELAXED, Relaxed) \
DECL_NAND(args, __ATOMIC_CONSUME, Consume) \
DECL_NAND(args, __ATOMIC_ACQUIRE, Acquire) \
DECL_NAND(args, __ATOMIC_RELEASE, Release) \
DECL_NAND(args, __ATOMIC_ACQ_REL, AcqRel) \
DECL_NAND(args, __ATOMIC_SEQ_CST, SeqCst) \

#define DECL_TEST_AND_SET_FUNCS() \
DECL_TEST_AND_SET(__ATOMIC_RELAXED, Relaxed) \
DECL_TEST_AND_SET(__ATOMIC_CONSUME, Consume) \
DECL_TEST_AND_SET(__ATOMIC_ACQUIRE, Acquire) \
DECL_TEST_AND_SET(__ATOMIC_RELEASE, Release) \
DECL_TEST_AND_SET(__ATOMIC_ACQ_REL, AcqRel) \
DECL_TEST_AND_SET(__ATOMIC_SEQ_CST, SeqCst) \

#define DECL_CLEAR_FUNCS() \
DECL_CLEAR(__ATOMIC_RELAXED, Relaxed) \
DECL_CLEAR(__ATOMIC_RELEASE, Release) \
DECL_CLEAR(__ATOMIC_SEQ_CST, SeqCst) \

DECL_STORE_FUNCS(int32_t, Int32)
DECL_STORE_FUNCS(int64_t, Int64)
DECL_STORE_FUNCS(uint32_t, Uint32)
DECL_STORE_FUNCS(uint64_t, Uint64)
DECL_STORE_FUNCS(uintptr_t, Uintptr)
DECL_STORE_FUNCS(void*, Pointer)

DECL_LOAD_FUNCS(int32_t, Int32)
DECL_LOAD_FUNCS(int64_t, Int64)
DECL_LOAD_FUNCS(uint32_t, Uint32)
DECL_LOAD_FUNCS(uint64_t, Uint64)
DECL_LOAD_FUNCS(uintptr_t, Uintptr)
DECL_LOAD_FUNCS(void*, Pointer)

DECL_CAS_FUNCS(int32_t, Int32)
DECL_CAS_FUNCS(int64_t, Int64)
DECL_CAS_FUNCS(uint32_t, Uint32)
DECL_CAS_FUNCS(uint64_t, Uint64)
DECL_CAS_FUNCS(uintptr_t, Uintptr)
DECL_CAS_FUNCS(void*, Pointer)

DECL_SWAP_FUNCS(int32_t, Int32)
DECL_SWAP_FUNCS(int64_t, Int64)
DECL_SWAP_FUNCS(uint32_t, Uint32)
DECL_SWAP_FUNCS(uint64_t, Uint64)
DECL_SWAP_FUNCS(uintptr_t, Uintptr)
DECL_SWAP_FUNCS(void*, Pointer)

DECL_ADD_FUNCS(int32_t, Int32)
DECL_ADD_FUNCS(int64_t, Int64)
DECL_ADD_FUNCS(uint32_t, Uint32)
DECL_ADD_FUNCS(uint64_t, Uint64)
DECL_ADD_FUNCS(uintptr_t, Uintptr)

DECL_AND_FUNCS(int32_t, Int32)
DECL_AND_FUNCS(int64_t, Int64)
DECL_AND_FUNCS(uint32_t, Uint32)
DECL_AND_FUNCS(uint64_t, Uint64)

DECL_OR_FUNCS(int32_t, Int32)
DECL_OR_FUNCS(int64_t, Int64)
DECL_OR_FUNCS(uint32_t, Uint32)
DECL_OR_FUNCS(uint64_t, Uint64)

DECL_XOR_FUNCS(int32_t, Int32)
DECL_XOR_FUNCS(int64_t, Int64)
DECL_XOR_FUNCS(uint32_t, Uint32)
DECL_XOR_FUNCS(uint64_t, Uint64)

DECL_NAND_FUNCS(int32_t, Int32)
DECL_NAND_FUNCS(int64_t, Int64)
DECL_NAND_FUNCS(uint32_t, Uint32)
DECL_NAND_FUNCS(uint64_t, Uint64)

DECL_TEST_AND_SET_FUNCS()
DECL_CLEAR_FUNCS()
