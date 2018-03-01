// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"math"
)

// these operations are essentially no-ops.
// TODO(vibhavp): Add optimisations to package compiles that
// removes them from the original bytecode.

func (vm *VM) i32ReinterpretF32() error {
	vm.pushUint32(math.Float32bits(vm.popFloat32()))
	return nil
}

func (vm *VM) i64ReinterpretF64() error {
	vm.pushUint64(math.Float64bits(vm.popFloat64()))
	return nil
}

func (vm *VM) f32ReinterpretI32() error {
	vm.pushFloat32(math.Float32frombits(vm.popUint32()))
	return nil
}

func (vm *VM) f64ReinterpretI64() error {
	vm.pushFloat64(math.Float64frombits(vm.popUint64()))
	return nil
}
