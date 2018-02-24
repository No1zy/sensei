package main

import (
	"flag"
	"os"
	"fmt"
)

func usage() {
	fmt.Printf("%s <command> ... \n\n", os.Args[0])
	fmt.Println("command:")
}

func parseArg() {
	fileName := flag.String("file", "config.yaml", "file name")
	commands := flag.String("cmd", "", "exec commands")
	isRestart := flag.Bool("restart", false, "restart flag")
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("%s", fileName)
	fmt.Println("%s", commands)
	fmt.Println("%v", isRestart)
}