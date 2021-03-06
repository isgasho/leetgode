package leetgode

import (
	"context"
	"fmt"
)

type CmdName string

const (
	LIST     CmdName = "list"
	PICK             = "pick"
	GENERATE         = "generate"
	TEST             = "test"
	EXEC             = "exec"
	HELP             = "help"
)

type Cmd interface {
	Name() string
	Usage() string
	MaxArg() int
	Run(context.Context, []string) error
}

var CmdMap = map[CmdName]Cmd{
	EXEC:     &ExecCmd{},
	LIST:     &ListCmd{},
	GENERATE: &GenerateCmd{},
	TEST:     &TestCmd{},
	PICK:     &PickCmd{},
	HELP:     &HelpCmd{},
}

func buildPath(id, slug string) string {
	// TODO: changeable directory
	format := "%s.%s.go"
	return fmt.Sprintf(format, id, slug)
}
