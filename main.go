package main

import (
	monitor "github.com/No1zy/proc-monitor/monitor"
)

func main() {
	monitor, _ := monitor.Create("sample.yml")
	monitor.Run()
}
