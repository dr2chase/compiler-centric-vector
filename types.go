// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vector

type Signed interface {
	~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	Signed | Unsigned
}

type Number interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~complex64 | ~complex128
}

type NumberOrdered interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}
