//go:build !purego

package cast

import "unsafe"

func Strings[T ~string](tt []T) []string { return cast[string, T](tt) }

func Ints[T ~int](tt []T) []int       { return cast[int, T](tt) }
func Int8s[T ~int8](tt []T) []int8    { return cast[int8, T](tt) }
func Int16s[T ~int16](tt []T) []int16 { return cast[int16, T](tt) }
func Int32s[T ~int32](tt []T) []int32 { return cast[int32, T](tt) }
func Int64s[T ~int64](tt []T) []int64 { return cast[int64, T](tt) }

func Uints[T ~uint](tt []T) []uint       { return cast[uint, T](tt) }
func Uint8s[T ~uint8](tt []T) []uint8    { return cast[uint8, T](tt) }
func Uint16s[T ~uint16](tt []T) []uint16 { return cast[uint16, T](tt) }
func Uint32s[T ~uint32](tt []T) []uint32 { return cast[uint32, T](tt) }
func Uint64s[T ~uint64](tt []T) []uint64 { return cast[uint64, T](tt) }

func Bytes[T ~byte](tt []T) []byte { return cast[byte, T](tt) }

func Float32s[T ~float32](tt []T) []float32 { return cast[float32, T](tt) }
func Float64s[T ~float64](tt []T) []float64 { return cast[float64, T](tt) }

func Complex64s[T ~complex64](tt []T) []complex64    { return cast[complex64, T](tt) }
func Complex128s[T ~complex128](tt []T) []complex128 { return cast[complex128, T](tt) }

func Bools[T ~bool](tt []T) []bool { return cast[bool, T](tt) }

func cast[T any, U any](uu []U) []T {
	if uu == nil {
		return nil
	}
	if cap(uu) == 0 {
		return []T{}
	}
	return unsafe.Slice((*T)(unsafe.Pointer(&uu[:cap(uu)][0])), cap(uu))[:len(uu)]
}
