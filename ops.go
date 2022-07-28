// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vector

// Slice/vector arithmetic

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

func Not(a []bool) []bool {
	r := make([]bool, len(a))
	for i := range a {
		r[i] = !a[i]
	}
	return r
}

// scalar fill
func Fill[T Number](n int, b T) []T {
	r := make([]T, n)
	for i := range r {
		r[i] = b
	}
	return r
}

// mask fill
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

// mask merge
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

func Truncate[S, T Int](a []S) []T {
	r := make([]T, len(a))
	for i, x := range a {
		r[i] = T(x)
	}
	return r
}

func UnpackLE[S, T Int](a []S) []T {
	sa, sb := bits[S](), bits[T]()
	if sa < sb {
		panic("Sizeof(sa) < sizeof(sb) [unpack into smaller type]")
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
