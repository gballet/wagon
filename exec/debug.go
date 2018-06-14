// Copyright 2018 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"container/list"
	"fmt"
	"strings"
)

type breakpoint struct {
	location int64
	count    int
}

// TODO feedback channel
var debuggerCommands = make(chan string)
var breakpoints *list.List

const maxDepth = 8

func InitDebugger() {
	breakpoints = list.New()
}

func isBreakpoint(pc int64) bool {
	for e := breakpoints.Front(); e != nil; e = e.Next() {
		bp, ok := e.Value.(*breakpoint)
		if !ok {
			panic("debugger: Invalid breakpoint!")
		}
		if bp.location == pc {
			fmt.Println(fmt.Sprintf("Hit breakpoint at %#x", pc))
			return true
		}
	}

	return false
}

func usage() {
	fmt.Println("List of debugger commands:")
	fmt.Println("h\tPrint this help message")
	fmt.Println("c\tContinue execution")
	fmt.Println("b [addr]\tSet a new breakpoint at address")
	fmt.Println("s\tDisplay the content stack")
	fmt.Println("d\tDisassemble at current location")
	fmt.Println("l\tShow locals")
}

func ConnectToDebugger() chan string {
	return debuggerCommands
}

func waitForDebuggerCommand(vm *VM) {
	for {
		stopMsg := fmt.Sprintf("Stopped at %#x", vm.ctx.pc)
		debuggerCommands <- stopMsg

		line := <-debuggerCommands
		if line == "" {
			continue
		}
		args := strings.Split(line, " ")
		command := args[0]

		switch command {
		case "h":
			usage()
		case "c":
			break
		case "b":
			if len(args) == 2 {
				var addr int64
				n, err := fmt.Sscanf(args[0], "%#x", &addr)
				if n != len(args[0]) || err != nil {
					fmt.Printf("Invalid hex address %s: %v", args[0], err)
				}

				breakpoints.PushFront(&breakpoint{location: addr, count: 1})
			} else if len(args) == 1 {
				breakpoints.PushFront(&breakpoint{location: vm.ctx.pc, count: 1})
			} else {
				fmt.Println("Invalid arguments")
				usage()
			}
		case "l":
			for i, local := range vm.ctx.locals {
				fmt.Printf("local[%d] = %#x", i, local)
			}
		case "s":
			for i := len(vm.ctx.stack) - 1; i > 0 && i > len(vm.ctx.stack)-maxDepth; i-- {
				fmt.Printf("%#x:\t%#x", i, vm.ctx.stack[i])
			}
		default:
			fmt.Println("Not implemented")
		}
	}
}
