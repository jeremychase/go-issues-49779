package main

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func main() {
	run()
}

func run() {
	cmd := exec.Command("./test.sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &unix.SysProcAttr{
		Cloneflags: unix.CLONE_NEWTIME,
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}
