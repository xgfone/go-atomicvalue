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

// Package atomicvalue provides an atomic value same to "sync/atomic".Value,
// but more lenient, which does not require that the underlying type
// is consistent when storing an interface. So it is suitable to store
// an interface with the different implementation. For example,
//
//	var errvalue atomicvalue.Value[error]
//	errvalue.Store(errors.New("err1"))
//	errvalue.Store(fmt.Errorf("%w", errors.New("err2")))
package atomicvalue

import "sync/atomic"

type valueWrapper[T any] struct{ Value T }

// Value is the same as sync/atomic.Value, but requires that the type
// of the stored value is consistent.
//
// Notice: the stored value is allowed to be nil or any interfaces.
type Value[T any] struct{ value atomic.Value }

// NewValue returns a new Value with the init value.
func NewValue[T any](initValue T) (v Value[T]) {
	v.Store(initValue)
	return
}

// Load refers to sync/atomic.Value#Load.
//
// NOTICE: return ZERO if it does not store any.
func (v *Value[T]) Load() (t T) {
	if _v := v.value.Load(); _v != nil {
		return _v.(valueWrapper[T]).Value
	}
	return
}

// Store refers to sync/atomic.Value#Store.
func (v *Value[T]) Store(val T) {
	v.value.Store(valueWrapper[T]{val})
}

// CompareAndSwap refers to sync/atomic.Value#CompareAndSwap.
func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.value.CompareAndSwap(valueWrapper[T]{old}, valueWrapper[T]{new})
}

// Swap refers to sync/atomic.Value#Swap.
//
// NOTICE: return ZERO if it does not store any.
func (v *Value[T]) Swap(new T) (old T) {
	if _v := v.value.Swap(valueWrapper[T]{new}); _v != nil {
		return _v.(valueWrapper[T]).Value
	}
	return
}
