// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vector

import (
	"fmt"
	"testing"
)

func TestClampS2S(t *testing.T) {
	i32s := []int32{-100000, -200000, -10, -20, 100000, 200000, 10, 20}
	want := []int8{-128, -128, -10, -20, 127, 127, 10, 20}
	got := ClampS2S[int32, int8](i32s)
	expectSlice(t, want, got)
}

func expect[T Int](t *testing.T, want, got T) {
	if want != got {
		t.Errorf("Wanted %d, got %d\n", want, got)
	}
}

func expectSlice[T Int](t *testing.T, want, got []T) {
	if len(want) != len(got) {
		t.Fatalf("Mismatched lengths, len(want)=%d, len(got)=%d, want=%v, got=%v", len(want), len(got), want, got)
	}
	for i, g := range got {
		if g != want[i] {
			t.Fatalf("Mismatch at index %d, want=%v, got=%v", i, want, got)
		}
	}
}

func TestMaxMin(t *testing.T) {
	expect(t, -128, signedMin[int8, int64]())
	expect(t, 127, signedMax[int8, int64]())
	expect(t, -32768, signedMin[int16, int64]())
	expect(t, 32767, signedMax[int16, int64]())
}

func TestArithmetic(t *testing.T) {
	x := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ones := Fill(len(x), int8(1))
	y := Times(Plus(x, ones), Minus(x, ones))
	z := Minus(Times(x, x), ones)
	want := []int8{0, 3, 8, 15, 24, 35, 48, 63, 80}
	expectSlice(t, want, y)
	expectSlice(t, want, z)
}

func TestBits(t *testing.T) {
	expect(t, 8, bits[int8]())
	expect(t, 8, bits[uint8]())
	expect(t, 16, bits[int16]())
	expect(t, 16, bits[uint16]())
	expect(t, 32, bits[int32]())
	expect(t, 32, bits[uint32]())
	expect(t, 64, bits[int64]())
	expect(t, 64, bits[uint64]())
}

func TestUnpackULE(t *testing.T) {
	u32s := []uint32{0x01, 0x8202, 0x038303, 0x84048404}
	wantu8 := []uint8{1, 0, 0, 0, 2, 0x82, 0, 0, 3, 0x83, 3, 0, 4, 0x84, 4, 0x84}
	wanti8 := []int8{1, 0, 0, 0, 2, -126, 0, 0, 3, -125, 3, 0, 4, -124, 4, -124}
	wanti16 := []int16{1, 0, -32254, 0, -31997, 3, -31740, -31740}
	gotu8 := UnpackLE[uint32, uint8](u32s)
	goti8 := UnpackLE[uint32, int8](u32s)
	goti16 := UnpackLE[uint32, int16](u32s)
	expectSlice(t, wantu8, gotu8)
	expectSlice(t, wanti8, goti8)
	expectSlice(t, wanti16, goti16)
}

func TestUnpackILE(t *testing.T) {
	i32s := []int32{0x01, 0x8202, 0x038303, -0x80000000 + 0x4048404}
	fmt.Printf("i32s=%v\n", i32s)
	wantu8 := []uint8{1, 0, 0, 0, 2, 0x82, 0, 0, 3, 0x83, 3, 0, 4, 0x84, 4, 0x84}
	wanti8 := []int8{1, 0, 0, 0, 2, -126, 0, 0, 3, -125, 3, 0, 4, -124, 4, -124}
	wanti16 := []int16{1, 0, -32254, 0, -31997, 3, -31740, -31740}
	gotu8 := UnpackLE[int32, uint8](i32s)
	goti8 := UnpackLE[int32, int8](i32s)
	goti16 := UnpackLE[int32, int16](i32s)
	expectSlice(t, wantu8, gotu8)
	expectSlice(t, wanti8, goti8)
	expectSlice(t, wanti16, goti16)
}
