// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vector

// Slice/vector arithmetic

// Plus returns a slice of the same size as a and b containing their elementwise sum.
// a and b must have the same length
func Plus[T Number](a, b []T) []T {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]T, len(a))
	for i := range a {
		r[i] = a[i] + b[i]
	}
	return r
}

// Minus returns a slice of the same size as a and b containing their elementwise difference.
// a and b must have the same length
func Minus[T Number](a, b []T) []T {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]T, len(a))
	for i := range a {
		r[i] = a[i] - b[i]
	}
	return r
}

// Times returns a slice of the same size as a and b containing their elementwise product.
// a and b must have the same length
func Times[T Number](a, b []T) []T {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]T, len(a))
	for i := range a {
		r[i] = a[i] * b[i]
	}
	return r
}

// Divide returns a slice of the same size as a and b containing their elementwise quotient.
// a and b must have the same length
func Divide[T Number](a, b []T) []T {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]T, len(a))
	for i := range a {
		r[i] = a[i] / b[i]
	}
	return r
}

// boolean ops

// And returns a slice of the same size as a and b containing their elementwise boolean AND.
// a and b must have the same length
func And(a, b []bool) []bool {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] && b[i]
	}
	return r
}

// Or returns a slice of the same size as a and b containing their elementwise boolean OR.
// a and b must have the same length
func Or(a, b []bool) []bool {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] || b[i]
	}
	return r
}

// Xor returns a slice of the same size as a and b containing their elementwise boolean XOR.
// a and b must have the same length
func Xor(a, b []bool) []bool {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] != b[i]
	}
	return r
}

// And returns a slice of the same size as a a's elementwise boolean NOT.
// a and b must have the same length
func Not(a []bool) []bool {
	r := make([]bool, len(a))
	for i := range a {
		r[i] = !a[i]
	}
	return r
}

// Fill returns a slice of length n with elements initialized to scalar b
func Fill[T Number](n int, b T) []T {
	r := make([]T, n)
	for i := range r {
		r[i] = b
	}
	return r
}

// Select returns a slice with the same length as mask with elements initialized to scalars ifTrue and ifFalse
// corresponding to the values in mask.
func Select[T Number](mask []bool, ifTrue, ifFalse T) []T {
	r := make([]T, len(mask))
	for i, b := range mask {
		if b {
			r[i] = ifTrue
		} else {
			r[i] = ifFalse
		}
	}
	return r
}

// Merge returns a slice with the same length as mask with elements initialized to
// elements from corresponding element of ifTrue and ifFalse, depending on the corresponding
// value in mask.
// mask, ifTrue, and ifFalse must all have the same length.
func Merge[T Number](mask []bool, ifTrue, ifFalse []T) []T {
	if len(ifTrue) != len(ifFalse) || len(ifTrue) != len(mask) {
		panic("vector length mismatch")
	}
	r := make([]T, len(mask))
	for i, b := range mask {
		if b {
			r[i] = ifTrue[i]
		} else {
			r[i] = ifFalse[i]
		}
	}
	return r
}

// Vector compares

// LT returns a slice of boolean, are the elements of a less than the elements of b?
func LT[T NumberOrdered](a, b []T) []bool {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] < b[i]
	}
	return r
}

// GT returns a slice of boolean, are the elements of a greater than the elements of b?
func GT[T NumberOrdered](a, b []T) []bool {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] > b[i]
	}
	return r
}

// LE returns a slice of boolean, are the elements of a less than or equal to the elements of b?
func LE[T NumberOrdered](a, b []T) []bool {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] <= b[i]
	}
	return r
}

// GE returns a slice of boolean, are the elements of a greater than or equal to the elements of b?
func GE[T NumberOrdered](a, b []T) []bool {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] >= b[i]
	}
	return r
}

// EQ returns a slice of boolean, are the elements of a equal to the elements of b?
func EQ[T Number](a, b []T) []bool {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] == b[i]
	}
	return r
}

// NE returns a slice of boolean, are the elements of a unequal to the elements of b?
func NE[T Number](a, b []T) []bool {
	if len(a) != len(b) {
		panic("vector length mismatch")
	}
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] != b[i]
	}
	return r
}

// Scalar operand comparison

// LTS returns a slice of boolean, are the elements of a less than the scalar b?
func LTS[T NumberOrdered](a []T, b T) []bool {
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] < b
	}
	return r
}

// GTS returns a slice of boolean, are the elements of a greater than the scalar b?
func GTS[T NumberOrdered](a []T, b T) []bool {
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] > b
	}
	return r
}

// LES returns a slice of boolean, are the elements of a less than or equal to the scalar b?
func LES[T NumberOrdered](a []T, b T) []bool {
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] <= b
	}
	return r
}

// GES returns a slice of boolean, are the elements of a greater than or equal to the scalar b?
func GES[T NumberOrdered](a []T, b T) []bool {
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] >= b
	}
	return r
}

// EQS returns a slice of boolean, are the elements of a equal to the scalar b?
func EQS[T Number](a []T, b T) []bool {
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] == b
	}
	return r
}

// NES returns a slice of boolean, are the elements of a unequal to the scalar b?
func NES[T Number](a []T, b T) []bool {
	r := make([]bool, len(a))
	for i := range a {
		r[i] = a[i] != b
	}
	return r
}

// Conversions

// bits returns 8, 16, 32, or 64, for T = [u]int{8,16,32,64}
func bits[T Int]() uint64 {
	c := uint64(0x2000100808)
	y := uint64(T(c))
	// return (y + y>>8 + y>>16 + y>>32) & 255
	return ((y * 0x101010001) >> 32) & 255
}

// signedMax returns the maximum signed value that can fit in a T, expressed as a U
func signedMax[T, U Int]() U {
	return U(1)<<(bits[T]()-1) - 1
}

// signedMin returns the minimum signed value that can fit in a T, expressed as a U
// if U is unsigned, the result is the corresponding positive value
func signedMin[T, U Int]() U {
	return -(U(1) << (bits[T]() - 1))
}

// unsignedMax returns the maximum unsigned value that can fit in a T, expressed as a U
// if T and U are the same size and U is unsigned, this will underflow.
func unsignedMax[T, U Int]() U {
	return U(1)<<(bits[T]()) - 1
}

// ClampS2U converts a slice of signed S into a slice of unsigned U,
// clamping values to what can be expressed in a U.
func ClampS2U[S, U Int](a []S) []U {
	r := make([]U, len(a))
	M := unsignedMax[U, S]()
	for i, x := range a {
		if x < 0 {
			x = 0
		} else if x > M {
			x = M
		}
		r[i] = U(x)
	}
	return r
}

// ClampU2S converts a slice of unsigned U into a slice of signed S,
// clamping values to what can be expressed in an S.
func ClampU2S[U, S Int](a []U) []S {
	r := make([]S, len(a))
	M := signedMax[S, U]()
	for i, x := range a {
		if x < 0 {
			x = 0
		} else if x > M {
			x = M
		}
		r[i] = S(x)
	}
	return r
}

// ClampS2S converts a slice of signed S into a slice of signed T,
// clamping values to what can be expressed in a T.
func ClampS2S[S, T Int](a []S) []T {
	r := make([]T, len(a))
	M := signedMax[T, S]()
	m := signedMin[T, S]()
	for i, x := range a {
		if x < m {
			x = m
		} else if x > M {
			x = M
		}
		r[i] = T(x)
	}
	return r
}

// Convert converts a slice of S into a slice of T,
// using the Go-defined conversion from type S to T.
func Convert[S, T Int](a []S) []T {
	r := make([]T, len(a))
	for i, x := range a {
		r[i] = T(x)
	}
	return r
}

// UnpackLE unpacks the contents of a slice of S into a slice of T,
// bitwise, using little-endian ordering.  S must not be smaller than T.
func UnpackLE[S, T Int](a []S) []T {
	sa, sb := bits[S](), bits[T]()
	if sa < sb {
		panic("Sizeof(sa) < sizeof(sb) [cannot unpack into a larger type]")
	}
	bPerA := int(sa / sb)
	b := make([]T, len(a)*bPerA)
	for i, x := range a {
		for k := 0; k < bPerA; k++ {
			b[bPerA*i+k] = T(x >> (k * int(sb)))
		}
	}
	return b
}

// PackLE packs the contents of a slice of S into a slice of T,
// bitwise, using little-endian ordering.  S must not be larger than T.
func PackLE[S, T Int](a []S) []T {
	sa, sb := bits[S](), bits[T]()
	mask := unsignedMax[S, uint64]()
	if sa > sb {
		panic("Sizeof(sa) > sizeof(sb) [cannot pack into a smaller type]")
	}
	aPerB := int(sb / sa)
	b := make([]T, (len(a)+aPerB-1)/aPerB)
	if len(a) == 0 {
		return b
	}
	var i int
	for i = 0; i < len(b)-1; i++ {
		var x T
		for k := 0; k < aPerB; k++ {
			x += T(mask & uint64(a[k+i*aPerB]) << (k * int(sa)))
		}
		b[i] = x
	}
	var x T
	for k := 0; k < len(a)-i*aPerB; k++ {
		x += T(mask & uint64(a[k+i*aPerB]) << (k * int(sa)))
	}
	b[i] = x
	return b
}
