package cast

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	AString string

	AInt   int
	AInt8  int8
	AInt16 int16
	AInt32 int32
	AInt64 int64

	AUint   uint
	AUint8  uint8
	AUint16 uint16
	AUint32 uint32
	AUint64 uint64

	AByte byte

	AFloat32 float32
	AFloat64 float64

	AComplex64  complex64
	AComplex128 complex128

	ABool bool
)

func ExampleStrings() {
	type MyString string

	show := func(ss []string) {
		fmt.Println(ss)
	}

	mm := []MyString{"hello", "world", "list", "of", "strings"}
	ss := Strings(mm)

	show(ss)

	// Output: [hello world list of strings]
}

func TestCast(t *testing.T) {
	t.Parallel()
	type args struct {
		u any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "nil string",
			args: args{u: []AString(nil)},
			want: ([]string)(nil),
		},
		{
			name: "string",
			args: args{u: []AString{"hello", "world", "list", "of", "strings"}},
			want: []string{"hello", "world", "list", "of", "strings"},
		},
		{
			name: "nil int",
			args: args{u: ([]AInt)(nil)},
			want: ([]int)(nil),
		},
		{
			name: "int",
			args: args{u: []AInt{1, 2, 3, 4, 5, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil int8",
			args: args{u: ([]AInt8)(nil)},
			want: ([]int8)(nil),
		},
		{
			name: "int8",
			args: args{u: []AInt8{1, 2, 3, 4, 5, 6}},
			want: []int8{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil int16",
			args: args{u: ([]AInt16)(nil)},
			want: ([]int16)(nil),
		},
		{
			name: "int16",
			args: args{u: []AInt16{1, 2, 3, 4, 5, 6}},
			want: []int16{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil int32",
			args: args{u: ([]AInt32)(nil)},
			want: ([]int32)(nil),
		},
		{
			name: "int32",
			args: args{u: []AInt32{1, 2, 3, 4, 5, 6}},
			want: []int32{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil int64",
			args: args{u: ([]AInt64)(nil)},
			want: ([]int64)(nil),
		},
		{
			name: "int64",
			args: args{u: []AInt64{1, 2, 3, 4, 5, 6}},
			want: []int64{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil uint",
			args: args{u: ([]AUint)(nil)},
			want: ([]uint)(nil),
		},
		{
			name: "uint",
			args: args{u: []AUint{1, 2, 3, 4, 5, 6}},
			want: []uint{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil uint8",
			args: args{u: ([]AUint8)(nil)},
			want: ([]uint8)(nil),
		},
		{
			name: "uint8",
			args: args{u: []AUint8{1, 2, 3, 4, 5, 6}},
			want: []uint8{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil uint16",
			args: args{u: ([]AUint16)(nil)},
			want: ([]uint16)(nil),
		},
		{
			name: "uint16",
			args: args{u: []AUint16{1, 2, 3, 4, 5, 6}},
			want: []uint16{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil uint32",
			args: args{u: ([]AUint32)(nil)},
			want: ([]uint32)(nil),
		},
		{
			name: "uint32",
			args: args{u: []AUint32{1, 2, 3, 4, 5, 6}},
			want: []uint32{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil uint64",
			args: args{u: ([]AUint64)(nil)},
			want: ([]uint64)(nil),
		},
		{
			name: "uint64",
			args: args{u: []AUint64{1, 2, 3, 4, 5, 6}},
			want: []uint64{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil byte",
			args: args{u: ([]AByte)(nil)},
			want: ([]byte)(nil),
		},
		{
			name: "byte",
			args: args{u: []AByte{1, 2, 3, 4, 5, 6}},
			want: []byte{1, 2, 3, 4, 5, 6},
		},
		{
			name: "nil float32",
			args: args{u: ([]AFloat32)(nil)},
			want: ([]float32)(nil),
		},
		{
			name: "float32",
			args: args{u: []AFloat32{1.6, 2.5, 3.4, 4.3, 5.2, 6.1}},
			want: []float32{1.6, 2.5, 3.4, 4.3, 5.2, 6.1},
		},
		{
			name: "nil float64",
			args: args{u: ([]AFloat64)(nil)},
			want: ([]float64)(nil),
		},
		{
			name: "float64",
			args: args{u: []AFloat64{1.6, 2.5, 3.4, 4.3, 5.2, 6.1}},
			want: []float64{1.6, 2.5, 3.4, 4.3, 5.2, 6.1},
		},
		{
			name: "nil complex64",
			args: args{u: ([]AComplex64)(nil)},
			want: ([]complex64)(nil),
		},
		{
			name: "complex64",
			args: args{u: []AComplex64{complex(1, 6), complex(2, 5), complex(3, 4), complex(4, 3), complex(5, 2), complex(6, 2)}},
			want: []complex64{complex(1, 6), complex(2, 5), complex(3, 4), complex(4, 3), complex(5, 2), complex(6, 2)},
		},
		{
			name: "nil complex128",
			args: args{u: ([]AComplex128)(nil)},
			want: ([]complex128)(nil),
		},
		{
			name: "complex128",
			args: args{u: []AComplex128{complex(1, 6), complex(2, 5), complex(3, 4), complex(4, 3), complex(5, 2), complex(6, 2)}},
			want: []complex128{complex(1, 6), complex(2, 5), complex(3, 4), complex(4, 3), complex(5, 2), complex(6, 2)},
		},
		{
			name: "nil bool",
			args: args{u: ([]ABool)(nil)},
			want: ([]bool)(nil),
		},
		{
			name: "bool",
			args: args{u: []ABool{true, false, true, false}},
			want: []bool{true, false, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got any

			switch u := tt.args.u.(type) {
			case []AString:
				got = Strings(u)
			case []AInt:
				got = Ints(u)
			case []AInt8:
				got = Int8s(u)
			case []AInt16:
				got = Int16s(u)
			case []AInt32:
				got = Int32s(u)
			case []AInt64:
				got = Int64s(u)
			case []AUint:
				got = Uints(u)
			case []AUint8:
				got = Uint8s(u)
			case []AUint16:
				got = Uint16s(u)
			case []AUint32:
				got = Uint32s(u)
			case []AUint64:
				got = Uint64s(u)
			case []AByte:
				got = Bytes(u)
			case []AFloat32:
				got = Float32s(u)
			case []AFloat64:
				got = Float64s(u)
			case []AComplex64:
				got = Complex64s(u)
			case []AComplex128:
				got = Complex128s(u)
			case []ABool:
				got = Bools(u)
			default:
				t.Fatalf("unknown type")
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
