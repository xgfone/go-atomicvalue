# Go Atomic Value [![Build Status](https://github.com/xgfone/go-atomicvalue/actions/workflows/go.yml/badge.svg)](https://github.com/xgfone/go-atomicvalue/actions/workflows/go.yml) [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/go-atomicvalue)](https://pkg.go.dev/github.com/xgfone/go-atomicvalue) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-atomicvalue/master/LICENSE)

Provide an atomic value, supporting Go `1.18+`, same to `"sync/atomic".Value`, but more lenient, which does not require that the underlying type is consistent when storing an interface. So it is suitable to store an interface with the different implementation. For example,

```go
package atomicvalue

import (
	"errors"
	"fmt"
)

func main() {
	// We declare an atomic value to store the different errors.
	var errvalue Value[error]

	err1 := errors.New("err1")
	errvalue.Store(err1)
	fmt.Printf("%T\n", errvalue.Load())

	err2 := fmt.Errorf("%w", err1)
	errvalue.Store(err2)
	fmt.Printf("%T\n", errvalue.Load())

	// Output:
	// *errors.errorString
	// *fmt.wrapError
}
```
