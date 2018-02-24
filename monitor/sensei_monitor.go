package monitor

import (
	//"errors"
	"bufio"
	"fmt"
	"github.com/kr/pty"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const (
	O_WRONLY = syscall.O_WRONLY
	O_APPEND = syscall.O_APPEND
	O_CREATE = syscall.O_CREAT
)

type Monitor struct {
	Config         Config `yaml:"config"`
	writeToConsole bool
}

type Config struct {
	Name      string   `yaml:"name"`
	Command   string   `yaml:"command"`
	Args      []string `yaml:"args"`
	IsRestart bool     `yaml:"is_restart"`
}

func writeLog(dst string, src io.Reader) {
	go func() {
		log.Printf("Write log: %s", dst)
		f, err := os.OpenFile(dst, O_WRONLY|O_APPEND|O_CREATE, 0666)
		if err != nil {
			log.Println("Open file failed: %s", dst)
		}
		defer f.Close()

		writer := bufio.NewWriter(f)
		defer writer.Flush()

		if _, err := io.Copy(writer, src); err != nil {
			log.Fatal(err)
		}
		return
	}()
}

func (monitor *Monitor) Run() {
	func() {
		for {
			monitor.Printf("exec %s: ", monitor.Config.Command)
			cmd := exec.Command(monitor.Config.Command, monitor.Config.Args...)
			f, err := pty.Start(cmd)
			fmt.Println("======== Process start ========")

			//writeLog("error.log", e)
			writeLog("output.log", f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "exec error. please confirm commads")
			}

			state, err := cmd.Process.Wait()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error termination: %s", err)
			} else if state.Exited() {
				fmt.Println("======== Process exited ========")
				if monitor.Config.IsRestart {
					continue
				} else {
					return
				}
			}
			return
		}
	}()
}

func (monitor *Monitor) Println(message string) {
	if monitor.writeToConsole {
		fmt.Println(message)
	}
}

func (monitor *Monitor) Printf(message string, a ...interface{}) {
	if monitor.writeToConsole {
		fmt.Printf(message, a)
	}
}

func Create(configFile string, command string, isRestart bool) (monitor *Monitor, err error) {
	if configFile != "" {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatal(err)
		}
		yaml.Unmarshal(data, &monitor)
	} else {
		monitor = &Monitor{}
		monitor.Config.Command = command
		monitor.Config.IsRestart = isRestart
	}
	return
}

