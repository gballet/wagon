// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

func (vm *VM) i32Const() error {
	vm.pushUint32(vm.fetchUint32())
	return nil
}

func (vm *VM) i64Const() error {
	vm.pushUint64(vm.fetchUint64())
	return nil
}

func (vm *VM) f32Const() error {
	vm.pushFloat32(vm.fetchFloat32())
	return nil
}

func (vm *VM) f64Const() error {
	vm.pushFloat64(vm.fetchFloat64())
	return nil
}
