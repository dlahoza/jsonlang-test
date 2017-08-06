package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DLag/jsonlang-test/interpreter"
	"github.com/DLag/jsonlang-test/parser"
)

func main() {
	var maxDepth int
	var help bool
	flag.IntVar(&maxDepth, "maxdepth", 10, "maximum stack depth")
	flag.IntVar(&maxDepth, "d", 10, "maximum stack depth")
	flag.BoolVar(&help, "help", false, "show this help")
	flag.BoolVar(&help, "h", false, "show this help")
	flag.Parse()
	if help {
		fmt.Println("JSON FSL interpreter")
		fmt.Println("Usage: " + flag.Arg(0) + " [OPTIONS]... [FILES]...")
		fmt.Println("\t-d, --maxdepth <num>\tSet maximum execution stack depth")
		fmt.Println("\t-h, --help\t\tShow this help")
		os.Exit(0)
	}
	scripts := flag.Args()
	internalFuncs := interpreter.NewInternalFuncScope(maxDepth)
	globalFuncs := interpreter.NewFuncScope(maxDepth)
	globalVars := interpreter.NewVarScope()
	for _, script := range scripts {
		fmt.Printf("Executing script %q\n", script)
		f, err := os.Open(script)
		defer f.Close()
		localVars := interpreter.NewVarScope()
		if err = parser.Parse(f, globalVars, globalFuncs); err != nil {
			fmt.Println("Error parsing script: ", err)
			os.Exit(1)
		}
		if err = globalFuncs.Execute("init", globalVars, localVars, internalFuncs, globalFuncs, 0); err != nil {
			fmt.Println("Error executing script: ", err)
			os.Exit(1)
		}
	}
}
