atomicx
===

Experimetal package that provides `sync/atomic` compatible API (except Value type) with relaxed memory orders.

__It can only work on gccgo!__

All C++11 memory orders are supported: Relaxed, Consume, Acquire, Release, Acquire-Release and
Sequential Consistency. All CASes can be optionally weak.

Package calls compilers' built-in `__atomic_*` [functions](https://gcc.gnu.org/onlinedocs/gcc/_005f_005fatomic-Builtins.html) 
that replaced older `__sync_*` atomic functions. These functions approximately match the requirements for the C++11 memory model.

C compiler is required to build the package. GCC or clang under Linux/OS X that supports C++11/C11 should be enough.

__Important Notes__

- Currently pointers are not compared to `nil` before deferencing. Since they are dereferenced within unmanaged C code it can
cause segmentation faults, so please check your pointers manually before calling atomic functions.
- API consists of ~400 individual functions. Go file with C wrappers is created automatically by Python script (current solution).

__Available Functions__

Functions compatible with `sync/atomic` API:

```
// T is int32, int64, uint32, uint64, uintptr
// O can be any memory order
AddTO(addr *T, delta T) T

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
// O1 can be any memory order
// O2 can't be AcqRel or Release or stronger than O1
CompareAndSwapStrongTO1O2(addr *T, old, new T) bool
CompareAndSwapWeakTO1O2(addr *T, old, new T) bool

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
// O is Relaxed, Consume, Acquire or SeqCst
LoadTO(addr *T) T

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
// O is Relaxed, Release or SeqCst
StoreTO(addr *T) T

// T is int32, int64, uint32, uint64, uintptr, unsafe.Pointer
// O can't be Consume
SwapTO(addr *T, new T) T
```

Extensions:

```
// O can be any memory order
TestAndSetO(addr *bool) bool

// O can be Relaxed, Release or SeqCst
ClearO(addr *bool)

// T is int32, int64, uint32, uint64
// O can be any memory order
AndTO(addr *T, delta T) T

// T is int32, int64, uint32, uint64
// O can be any memory order
OrTO(addr *T, delta T) T

// T is int32, int64, uint32, uint64
// O can be any memory order
XorTO(addr *T, delta T) T

// T is int32, int64, uint32, uint64
// O can be any memory order
NandTO(addr *T, delta T) T
```

_Copyright (c) 2018 Ivan Kavaliou_

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
