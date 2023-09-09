//go:build !purego

package cast

import (
	"testing"
)

var benchList = []AString{"hello", "world", "list", "of", "strings"}

func Benchmark_Strings(b *testing.B) {
	var t []string

	u := benchList

	for i := 0; i < b.N; i++ {
		t = Strings(u)
	}

	assertEqual(b, u, t)
}

func Benchmark_cast(b *testing.B) {
	var t []string

	u := benchList

	for i := 0; i < b.N; i++ {
		t = cast[string, AString](u)
	}

	assertEqual(b, u, t)
}

func Benchmark_castStringsAppend(b *testing.B) {
	var t []string

	u := benchList

	for i := 0; i < b.N; i++ {
		t = castStringsAppend(u)
	}

	assertEqual(b, u, t)
}

func Benchmark_castStringsInsert(b *testing.B) {
	var t []string

	u := benchList

	for i := 0; i < b.N; i++ {
		t = castStringsInsert(u)
	}

	assertEqual(b, u, t)
}

func castStringsAppend[T ~string](tt []T) []string {
	if tt == nil {
		return nil
	}
	vv := make([]string, 0, len(tt))
	for _, value := range tt {
		vv = append(vv, string(value))
	}
	return vv
}

func castStringsInsert[T ~string](tt []T) []string {
	if tt == nil {
		return nil
	}
	vv := make([]string, len(tt))
	if len(tt) != len(vv) {
		panic("must be equal")
	}
	for _, value := range tt {
		vv = append(vv, string(value))
	}
	return vv
}

func assertEqual(tb testing.TB, u []AString, t []string) {
	tb.Helper()

	for i := range t {
		if string(u[i]) != t[i] {
			tb.Errorf("u[%d] = %v, want %v", i, u[i], t[i])
			return
		}
	}
}
