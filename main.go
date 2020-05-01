package main

import (
	"fmt"
	"hurriecrane/shell/conf"
	"hurriecrane/shell/sh"
)

func main() {
	var globalShell sh.Shell

	// get default shell configuration
	defaultShellConf := conf.GetDefaultConf()

	// Initilise shell
	globalShell.InitiliseShell(&defaultShellConf)

	status, err := globalShell.StartShell()
	if err != nil {
		fmt.Printf(err.Error())
	}

	switch status {
	case 0:
		fmt.Println("Status code 0")
	case 1:
		fmt.Println("Status code 1")
	default:
		fmt.Printf("Unknown statuscode code: %v", status)
	}
}
