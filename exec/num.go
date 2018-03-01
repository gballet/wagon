// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"math"
	"math/bits"
)

// int32 operators

func (vm *VM) i32Clz() error {
	vm.pushUint64(uint64(bits.LeadingZeros32(vm.popUint32())))
	return nil
}

func (vm *VM) i32Ctz() error {
	vm.pushUint64(uint64(bits.TrailingZeros32(vm.popUint32())))
	return nil
}

func (vm *VM) i32Popcnt() error {
	vm.pushUint64(uint64(bits.OnesCount32(vm.popUint32())))
	return nil
}

func (vm *VM) i32Add() error {
	vm.pushUint32(vm.popUint32() + vm.popUint32())
	return nil
}

func (vm *VM) i32Mul() error {
	vm.pushUint32(vm.popUint32() * vm.popUint32())
	return nil
}

func (vm *VM) i32DivS() error {
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	vm.pushInt32(v1 / v2)
	return nil
}

func (vm *VM) i32DivU() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushUint32(v1 / v2)
	return nil
}

func (vm *VM) i32RemS() error {
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	vm.pushInt32(v1 % v2)
	return nil
}

func (vm *VM) i32RemU() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushUint32(v1 % v2)
	return nil
}

func (vm *VM) i32Sub() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushUint32(v1 - v2)
	return nil
}

func (vm *VM) i32And() error {
	vm.pushUint32(vm.popUint32() & vm.popUint32())
	return nil
}

func (vm *VM) i32Or() error {
	vm.pushUint32(vm.popUint32() | vm.popUint32())
	return nil
}

func (vm *VM) i32Xor() error {
	vm.pushUint32(vm.popUint32() ^ vm.popUint32())
	return nil
}

func (vm *VM) i32Shl() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushUint32(v1 << v2)
	return nil
}

func (vm *VM) i32ShrU() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushUint32(v1 >> v2)
	return nil
}

func (vm *VM) i32ShrS() error {
	v2 := vm.popUint32()
	v1 := vm.popInt32()
	vm.pushInt32(v1 >> v2)
	return nil
}

func (vm *VM) i32Rotl() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushUint32(bits.RotateLeft32(v1, int(v2)))
	return nil
}

func (vm *VM) i32Rotr() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushUint32(bits.RotateLeft32(v1, -int(v2)))
	return nil
}

func (vm *VM) i32LeS() error {
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	vm.pushBool(v1 <= v2)
	return nil
}

func (vm *VM) i32LeU() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushBool(v1 <= v2)
	return nil
}

func (vm *VM) i32LtS() error {
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	vm.pushBool(v1 < v2)
	return nil
}

func (vm *VM) i32LtU() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushBool(v1 < v2)
	return nil
}

func (vm *VM) i32GtS() error {
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	vm.pushBool(v1 > v2)
	return nil
}

func (vm *VM) i32GeS() error {
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	vm.pushBool(v1 >= v2)
	return nil
}

func (vm *VM) i32GtU() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushBool(v1 > v2)
	return nil
}

func (vm *VM) i32GeU() error {
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	vm.pushBool(v1 >= v2)
	return nil
}

func (vm *VM) i32Eqz() error {
	vm.pushBool(vm.popUint32() == 0)
	return nil
}

func (vm *VM) i32Eq() error {
	vm.pushBool(vm.popUint32() == vm.popUint32())
	return nil
}

func (vm *VM) i32Ne() error {
	vm.pushBool(vm.popUint32() != vm.popUint32())
	return nil
}

// int64 operators

func (vm *VM) i64Clz() error {
	vm.pushUint64(uint64(bits.LeadingZeros64(vm.popUint64())))
	return nil
}

func (vm *VM) i64Ctz() error {
	vm.pushUint64(uint64(bits.TrailingZeros64(vm.popUint64())))
	return nil
}

func (vm *VM) i64Popcnt() error {
	vm.pushUint64(uint64(bits.OnesCount64(vm.popUint64())))
	return nil
}

func (vm *VM) i64Add() error {
	vm.pushUint64(vm.popUint64() + vm.popUint64())
	return nil
}

func (vm *VM) i64Sub() error {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 - v2)
	return nil
}

func (vm *VM) i64Mul() error {
	vm.pushUint64(vm.popUint64() * vm.popUint64())
	return nil
}

func (vm *VM) i64DivS() error {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushInt64(v1 / v2)
	return nil
}

func (vm *VM) i64DivU() error {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 / v2)
	return nil
}

func (vm *VM) i64RemS() error {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushInt64(v1 % v2)
	return nil
}

func (vm *VM) i64RemU() error {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 % v2)
	return nil
}

func (vm *VM) i64And() error {
	vm.pushUint64(vm.popUint64() & vm.popUint64())
	return nil
}

func (vm *VM) i64Or() error {
	vm.pushUint64(vm.popUint64() | vm.popUint64())
	return nil
}

func (vm *VM) i64Xor() error {
	vm.pushUint64(vm.popUint64() ^ vm.popUint64())
	return nil
}

func (vm *VM) i64Shl() error {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 << v2)
	return nil
}

func (vm *VM) i64ShrS() error {
	v2 := vm.popUint64()
	v1 := vm.popInt64()
	vm.pushInt64(v1 >> v2)
	return nil
}

func (vm *VM) i64ShrU() error {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 >> v2)
	return nil
}

func (vm *VM) i64Rotl() error {
	v2 := vm.popInt64()
	v1 := vm.popUint64()
	vm.pushUint64(bits.RotateLeft64(v1, int(v2)))
	return nil
}

func (vm *VM) i64Rotr() error {
	v2 := vm.popInt64()
	v1 := vm.popUint64()
	vm.pushUint64(bits.RotateLeft64(v1, -int(v2)))
	return nil
}

func (vm *VM) i64Eq() error {
	vm.pushBool(vm.popUint64() == vm.popUint64())
	return nil
}

func (vm *VM) i64Eqz() error {
	vm.pushBool(vm.popUint64() == 0)
	return nil
}

func (vm *VM) i64Ne() error {
	vm.pushBool(vm.popUint64() != vm.popUint64())
	return nil
}

func (vm *VM) i64LtS() error {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushBool(v1 < v2)
	return nil
}

func (vm *VM) i64LtU() error {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushBool(v1 < v2)
	return nil
}

func (vm *VM) i64GtS() error {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushBool(v1 > v2)
	return nil
}

func (vm *VM) i64GtU() error {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushBool(v1 > v2)
	return nil
}

func (vm *VM) i64LeU() error {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushBool(v1 <= v2)
	return nil
}

func (vm *VM) i64LeS() error {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushBool(v1 <= v2)
	return nil
}

func (vm *VM) i64GeS() error {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushBool(v1 >= v2)
	return nil
}

func (vm *VM) i64GeU() error {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushBool(v1 >= v2)
	return nil
}

// float32 operators

func (vm *VM) f32Abs() error {
	vm.pushFloat32(float32(math.Abs(float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) f32Neg() error {
	vm.pushFloat32(-vm.popFloat32())
	return nil
}

func (vm *VM) f32Ceil() error {
	vm.pushFloat32(float32(math.Ceil(float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) f32Floor() error {
	vm.pushFloat32(float32(math.Floor(float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) f32Trunc() error {
	vm.pushFloat32(float32(math.Trunc(float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) f32Nearest() error {
	f := vm.popFloat32()
	vm.pushFloat32(float32(int32(f + float32(math.Copysign(0.5, float64(f))))))
	return nil
}

func (vm *VM) f32Sqrt() error {
	vm.pushFloat32(float32(math.Sqrt(float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) f32Add() error {
	vm.pushFloat32(vm.popFloat32() + vm.popFloat32())
	return nil
}

func (vm *VM) f32Sub() error {
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	vm.pushFloat32(v1 - v2)
	return nil
}

func (vm *VM) f32Mul() error {
	vm.pushFloat32(vm.popFloat32() * vm.popFloat32())
	return nil
}

func (vm *VM) f32Div() error {
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	vm.pushFloat32(v1 / v2)
	return nil
}

func (vm *VM) f32Min() error {
	vm.pushFloat32(float32(math.Min(float64(vm.popFloat32()), float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) f32Max() error {
	vm.pushFloat32(float32(math.Max(float64(vm.popFloat32()), float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) f32Copysign() error {
	vm.pushFloat32(float32(math.Copysign(float64(vm.popFloat32()), float64(vm.popFloat32()))))
	return nil
}

func (vm *VM) f32Eq() error {
	vm.pushBool(vm.popFloat32() == vm.popFloat32())
	return nil
}

func (vm *VM) f32Ne() error {
	vm.pushBool(vm.popFloat32() != vm.popFloat32())
	return nil
}

func (vm *VM) f32Lt() error {
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	vm.pushBool(v1 < v2)
	return nil
}

func (vm *VM) f32Gt() error {
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	vm.pushBool(v1 > v2)
	return nil
}

func (vm *VM) f32Le() error {
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	vm.pushBool(v1 <= v2)
	return nil
}

func (vm *VM) f32Ge() error {
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	vm.pushBool(v1 >= v2)
	return nil
}

// float64 operators

func (vm *VM) f64Abs() error {
	vm.pushFloat64(math.Abs(vm.popFloat64()))
	return nil
}

func (vm *VM) f64Neg() error {
	vm.pushFloat64(-vm.popFloat64())
	return nil
}

func (vm *VM) f64Ceil() error {
	vm.pushFloat64(math.Ceil(vm.popFloat64()))
	return nil
}

func (vm *VM) f64Floor() error {
	vm.pushFloat64(math.Floor(vm.popFloat64()))
	return nil
}

func (vm *VM) f64Trunc() error {
	vm.pushFloat64(math.Trunc(vm.popFloat64()))
	return nil
}

func (vm *VM) f64Nearest() error {
	f := vm.popFloat64()
	vm.pushFloat64(float64(int64(f + math.Copysign(0.5, f))))
	return nil
}

func (vm *VM) f64Sqrt() error {
	vm.pushFloat64(math.Sqrt(vm.popFloat64()))
	return nil
}

func (vm *VM) f64Add() error {
	vm.pushFloat64(vm.popFloat64() + vm.popFloat64())
	return nil
}

func (vm *VM) f64Sub() error {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushFloat64(v1 - v2)
	return nil
}

func (vm *VM) f64Mul() error {
	vm.pushFloat64(vm.popFloat64() * vm.popFloat64())
	return nil
}

func (vm *VM) f64Div() error {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushFloat64(v1 / v2)
	return nil
}

func (vm *VM) f64Min() error {
	vm.pushFloat64(math.Min(vm.popFloat64(), vm.popFloat64()))
	return nil
}

func (vm *VM) f64Max() error {
	vm.pushFloat64(math.Max(vm.popFloat64(), vm.popFloat64()))
	return nil
}

func (vm *VM) f64Copysign() error {
	vm.pushFloat64(math.Copysign(vm.popFloat64(), vm.popFloat64()))
	return nil
}

func (vm *VM) f64Eq() error {
	vm.pushBool(vm.popFloat64() == vm.popFloat64())
	return nil
}

func (vm *VM) f64Ne() error {
	vm.pushBool(vm.popFloat64() != vm.popFloat64())
	return nil
}

func (vm *VM) f64Lt() error {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushBool(v1 < v2)
	return nil
}

func (vm *VM) f64Gt() error {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushBool(v1 > v2)
	return nil
}

func (vm *VM) f64Le() error {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushBool(v1 <= v2)
	return nil
}

func (vm *VM) f64Ge() error {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushBool(v1 >= v2)
	return nil
}
