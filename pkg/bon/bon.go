package bon

import (
	_ "embed"
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

//go:embed snippets/help.md
var helpdoc string

var BonCmd = &Z.Cmd{
	Name:        "bonzai",
	Summary:     "bonzai helper",
	Aliases:     []string{"bon"},
	Description: helpdoc,
	Commands: []*Z.Cmd{
		help.Cmd,
		commandCmd,
	},
}

//go:embed snippets/command.md
var commandDoc string

//go:embed snippets/command_template.gotxt
var commandTemplate string

var commandCmd = &Z.Cmd{
	Name:        "command",
	Aliases:     []string{"cmd"},
	Summary:     "create bonzai stubs",
	Description: commandDoc,
	Commands:    []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, _ ...string) error {
		fmt.Println(commandTemplate)
		return nil
	},
}
