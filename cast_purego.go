//go:build purego

package cast

func Strings[T ~string](tt []T) []string {
	if tt == nil {
		return nil
	}
	vv := make([]string, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, string(value))
	}
	return vv
}

func Ints[T ~int](tt []T) []int {
	if tt == nil {
		return nil
	}
	vv := make([]int, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, int(value))
	}
	return vv
}

func Int8s[T ~int8](tt []T) []int8 {
	if tt == nil {
		return nil
	}
	vv := make([]int8, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, int8(value))
	}
	return vv
}

func Int16s[T ~int16](tt []T) []int16 {
	if tt == nil {
		return nil
	}
	vv := make([]int16, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, int16(value))
	}
	return vv
}

func Int32s[T ~int32](tt []T) []int32 {
	if tt == nil {
		return nil
	}
	vv := make([]int32, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, int32(value))
	}
	return vv
}

func Int64s[T ~int64](tt []T) []int64 {
	if tt == nil {
		return nil
	}
	vv := make([]int64, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, int64(value))
	}
	return vv
}

func Uints[T ~uint](tt []T) []uint {
	if tt == nil {
		return nil
	}
	vv := make([]uint, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, uint(value))
	}
	return vv
}

func Uint8s[T ~uint8](tt []T) []uint8 {
	if tt == nil {
		return nil
	}
	vv := make([]uint8, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, uint8(value))
	}
	return vv
}

func Uint16s[T ~uint16](tt []T) []uint16 {
	if tt == nil {
		return nil
	}
	vv := make([]uint16, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, uint16(value))
	}
	return vv
}

func Uint32s[T ~uint32](tt []T) []uint32 {
	if tt == nil {
		return nil
	}
	vv := make([]uint32, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, uint32(value))
	}
	return vv
}

func Uint64s[T ~uint64](tt []T) []uint64 {
	if tt == nil {
		return nil
	}
	vv := make([]uint64, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, uint64(value))
	}
	return vv
}

func Bytes[T ~byte](tt []T) []byte {
	if tt == nil {
		return nil
	}
	vv := make([]byte, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, byte(value))
	}
	return vv
}

func Float32s[T ~float32](tt []T) []float32 {
	if tt == nil {
		return nil
	}
	vv := make([]float32, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, float32(value))
	}
	return vv
}

func Float64s[T ~float64](tt []T) []float64 {
	if tt == nil {
		return nil
	}
	vv := make([]float64, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, float64(value))
	}
	return vv
}

func Complex64s[T ~complex64](tt []T) []complex64 {
	if tt == nil {
		return nil
	}
	vv := make([]complex64, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, complex64(value))
	}
	return vv
}

func Complex128s[T ~complex128](tt []T) []complex128 {
	if tt == nil {
		return nil
	}
	vv := make([]complex128, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, complex128(value))
	}
	return vv
}

func Bools[T ~bool](tt []T) []bool {
	if tt == nil {
		return nil
	}
	vv := make([]bool, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, bool(value))
	}
	return vv
}
