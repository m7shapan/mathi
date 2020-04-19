# mathi
mathi going to be a wrapper for golang [math](https://golang.org/pkg/math) package but it going to accept or return int64 instead of float64

## how to use 

```go
package main

import (
	"fmt"

	"github.com/m7shapan/mathi"
)

func main() {
	fmt.Println(mathi.RoundN(5.12355, 3)) // 5.124
}

```

## supported math func 
- [x] Abs
- [x] Ceil
- [x] Copysign
- [x] Dim
- [x] Floor
- [x] Max
- [x] Min
- [x] Pow
- [x] Pow10
- [x] Round
- [x] RoundToEven

## new mathi func
- [x] RoundN
