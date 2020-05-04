package sh

import (
	"bufio"
	"fmt"
	"hurriecrane/shell/conf"
	"hurriecrane/shell/parser"
	"os"
	"strings"
)

// Shell struct contains information
// about the shell
type Shell struct {
	version string
}

// Config is a global configuration
// object
var Config *conf.ShellConf

var p parser.Parser
var bufReader *bufio.Reader

// InitiliseShell will initilise and configure everything
// required to run the shell
//
func (s Shell) InitiliseShell(conf *conf.ShellConf) {
	// create reference to config
	Config = conf
	// create reader
	bufReader = bufio.NewReader(os.Stdin)
}

// StartShell s
func (s Shell) StartShell() (int, error) {
	var status int64 = -1
	for status < 0 {
		// print ps1
		fmt.Print(Config.Ps1)

		// read in line
		line, err := bufReader.ReadString('\n')
		if err != nil {
			errStr := err.Error()
			if errStr != "EOF" {
				fmt.Println("ERROR - Could read from stdin")
				fmt.Println(err.Error())
				status = 1
			}

			status = 0

			break
		}

		// actually process the input
		pline := strings.Trim(line, "\n\r")
		parserOutput, err := p.Parse(Config, &pline)
		fmt.Println(parserOutput)
	}

	return 0, nil
}
