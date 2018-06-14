// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-interpreter/wagon/exec"
	"github.com/go-interpreter/wagon/validate"
	"github.com/go-interpreter/wagon/wasm"
)

func main() {
	log.SetPrefix("wasm-run: ")
	log.SetFlags(0)

	verbose := flag.Bool("v", false, "enable/disable verbose mode")
	verify := flag.Bool("verify-module", false, "run module verification")
	debugger := flag.Bool("d", false, "debug mode")

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		flag.PrintDefaults()
		os.Exit(1)
	}

	wasm.SetVerboseMode(*verbose)

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m, err := wasm.ReadModule(f, importer)
	if err != nil {
		log.Fatalf("could not read module: %v", err)
	}

	if *verify {
		err = validate.VerifyModule(m)
		if err != nil {
			log.Fatalf("could not verify module: %v", err)
		}
	}

	if m.Export == nil {
		log.Fatalf("module has no export section")
	}

	if *debugger {
		exec.InitDebugger()
		commandChannel := exec.ConnectToDebugger()
		input := bufio.NewReader(os.Stdin)

		// Stop at the start of the module
		commandChannel <- "b 0x0"

		go func() {
			for {
				msg := <-commandChannel
				fmt.Println(msg)
				fmt.Print("> ")
				txt, err := input.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading command")
					commandChannel <- "h"
				} else {
					command := strings.TrimRight(txt, "\n\r")
					commandChannel <- command
				}
			}
		}()
	}

	vm, err := exec.NewVM(m)
	if err != nil {
		log.Fatalf("could not create VM: %v", err)
	}

	for name, e := range m.Export.Entries {
		i := int64(e.Index)
		fidx := m.Function.Types[int(i)]
		ftype := m.Types.Entries[int(fidx)]
		switch len(ftype.ReturnTypes) {
		case 1:
			fmt.Printf("%s() %s => ", name, ftype.ReturnTypes[0])
		case 0:
			fmt.Printf("%s() => ", name)
		default:
			log.Printf("running exported functions with more than one return value is not supported")
			continue
		}
		if len(ftype.ParamTypes) > 0 {
			log.Printf("running exported functions with input parameters is not supported")
			continue
		}
		o, err := vm.ExecCode(i)
		if err != nil {
			fmt.Printf("\n")
			log.Printf("err=%v", err)
		}
		if len(ftype.ReturnTypes) == 0 {
			fmt.Printf("\n")
			continue
		}
		fmt.Printf("%[1]v (%[1]T)\n", o)
	}
}

func importer(name string) (*wasm.Module, error) {
	f, err := os.Open(name + ".wasm")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m, err := wasm.ReadModule(f, nil)
	if err != nil {
		return nil, err
	}
	err = validate.VerifyModule(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
