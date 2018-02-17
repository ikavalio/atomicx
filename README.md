atomicx
===

Experimetal package that provides "sync/atomic" compatible API (except Value type) with configurable memory models.

All C++11 memory models are supported: Relaxed, Consume, Acquire, Release, Acquire-Release and
Sequential Consistency. All CASes can be optionally weak.

Package is implemented as `cgo` function wrappers over compilers' built-in `__atomic_*` functions (https://gcc.gnu.org/onlinedocs/gcc/_005f_005fatomic-Builtins.html) that replaced 
older `__sync_*` atomic functions. These functions approximately match the requirements for the C++11 memory model.

C compiler is required to build the package. GCC or clang under Linux/OS X that support C++11/C11 should be enough.

__Important Notes__

- Currently pointers are not compared to `nil` before deferencing. Since they are dereferenced within unmanaged C code it can
cause segmentation faults, so please check your pointers manually before calling atomic functions.   
- Atomic functions accept all memory orders as their arguments though only certain orders are valid for them. Passing incorrect
memory order will not break function's atomicity, but can cause undesired side-effects or undefined behavior. Please check 
documentation and make sure that correct memory orders are used in appropriate context.


_Copyright (c) 2018 Ivan Kavaliou_

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
