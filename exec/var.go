// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

func (vm *VM) getLocal() error {
	index := vm.fetchUint32()
	vm.pushUint64(vm.ctx.locals[int(index)])
	return nil
}

func (vm *VM) setLocal() error {
	index := vm.fetchUint32()
	vm.ctx.locals[int(index)] = vm.popUint64()
	return nil
}

func (vm *VM) teeLocal() error {
	index := vm.fetchUint32()
	val := vm.ctx.stack[len(vm.ctx.stack)-1]
	vm.ctx.locals[int(index)] = val
	return nil
}

func (vm *VM) getGlobal() error {
	index := vm.fetchUint32()
	vm.pushUint64(vm.globals[int(index)])
	return nil
}

func (vm *VM) setGlobal() error {
	index := vm.fetchUint32()
	vm.globals[int(index)] = vm.popUint64()
	return nil
}
