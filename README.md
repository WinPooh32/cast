# Cast

![test](https://github.com/WinPooh32/cast/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/WinPooh32/cast.svg)](https://pkg.go.dev/github.com/WinPooh32/cast)

Set of generic helpers for converting slices to their underlying type.

It's fast and allocs free when `purego` build tag is disabled.

<details> 
<summary>Supported types</summary>

```text
string
uint8
uint16
uint32
uint64
int8
int16
int32
int64
float32
float64
complex64
complex128
int
uint
```

</details>

## Example

```Go
package main

import (
	"fmt"

	"github.com/WinPooh32/cast"
)

type MyString string

func show(ss []string) {
	fmt.Println(ss)
}

func main() {
	mm := []MyString{"hello", "world", "list", "of", "strings"}

	ss := cast.Strings(mm)

	show(ss)
}
```

Output:

```text
[hello world list of strings]
```

## Benchmarks

```text
goos: linux
goarch: amd64
pkg: github.com/WinPooh32/cast
cpu: AMD Ryzen 7 3700X 8-Core Processor             
Benchmark_Strings-16            559789944  2.183 ns/op   0 B/op   0 allocs/op
Benchmark_castStringsAppend-16  20688170   57.98 ns/op   80 B/op  1 allocs/op
PASS
ok      github.com/WinPooh32/cast       2.700s
```

castStringsAppend implementation:

```Go
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
```
