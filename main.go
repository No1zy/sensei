package main

import (
	"github.com/No1zy/sensei/monitor"
	"log"
)

func main() {
	fileName, command, isRestart := parseArg()

	m, err := monitor.Create(fileName, command, isRestart)
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}
