package conf

import (
	"log"
	"os"
	"path/filepath"
)

// ShellConf encapsulates the configurable
// aspects of the shell
type ShellConf struct {
	Ps1          string
	ExecutionDir string
	CaptureMeta  bool
}

// GetDefaultConf returns a default shell
// configuration
func GetDefaultConf() ShellConf {
	return ShellConf{"$ > ", getCurrentDir(), true}
}

func getCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	return dir
}
