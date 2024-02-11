package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	rootCmd := flag.NewFlagSet("test", flag.ExitOnError)
	rootPrint := rootCmd.Bool("print", false, "print")

	if len(os.Args) < 3 {
		fmt.Println("Expected the 'printStuff' command")
		os.Exit(1)
	}
	switch os.Args[2] {
	case "-print":
		rootCmd.Parse(os.Args[2:])
		fmt.Println("printStuff", *rootPrint)
	}
}
