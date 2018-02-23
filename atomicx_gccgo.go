package atomicx

import "unsafe"

//StoreInt32Relaxed atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreInt32Relaxed
func StoreInt32Relaxed(addr *int32, val int32)


//StoreInt32Release atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreInt32Release
func StoreInt32Release(addr *int32, val int32)


//StoreInt32SeqCst atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreInt32SeqCst
func StoreInt32SeqCst(addr *int32, val int32)


//StoreInt64Relaxed atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreInt64Relaxed
func StoreInt64Relaxed(addr *int64, val int64)


//StoreInt64Release atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreInt64Release
func StoreInt64Release(addr *int64, val int64)


//StoreInt64SeqCst atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreInt64SeqCst
func StoreInt64SeqCst(addr *int64, val int64)


//StoreUint32Relaxed atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreUint32Relaxed
func StoreUint32Relaxed(addr *uint32, val uint32)


//StoreUint32Release atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreUint32Release
func StoreUint32Release(addr *uint32, val uint32)


//StoreUint32SeqCst atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreUint32SeqCst
func StoreUint32SeqCst(addr *uint32, val uint32)


//StoreUint64Relaxed atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreUint64Relaxed
func StoreUint64Relaxed(addr *uint64, val uint64)


//StoreUint64Release atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreUint64Release
func StoreUint64Release(addr *uint64, val uint64)


//StoreUint64SeqCst atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreUint64SeqCst
func StoreUint64SeqCst(addr *uint64, val uint64)


//StoreUintptrRelaxed atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreUintptrRelaxed
func StoreUintptrRelaxed(addr *uintptr, val uintptr)


//StoreUintptrRelease atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreUintptrRelease
func StoreUintptrRelease(addr *uintptr, val uintptr)


//StoreUintptrSeqCst atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StoreUintptrSeqCst
func StoreUintptrSeqCst(addr *uintptr, val uintptr)


//StorePointerRelaxed atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StorePointerRelaxed
func StorePointerRelaxed(addr *unsafe.Pointer, val unsafe.Pointer)


//StorePointerRelease atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StorePointerRelease
func StorePointerRelease(addr *unsafe.Pointer, val unsafe.Pointer)


//StorePointerSeqCst atomically stores val into *addr.
//extern github_com_ikavalio_atomicx.StorePointerSeqCst
func StorePointerSeqCst(addr *unsafe.Pointer, val unsafe.Pointer)


//LoadInt32Relaxed atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadInt32Relaxed
func LoadInt32Relaxed(addr *int32) int32


//LoadInt32Consume atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadInt32Consume
func LoadInt32Consume(addr *int32) int32


//LoadInt32Acquire atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadInt32Acquire
func LoadInt32Acquire(addr *int32) int32


//LoadInt32SeqCst atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadInt32SeqCst
func LoadInt32SeqCst(addr *int32) int32


//LoadInt64Relaxed atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadInt64Relaxed
func LoadInt64Relaxed(addr *int64) int64


//LoadInt64Consume atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadInt64Consume
func LoadInt64Consume(addr *int64) int64


//LoadInt64Acquire atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadInt64Acquire
func LoadInt64Acquire(addr *int64) int64


//LoadInt64SeqCst atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadInt64SeqCst
func LoadInt64SeqCst(addr *int64) int64


//LoadUint32Relaxed atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUint32Relaxed
func LoadUint32Relaxed(addr *uint32) uint32


//LoadUint32Consume atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUint32Consume
func LoadUint32Consume(addr *uint32) uint32


//LoadUint32Acquire atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUint32Acquire
func LoadUint32Acquire(addr *uint32) uint32


//LoadUint32SeqCst atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUint32SeqCst
func LoadUint32SeqCst(addr *uint32) uint32


//LoadUint64Relaxed atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUint64Relaxed
func LoadUint64Relaxed(addr *uint64) uint64


//LoadUint64Consume atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUint64Consume
func LoadUint64Consume(addr *uint64) uint64


//LoadUint64Acquire atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUint64Acquire
func LoadUint64Acquire(addr *uint64) uint64


//LoadUint64SeqCst atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUint64SeqCst
func LoadUint64SeqCst(addr *uint64) uint64


//LoadUintptrRelaxed atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUintptrRelaxed
func LoadUintptrRelaxed(addr *uintptr) uintptr


//LoadUintptrConsume atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUintptrConsume
func LoadUintptrConsume(addr *uintptr) uintptr


//LoadUintptrAcquire atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUintptrAcquire
func LoadUintptrAcquire(addr *uintptr) uintptr


//LoadUintptrSeqCst atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadUintptrSeqCst
func LoadUintptrSeqCst(addr *uintptr) uintptr


//LoadPointerRelaxed atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadPointerRelaxed
func LoadPointerRelaxed(addr *unsafe.Pointer) unsafe.Pointer


//LoadPointerConsume atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadPointerConsume
func LoadPointerConsume(addr *unsafe.Pointer) unsafe.Pointer


//LoadPointerAcquire atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadPointerAcquire
func LoadPointerAcquire(addr *unsafe.Pointer) unsafe.Pointer


//LoadPointerSeqCst atomically loads *addr.
//extern github_com_ikavalio_atomicx.LoadPointerSeqCst
func LoadPointerSeqCst(addr *unsafe.Pointer) unsafe.Pointer


//CompareAndSwapStrongInt32RelaxedRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32RelaxedRelaxed
func CompareAndSwapStrongInt32RelaxedRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32RelaxedRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32RelaxedRelaxed
func CompareAndSwapWeakInt32RelaxedRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32ConsumeRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32ConsumeRelaxed
func CompareAndSwapStrongInt32ConsumeRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32ConsumeRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32ConsumeRelaxed
func CompareAndSwapWeakInt32ConsumeRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32ConsumeConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32ConsumeConsume
func CompareAndSwapStrongInt32ConsumeConsume(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32ConsumeConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32ConsumeConsume
func CompareAndSwapWeakInt32ConsumeConsume(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32AcquireRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32AcquireRelaxed
func CompareAndSwapStrongInt32AcquireRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32AcquireRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32AcquireRelaxed
func CompareAndSwapWeakInt32AcquireRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32AcquireConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32AcquireConsume
func CompareAndSwapStrongInt32AcquireConsume(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32AcquireConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32AcquireConsume
func CompareAndSwapWeakInt32AcquireConsume(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32AcquireAcquire executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32AcquireAcquire
func CompareAndSwapStrongInt32AcquireAcquire(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32AcquireAcquire executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32AcquireAcquire
func CompareAndSwapWeakInt32AcquireAcquire(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32ReleaseRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32ReleaseRelaxed
func CompareAndSwapStrongInt32ReleaseRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32ReleaseRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32ReleaseRelaxed
func CompareAndSwapWeakInt32ReleaseRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32ReleaseConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32ReleaseConsume
func CompareAndSwapStrongInt32ReleaseConsume(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32ReleaseConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32ReleaseConsume
func CompareAndSwapWeakInt32ReleaseConsume(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32ReleaseAcquire executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32ReleaseAcquire
func CompareAndSwapStrongInt32ReleaseAcquire(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32ReleaseAcquire executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32ReleaseAcquire
func CompareAndSwapWeakInt32ReleaseAcquire(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32AcqRelRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32AcqRelRelaxed
func CompareAndSwapStrongInt32AcqRelRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32AcqRelRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32AcqRelRelaxed
func CompareAndSwapWeakInt32AcqRelRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32AcqRelConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32AcqRelConsume
func CompareAndSwapStrongInt32AcqRelConsume(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32AcqRelConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32AcqRelConsume
func CompareAndSwapWeakInt32AcqRelConsume(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32AcqRelAcquire executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32AcqRelAcquire
func CompareAndSwapStrongInt32AcqRelAcquire(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32AcqRelAcquire executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32AcqRelAcquire
func CompareAndSwapWeakInt32AcqRelAcquire(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32SeqCstRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32SeqCstRelaxed
func CompareAndSwapStrongInt32SeqCstRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32SeqCstRelaxed executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32SeqCstRelaxed
func CompareAndSwapWeakInt32SeqCstRelaxed(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32SeqCstConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32SeqCstConsume
func CompareAndSwapStrongInt32SeqCstConsume(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32SeqCstConsume executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32SeqCstConsume
func CompareAndSwapWeakInt32SeqCstConsume(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32SeqCstAcquire executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32SeqCstAcquire
func CompareAndSwapStrongInt32SeqCstAcquire(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32SeqCstAcquire executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32SeqCstAcquire
func CompareAndSwapWeakInt32SeqCstAcquire(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt32SeqCstSeqCst executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt32SeqCstSeqCst
func CompareAndSwapStrongInt32SeqCstSeqCst(addr *int32, old, new int32) bool


//CompareAndSwapWeakInt32SeqCstSeqCst executes the compare-and-swap operation for an int32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt32SeqCstSeqCst
func CompareAndSwapWeakInt32SeqCstSeqCst(addr *int32, old, new int32) bool


//CompareAndSwapStrongInt64RelaxedRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64RelaxedRelaxed
func CompareAndSwapStrongInt64RelaxedRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64RelaxedRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64RelaxedRelaxed
func CompareAndSwapWeakInt64RelaxedRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64ConsumeRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64ConsumeRelaxed
func CompareAndSwapStrongInt64ConsumeRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64ConsumeRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64ConsumeRelaxed
func CompareAndSwapWeakInt64ConsumeRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64ConsumeConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64ConsumeConsume
func CompareAndSwapStrongInt64ConsumeConsume(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64ConsumeConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64ConsumeConsume
func CompareAndSwapWeakInt64ConsumeConsume(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64AcquireRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64AcquireRelaxed
func CompareAndSwapStrongInt64AcquireRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64AcquireRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64AcquireRelaxed
func CompareAndSwapWeakInt64AcquireRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64AcquireConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64AcquireConsume
func CompareAndSwapStrongInt64AcquireConsume(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64AcquireConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64AcquireConsume
func CompareAndSwapWeakInt64AcquireConsume(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64AcquireAcquire executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64AcquireAcquire
func CompareAndSwapStrongInt64AcquireAcquire(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64AcquireAcquire executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64AcquireAcquire
func CompareAndSwapWeakInt64AcquireAcquire(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64ReleaseRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64ReleaseRelaxed
func CompareAndSwapStrongInt64ReleaseRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64ReleaseRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64ReleaseRelaxed
func CompareAndSwapWeakInt64ReleaseRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64ReleaseConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64ReleaseConsume
func CompareAndSwapStrongInt64ReleaseConsume(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64ReleaseConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64ReleaseConsume
func CompareAndSwapWeakInt64ReleaseConsume(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64ReleaseAcquire executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64ReleaseAcquire
func CompareAndSwapStrongInt64ReleaseAcquire(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64ReleaseAcquire executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64ReleaseAcquire
func CompareAndSwapWeakInt64ReleaseAcquire(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64AcqRelRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64AcqRelRelaxed
func CompareAndSwapStrongInt64AcqRelRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64AcqRelRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64AcqRelRelaxed
func CompareAndSwapWeakInt64AcqRelRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64AcqRelConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64AcqRelConsume
func CompareAndSwapStrongInt64AcqRelConsume(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64AcqRelConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64AcqRelConsume
func CompareAndSwapWeakInt64AcqRelConsume(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64AcqRelAcquire executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64AcqRelAcquire
func CompareAndSwapStrongInt64AcqRelAcquire(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64AcqRelAcquire executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64AcqRelAcquire
func CompareAndSwapWeakInt64AcqRelAcquire(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64SeqCstRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64SeqCstRelaxed
func CompareAndSwapStrongInt64SeqCstRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64SeqCstRelaxed executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64SeqCstRelaxed
func CompareAndSwapWeakInt64SeqCstRelaxed(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64SeqCstConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64SeqCstConsume
func CompareAndSwapStrongInt64SeqCstConsume(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64SeqCstConsume executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64SeqCstConsume
func CompareAndSwapWeakInt64SeqCstConsume(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64SeqCstAcquire executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64SeqCstAcquire
func CompareAndSwapStrongInt64SeqCstAcquire(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64SeqCstAcquire executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64SeqCstAcquire
func CompareAndSwapWeakInt64SeqCstAcquire(addr *int64, old, new int64) bool


//CompareAndSwapStrongInt64SeqCstSeqCst executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongInt64SeqCstSeqCst
func CompareAndSwapStrongInt64SeqCstSeqCst(addr *int64, old, new int64) bool


//CompareAndSwapWeakInt64SeqCstSeqCst executes the compare-and-swap operation for an int64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakInt64SeqCstSeqCst
func CompareAndSwapWeakInt64SeqCstSeqCst(addr *int64, old, new int64) bool


//CompareAndSwapStrongUint32RelaxedRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32RelaxedRelaxed
func CompareAndSwapStrongUint32RelaxedRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32RelaxedRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32RelaxedRelaxed
func CompareAndSwapWeakUint32RelaxedRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32ConsumeRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32ConsumeRelaxed
func CompareAndSwapStrongUint32ConsumeRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32ConsumeRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32ConsumeRelaxed
func CompareAndSwapWeakUint32ConsumeRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32ConsumeConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32ConsumeConsume
func CompareAndSwapStrongUint32ConsumeConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32ConsumeConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32ConsumeConsume
func CompareAndSwapWeakUint32ConsumeConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32AcquireRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32AcquireRelaxed
func CompareAndSwapStrongUint32AcquireRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32AcquireRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32AcquireRelaxed
func CompareAndSwapWeakUint32AcquireRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32AcquireConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32AcquireConsume
func CompareAndSwapStrongUint32AcquireConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32AcquireConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32AcquireConsume
func CompareAndSwapWeakUint32AcquireConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32AcquireAcquire executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32AcquireAcquire
func CompareAndSwapStrongUint32AcquireAcquire(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32AcquireAcquire executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32AcquireAcquire
func CompareAndSwapWeakUint32AcquireAcquire(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32ReleaseRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32ReleaseRelaxed
func CompareAndSwapStrongUint32ReleaseRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32ReleaseRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32ReleaseRelaxed
func CompareAndSwapWeakUint32ReleaseRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32ReleaseConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32ReleaseConsume
func CompareAndSwapStrongUint32ReleaseConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32ReleaseConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32ReleaseConsume
func CompareAndSwapWeakUint32ReleaseConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32ReleaseAcquire executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32ReleaseAcquire
func CompareAndSwapStrongUint32ReleaseAcquire(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32ReleaseAcquire executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32ReleaseAcquire
func CompareAndSwapWeakUint32ReleaseAcquire(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32AcqRelRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32AcqRelRelaxed
func CompareAndSwapStrongUint32AcqRelRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32AcqRelRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32AcqRelRelaxed
func CompareAndSwapWeakUint32AcqRelRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32AcqRelConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32AcqRelConsume
func CompareAndSwapStrongUint32AcqRelConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32AcqRelConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32AcqRelConsume
func CompareAndSwapWeakUint32AcqRelConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32AcqRelAcquire executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32AcqRelAcquire
func CompareAndSwapStrongUint32AcqRelAcquire(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32AcqRelAcquire executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32AcqRelAcquire
func CompareAndSwapWeakUint32AcqRelAcquire(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32SeqCstRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32SeqCstRelaxed
func CompareAndSwapStrongUint32SeqCstRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32SeqCstRelaxed executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32SeqCstRelaxed
func CompareAndSwapWeakUint32SeqCstRelaxed(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32SeqCstConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32SeqCstConsume
func CompareAndSwapStrongUint32SeqCstConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32SeqCstConsume executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32SeqCstConsume
func CompareAndSwapWeakUint32SeqCstConsume(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32SeqCstAcquire executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32SeqCstAcquire
func CompareAndSwapStrongUint32SeqCstAcquire(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32SeqCstAcquire executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32SeqCstAcquire
func CompareAndSwapWeakUint32SeqCstAcquire(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint32SeqCstSeqCst executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint32SeqCstSeqCst
func CompareAndSwapStrongUint32SeqCstSeqCst(addr *uint32, old, new uint32) bool


//CompareAndSwapWeakUint32SeqCstSeqCst executes the compare-and-swap operation for an uint32 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint32SeqCstSeqCst
func CompareAndSwapWeakUint32SeqCstSeqCst(addr *uint32, old, new uint32) bool


//CompareAndSwapStrongUint64RelaxedRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64RelaxedRelaxed
func CompareAndSwapStrongUint64RelaxedRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64RelaxedRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64RelaxedRelaxed
func CompareAndSwapWeakUint64RelaxedRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64ConsumeRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64ConsumeRelaxed
func CompareAndSwapStrongUint64ConsumeRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64ConsumeRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64ConsumeRelaxed
func CompareAndSwapWeakUint64ConsumeRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64ConsumeConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64ConsumeConsume
func CompareAndSwapStrongUint64ConsumeConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64ConsumeConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64ConsumeConsume
func CompareAndSwapWeakUint64ConsumeConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64AcquireRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64AcquireRelaxed
func CompareAndSwapStrongUint64AcquireRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64AcquireRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64AcquireRelaxed
func CompareAndSwapWeakUint64AcquireRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64AcquireConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64AcquireConsume
func CompareAndSwapStrongUint64AcquireConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64AcquireConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64AcquireConsume
func CompareAndSwapWeakUint64AcquireConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64AcquireAcquire executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64AcquireAcquire
func CompareAndSwapStrongUint64AcquireAcquire(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64AcquireAcquire executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64AcquireAcquire
func CompareAndSwapWeakUint64AcquireAcquire(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64ReleaseRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64ReleaseRelaxed
func CompareAndSwapStrongUint64ReleaseRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64ReleaseRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64ReleaseRelaxed
func CompareAndSwapWeakUint64ReleaseRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64ReleaseConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64ReleaseConsume
func CompareAndSwapStrongUint64ReleaseConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64ReleaseConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64ReleaseConsume
func CompareAndSwapWeakUint64ReleaseConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64ReleaseAcquire executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64ReleaseAcquire
func CompareAndSwapStrongUint64ReleaseAcquire(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64ReleaseAcquire executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64ReleaseAcquire
func CompareAndSwapWeakUint64ReleaseAcquire(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64AcqRelRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64AcqRelRelaxed
func CompareAndSwapStrongUint64AcqRelRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64AcqRelRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64AcqRelRelaxed
func CompareAndSwapWeakUint64AcqRelRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64AcqRelConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64AcqRelConsume
func CompareAndSwapStrongUint64AcqRelConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64AcqRelConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64AcqRelConsume
func CompareAndSwapWeakUint64AcqRelConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64AcqRelAcquire executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64AcqRelAcquire
func CompareAndSwapStrongUint64AcqRelAcquire(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64AcqRelAcquire executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64AcqRelAcquire
func CompareAndSwapWeakUint64AcqRelAcquire(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64SeqCstRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64SeqCstRelaxed
func CompareAndSwapStrongUint64SeqCstRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64SeqCstRelaxed executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64SeqCstRelaxed
func CompareAndSwapWeakUint64SeqCstRelaxed(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64SeqCstConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64SeqCstConsume
func CompareAndSwapStrongUint64SeqCstConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64SeqCstConsume executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64SeqCstConsume
func CompareAndSwapWeakUint64SeqCstConsume(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64SeqCstAcquire executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64SeqCstAcquire
func CompareAndSwapStrongUint64SeqCstAcquire(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64SeqCstAcquire executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64SeqCstAcquire
func CompareAndSwapWeakUint64SeqCstAcquire(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUint64SeqCstSeqCst executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUint64SeqCstSeqCst
func CompareAndSwapStrongUint64SeqCstSeqCst(addr *uint64, old, new uint64) bool


//CompareAndSwapWeakUint64SeqCstSeqCst executes the compare-and-swap operation for an uint64 value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUint64SeqCstSeqCst
func CompareAndSwapWeakUint64SeqCstSeqCst(addr *uint64, old, new uint64) bool


//CompareAndSwapStrongUintptrRelaxedRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrRelaxedRelaxed
func CompareAndSwapStrongUintptrRelaxedRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrRelaxedRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrRelaxedRelaxed
func CompareAndSwapWeakUintptrRelaxedRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrConsumeRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrConsumeRelaxed
func CompareAndSwapStrongUintptrConsumeRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrConsumeRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrConsumeRelaxed
func CompareAndSwapWeakUintptrConsumeRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrConsumeConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrConsumeConsume
func CompareAndSwapStrongUintptrConsumeConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrConsumeConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrConsumeConsume
func CompareAndSwapWeakUintptrConsumeConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrAcquireRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrAcquireRelaxed
func CompareAndSwapStrongUintptrAcquireRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrAcquireRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrAcquireRelaxed
func CompareAndSwapWeakUintptrAcquireRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrAcquireConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrAcquireConsume
func CompareAndSwapStrongUintptrAcquireConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrAcquireConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrAcquireConsume
func CompareAndSwapWeakUintptrAcquireConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrAcquireAcquire executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrAcquireAcquire
func CompareAndSwapStrongUintptrAcquireAcquire(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrAcquireAcquire executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrAcquireAcquire
func CompareAndSwapWeakUintptrAcquireAcquire(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrReleaseRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrReleaseRelaxed
func CompareAndSwapStrongUintptrReleaseRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrReleaseRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrReleaseRelaxed
func CompareAndSwapWeakUintptrReleaseRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrReleaseConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrReleaseConsume
func CompareAndSwapStrongUintptrReleaseConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrReleaseConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrReleaseConsume
func CompareAndSwapWeakUintptrReleaseConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrReleaseAcquire executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrReleaseAcquire
func CompareAndSwapStrongUintptrReleaseAcquire(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrReleaseAcquire executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrReleaseAcquire
func CompareAndSwapWeakUintptrReleaseAcquire(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrAcqRelRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrAcqRelRelaxed
func CompareAndSwapStrongUintptrAcqRelRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrAcqRelRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrAcqRelRelaxed
func CompareAndSwapWeakUintptrAcqRelRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrAcqRelConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrAcqRelConsume
func CompareAndSwapStrongUintptrAcqRelConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrAcqRelConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrAcqRelConsume
func CompareAndSwapWeakUintptrAcqRelConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrAcqRelAcquire executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrAcqRelAcquire
func CompareAndSwapStrongUintptrAcqRelAcquire(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrAcqRelAcquire executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrAcqRelAcquire
func CompareAndSwapWeakUintptrAcqRelAcquire(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrSeqCstRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrSeqCstRelaxed
func CompareAndSwapStrongUintptrSeqCstRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrSeqCstRelaxed executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrSeqCstRelaxed
func CompareAndSwapWeakUintptrSeqCstRelaxed(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrSeqCstConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrSeqCstConsume
func CompareAndSwapStrongUintptrSeqCstConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrSeqCstConsume executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrSeqCstConsume
func CompareAndSwapWeakUintptrSeqCstConsume(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrSeqCstAcquire executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrSeqCstAcquire
func CompareAndSwapStrongUintptrSeqCstAcquire(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrSeqCstAcquire executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrSeqCstAcquire
func CompareAndSwapWeakUintptrSeqCstAcquire(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongUintptrSeqCstSeqCst executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongUintptrSeqCstSeqCst
func CompareAndSwapStrongUintptrSeqCstSeqCst(addr *uintptr, old, new uintptr) bool


//CompareAndSwapWeakUintptrSeqCstSeqCst executes the compare-and-swap operation for an uintptr value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakUintptrSeqCstSeqCst
func CompareAndSwapWeakUintptrSeqCstSeqCst(addr *uintptr, old, new uintptr) bool


//CompareAndSwapStrongPointerRelaxedRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerRelaxedRelaxed
func CompareAndSwapStrongPointerRelaxedRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerRelaxedRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerRelaxedRelaxed
func CompareAndSwapWeakPointerRelaxedRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerConsumeRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerConsumeRelaxed
func CompareAndSwapStrongPointerConsumeRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerConsumeRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerConsumeRelaxed
func CompareAndSwapWeakPointerConsumeRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerConsumeConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerConsumeConsume
func CompareAndSwapStrongPointerConsumeConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerConsumeConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerConsumeConsume
func CompareAndSwapWeakPointerConsumeConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerAcquireRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerAcquireRelaxed
func CompareAndSwapStrongPointerAcquireRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerAcquireRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerAcquireRelaxed
func CompareAndSwapWeakPointerAcquireRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerAcquireConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerAcquireConsume
func CompareAndSwapStrongPointerAcquireConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerAcquireConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerAcquireConsume
func CompareAndSwapWeakPointerAcquireConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerAcquireAcquire executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerAcquireAcquire
func CompareAndSwapStrongPointerAcquireAcquire(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerAcquireAcquire executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerAcquireAcquire
func CompareAndSwapWeakPointerAcquireAcquire(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerReleaseRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerReleaseRelaxed
func CompareAndSwapStrongPointerReleaseRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerReleaseRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerReleaseRelaxed
func CompareAndSwapWeakPointerReleaseRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerReleaseConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerReleaseConsume
func CompareAndSwapStrongPointerReleaseConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerReleaseConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerReleaseConsume
func CompareAndSwapWeakPointerReleaseConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerReleaseAcquire executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerReleaseAcquire
func CompareAndSwapStrongPointerReleaseAcquire(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerReleaseAcquire executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerReleaseAcquire
func CompareAndSwapWeakPointerReleaseAcquire(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerAcqRelRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerAcqRelRelaxed
func CompareAndSwapStrongPointerAcqRelRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerAcqRelRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerAcqRelRelaxed
func CompareAndSwapWeakPointerAcqRelRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerAcqRelConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerAcqRelConsume
func CompareAndSwapStrongPointerAcqRelConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerAcqRelConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerAcqRelConsume
func CompareAndSwapWeakPointerAcqRelConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerAcqRelAcquire executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerAcqRelAcquire
func CompareAndSwapStrongPointerAcqRelAcquire(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerAcqRelAcquire executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerAcqRelAcquire
func CompareAndSwapWeakPointerAcqRelAcquire(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerSeqCstRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerSeqCstRelaxed
func CompareAndSwapStrongPointerSeqCstRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerSeqCstRelaxed executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerSeqCstRelaxed
func CompareAndSwapWeakPointerSeqCstRelaxed(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerSeqCstConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerSeqCstConsume
func CompareAndSwapStrongPointerSeqCstConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerSeqCstConsume executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerSeqCstConsume
func CompareAndSwapWeakPointerSeqCstConsume(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerSeqCstAcquire executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerSeqCstAcquire
func CompareAndSwapStrongPointerSeqCstAcquire(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerSeqCstAcquire executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerSeqCstAcquire
func CompareAndSwapWeakPointerSeqCstAcquire(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapStrongPointerSeqCstSeqCst executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapStrongPointerSeqCstSeqCst
func CompareAndSwapStrongPointerSeqCstSeqCst(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//CompareAndSwapWeakPointerSeqCstSeqCst executes the compare-and-swap operation for an unsafe.Pointer value.
//extern github_com_ikavalio_atomicx.CompareAndSwapWeakPointerSeqCstSeqCst
func CompareAndSwapWeakPointerSeqCstSeqCst(addr *unsafe.Pointer, old, new unsafe.Pointer) bool


//SwapInt32Relaxed atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt32Relaxed
func SwapInt32Relaxed(addr *int32, value int32) int32


//SwapInt32Acquire atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt32Acquire
func SwapInt32Acquire(addr *int32, value int32) int32


//SwapInt32Release atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt32Release
func SwapInt32Release(addr *int32, value int32) int32


//SwapInt32AcqRel atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt32AcqRel
func SwapInt32AcqRel(addr *int32, value int32) int32


//SwapInt32SeqCst atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt32SeqCst
func SwapInt32SeqCst(addr *int32, value int32) int32


//SwapInt64Relaxed atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt64Relaxed
func SwapInt64Relaxed(addr *int64, value int64) int64


//SwapInt64Acquire atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt64Acquire
func SwapInt64Acquire(addr *int64, value int64) int64


//SwapInt64Release atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt64Release
func SwapInt64Release(addr *int64, value int64) int64


//SwapInt64AcqRel atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt64AcqRel
func SwapInt64AcqRel(addr *int64, value int64) int64


//SwapInt64SeqCst atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapInt64SeqCst
func SwapInt64SeqCst(addr *int64, value int64) int64


//SwapUint32Relaxed atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint32Relaxed
func SwapUint32Relaxed(addr *uint32, value uint32) uint32


//SwapUint32Acquire atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint32Acquire
func SwapUint32Acquire(addr *uint32, value uint32) uint32


//SwapUint32Release atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint32Release
func SwapUint32Release(addr *uint32, value uint32) uint32


//SwapUint32AcqRel atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint32AcqRel
func SwapUint32AcqRel(addr *uint32, value uint32) uint32


//SwapUint32SeqCst atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint32SeqCst
func SwapUint32SeqCst(addr *uint32, value uint32) uint32


//SwapUint64Relaxed atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint64Relaxed
func SwapUint64Relaxed(addr *uint64, value uint64) uint64


//SwapUint64Acquire atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint64Acquire
func SwapUint64Acquire(addr *uint64, value uint64) uint64


//SwapUint64Release atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint64Release
func SwapUint64Release(addr *uint64, value uint64) uint64


//SwapUint64AcqRel atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint64AcqRel
func SwapUint64AcqRel(addr *uint64, value uint64) uint64


//SwapUint64SeqCst atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUint64SeqCst
func SwapUint64SeqCst(addr *uint64, value uint64) uint64


//SwapUintptrRelaxed atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUintptrRelaxed
func SwapUintptrRelaxed(addr *uintptr, value uintptr) uintptr


//SwapUintptrAcquire atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUintptrAcquire
func SwapUintptrAcquire(addr *uintptr, value uintptr) uintptr


//SwapUintptrRelease atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUintptrRelease
func SwapUintptrRelease(addr *uintptr, value uintptr) uintptr


//SwapUintptrAcqRel atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUintptrAcqRel
func SwapUintptrAcqRel(addr *uintptr, value uintptr) uintptr


//SwapUintptrSeqCst atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapUintptrSeqCst
func SwapUintptrSeqCst(addr *uintptr, value uintptr) uintptr


//SwapPointerRelaxed atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapPointerRelaxed
func SwapPointerRelaxed(addr *unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer


//SwapPointerAcquire atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapPointerAcquire
func SwapPointerAcquire(addr *unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer


//SwapPointerRelease atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapPointerRelease
func SwapPointerRelease(addr *unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer


//SwapPointerAcqRel atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapPointerAcqRel
func SwapPointerAcqRel(addr *unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer


//SwapPointerSeqCst atomically stores new into *addr and returns the previous *addr value.
//extern github_com_ikavalio_atomicx.SwapPointerSeqCst
func SwapPointerSeqCst(addr *unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer


//AddInt32Relaxed applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt32Relaxed
func AddInt32Relaxed(addr *int32, value int32) int32


//AddInt32Consume applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt32Consume
func AddInt32Consume(addr *int32, value int32) int32


//AddInt32Acquire applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt32Acquire
func AddInt32Acquire(addr *int32, value int32) int32


//AddInt32Release applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt32Release
func AddInt32Release(addr *int32, value int32) int32


//AddInt32AcqRel applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt32AcqRel
func AddInt32AcqRel(addr *int32, value int32) int32


//AddInt32SeqCst applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt32SeqCst
func AddInt32SeqCst(addr *int32, value int32) int32


//AddInt64Relaxed applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt64Relaxed
func AddInt64Relaxed(addr *int64, value int64) int64


//AddInt64Consume applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt64Consume
func AddInt64Consume(addr *int64, value int64) int64


//AddInt64Acquire applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt64Acquire
func AddInt64Acquire(addr *int64, value int64) int64


//AddInt64Release applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt64Release
func AddInt64Release(addr *int64, value int64) int64


//AddInt64AcqRel applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt64AcqRel
func AddInt64AcqRel(addr *int64, value int64) int64


//AddInt64SeqCst applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddInt64SeqCst
func AddInt64SeqCst(addr *int64, value int64) int64


//AddUint32Relaxed applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint32Relaxed
func AddUint32Relaxed(addr *uint32, value uint32) uint32


//AddUint32Consume applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint32Consume
func AddUint32Consume(addr *uint32, value uint32) uint32


//AddUint32Acquire applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint32Acquire
func AddUint32Acquire(addr *uint32, value uint32) uint32


//AddUint32Release applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint32Release
func AddUint32Release(addr *uint32, value uint32) uint32


//AddUint32AcqRel applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint32AcqRel
func AddUint32AcqRel(addr *uint32, value uint32) uint32


//AddUint32SeqCst applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint32SeqCst
func AddUint32SeqCst(addr *uint32, value uint32) uint32


//AddUint64Relaxed applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint64Relaxed
func AddUint64Relaxed(addr *uint64, value uint64) uint64


//AddUint64Consume applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint64Consume
func AddUint64Consume(addr *uint64, value uint64) uint64


//AddUint64Acquire applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint64Acquire
func AddUint64Acquire(addr *uint64, value uint64) uint64


//AddUint64Release applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint64Release
func AddUint64Release(addr *uint64, value uint64) uint64


//AddUint64AcqRel applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint64AcqRel
func AddUint64AcqRel(addr *uint64, value uint64) uint64


//AddUint64SeqCst applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUint64SeqCst
func AddUint64SeqCst(addr *uint64, value uint64) uint64


//AddUintptrRelaxed applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUintptrRelaxed
func AddUintptrRelaxed(addr *uintptr, value uintptr) uintptr


//AddUintptrConsume applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUintptrConsume
func AddUintptrConsume(addr *uintptr, value uintptr) uintptr


//AddUintptrAcquire applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUintptrAcquire
func AddUintptrAcquire(addr *uintptr, value uintptr) uintptr


//AddUintptrRelease applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUintptrRelease
func AddUintptrRelease(addr *uintptr, value uintptr) uintptr


//AddUintptrAcqRel applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUintptrAcqRel
func AddUintptrAcqRel(addr *uintptr, value uintptr) uintptr


//AddUintptrSeqCst applies add to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AddUintptrSeqCst
func AddUintptrSeqCst(addr *uintptr, value uintptr) uintptr


//AndInt32Relaxed applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt32Relaxed
func AndInt32Relaxed(addr *int32, value int32) int32


//AndInt32Consume applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt32Consume
func AndInt32Consume(addr *int32, value int32) int32


//AndInt32Acquire applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt32Acquire
func AndInt32Acquire(addr *int32, value int32) int32


//AndInt32Release applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt32Release
func AndInt32Release(addr *int32, value int32) int32


//AndInt32AcqRel applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt32AcqRel
func AndInt32AcqRel(addr *int32, value int32) int32


//AndInt32SeqCst applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt32SeqCst
func AndInt32SeqCst(addr *int32, value int32) int32


//AndInt64Relaxed applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt64Relaxed
func AndInt64Relaxed(addr *int64, value int64) int64


//AndInt64Consume applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt64Consume
func AndInt64Consume(addr *int64, value int64) int64


//AndInt64Acquire applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt64Acquire
func AndInt64Acquire(addr *int64, value int64) int64


//AndInt64Release applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt64Release
func AndInt64Release(addr *int64, value int64) int64


//AndInt64AcqRel applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt64AcqRel
func AndInt64AcqRel(addr *int64, value int64) int64


//AndInt64SeqCst applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndInt64SeqCst
func AndInt64SeqCst(addr *int64, value int64) int64


//AndUint32Relaxed applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint32Relaxed
func AndUint32Relaxed(addr *uint32, value uint32) uint32


//AndUint32Consume applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint32Consume
func AndUint32Consume(addr *uint32, value uint32) uint32


//AndUint32Acquire applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint32Acquire
func AndUint32Acquire(addr *uint32, value uint32) uint32


//AndUint32Release applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint32Release
func AndUint32Release(addr *uint32, value uint32) uint32


//AndUint32AcqRel applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint32AcqRel
func AndUint32AcqRel(addr *uint32, value uint32) uint32


//AndUint32SeqCst applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint32SeqCst
func AndUint32SeqCst(addr *uint32, value uint32) uint32


//AndUint64Relaxed applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint64Relaxed
func AndUint64Relaxed(addr *uint64, value uint64) uint64


//AndUint64Consume applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint64Consume
func AndUint64Consume(addr *uint64, value uint64) uint64


//AndUint64Acquire applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint64Acquire
func AndUint64Acquire(addr *uint64, value uint64) uint64


//AndUint64Release applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint64Release
func AndUint64Release(addr *uint64, value uint64) uint64


//AndUint64AcqRel applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint64AcqRel
func AndUint64AcqRel(addr *uint64, value uint64) uint64


//AndUint64SeqCst applies and to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.AndUint64SeqCst
func AndUint64SeqCst(addr *uint64, value uint64) uint64


//OrInt32Relaxed applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt32Relaxed
func OrInt32Relaxed(addr *int32, value int32) int32


//OrInt32Consume applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt32Consume
func OrInt32Consume(addr *int32, value int32) int32


//OrInt32Acquire applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt32Acquire
func OrInt32Acquire(addr *int32, value int32) int32


//OrInt32Release applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt32Release
func OrInt32Release(addr *int32, value int32) int32


//OrInt32AcqRel applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt32AcqRel
func OrInt32AcqRel(addr *int32, value int32) int32


//OrInt32SeqCst applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt32SeqCst
func OrInt32SeqCst(addr *int32, value int32) int32


//OrInt64Relaxed applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt64Relaxed
func OrInt64Relaxed(addr *int64, value int64) int64


//OrInt64Consume applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt64Consume
func OrInt64Consume(addr *int64, value int64) int64


//OrInt64Acquire applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt64Acquire
func OrInt64Acquire(addr *int64, value int64) int64


//OrInt64Release applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt64Release
func OrInt64Release(addr *int64, value int64) int64


//OrInt64AcqRel applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt64AcqRel
func OrInt64AcqRel(addr *int64, value int64) int64


//OrInt64SeqCst applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrInt64SeqCst
func OrInt64SeqCst(addr *int64, value int64) int64


//OrUint32Relaxed applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint32Relaxed
func OrUint32Relaxed(addr *uint32, value uint32) uint32


//OrUint32Consume applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint32Consume
func OrUint32Consume(addr *uint32, value uint32) uint32


//OrUint32Acquire applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint32Acquire
func OrUint32Acquire(addr *uint32, value uint32) uint32


//OrUint32Release applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint32Release
func OrUint32Release(addr *uint32, value uint32) uint32


//OrUint32AcqRel applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint32AcqRel
func OrUint32AcqRel(addr *uint32, value uint32) uint32


//OrUint32SeqCst applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint32SeqCst
func OrUint32SeqCst(addr *uint32, value uint32) uint32


//OrUint64Relaxed applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint64Relaxed
func OrUint64Relaxed(addr *uint64, value uint64) uint64


//OrUint64Consume applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint64Consume
func OrUint64Consume(addr *uint64, value uint64) uint64


//OrUint64Acquire applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint64Acquire
func OrUint64Acquire(addr *uint64, value uint64) uint64


//OrUint64Release applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint64Release
func OrUint64Release(addr *uint64, value uint64) uint64


//OrUint64AcqRel applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint64AcqRel
func OrUint64AcqRel(addr *uint64, value uint64) uint64


//OrUint64SeqCst applies or to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.OrUint64SeqCst
func OrUint64SeqCst(addr *uint64, value uint64) uint64


//XorInt32Relaxed applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt32Relaxed
func XorInt32Relaxed(addr *int32, value int32) int32


//XorInt32Consume applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt32Consume
func XorInt32Consume(addr *int32, value int32) int32


//XorInt32Acquire applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt32Acquire
func XorInt32Acquire(addr *int32, value int32) int32


//XorInt32Release applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt32Release
func XorInt32Release(addr *int32, value int32) int32


//XorInt32AcqRel applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt32AcqRel
func XorInt32AcqRel(addr *int32, value int32) int32


//XorInt32SeqCst applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt32SeqCst
func XorInt32SeqCst(addr *int32, value int32) int32


//XorInt64Relaxed applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt64Relaxed
func XorInt64Relaxed(addr *int64, value int64) int64


//XorInt64Consume applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt64Consume
func XorInt64Consume(addr *int64, value int64) int64


//XorInt64Acquire applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt64Acquire
func XorInt64Acquire(addr *int64, value int64) int64


//XorInt64Release applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt64Release
func XorInt64Release(addr *int64, value int64) int64


//XorInt64AcqRel applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt64AcqRel
func XorInt64AcqRel(addr *int64, value int64) int64


//XorInt64SeqCst applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorInt64SeqCst
func XorInt64SeqCst(addr *int64, value int64) int64


//XorUint32Relaxed applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint32Relaxed
func XorUint32Relaxed(addr *uint32, value uint32) uint32


//XorUint32Consume applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint32Consume
func XorUint32Consume(addr *uint32, value uint32) uint32


//XorUint32Acquire applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint32Acquire
func XorUint32Acquire(addr *uint32, value uint32) uint32


//XorUint32Release applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint32Release
func XorUint32Release(addr *uint32, value uint32) uint32


//XorUint32AcqRel applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint32AcqRel
func XorUint32AcqRel(addr *uint32, value uint32) uint32


//XorUint32SeqCst applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint32SeqCst
func XorUint32SeqCst(addr *uint32, value uint32) uint32


//XorUint64Relaxed applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint64Relaxed
func XorUint64Relaxed(addr *uint64, value uint64) uint64


//XorUint64Consume applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint64Consume
func XorUint64Consume(addr *uint64, value uint64) uint64


//XorUint64Acquire applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint64Acquire
func XorUint64Acquire(addr *uint64, value uint64) uint64


//XorUint64Release applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint64Release
func XorUint64Release(addr *uint64, value uint64) uint64


//XorUint64AcqRel applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint64AcqRel
func XorUint64AcqRel(addr *uint64, value uint64) uint64


//XorUint64SeqCst applies xor to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.XorUint64SeqCst
func XorUint64SeqCst(addr *uint64, value uint64) uint64


//NandInt32Relaxed applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt32Relaxed
func NandInt32Relaxed(addr *int32, value int32) int32


//NandInt32Consume applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt32Consume
func NandInt32Consume(addr *int32, value int32) int32


//NandInt32Acquire applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt32Acquire
func NandInt32Acquire(addr *int32, value int32) int32


//NandInt32Release applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt32Release
func NandInt32Release(addr *int32, value int32) int32


//NandInt32AcqRel applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt32AcqRel
func NandInt32AcqRel(addr *int32, value int32) int32


//NandInt32SeqCst applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt32SeqCst
func NandInt32SeqCst(addr *int32, value int32) int32


//NandInt64Relaxed applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt64Relaxed
func NandInt64Relaxed(addr *int64, value int64) int64


//NandInt64Consume applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt64Consume
func NandInt64Consume(addr *int64, value int64) int64


//NandInt64Acquire applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt64Acquire
func NandInt64Acquire(addr *int64, value int64) int64


//NandInt64Release applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt64Release
func NandInt64Release(addr *int64, value int64) int64


//NandInt64AcqRel applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt64AcqRel
func NandInt64AcqRel(addr *int64, value int64) int64


//NandInt64SeqCst applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandInt64SeqCst
func NandInt64SeqCst(addr *int64, value int64) int64


//NandUint32Relaxed applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint32Relaxed
func NandUint32Relaxed(addr *uint32, value uint32) uint32


//NandUint32Consume applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint32Consume
func NandUint32Consume(addr *uint32, value uint32) uint32


//NandUint32Acquire applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint32Acquire
func NandUint32Acquire(addr *uint32, value uint32) uint32


//NandUint32Release applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint32Release
func NandUint32Release(addr *uint32, value uint32) uint32


//NandUint32AcqRel applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint32AcqRel
func NandUint32AcqRel(addr *uint32, value uint32) uint32


//NandUint32SeqCst applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint32SeqCst
func NandUint32SeqCst(addr *uint32, value uint32) uint32


//NandUint64Relaxed applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint64Relaxed
func NandUint64Relaxed(addr *uint64, value uint64) uint64


//NandUint64Consume applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint64Consume
func NandUint64Consume(addr *uint64, value uint64) uint64


//NandUint64Acquire applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint64Acquire
func NandUint64Acquire(addr *uint64, value uint64) uint64


//NandUint64Release applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint64Release
func NandUint64Release(addr *uint64, value uint64) uint64


//NandUint64AcqRel applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint64AcqRel
func NandUint64AcqRel(addr *uint64, value uint64) uint64


//NandUint64SeqCst applies nand to delta and *addr and returns the new value.
//extern github_com_ikavalio_atomicx.NandUint64SeqCst
func NandUint64SeqCst(addr *uint64, value uint64) uint64


//TestAndSetRelaxed does atomic test-and-set operation on the bool *addr (atomically *addr = *addr ? *addr : true).
//extern github_com_ikavalio_atomicx.TestAndSetRelaxed
func TestAndSetRelaxed(addr *bool) bool


//TestAndSetConsume does atomic test-and-set operation on the bool *addr (atomically *addr = *addr ? *addr : true).
//extern github_com_ikavalio_atomicx.TestAndSetConsume
func TestAndSetConsume(addr *bool) bool


//TestAndSetAcquire does atomic test-and-set operation on the bool *addr (atomically *addr = *addr ? *addr : true).
//extern github_com_ikavalio_atomicx.TestAndSetAcquire
func TestAndSetAcquire(addr *bool) bool


//TestAndSetRelease does atomic test-and-set operation on the bool *addr (atomically *addr = *addr ? *addr : true).
//extern github_com_ikavalio_atomicx.TestAndSetRelease
func TestAndSetRelease(addr *bool) bool


//TestAndSetAcqRel does atomic test-and-set operation on the bool *addr (atomically *addr = *addr ? *addr : true).
//extern github_com_ikavalio_atomicx.TestAndSetAcqRel
func TestAndSetAcqRel(addr *bool) bool


//TestAndSetSeqCst does atomic test-and-set operation on the bool *addr (atomically *addr = *addr ? *addr : true).
//extern github_com_ikavalio_atomicx.TestAndSetSeqCst
func TestAndSetSeqCst(addr *bool) bool


//ClearRelaxed does atomic clear operation on the bool *addr. Should be used in conjunction with TestAndSet.
//extern github_com_ikavalio_atomicx.ClearRelaxed
func ClearRelaxed(addr *bool)


//ClearRelease does atomic clear operation on the bool *addr. Should be used in conjunction with TestAndSet.
//extern github_com_ikavalio_atomicx.ClearRelease
func ClearRelease(addr *bool)


//ClearSeqCst does atomic clear operation on the bool *addr. Should be used in conjunction with TestAndSet.
//extern github_com_ikavalio_atomicx.ClearSeqCst
func ClearSeqCst(addr *bool)

