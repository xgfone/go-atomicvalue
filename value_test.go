// Copyright 2023 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package atomicvalue

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleValue() {
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

func TestValue(t *testing.T) {
	var value Value[int]
	if v := value.Load(); v != 0 {
		t.Errorf("expect %v, but got %v", 0, v)
	}

	value = Value[int]{}
	if v := value.Swap(123); v != 0 {
		t.Errorf("expect %v, but got %v", 0, v)
	}

	value = NewValue(123)
	if v := value.Swap(456); v != 123 {
		t.Errorf("expect %v, but got %v", 123, v)
	}
	if v := value.Load(); v != 456 {
		t.Errorf("expect %v, but got %v", 456, v)
	}

	if value.CompareAndSwap(123, 789) {
		t.Errorf("unexpect compare and swap, but fail")
	} else if v := value.Load(); v != 456 {
		t.Errorf("expect %v, but got %v", 456, v)
	}

	if !value.CompareAndSwap(456, 789) {
		t.Errorf("expect compare and swap, but fail")
	} else if v := value.Load(); v != 789 {
		t.Errorf("expect %v, but got %v", 789, v)
	}
}
