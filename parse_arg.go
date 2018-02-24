package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("%s <command> ... \n\n", os.Args[0])
	fmt.Println("command:")
}

func parseArg() (string, string, bool) {
	fileName  := flag.String("file", "", "file name")
	command   := flag.String("cmd", "", "exec commands")
	isRestart := flag.Bool("restart", false, "restart flag")
	flag.Parse()

	if flag.NFlag() == 0 {
		usage()
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *fileName == "" && *command == "" {
		fmt.Fprintln(os.Stderr, "filename or command require input")
	}
	return *fileName, *command, *isRestart
}
