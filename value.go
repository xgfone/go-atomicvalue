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

// Package atomicvalue provides an atomic value equivalent to
// "sync/atomic".Value, but more lenient, which does not require
// that the type of the stored value is consistent and is suitable
// to store an interface with the different implementation.
// For example,
//
//	var errvalue atomicvalue.Value
//	errvalue.Store(errors.New("err1"))
//	errvalue.Store(fmt.Errorf("%w", errors.New("err2")))
package atomicvalue

import "sync/atomic"

type valueWrapper struct{ Value interface{} }

// Value is the same as sync/atomic.Value, but requires that the type
// of the stored value is consistent.
//
// Notice: the stored value is allowed to be nil or any interfaces.
type Value struct{ value atomic.Value }

// NewValue returns a new Value with the init value.
func NewValue(initValue interface{}) (v Value) {
	v.Store(initValue)
	return
}

// Load refers to sync/atomic.Value#Load.
func (v *Value) Load() interface{} {
	if _v := v.value.Load(); _v != nil {
		return _v.(valueWrapper).Value
	}
	return nil
}

// Store refers to sync/atomic.Value#Store.
func (v *Value) Store(val interface{}) {
	v.value.Store(valueWrapper{val})
}

// CompareAndSwap refers to sync/atomic.Value#CompareAndSwap.
func (v *Value) CompareAndSwap(old, new interface{}) (swapped bool) {
	return v.value.CompareAndSwap(valueWrapper{old}, valueWrapper{new})
}

// Swap refers to sync/atomic.Value#Swap.
func (v *Value) Swap(new interface{}) (old interface{}) {
	if _v := v.value.Swap(valueWrapper{new}); _v != nil {
		return _v.(valueWrapper).Value
	}
	return nil
}
