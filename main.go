package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DLag/jsonlang-test/interpreter"
	"github.com/DLag/jsonlang-test/parser"
	"log"
	"io"
)

var getFiles = func() []string {
	return flag.Args()
}

var output = io.Writer(os.Stdout)

func main() {
	var maxDepth int
	var help bool
	flag.IntVar(&maxDepth, "maxdepth", 10, "maximum stack depth")
	flag.IntVar(&maxDepth, "d", 10, "maximum stack depth")
	flag.BoolVar(&help, "help", false, "show this help")
	flag.BoolVar(&help, "h", false, "show this help")
	flag.Parse()
	scripts := getFiles()
	if help {
		fmt.Println("JSON FSL interpreter")
		fmt.Println("Usage: " + flag.Arg(0) + " [OPTIONS]... [FILES]...")
		fmt.Println("\t-d, --maxdepth <num>\tSet maximum execution stack depth")
		fmt.Println("\t-h, --help\t\tShow this help")
		os.Exit(0)
	}
	globalFuncs := interpreter.NewFuncScope(maxDepth)
	globalVars := interpreter.NewVarScope()
	for _, script := range scripts {
		fmt.Printf("Executing script %q\n", script)
		f, err := os.Open(script)
		if err != nil {
			log.Fatalf("Can't open file %q", script)
		}
		if err = parser.Parse(f, globalVars, globalFuncs); err != nil {
			log.Fatal("Error parsing script: ", err)
		}
		if err = f.Close(); err != nil {
			log.Fatal("Error when closing file: ", err)
		}
		i := interpreter.NewInterpreter(globalVars, globalFuncs, output, maxDepth)
		if err = i.Run(); err != nil {
			log.Fatal("Error executing script: ", err)
		}
	}
}
