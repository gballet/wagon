// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"math"
)

func (vm *VM) i32Wrapi64() error {
	vm.pushUint32(uint32(vm.popUint64()))
	return nil
}

func (vm *VM) i32TruncSF32() error {
	vm.pushInt32(int32(math.Trunc(float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) i32TruncUF32() error {
	vm.pushUint32(uint32(math.Trunc(float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) i32TruncSF64() error {
	vm.pushInt32(int32(math.Trunc(vm.popFloat64())))
	return nil
}

func (vm *VM) i32TruncUF64() error {
	vm.pushUint32(uint32(math.Trunc(vm.popFloat64())))
	return nil
}

func (vm *VM) i64ExtendSI32() error {
	vm.pushInt64(int64(vm.popInt32()))
	return nil
}

func (vm *VM) i64ExtendUI32() error {
	vm.pushUint64(uint64(vm.popUint32()))
	return nil
}

func (vm *VM) i64TruncSF32() error {
	vm.pushInt64(int64(math.Trunc(float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) i64TruncUF32() error {
	vm.pushUint64(uint64(math.Trunc(float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) i64TruncSF64() error {
	vm.pushInt64(int64(math.Trunc(vm.popFloat64())))
	return nil
}

func (vm *VM) i64TruncUF64() error {
	vm.pushUint64(uint64(math.Trunc(vm.popFloat64())))
	return nil
}

func (vm *VM) f32ConvertSI32() error {
	vm.pushFloat32(float32(vm.popInt32()))
	return nil
}

func (vm *VM) f32ConvertUI32() error {
	vm.pushFloat32(float32(vm.popUint32()))
	return nil
}

func (vm *VM) f32ConvertSI64() error {
	vm.pushFloat32(float32(vm.popInt64()))
	return nil
}

func (vm *VM) f32ConvertUI64() error {
	vm.pushFloat32(float32(vm.popUint64()))
	return nil
}

func (vm *VM) f32DemoteF64() error {
	vm.pushFloat32(float32(vm.popFloat64()))
	return nil
}

func (vm *VM) f64ConvertSI32() error {
	vm.pushFloat64(float64(vm.popInt32()))
	return nil
}

func (vm *VM) f64ConvertUI32() error {
	vm.pushFloat64(float64(vm.popUint32()))
	return nil
}

func (vm *VM) f64ConvertSI64() error {
	vm.pushFloat64(float64(vm.popInt64()))
	return nil
}

func (vm *VM) f64ConvertUI64() error {
	vm.pushFloat64(float64(vm.popUint64()))
	return nil
}

func (vm *VM) f64PromoteF32() error {
	vm.pushFloat64(float64(vm.popFloat32()))
	return nil
}
