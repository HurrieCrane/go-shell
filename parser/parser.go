package parser

import (
	"fmt"
	"hurriecrane/shell/conf"
	"hurriecrane/shell/grammar"
	"hurriecrane/shell/meta"
	"hurriecrane/shell/parser/ast"
	"strings"
)

// Parser struct contains
// information about the
// compiler
type Parser struct {
	Name string
}

// Parsed s
type Parsed struct {
	// The current command
	Command grammar.CMD
	Meta    meta.ParsingMeta
}

// Parse causes a string of text
// to be parsed into
func (p Parser) Parse(conf *conf.ShellConf, str *string) (Parsed, error) {
	var s ast.Token
	fmt.Println(s.Value)

	parsed := Parsed{}
	if conf.CaptureMeta {
		parsed.Meta.StartRecording()
	}

	parts := strings.Split(*str, " ")
	for i := 0; i < len(parts); i++ {
		cmd := grammar.ReverseLookupString(&parts[i])
		fmt.Println(cmd)
	}

	if conf.CaptureMeta {
		parsed.Meta.StopRecording()
		fmt.Println("Parsing took " + string(parsed.Meta.GetParsingTime()))
	}
	return Parsed{}, nil
}

func ProcessLexeme() {

}
