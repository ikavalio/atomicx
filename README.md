atomicx
===

Experimetal package that provides `sync/atomic` compatible API (except Value type) with configurable memory orders.

__It can only work on gccgo!__

All C++11 memory orders are supported: Relaxed, Consume, Acquire, Release, Acquire-Release and
Sequential Consistency. All CASes can be optionally weak.

Package calls compilers' built-in `__atomic_*` [functions](https://gcc.gnu.org/onlinedocs/gcc/_005f_005fatomic-Builtins.html) 
that replaced older `__sync_*` atomic functions. These functions approximately match the requirements for the C++11 memory model.

C compiler is required to build the package. GCC or clang under Linux/OS X that supports C++11/C11 should be enough.

__Important Notes__

- Currently pointers are not compared to `nil` before deferencing. Since they are dereferenced within unmanaged C code it can
cause segmentation faults, so please check your pointers manually before calling atomic functions.   
- Atomic functions accept all memory orders as their arguments though only certain orders are valid for them. Passing incorrect
memory order will not break function's atomicity, but can cause undesired side-effects or undefined behavior. Please check 
documentation and make sure that correct memory orders are used in appropriate context.

__Available Functions__

Functions compatible with `sync/atomic` API:

```
// T is int32, int64, uint32, uint64, uintptr
AddT(addr *T, delta T, order MemoryOrder) T

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
CompareAndSwapStrongT(addr *T, old, new T, order MemoryOrder) bool

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
LoadT(addr *T, order MemoryOrder) T

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
StoreT(addr *T, order MemoryOrder) T

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
SwapT(addr *T, new T, order MemoryOrder) T
```

Extensions:

```
// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
CompareAndSwapStrong2T(addr *T, old, new T, orderSuccess MemoryOrder, orderFailure MemoryOrder) bool

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
CompareAndSwapWeakT(addr *T, old, new T, order MemoryOrder) bool

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
CompareAndSwapWeak2T(addr *T, old, new T, orderSuccess MemoryOrder, orderFailure MemoryOrder) bool

TestAndSet(addr *bool, order MemoryOrder) bool

Clear(addr *bool, order MemoryOrder)

// T is int32, int64, uint32, uint64
AndT(addr *T, delta T, order MemoryOrder) T

// T is int32, int64, uint32, uint64
OrT(addr *T, delta T, order MemoryOrder) T

// T is int32, int64, uint32, uint64
XorT(addr *T, delta T, order MemoryOrder) T

// T is int32, int64, uint32, uint64
NandT(addr *T, delta T, order MemoryOrder) T
```

_Copyright (c) 2018 Ivan Kavaliou_

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
